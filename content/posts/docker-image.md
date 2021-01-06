---
title: "Docker Image"
date: 2021-01-04T15:36:27+08:00
draft: false
---

1. 解决apache启动错误：Could not reliably determine the server's fully qualified domain name

1）进入apache的安装目录：(视个人安装情况而不同) [root@server ~]# cd /usr/local/apache/conf
 
2）编辑httpd.conf文件，搜索"#ServerName"，添加ServerName localhost:80
[root@server conf]# ls
extra  httpd.conf  magic  mime.types  original
[root@server conf]# vi httpd.conf
#ServerName www.example.com:80
ServerName localhost:80
3）再重新启动apache 即可。
[root@server ~]# /usr/local/apache/bin/apachectl restart