---
title: "Http Ip"
date: 2020-12-22T19:53:55+08:00
draft: false
---

公网IP无法访问，显示服务器拒绝了请求？

1. 你的网站能ping通，80端口也是开放的，那就是你服务器没提相应的HTTP服务，不知道你是Apache还是Nginx，先检查下这两个服务是否正常。

应该是没安装http服务

2. 启动http服务

httpd -version #检查是否安装 

yum install httpd #安装http服务

service httpd start #启动http服务

service httpd status #检查http服务状态（会显示绿色的active(running)表示启动成功）

service httpd restart #重启http服务