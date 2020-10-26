---
title: "Rpc and Http"
date: 2020-09-16T10:39:40+08:00
draft: true
---

rpc

### 标准化
为了允许不同的客户端均能访问服务器，许多标准化的 RPC 系统应运而生了。其中大部分采用IDL(Interface Description Language)，方便跨平台的远程过程调用。



### OSI网络七层模型
OSI的七层网络结构模型（虽然实际应用中基本上都是五层），它可以分为以下几层： （从上到下）

第一层：应用层。定义了用于在网络中进行通信和传输数据的接口； (HTTP)
第二层：表示层。定义不同的系统中数据的传输格式，编码和解码规范等；
第三层：会话层。管理用户的会话，控制用户间逻辑连接的建立和中断；

第四层：传输层。管理着网络中的端到端的数据传输；（TCP/IP）
第五层：网络层。定义网络设备间如何传输数据；
第六层：链路层。将上面的网络层的数据包封装成数据帧，便于物理层传输；
第七层：物理层。这一层主要就是传输这些二进制数据。

实际应用过程中，五层协议结构里面是没有表示层和会话层的。应该说它们和应用层合并了。我们应该将重点放在应用层和传输层这两个层面。因为HTTP是应用层协议，而TCP是传输层协议。好，知道了网络的分层模型以后我们可以更好地理解为什么RPC服务相比HTTP服务要Nice一些！


### RPC服务

分别是RPC架构，同步异步调用以及流行的RPC框架。


### References
[Remote Procedure Call](/pages/remote-procedure-call)
[gRPC](https://en.wikipedia.org/wiki/GRPC)
[深入剖析通信层和 RPC 调用的异步化](https://juejin.im/post/6844903761794564104)
[](https://mp.weixin.qq.com/mp/appmsgalbum?action=getalbum&album_id=1358237826197962753&__biz=MzUzNTY5MzU2MA==#wechat_redirect)