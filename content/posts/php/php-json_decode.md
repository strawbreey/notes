---
title: "Php Json_decode"
date: 2020-12-07T17:31:25+08:00
draft: false
---

 对 JSON 格式的字符串进行解码

 ```php
 json_decode ( string $json [, bool $assoc = FALSE [, int $depth = 512 [, int $options = 0 ]]] ) : mixed
 ```

 接受一个 JSON 编码的字符串并且把它转换为 PHP 变量

### 参数 

- json 待解码的 json string 格式的字符串。

- assoc 当该参数为 TRUE 时，将返回 array 而非 object 。

- depth 指定递归深度。

- options 由 JSON_BIGINT_AS_STRING, JSON_INVALID_UTF8_IGNORE, JSON_INVALID_UTF8_SUBSTITUTE, JSON_OBJECT_AS_ARRAY, JSON_THROW_ON_ERROR 组成的掩码。


### 返回值

通过恰当的 PHP 类型返回在 json 中编码的数据。值true, false 和 null 会相应地返回 TRUE, FALSE 和 NULL。 如果 json 无法被解码， 或者编码数据深度超过了递归限制的话，将会返回NULL 。