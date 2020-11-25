---
title: "Php Find Php"
date: 2020-10-10T10:05:08+08:00
draft: false
---

Linux下查找及修改PHP配置文件ini的路径

一、查找PHP配置文件
说到查找，当然首先想到的是find命令。执行如下命令，即可查找到php.ini文件

1
sudo find / -name php.ini
01.png

可是，找到三个php.ini文件，具体哪个是当前正在运行的PHP使用的配置文件呢？ PHP提供了两种方式，可供使用。

方法一

这个是比较简单的方法，使用如下命令，可以清楚的看出当前的php使用的配置文件。

1
php --ini
01.png

可以使用php --help查看此命令

01.png

方法二

打印出phpinfo()，然后就可以看出了，如下：

01.png

针对这个方法，也可以在命令行中查看，有如下两种方式：

1
2
3
php -i  |grep php.ini  //php -i其实就是输出phpinfo
 
php -r "phpinfo();" |grep php.ini
01.png



二、修改PHP配置文件
很遗憾，尝试了一下，发现貌似只能在安装php的时候，通过--with-config-file-path来指定，后续进一步探索吧。