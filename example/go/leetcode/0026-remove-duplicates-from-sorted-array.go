package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	left := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func main() {
	nums := []int{1, 1, 2}

	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	num := removeDuplicates(nums)

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。

	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(nums[i])
	}
}
