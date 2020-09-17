---
title: "Php Remote Procedure"
date: 2020-09-16T11:11:43+08:00
draft: true
---

## php 请求远程

php请求远程url有两个方法`fopen`/`file_get_contents` 和 `curl`

1，fopen/file_get_contents与curl的差异

（1）fopen /file_get_contents 每次请求都会重新做DNS查询，并不对DNS信息进行缓存。但是CURL会自动对DNS信息进行缓存。对同一域名下的网页或者图片的请求只需要一次DNS查询。这大大减少了DNS查询的次数。所以CURL的性能比fopen /file_get_contents 好很多。
（2）fopen /file_get_contents在请求HTTP时，使用的是http_fopen_wrapper，不会keeplive。而curl却可以。这样在多次请求多个链接时，curl效率会好一些。
（3）curl可以模拟多种请求，例如：POST数据，表单提交等，用户可以按照自己的需求来定制请求。而fopen / file_get_contents只能使用get方式获取数据。


### php中流行的rpc框架

php中流行的rpc框架 phprpc，yar, thrift, gRPC, swoole, hprose 等

#### phprpc 简介
是一个轻型的、安全的、跨网际的、跨语言的、跨平台的、跨环境的、跨域的、支持复杂对象传输的、支持引用参数传递的、支持内容输出重定向的、支持分级错误处理的、支持会话的、面向服务的高性能远程过程调用协议。

```php
// 服务端
include('../phprpc_server.php');

class Hello{
    static function helloWorld(){
        return "hello world";
    }
}

$server = new PHPRPC_Server();
$server->add('helloWorld','hello');
$server->start()

// 客户端
include("../phprpc_client.php");
$client = new PHPRPC_Client('http://www.test.com/phprpc/test/server.php');
echo $client->helloWorld();
```

[](https://www.php.net/manual/en/ref.xmlrpc.php)