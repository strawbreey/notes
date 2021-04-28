---
title: "Leetcode 0004 Median of Two Sorted Arrays"
date: 2020-11-26T23:52:19+08:00
draft: false
---

给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

 
```
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000

输入：nums1 = [], nums2 = [1]
输出：1.00000

输入：nums1 = [2], nums2 = []
输出：2.00000
```

提示：

    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106


solution

```js
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {

    let len = nums1.length + nums2.length

    let size = len % 2 == 0 ? len / 2 - 1 :  (len + 1) / 2 -1

    let arr = nums1.concat(nums2).sort((a,b) => Number(a) -Number(b))

    return len % 2 == 0 ? (arr[size] + arr[size + 1]) / 2 : arr[size]
};

```