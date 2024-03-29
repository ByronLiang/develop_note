# IM消息处理

## 处理高频率的消息接发

1. 对发送消息进行缓存化，可采用进程缓存（LRU策略，易丢失数据）Redis存储：有序集合 member: 消息结构体 score: 时间戳
更高阶设计：进程缓存与Redis双结合方式

2. 每个对话，每个群组，使用一个集合存放数据

3. 定时从缓存区持久化数据；识别最近发送消息的集合，并从集合消费数据

### 消息结构体

1. 消息时间戳+消息Seq组合：确保消息发送顺序一致
   MsgSeq: 单线程自增取号器；确保高并发的消息发送下，相同时间戳下，MsgSeq能判断消息的先后顺序；MsgSeq越小，发送时刻越早

### 历史消息读取

1. 拉取历史消息时, 由于一直处于对话，消息数据是呈增量状态；使用分页`(offset/limit)`拉取数据，会产生重复消息数据

2. 应基于首次拉取历史消息的数据行为基准，进行分页查询；
   若首次拉取15条消息，最新消息的ID=1000, 则之后拉取消息从 ID<=1000, 进行分页查询，`where id <= 1000 offset=15, limit=15`

3. 每次基于MsgKey=消息时间戳+消息Seq组合, 发起查询：`where msg_timestamp <= 消息时间戳 and msg_seq < 消息Seq`
   基于上一次查询，再拉取更旧的消息记录
