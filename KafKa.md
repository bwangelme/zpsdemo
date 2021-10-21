## 启动 kafka

```
# 启动 zk
./bin/zookeeper-server-start.sh config/zookeeper.properties

# 启动 kafka
./bin/kafka-server-start.sh config/server.properties
```

## 管理 Topic

```sh
# 创建 topic lazy
./bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic lazy

# 列出所有 topic
./bin/kafka-topics.sh --list --bootstrap-server localhost:9092

# 查看正在同步的 Topic
./bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --under-replicated-partition

# 查看哪些分区不可用
./bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --unavailable-partitions


# 描述单个 topic
Topic: __consumer_offsets       TopicId: UyQtc03qTLKnvp_AbFHexg PartitionCount: 50      ReplicationFactor: 1    Configs: compression.type=producer,cleanup.policy=compact,segment.bytes=104857600
        Topic: __consumer_offsets       Partition: 0    Leader: 0       Replicas: 0     Isr: 0
        Topic: __consumer_offsets       Partition: 1    Leader: 0       Replicas: 0     Isr: 0
        ....
        Topic: __consumer_offsets       Partition: 49   Leader: 0       Replicas: 0     Isr: 0
        
# 输出的表格描述了主题中每个分区的副本所在的位置，0 表示是 broker 服务器的 ID
```

## 发送消息

```
# 创建生产者，发送消息到 lazy 上
kafka-console-producer --broker-list localhost:9092 --topic lazy

# 创建消费者，监听 lazy 的消息
kafka-console-consumer --bootstrap-server localhost:9092 --topic lazy --from-beginning
```

## Kafka 中 replication 和 partitions 指什么

replication: 副本数
partitions: 分区数
