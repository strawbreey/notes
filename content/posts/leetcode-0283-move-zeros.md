---
title: "Leetcode 0283 Move Zeros"
date: 2020-11-19T10:00:00+08:00
tags: ['leetcode'] 
draft: false
---

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。


Example:

```
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
```


Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.

说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

Solution

非零元素右移， 然后补零

```js
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    let count = 0
    for (let i = 0; i < nums.length; i++) {
        if (nums[i]) {
            nums[count] = nums[i] 
            count++  
        }
    }

    for (let i = 0; i < (nums.length - count); i++) {
        nums[nums.length - 1 - i] = 0
    }
    return nums
};
```


双重循环交换非零元素

```js
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    for (let i = 0; i < nums.length - 1; i++) {
        if (nums[i] == 0) {
            for (let j = i + 1; j < nums.length; j++) {
                if (nums[j]) {
                    nums[i] = nums[j]
                    nums[j] = 0
                    break
                }

            }
        }
    }
    return nums
};
```

双指针解法

使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。

右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

注意到以下性质：

  - 左指针左边均为非零数；

  - 右指针左边直到左指针处均为零

![Double pointer](/images/36d1ac5d689101cbf9947465e94753c626eab7fcb736ae2175f5d87ebc85fdf0-283_2.gif){:height="50%" width="50%"}

```go
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
func moveZeroes(nums []int) {
    left, right, n := 0, 0, len(nums)
    for right < n {
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
        right++
    }
}
```

