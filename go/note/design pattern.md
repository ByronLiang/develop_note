# 常用设计模式

## functional options 函数式选项配置

1. 能灵活对结构体的成员进行动态配置，新增成员与新增配置，无需改动内部代码

2. 配置参数无需按照顺序

```go
type ImClient struct {
    ClientId    string          // 标识ID
    SystemId    string          // 系统ID
    Socket      *websocket.Conn // 用户连接
    ConnectTime int64          // 首次连接时间
    IsDeleted   bool            // 是否删除或下线
    UserId      string          // 业务端标识用户ID
    Extend      string          // 扩展字段，用户可以自定义
    GroupList   []string
    Sub         *redis.PubSub
    Path        string
}

type ImClientOptinon func(client *ImClient)

// 配置ConnectTime成员
func SetConnectTime(time int64) ImClientOptinon {
    return func(client *ImClient) {
        client.ConnectTime = time
    }
}

// 遍历配置回调函数 对结构体参数进行设置
func (c *ImClient) InitOpts(opts... ImClientOptinon) *Client {
    for _, opt := range opts {
        // Call the option giving the instantiated
        // *ImClient as the argument
        opt(c)
    }
    return c
}

```