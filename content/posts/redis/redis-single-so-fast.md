---
title: "Redis Single So Fast"
date: 2021-08-05T17:17:53+08:00
draft: false
---

官方FAQ表示，因为Redis是基于内存的操作，CPU不是Redis的瓶颈，Redis的瓶颈最有可能是机器内存的大小或者网络带宽。既然单线程容易实现，而且CPU不会成为瓶颈，那就顺理成章地采用单线程的方案了（毕竟采用多线程会有很多麻烦！）。

### 参考资料

- https://stackoverflow.com/questions/48035646/if-redis-is-single-threaded-how-can-it-be-so-fast
- [为什么说Redis是单线程的以及Redis为什么这么快！](https://xuliugen.blog.csdn.net/article/details/79470556)
- https://redis.io/topics/faq