---
title: "Leetcode 0452 Mininum Number of Arrows to Burst Balloons"
date: 2020-11-23T19:36:30+08:00
draft: false
tags: ['leetcode']
---

在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。

一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。

给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。

```
输入：points = [[10,16],[2,8],[1,6],[7,12]]
输出：2
解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球


输入：points = [[1,2],[3,4],[5,6],[7,8]]
输出：4


输入：points = [[1,2],[2,3],[3,4],[4,5]]
输出：2


输入：points = [[1,2]]
输出：1


输入：points = [[2,3],[2,3]]
输出：1

```


提示：

```
  0 <= points.length <= 104
  points[i].length == 2
  -231 <= xstart < xend <= 231 - 1
```

Solution

```js
/**
 * @param {number[][]} points
 * @return {number}
 */
var findMinArrowShots = function(points) {
    points = points.sort((a,b) => (a[0] - b[0]) || (a[1] - b[1]))
    let mergeArr = [], count = 0;

    for (let i = 0; i< points.length; i++) {
        if (mergeArr.length == 0) {
            mergeArr.push(points[i])
            count++
        } else {
            
            if (points[i][0] <= mergeArr[count - 1][1]) {
                mergeArr[count-1][0] = points[i][0]
            } 

            if (points[i][1] < mergeArr[count - 1][1]) {
                mergeArr[count-1][1] = points[i][1]
            }

            if (points[i][0] > mergeArr[count - 1][1]) {
                mergeArr.push(points[i])
                count++
            }
        }

    }
    return mergeArr.length
};
```

```js
var findMinArrowShots = function(points) {

    if (points.length == 0) {
        return 0
    }

    points = points.sort((a,b) => a[1] - b[1])

    let pos = points[0][1]
    let count = 1
    points.forEach(item => {
        if (item[0] > pos) {
            count++
            pos = item[1]
        }
    })
    return count
};
```


### 参考链接 

- [用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)

