---
title: "Linux Firewall"
date: 2020-10-09T19:45:59+08:00
draft: false
---

###  curl: (7) Failed connect to 10.0.0.210:9200; 拒绝连接

###  没有到主机的路由这种问题很常见，多数是由机器的防火墙没有关闭。

###  Ubuntu

####  查看防火墙状态

     ufw status

####  关闭防火墙

     ufw disable

###  centos6

####  查看防火墙状态

     service iptables status

####  关闭防火墙

     chkconfig iptables off

###  centos7

####  查看防火墙状态

     firewall-cmd --state

####  临时关闭防火墙

     systemctl stop firewalld.service

####  永久关闭防火墙

     systemctl disable firewalld