---
title: "Apache Forbidden"
date: 2021-01-06T10:07:34+08:00
draft: false
---

apache服务出现问题的解决方法总结: 

在配置Linux的 Apache服务时，经常会遇到 Forbidden 403（拒绝访问）错误，总结了一下有以下4种原因:

- 访问的文档权限不够。要755以上权限。解决方法：用命令chmod 755 /var/www/ 或其他相应目录。
- SELinux或防火墙的原因。解决方法：先关闭SELinux和让防火墙通过WWW服务。
- 虚拟主机配置错误。例如我遇到过一次的：
- httpd.conf里加载了虚拟主机的配置文件：