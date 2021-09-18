---
title: "Common Commands"
date: 2020-08-31T10:36:43+08:00
draft: false
---

### linux common command

```bash

# 测试镜像是否有用
nslookup mirrors.tencent.com

rpm -Uvh https://mirrors.tencent.com/tlinux/rpm/tlinux-release-2-11.tl2.x86_64.rpm

# 下载php
wget php-7.2.15.tar.bz2 http://cn2.php.net/distributions/php-7.2.15.tar.bz2

# 安装
yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm

rpm -Uvh  https://mirrors.tencent.com/tlinux/rpm/epel-release-6-12.tl1.noarch.rpm

# linux cp 复制
sudo cp /opt/soft/php/sbin/php-fpm /usr/local/bin/        sudo php-fpm

# linux  杀死进程 
kill -INT cat /usr/local/php/var/run/php-fpm.pid

# linux  安装
yum -y install gcc openssl openssl-devel curl curl-devel libjpeg libjpeg-devel libpng libpng-devel freetype freetype-devel pcre pcre-devel libxslt libxslt-devel bzip2 bzip2-devel

# linux 杀死全部进程
killall php-fpm

# linux 查询进程 php
ps -ef|grep php

# 查看php-fpm的位置
whereis php-fpm

php-fpm: /usr/sbin/php-fpm /etc/php-fpm.d /etc/php-fpm.conf /usr/local/bin/php-fpm /usr/share/man/man8/php-fpm.8.gz

# 启动php-fpm
usr/sbin/php-fpm



# linux 重载
nginx -s reload

```

链接数据库异常 

```shell
# Can't connect to local MySQL server through socket '/var/lib/mysql/mysql.sock'
```


