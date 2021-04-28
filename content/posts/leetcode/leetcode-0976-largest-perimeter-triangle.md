---
title: "Leetcode 0976 Largest Perimeter Triangle"
date: 2020-11-29T00:51:35+08:00
draft: false
---

给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。

如果不能形成任何面积不为零的三角形，返回 0。

```
输入：[2,1,2]
输出：5

输入：[1,2,1]
输出：0

输入：[3,2,3,4]
输出：10

输入：[3,6,2,3]
输出：8

```

Solution
```js
/**
 * @param {number[]} A
 * @return {number}
 */
var largestPerimeter = function(A) {
    A = A.sort((a,b) => Number(b) - Number(a))
    let Max = 0
    for (let i = 2; i < A.length; i ++) {
        if (A[i - 2] < A[i - 1] + A[i] ) {
            Max = A[i] + A[i - 1] + A[i - 2]
            break
        }
    }
    return Max    
};
```

### 参考资料

- [三角形的最大周长](https://leetcode-cn.com/problems/largest-perimeter-triangle/)