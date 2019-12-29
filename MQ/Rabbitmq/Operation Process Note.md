# RabbitMQ 基本运维笔记

## 应用使用流程

1. 建立vhost、用户登录信息 [多租户特性]
2. 配置用户的访问权限(角色)

### 角色类型

1. none 无角色
2. management 能登陆到Web管理页
3. policymaker 管理policy及参数
4. monitoring 看到所有连接、信道及节点信息
5. administartor 管理用户、虚拟主机、权限、策略与参数
