package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, k}
		}
		m[v] = k
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// fmt.Print(twoSum(nums, target))
	fmt.Print(twoSum1(nums, target))

}
