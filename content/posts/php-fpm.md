---
title: "Php Fpm"
date: 2020-09-07T17:25:15+08:00
draft: true
---

PHP-FPM是一个PHP的PHPFastCGI管理器。

PHP-FPM提供了更好的PHP进程管理方式，可以有效控制内存和进程、可以平滑重载PHP配置，比spawn-fcgi具有更多优点，所以被PHP官方收录了。在./configure的时候带 –enable-fpm参数即可开启PHP-FPM。

kill -INT `cat /var/run/php-fpm/php-fpm.pid`

杀死进程
killall php-fpm

开启进程
/usr/local/php72/sbin/php-fpm