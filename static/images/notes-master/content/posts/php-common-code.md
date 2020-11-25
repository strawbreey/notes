---
title: "Php Common Code"
date: 2020-09-23T19:53:01+08:00
draft: false
---

1. 分割字符串

```php
$str = aaa,bbb,ccc
explode(',', $str) // [aaa, bbb, ccc]
```

2. 数组push

```php
array_push($arr, $new);
```

3. 获取对应列
```php
array_column($arr, 'key')
array_column($arr, 1)
```

4. array_diff 取差

```php
$arr = [1,2,3]
$newArr = [2,3,4]
array_diff($arr, $newArr) // [1]
```
