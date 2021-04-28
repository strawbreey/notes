---
title: "Leetcode 0204 Count Primes"
date: 2020-12-07T11:20:47+08:00
draft: false
---

统计所有小于非负整数 n 的质数的数量。

```
输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

输入：n = 0
输出：0

输入：n = 1
输出：0
```

提示：
`0 <= n <= 5 * 106`

Solution:
```js
/**
 * @param {number} n
 * @return {number}
 */


var countPrimes = function(n) {
    let i = 2
    let h = Array(n).fill(1), r = 0

    h[0] = 0
    h[1] = 0

    while (i * i <= n) {
        if (h[i]) {
            j = i
            while (i * j < n) {
                h[i * j] = 0
                j++
            }
        }
        i++
    }
    console.log(h)
    return h.filter(item => item == 1).length
};
```