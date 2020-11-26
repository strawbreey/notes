---
title: "Leetcode 0056 Merge Intervals"
date: 2020-11-23T00:26:57+08:00
draft: false
tags: ['leetcode']
---

给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

 
提示：

intervals[i][0] <= intervals[i][1]


```js
/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
    points = intervals.sort((a,b) => (a[0] - b[0]) || (a[1] - b[1]))
    let mergeArr = [], count = 0;

    for (let i = 0; i< points.length; i++) {
        if (mergeArr.length == 0) {
            mergeArr.push(points[i])
            count++
        } else {
            
            if (points[i][0] <= mergeArr[count - 1][1]) {
                if (points[i][1] > mergeArr[count - 1][1]) {
                    mergeArr[count-1][1] = points[i][1]
                }
            } else  {
                mergeArr.push(points[i])
                count++
            }
        }

    }

    return mergeArr
};
```

### 参考资料

- [合并区间](https://leetcode-cn.com/problems/merge-intervals/)
