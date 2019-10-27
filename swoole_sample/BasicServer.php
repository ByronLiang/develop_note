<?php

/**
 * 基本监听类
 */
class BasicServer
{
    private $server;

    public function __construct()
    {
        $this->server = new Swoole\Server('127.0.0.1', 9501);

        $this->multiListenPorts();

        $this->server->on('Receive', [$this, 'onReceive']);

        $this->server->start();
    }

    protected function multiListenPorts()
    {
        for($port = 9502; $port <= 9505; $port ++) {
            $this->server->listen("127.0.0.1", $port, SWOOLE_SOCK_TCP);
        }
    }

    protected function getAllClientList($chunk_size = 10)
    {
        $clients = [];
        $start_fd = 0;
        while(true) {
            $conn_list = $this->server->getClientList($start_fd, $chunk_size);
            if ($conn_list===false or count($conn_list) === 0) {
                break;
            }
            $clients = array_merge($clients, $conn_list);
            $start_fd = end($conn_list);
        }

        return $clients;
    }

    public function onReceive($serv, $fd, $reactor_id, $data)
    {
        $receive_data = trim($data);
        $info = $serv->getClientInfo($fd);
        if ($receive_data == 'xxx') {
            $res = $serv->send($fd, "hello ".$data. "\n");
            if ($res == false) {
                echo "error\n";
            }
        }
        $this->broadcastMsg($fd, $receive_data);
    }

    protected function broadcastMsg($expect_fd, $data)
    {
        $conn_list = $this->getAllClientList(1);
        foreach($conn_list as $fd_data) {
            // 发送者不被广播处理
            if ($fd_data != $expect_fd) {
                // 对连接的其他客户端进行广播
                $res = $this->server->send($fd_data, "broadcast: ".$data . "\n");
                if ($res == false) {
                    echo "error\n";
                }
            }
        }
    }
}

$server = new BasicServer();
