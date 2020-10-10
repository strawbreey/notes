---
title: "Nginx Rewrite"
date: 2020-10-10T10:54:06+08:00
draft: true
---

Rewrite 是nginx 服务器提供的一个重要的基本功能，用于URL的重写


#### 地址重写和转发

"地址重写"，实际上是为了实现地址标准化。那么，什么是地址标准化呢? 我们来举一个例子。比如在访问Google首页的时候，我们在地址栏中可以输入www.google.com，也可以输入google.cn，它们都能够准确地指向Google首页，从客户端来看，Google首页同时对应了两个地址，实际上，Google服务器是在不同的地址中选择了确定的一个，即www.google.com，进而返回服务器响应的。这个过程就是地址标准化的过程。google.cn这个地址在服务器中被改变为www.google.com的过程就是地址重定向的过程。

#### Rewrite规则

6.3.1 域名跳转

```
server {
  listen 80;
  server_name jump.web.name
  rewrite ^/ http://www.web.com # 域名跳转
}
```