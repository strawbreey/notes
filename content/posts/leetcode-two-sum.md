---
title: "Leetcode Two Sum"
date: 2020-11-06T17:03:36+08:00
draft: false
---

两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


### Example:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1] 

Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]
```
### solution

```c

```

```js

// 暴力解决方案
var twoSum = function(nums, target) {
    let arr = []
    for(let index = 0; index < nums.length; index ++) {
        let targetindex = nums.findIndex(item => item == target - nums[index])
        if (targetindex != -1 && targetindex != index ) {
            arr = [index, targetindex ]
            break
        }
    }
    return arr.sort()
};

// hash
let obj = {}
for(let i = 0; i < nums.length; i++) {
    let key = target - nums[i]
    if (key in obj) {
        return [obj[key], i]
    } else {
        obj[nums[i]] = i
    }
}
```

```go

```
```php

```

### 参考资料

- [两数之和](https://leetcode-cn.com/problems/two-sum/)
- [two sum](https://leetcode.com/problems/two-sum/)