---
title: "Apache Php"
date: 2020-12-28T19:47:59+08:00
draft: false
---

网页解析不了php，可以查看到源码
思路：
因为 apache 解析不了 php，所以要先看 php 有没有加载解析 php 模版
1. httpd -M 查看是否存在  php5_module

2. 如果存在的话， 添加AddType application/x-httpd-php .php 和  将目录的默认索引页面改为index.php 即  DirectoryIndex index.php

### Linux Apache 配置 PHP

httpd主配置文件  `/etc/httpd/conf/httpd.conf`

vim /usr/local/apache2.4/conf/httpd.conf //修改以下4个地方

2. ServerName

当启动服务时,会报这个提示

```
Could not reliably determine the server's fully qualified domain name, using 172.19.0.2. Set the 'ServerName' directive globally to suppress this message
```

解决方案: 

```bash
# 解决办法是:找到这一行,将行首的注释去掉

ServerName 127.0.0.1:80
```


2. Require all denied

打开网页提示403

解决办法:Require all denied将denied改为granted



三、AddType application/x-httpd-php .php

apache若想支持php,需要增加AddType application/x-httpd-php .php #增加到下图所示,如果不增加,php不能解析


四、DirectoryIndex index.html index.php #增加索引页,这样可以在不输入index.php的时候,可以直接访问

检测apache是否解析

网页目录:/usr/local/apache/htdocs/

写一个1.php,内容自定义
<?php
phpinfo();
?>

网站打开,出现全是 源代码的文件,需要

1. 检查是否加载php木块
httpd -M

#如果没有加载则需要检查是否存在此模块,
ls /usr/local/apache/modules/libphp5.so
如果有文件,但是没有加载,需要检查配置文件,
Linux-Apache和PHP结合
之后,检查配置文件是否有增加此行AddType application/x-httpd-php .php
/usr/local/apache/bin/apachectl -t #检查配置文件,语法是否正确,纠错
检查是否此行是否加入index.php(不重要,之前加了索引页)
Linux-Apache和PHP结合


### Apache服务器出现Forbidden 403错误提示的解决方法总结

1. 访问的文档权限不够。要755以上权限。解决方法：用命令chmod 755 /var/www/ 或其他相应目录。
2. SELinux或防火墙的原因。解决方法：先关闭SELinux和让防火墙通过WWW服务。
3. 虚拟主机配置错误。
4. DocumentRoot的设置错误。

