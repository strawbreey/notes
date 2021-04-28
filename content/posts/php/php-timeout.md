---
title: "Php Timeout"
date: 2020-10-20T18:16:26+08:00
draft: false
tags: ['php']

---

设置PHP脚本执行超时的时间有下面这样一些方法：

php.ini 中限定程序的最长执行时间是 30 秒，这是由 php.ini 配置文件中的 max_execution_time 变量指定，倘若你有一个需要颇多时间才能完成的程序代码，代码会由于超时而执行失败，例如要发送很多电子邮件给大量收件者，或者要进行繁重的数据分析工作，服务器会在 30 秒后强行中止正在执行的程序，如何解决这个问题呢。

1. 在php.ini里面设置 max_execution_time = 1800; 脚本被解析器中止之前允许的最大执行时间

2. 通过PHP的 ini_set 函数设置 ini_set("max_execution_time", "1800");

3. 通过set_time_limit 函数设置 set_time_limit(1800) ;

4. 在php-fpm.conf中设置 request_terminate_timeout 请求开始n秒

注意

- request_terminate_timeout 适用于，当max_execution_time由于某种原因无法终止脚本的时候，会把这个php-fpm请求干掉。

- web请求php执行时间受到2方面控制，一个是php.ini的max_execution_time（要注意的是sleep，http请求等待响应的时间是不算的，这里算的是真正的执行时间），另一个是php-fpm request_terminate_timeout 设置，这个算的是请求开始n秒。

- nginx的关键参数是 fastcgi 相关的 timeout，即：fastcgi_connect_timeout，fastcgi_read_timeout，fastcgi_send_timeout

这几个 nginx 参数的主语都是 nginx，所以 fastcgi_connect_timeout 的意思是 nginx 连接到 fastcgi 的超时时间，fastcgi_read_timeout 是 nginx 读取 fastcgi 的内容的超时时间，fastcgi_send_timeout 是 nginx 发送内容到 fastcgi 的超时时间。


- Nginx 504 Gateway Time-out的含义是没有请求到可以执行的PHP-CGI。

- Nginx 502 Bad Gateway的含义是请求的PHP-CGI已经执行，但是由于读取资源的等没有执行完毕而导致PHP-CGI进程终止。