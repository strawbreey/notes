---
title: "Php Http"
date: 2021-07-22T17:14:50+08:00
draft: false
---

1. 用 file_get_contents 以get方式获取内容：

```php
$url = 'https://wenda.shukaiming.com/';
echo file_get_contents($url);
```

2. 方法2: 用fopen打开url, 以get方式获取内容：
```php
//r标识read，即标识只读
$fp = fopen($url, 'r');
stream_get_meta_data($fp);
while (!feof($fp)) {
    $body.= fgets($fp, 1024);
}
echo $body;
fclose($fp);
```

3：用 file_get_contents函数,以post方式获取url

```php
$data = array(‘foo' => ‘bar');
$data = http_build_query($data);
$opts = array('http' => array('method' => 'POST', 'header' => 'Content-type: application/x-www-form-urlencodedrn' . 'Content-Length: ' . strlen($data) . '\r\n', 'content' => $data));
$context = stream_context_create($opts);
$html = file_get_contents('https://wenda.shukaming.com', false, $context);
echo $html;
```

4. 使用curl库，使用curl库之前，可能需要查看一下php.ini是否已经打开了curl扩展

```
<?php $ch = curl_init(); $timeout = 5; curl_setopt($ch, CURLOPT_URL, 'http://wenda.shukaiming.com'); curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1); curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, $timeout); $file_contents = curl_exec($ch); curl_close($ch); echo $file_contents; ?>
```