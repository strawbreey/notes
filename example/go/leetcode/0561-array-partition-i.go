package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[2*i] > nums[2*i+1] {
			sum += nums[2*i+1]
		} else {
			sum += nums[2*i]
		}
	}
	return sum
}

func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Print(maximumProduct(nums))
}
