---
title: "Leetcode 0118 Pascals Triangle"
date: 2020-12-07T11:27:26+08:00
draft: false
---

给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

[PascalTriangleAnimated2](/images/PascalTriangleAnimated2.gif)

在杨辉三角中，每个数是它左上方和右上方的数的和。

```
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

Solutions:

```js
/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
    let arr = []

    if (numRows == 0) {
        return []
    }

    if (numRows == 1) {
        return [[1]]
    }

    for (let i = 0; i < numRows;i++) {
        arr[i] = arr[i] ? arr[i] : []
        arr[i][0] = 1
        arr[i][i] = 1
        for (let j = 1; j < i; j++ ) {
            arr[i][j] = arr[i - 1][j - 1] + arr[i-1][j]
        }
    }
    return arr
};
```