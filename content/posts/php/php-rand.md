---
title: "Php Rand"
date: 2021-08-05T16:17:56+08:00
draft: false
---

## PHP的三种内置函数生成随机数的方法

1. rand函数

rand() 函数可以不加任何参数，就可以生成随机整数。如果要设置随机数范围，可以在函数中设置 min 和 max 的值。

```php
echo rand() 
echo rand(1, 10)
```

2. mt_rand函数

mt_rand() 函数是使用 Mersenne Twister算法 返回随机整数，与rand() 函数的主要区别是： mt_rand() 产生随机数值的平均速度比rand() 快4倍。（就是快了）

```php
echo mt_rand()
echo mt_rand(1, 10)
```

3. uniqid函数

uniqid() 函数基于以微秒计的当前时间，生成一个唯一的 ID。默认生成ID的长度为13位或者23位，由英文字母和数字组成。uniqid() 函数有两个参数，格式如下：

```php
uniqid(prefix, more_entropy) // 参数一： prefix: 生成ID的前缀 参数二： more_entropy: 是否添加额外的熵

echo uniqid();                    // 生成13位字符串，如：55f540e273e93
echo uniqid('one.');              // 生成前缀为one.加13位随机字符的字符串，如：one.55f540e273e93
echo uniqid('two.', true);        // 生成前缀为two.加23位随机字符的字符串（加了熵），如：two.55f540e273e
```