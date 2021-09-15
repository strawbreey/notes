package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func main() {
	nums := []int{1, 1, 2}
	targer := 1
	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	num := removeElement(nums, targer)

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
	fmt.Println(num)
	for i := 0; i < num; i++ {
		fmt.Println(nums[i])
	}
}
