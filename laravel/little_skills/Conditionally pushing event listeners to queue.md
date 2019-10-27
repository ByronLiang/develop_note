# 如何有选择地将事件监听器推送到队列中

## 需求&情景

- 在listener中进行相关判断，并决定是否执行此队列业务；

- 订单完成支付，推送事件进行订单包裹的分发，但对于拼团的订单，需要以成团成功后才完成包裹的分发；

## 解决方案

- 在Listener类里加入`shouldQueue`方法，并以布尔值返回，来有选择地是否执行此队列逻辑
- 有效避免队列中有很多无用的Job在执行；

### 代码

- 分发事件

```
class BuyEvent
{
    use Dispatchable, InteractsWithSockets, SerializesModels;

    public $m;

    public function __construct($m)
    {
        $this->m = $m;
    }
}

```

- 事件监听队列

```
class BuyListener implements ShouldQueue
{
    public function handle($event)
    {
        $m = $event->m;
        \Log::info('money is '. $m);
    }

    public function shouldQueue($event)
    {
        return $event->m > 100;
    }
}

```

- 在shouldQueue方法里加入对m变量的判断；当变量m大于100，才执行队列内容