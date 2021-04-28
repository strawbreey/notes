---
title: "Grpc"
date: 2020-12-09T15:01:11+08:00
draft: false
---

Why gRPC?


gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

- Simple service definition

Define your service using Protocol Buffers, a powerful binary serialization toolset and language

-  Start quickly and scale

Install runtime and dev environments with a single line and also scale to millions of RPCs per second with the framework

- Works across languages and platforms

Automatically generate idiomatic client and server stubs for your service in a variety of languages and platforms


- Bi-directional streaming and integrated auth

Bi-directional streaming and fully integrated pluggable authentication with HTTP/2-based transport


### The main usage scenarios
- Efficiently connecting polyglot services in microservices style architecture
- Connecting mobile devices, browser clients to backend services
- Generating efficient client libraries

### Core features that make it awesome
Idiomatic client libraries in 10 languages
Highly efficient on wire and with a simple service definition framework
Bi-directional streaming with http/2 based transport
Pluggable auth, tracing, load balancing and health checking


### RPC 有什么用
你可以使用它：

- 搭建多个端口支持多个协议的服务

- 搭建消息队列消费者服务，提供消息队列生产者客户端发送消息

- 搭建本地定时器，分布式定时器服务。

- 搭建流式服务，实现push，文件上传，消息下发等流式模型。

- 访问各种私有协议后端服务，调用各种存储

- 通过trpc工具生成桩代码和服务模板