---
title: "Leetcode 0007 Reverse Integer"
date: 2020-11-06T19:54:54+08:00
draft: false
---

Given a 32-bit signed integer, reverse digits of an integer.

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

Example:

```
Input: x = 123
Output: 321

Input: x = -123
Output: -321

Input: x = 120
Output: 21

Input: x = 0
Output: 0
```

```js
/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let num = 0
    let flag = x > 0 ? 1 : -1
    x = x > 0 ? x : -x
    while(x > 0) {
        num = num * 10 + x % 10
        x = Math.floor(x / 10)
    }
    
    if (num > 2147483648) {
        num = 0
    }
    return num * flag
};
```

```go

```

```php
```