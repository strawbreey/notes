---
title: "Php Header"
date: 2020-10-21T19:58:34+08:00
draft: false
---

fread

fread — 读取文件（可安全用于二进制文件）

```php
fread ( resource $handle , int $length ) : string

// handle 文件系统指针，是典型地由 fopen() 创建的 resource(资源)。
// length 最多读取 length 个字节。
```

fread() 从文件指针 handle 读取最多 length 个字节。 该函数在遇上以下几种情况时停止读取文件：

- 读取了 length 个字节
- 到达了文件末尾（EOF）
- a packet becomes available or the socket timeout occurs (for network streams)
- if the stream is read buffered and it does not represent a plain file, at most one read of up to a number of bytes equal to the chunk size (usually 8192) is made; depending on the previously buffered data, the size of the returned data may be larger than the chunk size.


```php
$filename = "/usr/local/something.txt";
$handle = fopen($filename, "r");
$contents = fread($handle, filesize($filename));
fclose($handle);
```

### 返回值

返回所读取的字符串， 或者在失败时返回 FALSE。

### 注意

如果只是想将一个文件的内容读入到一个字符串中，用 file_get_contents()，它的性能比上面的代码好得多。



### file_get_contents

file_get_contents — 将整个文件读入一个字符串

```php

file_get_contents ( string $filename [, bool $use_include_path = false [, resource $context [, int $offset = -1 [, int $maxlen ]]]] ) : string

// 参数 filename 要被写入数据的文件名。
// 参数 data 要写入的数据。类型可以是 string，array 或者是 stream 资源（如上面所说的那样）。 如果 data 指定为 stream 资源，这里 stream 中所保存的缓存数据将被写入到指定文件中，这种用法就相似于使用 stream_copy_to_stream() 函数。 参数 data 可以是数组（但不能为多维数组），这就相当于 file_put_contents($filename, join('', $array))。
// 参数 flags flags 的值可以是 以下 flag 使用 OR (|) 运算符进行的组合。
// 参数 context  一个 context 资源。
// 返回值 该函数将返回写入到文件内数据的字节数，失败时返回FALSE

```

和依次调用 fopen()，fwrite() 以及 fclose() 功能一样。

范例

```php

$file = 'people.txt';
// Open the file to get existing content
$current = file_get_contents($file);
// Append a new person to the file
$current .= "John Smith\n";
// Write the contents back to the file
file_put_contents($file, $current);
```


```php



$fp = fopen('somefile.txt', 'r');
if (!$fp) {
    echo 'Could not open file somefile.txt';
}
while (false !== ($char = fgetc($fp))) {
    echo "$char\n";
}

```


