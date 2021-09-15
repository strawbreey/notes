package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, v := range nums {
		for j := i + 1; j <= i+k; j++ {
			if j < len(nums) && v == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {

	m := make(map[int]int) // 声明map
	for i, v := range nums {
		if m[v] != 0 && i-m[v] < k {
			return true
		} else {
			m[v] = i + 1
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Print(containsNearbyDuplicate1(nums, k))
}
