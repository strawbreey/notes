package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums)+1)
	for _, v := range nums {
		arr[v] += 1
	}

	a := 0
	b := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] == 2 {
			a = i
		}
		if arr[i] == 0 {
			b = i
		}
	}

	return []int{a, b}

}
func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Print(findErrorNums(nums))
}
