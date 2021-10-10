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

+ https://shenbaise9527.com/opentracing-zipkin-guide/#storage
+ https://www.lixueduan.com/post/tracing/02-framework-compare/
