---
title: "Php Download File"
date: 2020-12-16T15:24:48+08:00
draft: false
---

```php 
// 1. 方案一
// file — 把整个文件读入一个数组中
$fp=file($url);

foreach ($fp as $fileLine) {
  echo $fileLine;
}

// 2. 方案二
// file_get_contents — 将整个文件读入一个字符串
echo file_get_content($url)

// 3. 方案三
// 注意fopen读取文件需要配合使用fgets和fclose函数。
// 可以使用 'b' 来强制使用二进制模式，这样就不会转换数据。要使用这些标记，要么用 'b' 或者用 't' 作为 mode 参数的最后一个字符。
$fp = fopen($file_name, 'rb');
while (!feof($fp)) {
  // fgets — 从文件指针中读取一行
  // 从 handle 指向的文件中读取一行并返回长度最多为 length - 1 字节的字符串。碰到换行符（包括在返回值中）、EOF 或者已经读取了 length - 1 字节后停止（看先碰到那一种情况）。
  // 如果没有指定 length，则默认为 1K，或者说 1024 字节。
  // 16384 字节 16k
  $buffer = fgets($fp, 16384);
  echo $buffer;
}
fclose($fp);

// 4. 方案四
$file = fopen($file_path, 'rb');

//打开缓冲区
ob_start();

// 如果设置为 true，则忽略与用户的断开，如果设置为 false，会导致脚本停止运行。
ignore_user_abort(TRUE);

// Send data in 16kb blocks
$block = 1024 * 16;

while (!feof($file))//AND ($pos = ftell($file)) <= $end
{
    if (connection_aborted()) {
        break;
    }
    echo fread($file, $block);
    ob_flush();
    flush();
}
ob_flush();
flush();
// 关闭缓冲区
// 清空（擦除）缓冲区并关闭输出缓冲
fclose($file);
ob_end_clean();
```

通过上面两个例子的对比，可以看出使用file_get_contents()打开URL，也许是更多人的选择，因为其比fopen()更简单便捷。

不过，如果是读取比较大的资源，则是用fopen()比较合适。

### fread 和 fget 的区别
### fread 和 fget 的区别
### fread 和 fget 的区别
### fread 和 fget 的区别


### 获取url的后缀

```php
//获取要下载的图片的MIME信息

function getMime($url) {
  if(preg_match("/\.(jpg|jpeg)$/",$url))
    return "image/jpeg";
  else if(preg_match("/\.(gif)$/",$url))
    return "image/gif";
  else if(preg_match("/\.(png)$/",$url))
    return "image/png";
  else if(preg_match("/\.(bmp)$/",$url))
    return "image/bmp";
  else
    return "err";
}

//获取要下载的图片后缀名

function getExt($url) {
  if(preg_match("/\.(jpg|jpeg)$/",$url))
    return "jpg";
  else if(preg_match("/\.(gif)$/",$url))
    return "gif";
  else if(preg_match("/\.(png)$/",$url))
    return "png";
  else if(preg_match("/\.(bmp)$/",$url))
    return "bmp";
  else
    return "err";
}
```

### 参考资料

- [file](https://www.php.net/manual/zh/function.file.php)
- [file_get_contents](https://www.php.net/manual/zh/function.file-get-contents.php)