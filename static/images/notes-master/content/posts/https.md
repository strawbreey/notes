---
title: "Https"
date: 2020-11-02T16:49:57+08:00
draft: false
---

![http protocol](/images/1_vraaL2BGqdT8tjtxJlzU-A.jpeg)

HTTPS实际上是HTTP穿上了SSL/TLS的外套，我们都知道HTTP属于应用层的协议，是离程序员比较近的，而SSL（安全套接字）与TLS（运输层安全）都属于运输层协议。



![http history](/images/v2-9659d122f0993ba815e3c637661687ed_b.jpg)

随着网络技术的发展，1999 年设计的 HTTP/1.1 已经不能满足需求，所以 Google 在 2009 年设计了基于 TCP 的 SPDY，后来 SPDY 的开发组推动 SPDY 成为正式标准，不过最终没能通过。不过 SPDY 的开发组全程参与了 HTTP/2 的制定过程，参考了 SPDY 的很多设计，所以我们一般认为 SPDY 就是 HTTP/2 的前身。无论 SPDY 还是 HTTP/2，都是基于 TCP 的，TCP 与 UDP 相比效率上存在天然的劣势，所以 2013 年 Google 开发了基于 UDP 的名为 QUIC 的传输层协议，QUIC 全称 Quick UDP Internet Connections，希望它能替代 TCP，使得网页传输更加高效。后经提议，互联网工程任务组正式将基于 QUIC 协议的 HTTP （HTTP over QUIC）重命名为 HTTP/3。


### SSL/TLS 简述


### SSL/TLS 加密解密


### 参考资料


[HTTP/3 原理实战](https://zhuanlan.zhihu.com/p/143464334)