---
title: "Php Error Reporting"
date: 2021-09-26T18:15:05+08:00
draft: false
---

实例
规定不同的错误级别报告：

```php
<?php
 // 关闭错误报告
 error_reporting(0);

 // 报告 runtime 错误
 error_reporting(E_ERROR | E_WARNING | E_PARSE);

 // 报告所有错误
 error_reporting(E_ALL);

 // 等同 error_reporting(E_ALL);
 ini_set("error_reporting", E_ALL);

 // 报告 E_NOTICE 之外的所有错误
 error_reporting(E_ALL & ~E_NOTICE);
?> 

```


### 定义和用法

error_reporting() 函数跪地你给应该报告何种 PHP 错误。

error_reporting() 函数能够在运行时设置 error_reporting 指令。

PHP 有诸多错误级别，使用该函数可以设置在脚本运行时的级别。如果没有设置可选参数 level，error_reporting() 仅会返回当前的错误报告级别。