---
title: "Php Fopen"
date: 2021-07-22T17:05:34+08:00
draft: false
---

只要在 php.ini 文件中激活了 allow_url_fopen 选项，就可以在大多数需要用文件名作为参数的函数中使用 HTTP 和 FTP 的 URL 来代替文件名。同时，也可以在 include、include_once、require 及 require_once 语句中使用 URL。

### 获取远程文件的标题

```php
$file = fopen ("http://www.example.com/", "r");
if (!$file) {
    echo "<p>Unable to open remote file.\n";
    exit;
}
while (!feof ($file)) {
    $line = fgets ($file, 1024);
    /* This only works if the title and its tags are on one line */
    if (eregi ("<title>(.*)</title>", $line, $out)) {
        $title = $out[1];
        break;
    }
}
fclose($file)
```

###  将数据保存到远程服务器

```php
$file = fopen ("ftp://ftp.example.com/incoming/outputfile", "w");
if (!$file) {
    echo "<p>Unable to open remote file for writing.\n";
    exit;
}
/* Write the data here. */
fwrite ($file, $_SERVER['HTTP_USER_AGENT'] . "\n");
fclose ($file);
```