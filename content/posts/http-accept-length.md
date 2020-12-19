---
title: "Http Accept Length"
date: 2020-12-16T15:21:03+08:00
draft: false
---

```php
$interval=120; //2分钟
if(isset($_SERVER['HTTP_IF_MODIFIED_SINCE'])){
    // HTTP_IF_MODIFIED_SINCE即下面的: Last-Modified,文档缓存时间.
    // 缓存时间+时长.
    
    $c_time = strtotime($_SERVER['HTTP_IF_MODIFIED_SINCE'])+$interval;

    // 当大于当前时间时, 表示还在缓存中... 释放304
    if($c_time > time()){
        header('HTTP/1.1 304 Not Modified');
        exit();
    }
}

// Cache-Control头，告诉让浏览器自己缓存个120秒
header('Cache-Control:max-age=120');

// 在当前时间的120秒后缓存失效，浏览器会在120秒后把缓存失效，再次请求时会到服务器端请求而不是本地
header("Expires: " . gmdate("D, d M Y H:i:s",time()+ 120)." GMT");

// Last-Modified则是说明这个文档最后修改时间
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
```


1. 当浏览器再次访问时（刷新页面），浏览器因为在第一次时接收到了Last-Modified，它便会根据这个时间去检查服务器的文档是否更新过，也就是会把这个时间用HTTP_IF_MODIFIED_SINCE这个东西带过去，于是也就有了第一个花括号内容的判断。而花括号内再次根据这个时间加上你设置的时间（120秒），看看是不是超出了当前的时间time()，如果超出了，那就是还在缓存期内，帮直接返回304，然后直接退出。这时候浏览器便会自动用缓存的内容。这样子做，服务器返回的内容就很短，只是一个304响应头而已，而内容不需要返回，这可大大节省网络流量，显然浏览器的响应速度也会感觉明显快了。

2. 按F5刷新或者按浏览器的刷新按钮时，你将发现2分钟内你的PHP请求将是304状态返回。
强制刷新（CTRL+F5）时，浏览器不会带HTTP_IF_MODIFIED_SINCE这个东西去判断，所以同正常浏览一样，返回200状态。这个逻辑是正常的，因为强制刷新本就相当于完全请求一遍最新内容。


![http 403缓存](/images/977033-20170410175310876-1823314432.png)

php中在做文件下载的时候，其中要加上这么一些header信息：

```php
header("Content-type: application/octet-stream");
header("Accept-Ranges: bytes");
header("Accept-Length:".$fileSize);
header("Content-Length:".$fileSize);
header("Content-Disposition: attachment; filename=".$fileName);
```
- Content-Length: 获取下载文件的大小, 但下载大小和真实大小不一定时, 会导致下载不完整

### http header 常见配置

1. max-age：

header('Cache-Control:max-age=120');

这个的输出是表明让浏览器缓存120秒。这个指令很有用，但是局限性在于，刷新这个页面时还是会重新请求，所以你会感觉这个头好像没什么作用，但是对于一些输出的JS、CSS内容时是有用的，也就是说你刷新的页面不是它本身，而只是引用了设置有此头的PHP文件的话，它就会在你指定的时间内从本地缓存中读取内容而不会请求服务器。同时通过超链接过来的已经请求过的页面时，你也会看到浏览器从本地cache里读取而不用通过服务器。

2. Expires

Expires是HTTP/1.0中的,它比max-age要麻烦点.Expires指定的时间分下面二种,这个主要考虑到apache中设置是A还是M.

- 相对文件的最后访问时间(Atime)
当Apache使用A时间来做Expires时.这样设置时.他就和max-age的值相等,因为max-age是相对文件的请求时间(Atime).

例如:ExpiresByType text/html A600

由上面我们得知,Apache设置Atime时,过期为600秒时.

Expires=18:00+600=18:10
max-age=18:00+600=18:10

得出:Expires=max-age

- 绝对修改时间(MTime)
这又分二种情况,我们来拿A.htm来讲
假设文件的建立时间为18:00.

当用户Request请求为18:00时,过期为600秒

Expires=18:00+600=18:10
max-age=18:00+600=18:10

得出:Expires等于max-age

3. Cache-Control

### 参考资料

- [浅谈http中的Cache-Control](https://blog.csdn.net/u012375924/article/details/82806617)