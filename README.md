## 启动 kafka

```
# 启动 zk
sudo /usr/local/Cellar/kafka/2.8.0/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties

# 启动 kafka
/usr/local/Cellar/kafka/2.8.0/bin/kafka-server-start /usr/local/etc/kafka/server.properties
```

## 管理 Topic

```
# 创建 topic lazy
kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic lazy

# 列出所有 topic
kafka-topics --list --zookeeper localhost:2181
```

## 发送消息

```
# 创建生产者，发送消息到 lazy 上
kafka-console-producer --broker-list localhost:9092 --topic lazy

# 创建消费者，监听 lazy 的消息
kafka-console-consumer --bootstrap-server localhost:9092 --topic lazy --from-beginning
```
