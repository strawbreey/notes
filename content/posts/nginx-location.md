---
title: "Nginx Location"
date: 2020-10-10T10:32:43+08:00
draft: false
---

nginx 指定文件路径有两种方式


listen参数决定Nginx服务如何监听端口。在listen后可以只加IP地址、端口或主机名，非常灵活，例如：[插图]

```
listen 127.0.0.1:8000
listen 127.0.0.1; # 不加端口时, 默认监听80端口
listen 8000 
listen *:8000
listen localhost: 8000
```


1. 以root方式设置资源路径

```shell
# 例如，定义资源文件相对于HTTP请求的根目录。
location /download/ {
  root /opt/web/html;
}
```

在上面的配置中，如果有一个请求的URI是/download/index/test.html，那么Web服务器将会返回服务器上/opt/web/html/download/index/test.html文件的内容。


2. 以alias方式设置资源路径

alias也是用来设置文件资源路径的，它与root的不同点主要在于如何解读紧跟location后面的uri参数，这将会致使alias与root以不同的方式将用户请求映射到真正的磁盘文件上。例如，如果有一个请求的URI是/conf/nginx.conf，而用户实际想访问的文件在/usr/local/nginx/conf/nginx.conf，那么想要使用alias来进行设置的话，可以采用如下方式：

```shell

location /conf {
  alias /usr/local/nginx/conf;
}

# 等同于

location /conf {
  root /usr/local/nginx
}
```

使用alias时，在URI向实际文件路径的映射过程中，已经把location后配置的/conf这部分字符串丢弃掉，因此，/conf/nginx.conf请求将根据alias path映射为path/nginx.conf。root则不然，它会根据完整的URI请求来映射，因此，/conf/nginx.conf请求会根据root path映射为path/conf/nginx.conf。这也是root可以放置到http、server、location或if块中，而alias只能放置到location块中的原因。

alias后面还可以添加正则表达式，例如：

```shell
location ~^/test/(\w+)\.(\w+)$ {
  alias /usr/local/nginx/$2/$1.$2;
}
```

访问首页

有时，访问站点时的URI是/，这时一般是返回网站的首页，而这与root和alias都不同。这里用ngx_http_index_module模块提供的index配置实现。index后可以跟多个文件参数，Nginx将会按照顺序来访问这些文件，例如：

```shell
location / {
  root path;
  index /index.html /html/index.php /index.php

}
```

根据HTTP返回码重定向页面

当对于某个请求返回错误码时，如果匹配上了error_page中设置的code，则重定向到新的URI中。例如：

```shell
error_page 404 /404.html;
error_page 502 503 504 /50x.html
error_page 403 http://exmaplte.com.forbidden.html
error_page 404 =@fetch
```

注意，虽然重定向了URI，但返回的HTTP错误码还是与原来的相同。用户可以通过“=”来更改返回的错误码，例如：

try_files

try_files后要跟若干路径，如path1 path2...，而且最后必须要有uri参数，意义如下：尝试按照顺序访问每一个path，如果可以有效地读取，就直接向用户返回这个path对应的文件结束请求，否则继续向下访问。如果所有的path都找不到有效的文件，就重定向到最后的参数uri上。因此，最后这个参数uri必须存在，而且它应该是可以有效重定向的。例如：

```shell

```