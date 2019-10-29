<?php

use Swoole\Coroutine\Scheduler;
use Swoole\Coroutine\Channel;

class CoScheduler
{
    private $scheduler, $channel, $list;
    
    public function __construct()
    {
        $args = func_get_args();
        $this->scheduler = new Scheduler();
        // 初始化通道容量
        $this->channel = new Channel(count($args));
        foreach ($args as $key => $arg) {
            /**
            * 利用通道存储业务逻辑
            */
            // $this->scheduler->add([$this, 'addChannel'], $key, $arg);
            $this->scheduler->add([$this, $arg], rand(1, 100));
        }
        /**
         * 取出通道内容进行处理
         */
        // $this->list = [];
        // $this->scheduler->add([$this, 'getChannel'], $args);
        $this->scheduler->start();
        echo "finished\n";
        // var_dump($this->list);
    }

    public function addChannel($key, $value)
    {
        $res = $this->{$value}(rand(1, 100));
        $this->channel->push([$key => $res]);
    }

    public function getChannel($args)
    {
        foreach ($args as $key => $chan) {
            $res = $this->channel->pop();
            $this->list[$key] = $res;
        }
    }

    public function task1($value)
    {
        Co::sleep(1);
        echo "id: " . $value . " task1 Done.\n";

        return "id: " . $value . " I'm task1";
    }

    public function task2($value)
    {
        Co::sleep(0.5);
        echo "id: " . $value . " task2 Done.\n";

        return "id: " . $value . " I'm task2";
    }
}

$co_schedule = new CoScheduler('task1', 'task2');