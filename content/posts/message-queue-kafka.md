---
title: "Message Queue Kafka"
date: 2020-11-12T11:29:25+08:00
draft: false
---

![ckafka](/images/ckafka.png)


- 生产者 Producer 可能是网页活动产生的消息、服务日志等信息。生产者通过 push 模式将消息发布到 Cloud Kafka 的 Broker 集群。
- 集群通过 Zookeeper 管理集群配置，进行 leader 选举，故障容错等。
- 消费者 Consumer 被划分为若干个 Consumer Group。消费者通过 pull 模式从 Broker 中消费消息。