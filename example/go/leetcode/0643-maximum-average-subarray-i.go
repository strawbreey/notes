package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	max := 0
	for _, v := range nums[:k] {
		sum += v
	}
	max = sum

	for n := k; n < len(nums); n++ {
		sum = sum - nums[n-k] + nums[n]
		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	k := 5
	fmt.Print(findMaxAverage(nums, k))
}
