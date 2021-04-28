---
title: "Php Fastcgi"
date: 2020-10-27T10:39:26+08:00
draft: false
tags: ['php']

---

### CGI

CGI全称"通用网关接口"（Common Gateway Interface），用于HTTP服务器与其它机器上的程序服务通信交流的一种工具，CGI程序须运行在网络服务器上。

传统CGI接口方式的主要缺点是性能较差，因为每次HTTP服务器遇到动态程序时都需要重启解析器来执行解析，然后结果被返回给HTTP服务器。这在处理高并发访问几乎是不可用的，因此就诞生了FastCGI。另外传统的CGI接口方式安全性也很差。


### FastCGI

FastCGI是一个可伸缩地、高速地在HTTP服务器和动态脚本语言间通信的接口（FastCGI接口在Linux下是socket（可以是文件socket，也可以是ip socket）），主要优点是把动态语言和HTTP服务器分离开来。多数流行的HTTP服务器都支持FastCGI，包括Apache、Nginx和lightpd。

同时，FastCGI也被许多脚本语言所支持，比较流行的脚本语言之一为PHP。FastCGI接口方式采用C/S架构，可以将HTTP服务器和脚本解析服务器分开，同时在脚本解析服务器上启动一个或多个脚本解析守护进程。当HTTP服务器每次遇到动态程序时，可以将其直接交付给FastCGI进程执行，然后将得到的结构返回给浏览器。这种方式可以让HTTP服务器专一地处理静态请求或者将动态脚本服务器的结果返回给客户端，这在很大程度上提高了整个应用系统的性能。

FastCGI的重要特点：

1、FastCGI是HTTP服务器和动态脚本语言间通信的接口或者工具。

2、FastCGI优点是把动态语言解析和HTTP服务器分离开来。

3、Nginx、Apache、Lighttpd以及多数动态语言都支持FastCGI。

4、FastCGI接口方式采用C/S架构，分为客户端（HTTP服务器）和服务端（动态语言解析服务器）。

5、PHP动态语言服务端可以启动多个FastCGI的守护进程。

6、HTTP服务器通过FastCGI客户端和动态语言FastCGI服务端通信。


![Nginx+FastCGI运作过程](/images/5350978-fcd13cc99e17f59a.webp)

FastCGI的主要优点是把动态语言和HTTP服务器分离开来，是Nginx专一处理静态请求和向后转发动态请求，而PHP/PHP-FPM服务器专一解析PHP动态请求。


### FPM
FPM (FastCGI 进程管理器) 是一个可选的 PHP FastCGI 实现并且附加了一些（主要是）对高负载网站很有用的特性。

### CGI、FastCGI、php-fpm 之间的关系

CGI是为了保证web server传递过来的数据是标准格式的，它是一个协议，方便CGI程序的编写者。Fastcgi是CGI的更高级的一种方式，是用来提高CGI程序性能的。

web server（如nginx）只是内容的分发者。比如，如果请求/index.html，那么web server会去文件系统中找到这个文件，发送给浏览器，这里分发的是静态资源。

那PHP-FPM又是什么呢？它是一个实现了Fastcgi协议的程序,用来管理Fastcgi起的进程的,即能够调度php-cgi进程的程序。现已在PHP内核中就集成了PHP-FPM，使用--enalbe-fpm这个编译参数即可。另外，修改了php.ini配置文件后，没办法平滑重启，需要重启php-fpm才可。此时新fork的worker会用新的配置，已经存在的worker继续处理完手上的活。

![cgi](/images/cgi.png)

### 参考链接 

- [FastCGI 进程管理器 ¶](https://www.php.net/manual/zh/book.fpm.php)
- [CGI、FastCGI和PHP-FPM关系图解](https://www.awaimai.com/371.html)
- [如何通俗地解释 CGI、FastCGI、php-fpm 之间的关系？](https://www.zhihu.com/question/30672017)