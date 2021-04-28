---
title: "Leetcode 0659 Split Array Into Consecutive Subsequences"
date: 2020-12-07T11:23:03+08:00
draft: false
---

给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个长度为 3 的子序列，其中每个子序列都由连续整数组成。

如果可以完成上述分割，则返回 true ；否则，返回 false 。

```
输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3
3, 4, 5

输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3, 4, 5
3, 4, 5

输入: [1,2,3,4,4,5]
输出: False

```

提示：

`1 <= nums.length <= 10000`

Solution:

```js
func isPossible(nums []int) bool {
    if len(nums) < 3 {
		return false
	}

	var useNum bool

	var as = make([][]int, 0)
	for _, num := range nums {
		useNum = false

		for aIdx, _ := range as {
			start := len(as) - aIdx - 1
			a := as[start : len(as)-aIdx][0]
			if num == (a[(len(a) - 1):][0] + 1) {
				useNum = true
				as[start] = append(a, num)
				break
			}
		}

		if !useNum {
			as = append(as, []int{num})
		}
	}

	for _, a := range as {
		if len(a) < 3 {
			return false
		}
	}

	return true
}
```