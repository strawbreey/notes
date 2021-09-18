package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			num++
		}

		if nums[i] == 0 {
			num = 0
		}

		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Print(findMaxConsecutiveOnes(nums))
}
