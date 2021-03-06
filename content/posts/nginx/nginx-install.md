---
title: "Nginx Install"
date: 2020-09-08T17:32:46+08:00
draft: false
---

nginx [engine x] is an HTTP and reverse proxy server, a mail proxy server, and a generic TCP/UDP proxy server, originally written by Igor Sysoev. For a long time, it has been running on many heavily loaded Russian sites including Yandex, Mail.Ru, VK, and Rambler. According to Netcraft, nginx served or proxied 25.75% busiest sites in August 2020. Here are some of the success stories: Dropbox, Netflix, Wordpress.com, FastMail.FM.

## 安装Nginx


```shell

cd /usr/local/src/

# 下载nginx
wget http://nginx.org/download/nginx-1.6.2.tar.gz


# 解压nginx
tar zxvf nginx-1.6.2.tar.gz

# 编译安装
 cd nginx-1.6.2
 ./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
 make & make install

# 查看是否安装
nginx -v

# nginx 测试
/usr/local/webserver/nginx/sbin/nginx -t

# 启动nginx
/usr/local/webserver/nginx/sbin/nginx

```

常用命令

```shell
nginx -s reload            # 重新载入配置文件
nginx -s reopen            # 重启 Nginx
nginx -s stop              # 停止 Nginx
```

#### 相关文章
[nginx.org](http://nginx.org/en/)


