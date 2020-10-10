---
title: "Beego"
date: 2020-10-08T02:09:21+08:00
draft: true
---

### deploy

通过 bee 创建的项目，beego 默认情况下是开发模式。

我们可以通过如下的方式改变我们的模式：
```
beego.RunMode = "prod"
```
或者我们在 conf/app.conf 下面设置如下：
```
runmode = prod
```

发行部署

Go 语言的应用最后编译之后是一个二进制文件，你只需要 copy 这个应用到服务器上，运行起来就行。beego 由于带有几个静态文件、配置文件、模板文件三个目录，所以用户部署的时候需要同时 copy 这三个目录到相应的部署应用之下，下面以我实际的应用部署为例：

独立部署

独立部署即为在后端运行程序，让程序跑在后台。

在 linux 下面部署，我们可以利用 nohup 命令，把应用部署在后端，如下所示：

```shell
nohup ./beepkg &
```

Windows
在 Windows 系统中，设置开机自动，后台运行，有如下几种方式：

制作 bat 文件，放在“启动”里面
制作成服务


nginx 部署

```conf
server {
    listen       80;
    server_name  .a.com;

    charset utf-8;
    access_log  /home/a.com.access.log;

    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;

        root "/path/to/app_a/static";
        try_files $uri @backend;
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8080;
    }
}

server {
    listen       80;
    server_name  .b.com;

    charset utf-8;
    access_log  /home/b.com.access.log  main;

    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;

        root "/path/to/app_b/static";
        try_files $uri @backend;
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8081;
    }

```

Apache 配置

apache 和 nginx 的实现原理一样，都是做一个反向代理，把请求向后端传递，配置如下所示：


Supervisord

Supervisord 是用 Python 实现的一款非常实用的进程管理工具，supervisord 还要求管理的程序是非 daemon 程序，supervisord 会帮你把它转成 dae