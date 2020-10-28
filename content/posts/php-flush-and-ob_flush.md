---
title: "Php Flush and Ob_flush"
date: 2020-10-21T17:45:52+08:00
draft: true
---

### flush

flush — 刷新输出缓冲

```php
flush ( void ) : void
```

刷新PHP程序的缓冲，而不论PHP执行在何种情况下（CGI ，web服务器等等）。该函数将当前为止程序的所有输出发送到用户的浏览器。

flush() 函数不会对服务器或客户端浏览器的缓存模式产生影响。因此，必须同时使用 ob_flush() 和flush() 函数来刷新输出缓冲。

个别web服务器程序，特别是Win32下的web服务器程序，在发送结果到浏览器之前，仍然会缓存脚本的输出，直到程序结束为止。

有些Apache的模块，比如mod_gzip，可能自己进行输出缓存，这将导致flush()函数产生的结果不会立即被发送到客户端浏览器。

甚至浏览器也会在显示之前，缓存接收到的内容。例如 Netscape 浏览器会在接受到换行或 html 标记的开头之前缓存内容，并且在接受到 </table> 标记之前，不会显示出整个表格。

### ob_flush

ob_flush — 冲刷出（送出）输出缓冲区中的内容

```php
ob_flush ( void ) : void
```

这个函数将送出缓冲区的内容（如果里边有内容的话）。如果想进一步处理缓冲区中的内容，必须在ob_flush()之前调用ob_get_contents() ，因为在调用ob_flush()之后缓冲区内容将被丢弃。

此函数不会销毁输出缓冲区，而像ob_end_flush() 函数会销毁缓冲区。


### ob_clean

 ob_clean — 清空（擦掉）输出缓冲区

 ```php
  ob_clean ( void ) : void
 ```

 此函数用来丢弃输出缓冲区中的内容。

 此函数不会像 ob_end_clean() 函数那样销毁输出缓冲区。

 输出缓冲必须已被 ob_start() 以 PHP_OUTPUT_HANDLER_CLEANABLE 标记启动。否则 ob_clean() 不会有效果。


### ob_end_clean

ob_end_clean — 清空（擦除）缓冲区并关闭输出缓冲

```php
ob_end_clean ( void ) : bool
```

此函数丢弃最顶层输出缓冲区的内容并关闭这个缓冲区。如果想要进一步处理缓冲区的内容，必须在ob_end_clean()之前调用ob_get_contents()，因为当调用ob_end_clean()时缓冲区内容将被丢弃。

### ob_end_flush

ob_end_flush — 冲刷出（送出）输出缓冲区内容并关闭缓冲

```php
ob_end_flush ( void ) : bool
```

这个函数将送出最顶层缓冲区的内容（如果里边有内容的话），并关闭缓冲区。如果想进一步处理缓冲区中的内容，必须在ob_end_flush()之前调用 ob_get_contents()，因为在调用ob_end_flush()后缓冲区内容被丢弃。


### flush和ob_flush的使用上有一些特别注意的地方，造成无法刷新输出缓冲

1. flush和ob_flush的正确顺序，先ob_flush再flush

2.  使用ob_flush()前，确保前面的内容大小足够4069字符。

### php.ini中有两个关键参数会影响到php的缓存输出控制

1. output_buffering ：on/off 或者整数 。设置为on时，将在所有脚本中使用输出缓存控制，不限制缓存的大小。而设置为整数时，如output_buffering=4096，当缓存数据达到4096字节时会自动输出刷新缓存。而这个参数的不同正是导致以上代码在不同时候执行结果不同的原因。当output_buffering关闭时，脚本所有的输出（echo）都会即时发送到客户端，执行上面代码时就是每秒输出一个数字。而开启output_buffering后，输出内容就会先缓存在服务端，直到脚本结束时才一起发送给客户端。

2. implicit_flush：on/off。设定ON意味着，当脚本有输出时，自动立即发送到客户端。相当于在echo后自动加flush()。


### 总结

ob_end_flush()：输出当前服务器端缓存的输出数据，并关闭缓存。

ob_end_clean()：清空缓存内容，并关闭缓存。

ob_get_flush()：将当前服务器端缓存的输出数据以字符串形式返回，并关闭缓存

ob_get_contents()：将缓存中保存的内容以字符串形式返回，并保留缓存。

ob_get_length():返回缓存中数据的长度。

ob_get_clean()：获取缓存中的数据，请清空缓存，相当于依次执行ob_get_contents()和ob_end_clean()。

ob_implicit_flush()：相当于开启php.ini中的implicit_flush参数，立即发送脚本的输出。

ob_gzhandler()：使用gzip压缩缓存数据。用于将文本数据压缩后再发送到客户端，可以极大减少数据传输所用的时间，对于提高网站浏览速度帮助很大。通常作为ob_start()的回调函数来使用。

ob_list_handlers()：列出所有输出使用的操作方法。

### 参考文献
- [细说flush、ob_flush的区别](https://www.cnblogs.com/phpper/p/7750104.html)
- []()
