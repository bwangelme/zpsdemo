部署 zipkin 的几个关键问题
===

## 目标

### 低消耗

go client 发送数据到 kafka 的 send time < 20ms
python client 发送数据到 kafka 的 sen time < 10ms (python 是同步发送的逻辑，所以时延要求更低)

### 应用级的透明

Python Or Go 应用程序在执行时，能够记录以下操作的指标和操作时间

+ redis operation, time, key
+ kv db operation, time, key
+ mysql sql, time
+ http url, time, method
+ pidl interface, time
+ thrift interface, time

此外，dae 还提供了接口，支持添加自定义的 span

### 延展性

+ kafka 支持不同 app 使用不同的 server 或 topic，可以横向扩展
+ es 是每天将数据存储到同一个 index 中，不知道扩展性是否会有问题

## 部署

### 数据在 es 中是如何存储的

按天创建索引，索引的名称是 `zipkin-span-2021-10-10`

+ `ES_INDEX` 设置索引前缀，默认是 `zipkin`
+ `ES_DATE_SEPARATOR` 索引的日期分割符，默认是 `-`
+ `ES_INDEX_SHARDS` 分片数
+ `ES_INDEX_REPLICAS` 副本数

### 传输时不同 App 使用不同的 Topic

+ zipkin server
    + 启动时 `KAFKA_TOPIC` 环境变量可以设置以`,`分隔的 topic 列表
    + 启动时 `KAFKA_BOOTSTRAP_SERVERS` 环境变量可以设置以`,`分隔的 kafka server 列表
+ golang zipkin client
    + zipkin-go/reporter/kafka 可以设置 kafka 的 server 地址和 topic
+ python zipkin client (TODO)

## 规模

+ zipkin 的 kafka 支持多大的并发量 ?
+ 存储 10W 个 span，es 需要多大的磁盘空间 ?

## 监控

+ kafka 的 topic 中消息的数量
+ zipkin collector 记录收到的 span / trace 数量
+ 按照 app + url 分别记录发送的 span/trace 数量
+ client 发送数据到 kafka 的时间
