---
title: "Nginx Try Files"
date: 2020-09-15T11:29:43+08:00
draft: true
---

### 基本配置
```shell
  location / {
    root         /data/website/baidu;
    try_files $uri $uri/ /index.php? $query_string;
  }
```

Nginx的配置语法灵活，可控制度非常高。在0.7以后的版本中加入了一个try_files指令，配合命名location，可以部分替代原本常用的rewrite配置方式，提高解析效率。

当用户请求 http://localhost/example 时，这里的 $uri 就是 /example。 
try_files 会尝试寻找这个文件, 如果存在名为 /$root/example的文件，就直接把这个文件的内容发送给用户。 
显然，目录中没有叫 example 的文件。然后就看 $uri/，增加了一个 /，也就是看有没有名为 /$root/example/ 的目录。 
又找不到，就会 fall back 到 try_files 的最后一个选项 /index.php，发起一个内部 “子请求”，也就是相当于 nginx 发起一个 HTTP 请求到 http://localhost/index.php。

### try_files指令说明

```
try_files指令
语法：try_files file ... uri 或 try_files file ... = code
默认值：无
作用域：server location
```

其作用是按顺序检查文件是否存在，返回第一个找到的文件或文件夹(结尾加斜线表示为文件夹)，如果所有的文件或文件夹都找不到，会进行一个内部重定向到最后一个参数。

需要注意的是，只有最后一个参数可以引起一个内部重定向，之前的参数只设置内部URI的指向。最后一个参数是回退URI且必须存在，否则会出现内部500错误。命名的location也可以使用在最后一个参数中。与rewrite指令不同，如果回退URI不是命名的location那么$args不会自动保留，如果你想保留$args，则必须明确声明。