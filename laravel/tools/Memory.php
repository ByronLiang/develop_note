<?php

/**
 * 只可在Linux环境下进行
 * 进程锁避免并发
 */

class Memory
{
    /**
     * @var string
     */
    protected $sem_id = '';

    /**
     * @var string 信号量
     */
    protected $signal;

    /**
     * Memory constructor.
     *
     * @param array $config
     */
    public function __construct(array $config = [])
    {
        foreach ($config as $key => $value) {
            property_exists($this, $key) && $this->{$key} = $value;
        }

        $this->setSignal();
    }

    /**
     * 加锁
     * 非堵塞获取信号量
     * @return bool|mixed
     */
    public function lock()
    {
        return sem_acquire($this->signal, true);
    }

    /**
     * 解锁
     * 移除信号量
     *
     * @return bool|mixed
     */
    public function unlock()
    {
        return sem_remove($this->signal);
    }

    /**
     * 设置信号量
     */
    public function setSignal()
    {
        if (empty($this->sem_id)) {
            $this->sem_id = ftok(__FILE__, 's');
            echo "The System V IPC key: " . ftok(__FILE__, 's') . "\n";
        }
        $this->signal = sem_get($this->sem_id, 1, 0666, 0);
    }
}

$me = new Memory();
$res = $me->lock();
while (! $res) {
    echo "无法获取锁" . "\n";
    sleep(2);
    $me = new Memory();
    $res = $me->lock();
}
echo "获得锁了"."\n";
sleep(2);
echo "进行中"."\n";
sleep(4);
$me->unlock();
echo "结束"."\n";
