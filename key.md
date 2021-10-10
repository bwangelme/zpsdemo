部署 zipkin 的几个关键问题
===

## 部署

+ 数据在 es 中是如何存储的

### 传输时不同 App 使用不同的 Topic

+ zipkin server
    + 启动时 `KAFKA_TOPIC` 环境变量可以设置以`,`分隔的 topic 列表
    + 启动时 `KAFKA_BOOTSTRAP_SERVERS` 环境变量可以设置以`,`分隔的 kafka server 列表
+ golang zipkin client
    + zipkin-go/reporter/kafka 可以设置 kafka 的 server 地址和 topic
+ python zipkin client (TODO)

## 规模

+ zipkin 的 kafka 支持多大的并发量
+ 存储 10W 个 span，es 需要多大的磁盘空间

## 监控

+ kafka 的 topic 中消息的数量
+ zipkin collector 记录收到的 span / trace 数量
+ 按照 app + url 分别记录发送的 span/trace 数量
