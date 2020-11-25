---
title: "Php Quotation Marks"
date: 2020-09-23T16:18:58+08:00
draft: false
---

在PHP中

单引号代表纯字符串，不论里面有什么东西，当字符串处理

```php
$a = 'abc';
ehco 'i am $a'; // i am $a
```

双引号代表可以处理的字符串，如果字符串中有变量，那么，会优先解析变量

```php
$a = 'abc';
echo "i am $a"; // i am abc
```
理论上, 单引号的速度 > 双引号的速度，因为不用解析变量