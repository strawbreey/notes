---
title: "Nignx X Sendfile"
date: 2020-10-21T19:28:35+08:00
draft: false
---

很多时候用户需要从网站下载文件，如果文件是可以通过一个固定链接公开获取的，那么我们只需将文件存放到 webroot下的目录里就好。但大多数情况下，我们需要做权限控制，例如下载 PDF 账单，又例如下载网盘里的档案。这时，我们通常借助于脚本代码来实现，而这无疑会增加服务器的负担。

```php

  // 用户身份认证，若验证失败跳转
  authenticate();

  // 获取需要下载的文件，若文件不存在跳转
  $file = determine_file();

  // 读取文件内容
  $content=file_get_contents($file);

  // 发送合适的 HTTP 头
  header("Content-type: application/octet-stream");
  header('Content-Disposition: attachment; filename="' . basename($file) . '"');
  header("Content-Length: ". filesize($file));
  echo $content; // 或者 readfile($file);

```



这样做意味着我们的程序需要将文件内容从磁盘经过一个固定的 buffer 去循环读取到内存，再发送给前端 web 服务器，最后才到达用户。当需要下载的文件很大的时候，这种方式将消耗大量内存，甚至引发 php 进程超时或崩溃。Cache 也很头疼，更不用说中断重连的情况了。
一个理想的解决方式应该是，由 php 程序进行权限检查等逻辑判断，一切通过后，让前台的 web 服务器直接将文件发送给用户——像 Nginx 这样的前台更善于处理静态文件。这样一来 php 脚本就不会被 I/O 阻塞了。

### X-Sendfile

X-Sendfile 是一种将文件下载请求由后端应用转交给前端 web 服务器处理的机制，它可以消除后端程序既要读文件又要处理发送的压力，从而显著提高服务器效率，特别是处理大文件下载的情形下。

X-Sendfile 通过一个特定的 HTTP header 来实现：在 X-Sendfile 头中指定一个文件的地址来通告前端 web 服务器。当 web 服务器检测到后端发送的这个 header 后，它将忽略后端的其他输出，而使用自身的组件（包括 缓存头 和 断点重连 等优化）机制将文件发送给用户。

不过，在使用 X-Sendfile 之前，我们必须明白这并不是一个标准特性，在默认情况下它是被大多数 web 服务器禁用的。而不同的 web 服务器的实现也不一样，包括规定了不同的 X-Sendfile 头格式。如果配置失当，用户可能下载到 0 字节的文件。

使用 X-Sendfile 将允许下载非 web 目录中的文件（例如/root/），即使文件在 .htaccess 保护下禁止访问，也会被下载。

不同的 web 服务器实现了不同的 HTTP 头

#### 使用 X-SendFile 的缺点: 

失去了对文件传输机制的控制。例如如果你希望在完成文件下载后执行某些操作，比如只允许用户下载文件一次，这个 X-Sendfile 是没法做到的，因为后台的 php 脚本并不知道下载是否成功。


### 使用

Apache 请参考mod_xsendfile模块。下面我介绍 Nginx 的用法。

Nginx 默认支持该特性，不需要加载额外的模块。只是实现有些不同，需要发送的 HTTP 头为 X-Accel-Redirect。另外，需要在配置文件中做以下设定

```shell
# 表示这个路径只能在 Nginx 内部访问，不能用浏览器直接访问防止未授权的下载。
location /protected/ {
  # 在 location 中加入 "internal", 声明仅限内部调用
  internal; 
  root  /some/path;
}
```

PHP 发送 X-Accel-Redirect 给 Nginx：

```php
  $filePath = '/protected/iso.img';
  header('Content-type: application/octet-stream');
  header('Content-Disposition: attachment; filename="' . basename($file) . '"');
  // 让Xsendfile发送文件
  header('X-Accel-Redirect: '.$filePath);
```

client -> php -> client - > nginx

客户端访问php, php让client重定向访问nginx


这样用户就会下载到 /some/path/protected/iso.img 这个路径下的文件。

如果你想发送的是 /some/path/iso.img 文件，那么 Nginx 配置应该是
```conf
location /protected/ {
 internal;
 alias  /some/path/; # 注意最後的斜杠
}
```

### 使用代理下载文件

```conf
location /download {
  internal;
  proxy_pass https://127.0.0.8/file/; #使用代理
}
```

```php
//发送远程文件 最终会转发到https://127.0.0.8/file/file?path=xxx.mp4
header("X-Accel-Redirect:/protected_files/file?path=xxx.mp4");
```

但是经过测试这种转发远程文件的速度，远不及，HTTP 302  (重定向) 的速度快。

为了权限认证下载只能是妥协一下了。

HTTP 302 Found 重定向状态码表明请求的资源被暂时的移动到了由Location 头部指定的 URL 上。浏览器会重定向到这个URL， 但是搜索引擎不会对该资源的链接进行更新 (In SEO-speak, it is said that the link-juice is not sent to the new URL)。

即使规范要求浏览器在重定向时保证请求方法和请求主体不变，但并不是所有的用户代理都会遵循这一点，你依然可以看到有缺陷的软件的存在。所以推荐仅在响应 GET 或 HEAD 方法时采用 302 状态码，而在其他时候使用 307 Temporary Redirect 来替代，因为在这些场景下方法变换是明确禁止的。

### 参考链接 
- [nginx-xsendfile](https://www.nginx.com/resources/wiki/start/topics/examples/xsendfile/)
- [在Nginx中使用X-Sendfile头提升PHP文件下载的性能（针对大文件下载）](https://www.jb51.net/article/51854.htm)
- [php+nginx做下载服务器配置—X-Accel-Redirect](http://www.phpweblog.net/phpbaby/archive/2012/07/10/7709.html)

