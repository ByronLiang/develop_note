<?php

class ShareData
{
    private $serv;
    private $fds = [];

    public function __construct()
    {
        $this->serv = new Swoole\Server('127.0.0.1', 9501);
        $this->serv->set([
            'worker_num' => 2,
        ]);
        $this->serv->on('Connect', [$this, 'connectHandle']);

        $this->serv->on('Receive', [$this, 'onReceive']);

        $this->serv->start();
    }

    public function connectHandle($server, $fd, $reactor_id)
    {
        $info = $server->connection_info($fd);
        $data = "INFO: fd=$fd, reactor_id=$reactor_id, addr={$info['remote_ip']}:{$info['remote_port']}\n";
        $server->send($fd, $data);
        echo "connection open: {$fd}\n";
        // 只能在当前进程进行共享, 无法跨进程共享数据
        $this->fds[] = $fd;
        var_dump($this->fds);
    }

    public function onReceive($serv, $fd, $reactor_id, $data)
    {
        $receive_data = trim($data);
    }

    // 初始化swoole 存储
    protected function initSwooleTable()
    {
        $table = new swoole_table(1024);
        $table->column('fd', swoole_table::TYPE_INT);
        $table->column('reactor_id', swoole_table::TYPE_INT);
        $table->column('data', swoole_table::TYPE_STRING, 64);
        $table->create();
        $this->serv->table = $table;
    }

    // 以fd为key 存储以set开头的命令
    protected function setAndGetData($serv, $fd, $reactor_id, $data)
    {
        $cmd = explode(" ", trim($data));
	    //get
        if ($cmd[0] == 'get') {
            // get self
            if (count($cmd) < 2) {
                $cmd[1] = $fd;
            }
            $get_fd = intval($cmd[1]);
            $info = $serv->table->get($get_fd);
            $serv->send($fd, var_export($info, true)."\n");
        } elseif ($cmd[0] == 'set') {
            // set
		    $ret = $serv->table->set($fd, array('reactor_id' => $reactor_id, 'fd' => $fd, 'data' => $cmd[1]));
		    if ($ret === false) {
			    $serv->send($fd, "ERROR\n");
		    } else {
			    $serv->send($fd, "OK\n");
	    	}
	    } else {
		    $serv->send($fd, "command error.\n");
	    }
    }
}

$server = new ShareData();