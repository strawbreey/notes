---
title: "Php Strpos"
date: 2021-05-06T14:45:53+08:00
draft: false
---

(PHP 4, PHP 5, PHP 7, PHP 8)
strpos — 查找字符串首次出现的位置

strpos ( string $haystack , mixed $needle , int $offset = 0 ) : int

参数

- haystack:  在该字符串中进行查找。

- needle:  Prior to PHP 8.0.0, if needle is not a string, it is converted to an integer and applied as the ordinal value of a character. This behavior is deprecated as of PHP 7.3.0, and relying on it is highly discouraged. Depending on the intended behavior, the needle should either be explicitly cast to string, or an explicit call to chr() should be performed.

- offset 如果提供了此参数，搜索会从字符串该字符数的起始位置开始统计。 如果是负数，搜索会从字符串结尾指定字符数开始。


### 返回值

返回 needle 存在于 haystack 字符串起始的位置(独立于 offset)。同时注意字符串位置是从0开始，而不是从1开始的。

如果没找到 needle，将返回 false。

### 参考

- stripos() - 查找字符串首次出现的位置（不区分大小写）
- strrpos() - 计算指定字符串在目标字符串中最后一次出现的位置
- strripos() - 计算指定字符串在目标字符串中最后一次出现的位置（不区分大小写）
- strstr() - 查找字符串的首次出现
- strpbrk() - 在字符串中查找一组字符的任何一个字符
- substr() - 返回字符串的子串
- preg_match() - 执行匹配正则表达式

