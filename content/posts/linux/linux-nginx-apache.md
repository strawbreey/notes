---
title: "Linux Nginx Apache"
date: 2020-12-28T15:41:53+08:00
draft: false
---

## Linux服务器下Nginx与Apache共存

将nginx作为代理服务器和web服务器使用，nginx监听80端口，Apache监听除80以外的端口，我这暂时使用8080端口。


在Linux 一经搭建好环境 先后安装了Nginx 和Apache 由于 默认端口都是：80

一般客户请求的服务器端口默认为80 所以Nginx作为静态页端口设置：80; 

Apache设置端口为:8080(在httpd.conf 文件中修改Listen：8080)
Apache下的网站：

在nginx.conf中 添加

```conf
server {
  listen       80;
  server_name  www.one.ityangs.cn one.ityangs.cn;

  location / {            
    proxy_pass              http://127.0.0.1:8080;              
    proxy_redirect          off;        
    proxy_set_header Host $host;       
    proxy_set_header X-Real-IP $remote_addr;       
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;           
  } 
}
```

在httpd.conf中 添加
```conf
<virtualhost *:8080>
    ServerName  www.one.ityangs.cn
    ServerAlias  www.one.ityangs.cn one.ityangs.cn 
    DocumentRoot  /www/one
    DirectoryIndex index.php index.html

  <Directory /www/one>
    Options +Includes +FollowSymLinks -Indexes
    AllowOverride All
    Order Deny,Allow
    Allow from All
  </Directory>
</virtualhost>
```
### Nginx下的网站：
在nginx.conf中 添加
```
　server {
    listen       80;
    server_name    two.ityangs.cn www.two.ityangs.cn;
    root   /www/two;
    location /{

      index  index.html index.htm index.php;
        if (!-e $request_filename) {
        rewrite  ^(.*)$  /index.php?s=$1  last;
        break;
      }
      error_page 404  /var/www/html/404.html;
    }

    location ~ \.php(.*)$ {
      fastcgi_pass   127.0.0.1:9000;
      fastcgi_index  index.php;
      fastcgi_split_path_info  ^((?U).+\.php)(/?.+)$;
      fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
      fastcgi_param  PATH_INFO  $fastcgi_path_info;
      fastcgi_param  PATH_TRANSLATED  $document_root$fastcgi_path_info;
      include        fastcgi_params;  
    }
}
```

### apache httpd 常用命令

httpd -k start

httpd -k restart

httpd -k stop