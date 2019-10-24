<?php
$serv = new Swoole\Server('127.0.0.1', 9001);

for($port = 9002; $port < 9005; $port++)
{
    $serv->listen("127.0.0.1", $port, SWOOLE_SOCK_TCP);
}

$serv->on("receive", function($serv, $fd, $reactor_id, $data) {
    $info = $serv->getClientInfo($fd);
    $conn_list = $serv->getClientList(0, 10);
    if (trim($data) == 'xxx') {
    	if ($serv->send($fd, "hello ".$data) == false) {
        	echo "error\n";
    	}
    }
    foreach($conn_list as $fd_data) {
    	// 发送者不被广播处理
    	if ($fd_data != $fd) {
    		// 对连接的其他客户端进行广播
    		if ($serv->send($fd_data, "broadcast: ".$data) == false) {
        		echo "error\n";
    		}
    	}
    }
});

$serv->start();
