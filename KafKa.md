## 启动 kafka

```
# 启动 zk
./bin/zookeeper-server-start.sh config/zookeeper.properties

# 启动 kafka
./bin/kafka-server-start.sh config/server.properties
```

## 管理 Topic

```
# 创建 topic lazy
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic lazy

# 列出所有 topic
kafka-topics --list --bootstrap-server localhost:9092
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