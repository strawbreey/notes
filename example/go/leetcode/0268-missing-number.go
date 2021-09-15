package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	arr := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		arr[nums[i]] = 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			return i
		}
	}

	return 0
}

func missingNumber1(nums []int) int {
	ans := 0
	for i := range nums {
		ans ^= (i + 1) ^ nums[i]
	}
	return ans
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Print(missingNumber(nums))

}
