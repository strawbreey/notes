---
title: "Kafka"
date: 2021-08-05T11:03:05+08:00
draft: false
---

## Kafka

### 快速开始

下载并解压

```bash
tar -xzf kafka_2.11-1.0.0.tgz
cd kafka_2.11-1.0.0
```

### 启动服务器

Kafka 使用 ZooKeeper 如果你还没有ZooKeeper服务器，你需要先启动一个ZooKeeper服务器。 您可以通过与kafka打包在一起的便捷脚本来快速简单地

创建一个单节点ZooKeeper实例。
```bash
bin/zookeeper-server-start.sh config/zookeeper.properties
```

现在启动Kafka服务器：

```bash
 bin/kafka-server-start.sh config/server.properties
```

创建一个 topic

```bash
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
```

现在我们可以运行list（列表）命令来查看这个topic：

```bash
bin/kafka-topics.sh --list --zookeeper localhost:2181
```

发送一些消息

```bash
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
```

5. 启动一个 consumer

Kafka 还有一个命令行consumer（消费者），将消息转储到标准输出。

```bash
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```