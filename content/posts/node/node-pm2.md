---
title: "Node Pm2"
date: 2021-07-22T19:23:54+08:00
draft: false
---

PM2是node进程管理工具，可以利用它来简化很多node应用管理的繁琐任务，如性能监控、自动重启、负载均衡等，而且使用非常简单。

下面就对PM2进行入门性的介绍，基本涵盖了PM2的常用的功能和配置。

### 安装

```bash
npm install -g pm2
```

目录介绍

pm2安装好后，会自动创建下面目录。看文件名基本就知道干嘛的了，就不翻译了。

$HOME/.pm2 will contain all PM2 related files
$HOME/.pm2/logs will contain all applications logs
$HOME/.pm2/pids will contain all applications pids
$HOME/.pm2/pm2.log PM2 logs
$HOME/.pm2/pm2.pid PM2 pid
$HOME/.pm2/rpc.sock Socket file for remote commands
$HOME/.pm2/pub.sock Socket file for publishable events
$HOME/.pm2/conf.js PM2 Configuration

### 常用命令

参数说明：
```bash
--watch：监听应用目录的变化，一旦发生变化，自动重启。如果要精确监听、不见听的目录，最好通过配置文件。
-i --instances：启用多少个实例，可用于负载均衡。如果-i 0或者-i max，则根据当前机器核数确定实例数目。
--ignore-watch：排除监听的目录/文件，可以是特定的文件名，也可以是正则。比如--ignore-watch="test node_modules "some scripts""
-n --name：应用的名称。查看应用信息的时候可以用到。
-o --output <path>：标准输出日志文件的路径。
-e --error <path>：错误输出日志文件的路径。
--interpreter <interpreter>：the interpreter pm2 should use for executing app (bash, python...)。比如你用的coffee script来编写应用。



# 重启
pm2 restart app.js

# 停止

pm2 stop app_name|app_id #停止特定的应用。可以先通过pm2 list获取应用的名字（--name指定的）或者进程id。


pm2 stop all #如果要停止所有应用，可以

# 删除
pm2 stop app_name|app_id
pm2 stop all

# 查看进程状态
pm2 list

pm2 describe 0
```