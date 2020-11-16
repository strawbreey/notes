---
title: "Curl"
date: 2020-10-19T19:55:15+08:00
draft: false
---

cURL是一个利用URL语法在命令行下工作的文件传输工具，1997年首次发行。它支持文件上传和下载，所以是综合传输工具，但按传统，习惯称cURL为下载工具。cURL还包含了用于程序开发的libcurl。

cURL是一个开源项目，主要的产品是curl（命令行工具）和libcurl（C语言的API库），两者功能均是：基于网络协议，对指定URL进行网络传输。[2][3]

cURL涉及是任何网络协议传输，不涉及对具体数据的具体处理。（如：html的渲染等）

```shell
# 简单模式
curl http://example.com

# 详细（verbose）模式：
curl --verbose http://example.com
curl -v http://example.com

# 下载（output）
curl --output output.html http://example.com/
curl -o output.html http://example.com/

# 重定向：（curl默认不会重定向）
curl --location output.html http://example.com/
curl -L output.html http://example.com/

curl http://httpbin.org/user-agent

curl https://httpbin.org/get?show_env=1

```


### curl 和 wget 区别

wget是个专职的下载利器，简单，专一，极致；而curl可以下载，但是长项不在于下载，而在于模拟提交web数据，POST/GET请求，调试网页，等等。

在下载上，也各有所长，wget可以递归，支持断点；而curl支持URL中加入变量，因此可以批量下载。

### libcurl

libcurl主要功能就是用不同的协议连接和沟通不同的服务器~也就是相当封装了的sockPHP 支持libcurl（允许你用不同的协议连接和沟通不同的服务器）。 libcurl当前支持http, https, ftp, gopher, telnet, dict, file, 和ldap 协议。libcurl同样支持HTTPS证书授权，HTTP POST, HTTP PUT, FTP 上传（当然你也可以使用PHP的ftp扩展）, HTTP基本表单上传，代理，cookies,和用户认证。

### php cURL

PHP 支持 Daniel Stenberg 创建的 libcurl 库，能够连接通讯各种服务器、使用各种协议。libcurl 目前支持的协议有 http、https、ftp、gopher、telnet、dict、file、ldap。 libcurl 同时支持 HTTPS 证书、HTTP POST、HTTP PUT、 FTP 上传(也能通过 PHP 的 FTP 扩展完成)、HTTP 基于表单的上传、代理、cookies、用户名+密码的认证。

### 参考资料
- [https://curl.trillworks.com/](https://curl.trillworks.com/)
- [Client URL Library ¶](https://www.php.net/manual/en/book.curl.php)
- [curl/curl](https://github.com/curl/curl)