package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	num := len(numbers)
	for i := 0; i < num-1; i++ {
		for j := i + 1; j < num; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// fmt.Print(twoSum(nums, target))
	fmt.Print(twoSum(nums, target))

}
