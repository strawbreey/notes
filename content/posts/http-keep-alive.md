---
title: "Http Keep Alive"
date: 2020-11-10T16:30:43+08:00
draft: true
---

HTTP/1.1和一部分的HTTP/1.0想出了持久连接（HTTP PersistentConnections，也称为HTTP keep-alive或HTTP connection reuse）的方法。持久连接的特点是，只要任意一端没有明确提出断开连接，则保持TCP连接状态。

![keep-alive](/images/epub_907764_46.jpg)



### 参考资料

- [MDN Keep Alive](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Keep-Alive)