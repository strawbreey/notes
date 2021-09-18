package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	arr := make([]int, len(nums)+1)

	for _, v := range nums {
		if v <= len(nums) {
			arr[v] = 1
		}
	}

	arr1 := make([]int, 0)

	for i := 1; i < len(arr); i++ {
		if arr[i] == 0 {
			arr1 = append(arr1, i)
		}
	}

	return arr1
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Print(findDisappearedNumbers(nums))
}
