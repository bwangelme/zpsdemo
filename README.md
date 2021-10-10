ZipKin Server Demo
===

## Build & Running ZipKin

### Build

```shell
cd /home/xuyundong/.local/zipkin-2.22.1
# 构建 Server 和 它的依赖
./mvnw -DskipTests --also-make -pl zipkin-server clean install
```

### Running

运行 zipkin Server，使用 kafka 作为 collector

+ `KAFKA_BOOTSTRAP_SERVERS=127.0.0.1:9092 STORAGE_TYPE=elasticsearch ES_HOSTS=http://localhost:9200 java -jar ./zipkin-server/target/zipkin-server-2.22.1-exec.jar`

+ __注意__: 构建的时候要关掉 maven 的镜像，使用 apache.org 源，否则会出现下载依赖包的错误: [Authorization failed](https://gist.github.com/bwangelme/af849ca7553f63fa433e593738377457)


## 待阅读

+ [全链路监控（一）：方案概述与比较](https://juejin.cn/post/6844903560732213261#heading-0)
+ [一文读懂微服务监控之分布式追踪](https://zhuanlan.zhihu.com/p/77139483)
+ [Twitter 的 RPC 框架 Finagle 简介](https://www.infoq.cn/article/2014/05/twitter-finagle-intro)
+ https://medium.com/oracledevs/setup-a-distributed-tracing-infrastructure-with-zipkin-kafka-and-cassandra-d0a68fb3eee6
+ https://www.confluent.io/blog/importance-of-distributed-tracing-for-apache-kafka-based-applications/

## 参考链接

+ [zipkin QuickStart](https://zipkin.io/pages/quickstart.html)
+ [Zipkin指南](https://shenbaise9527.com/opentracing-zipkin-guide/#storage)
+ [zipkin Architecture](https://zipkin.io/pages/architecture.html)
+ [zipkin server readme](https://github.com/openzipkin/zipkin/tree/master/zipkin-server#endpoints)
