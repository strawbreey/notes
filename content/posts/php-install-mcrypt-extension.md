---
title: "Php Install Mcrypt Extension"
date: 2020-11-27T15:23:47+08:00
draft: false
tags: ['php']
---

在 php 官网 下载 [mcrypt](http://pecl.php.net/package/mcrypt) 包

```bash
wget  http://pecl.php.net/get/mcrypt-1.0.3.tgz
tar xf mcrypt-1.0.3.tgz
cd mcrypt-1.0.3

# 编译安装 mcrypt
/usr/local/php/bin/phpize
./configure --with-php-config=/usr/local/php/bin/php-config  && make && make install

# php.ini 加上扩展 
extension=mcrypt.so

# 重启 php-fpm, 
/etc/init.d/php-fpm restart  # php-fpm 默认位置

/usr/local/php72/sbin/php-fpm # php 编译安装的位置

```



### 参考资料

- [mcrypt](http://pecl.php.net/package/mcrypt)
- [php 7.2 安装 mcrypt 扩展](https://www.jianshu.com/p/3427bc16127e)