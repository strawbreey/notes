---
title: "Php Implode"
date: 2021-08-06T16:12:45+08:00
draft: false
---

### implode

(PHP 4, PHP 5, PHP 7, PHP 8)

implode — 将一个一维数组的值转化为字符串

#### 说明 

implode(string $glue, array $pieces): string
implode(array $pieces): string

用 glue 将一维数组的值连接为一个字符串。

> 注意:

因为历史原因，implode() 可以接收两种参数顺序，但是为了和 explode() 内的顺序保持一致，不按文档的方式已被废弃。

#### 参数
- glue 默认为空的字符串。
- pieces 你想要转换的数组。

#### 返回值

返回一个字符串，其内容为由 glue 分割开的数组的值。