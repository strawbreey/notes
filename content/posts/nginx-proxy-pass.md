---
title: "Nginx Proxy Pass"
date: 2020-11-09T16:03:57+08:00
draft: true
---

```conf
  server {
    listen       80;
    server_name  cloudwave.cn;

    location /productlist {
      # 匹配所有以/productlist开头所有页面都被拦截
      # 将匹配到的页面请求转发到新的静态资源服务器
      proxy_pass http://new.cloudwave.cn;
    }
  }
  server {
    listen       443;
    server_name  cloudwave.cn;

    location /productlist {
      proxy_pass https://new.cloudwave.cn;
    }
  }
```
