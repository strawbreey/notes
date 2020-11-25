---
title: "Leetcode 1014 Best Sightseeing Pair"
date: 2020-11-23T22:54:56+08:00
draft: false
---

给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。

一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。

返回一对观光景点能取得的最高分。


```
输入：[8,1,5,2,6]
输出：11
解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
```

提示: 
```
2 <= A.length <= 50000
1 <= A[i] <= 1000
```


```js
/**
 * @param {number[]} A
 * @return {number}
 */
var maxScoreSightseeingPair = function(A) {
    let count = A[0] + A[1] - 1 
    for (let i = 0; i<A.length; i++) {
        for (let j = i + 1; j < A.length; j++) {
            if (A[i] + A[j] + i -j > count) {
                count = A[i] + A[j] + i -j
            }
        }
    }
    return count
};
```