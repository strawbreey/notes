---
title: "Php Fpm"
date: 2020-09-07T17:25:15+08:00
draft: false
tags: ['php']

---

PHP-FPM是一个PHP的PHPFastCGI管理器。

PHP-FPM提供了更好的PHP进程管理方式，可以有效控制内存和进程、可以平滑重载PHP配置，比spawn-fcgi具有更多优点，所以被PHP官方收录了。在./configure的时候带 –enable-fpm参数即可开启PHP-FPM。

```shell
kill -INT `cat /var/run/php-fpm/php-fpm.pid`

# 杀死进程
killall php-fpm

# 开启进程
/usr/local/php72/sbin/php-fpm

# 查看端口号
netstat -lnt | grep 9000

# 查看php-fpm
ps aux | grep php-fpm

# 重新启动并加载配置
/usr/local/php/sbin-php-fpm -c php.ini
````

1 修改php.ini或www.conf？
修改php.ini文件：

$ php --ini                               # 确定php.ini文件的位置
$ sudo vi /etc/php.ini                    # 修改php.ini文件
修改php-fpm.conf文件：

$ sudo vi /etc/php-fpm/www.conf
编辑之后保存。

2 CentOS/RHEL 7
$ sudo systemctl start php-fpm      # 启动php-fpm
$ sudo systemctl stop php-fpm       # 停止php-fpm
$ sudo systemctl reload php-fpm     # 重载php-fpm
$ sudo systemctl restart php-fpm    # 重启php-fpm
3 CentOS/RHEL 6.x等旧版本
$ sudo service php-fpm start        # 启动php-fpm
$ sudo service php-fpm stop         # 停止php-fpm
$ sudo service php-fpm restart      # 重启php-fpm
$ sudo service php-fpm reload       # 重载php-fpm
4 Ubuntu/Debian
$ sudo service php5-fpm start       # 启动
$ sudo service php5-fpm stop        # 停止
$ sudo service php5-fpm restart     # 重启
$ sudo service php5-fpm reload      # 重载
如果系统使用systemd，比如Ubuntu Linux 16.04+ LTS或者Debian Linux 8.x+，可以这样：

$ sudo systemctl start php5-fpm.service        # 启动
$ sudo systemctl stop php5-fpm.service         # 停止
$ sudo systemctl restart php5-fpm.service      # 重启
$ sudo systemctl reload php5-fpm.service       # 重载
5 Ubuntu/Debian操作php7.0-fpm
$ sudo service php7.0-fpm start
$ sudo service php7.0-fpm stop
$ sudo service php7.0-fpm restart
$ sudo service php7.0-fpm reload
如果系统使用systemd，比如Ubuntu Linux 16.04+ LTS或者Debian Linux 8.x+，可以这样：

$ sudo systemctl start php7.0-fpm.service
$ sudo systemctl stop php7.0-fpm.service
$ sudo systemctl restart php7.0-fpm.service
$ sudo systemctl reload php7.0-fpm.service
6 Alpine Linux
# /etc/init.d/php-fpm start
# /etc/init.d/php-fpm stop
# /etc/init.d/php-fpm restart
7 FreeBSD Unix
# /usr/local/etc/rc.d/php-fpm start
# /usr/local/etc/rc.d/php-fpm stop
# /usr/local/etc/rc.d/php-fpm reload
# /usr/local/etc/rc.d/php-fpm restart
或者用service命令：

# service php-fpm start
# service php-fpm stop
# service php-fpm restart
# service php-fpm reload

