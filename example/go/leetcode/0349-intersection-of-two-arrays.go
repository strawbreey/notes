package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {

	m := make(map[int]int) // 声明map
	arr := make([]int, 0)

	for _, v := range nums1 {
		m[v] = 1
	}

	for _, v := range nums2 {
		if m[v] == 1 {
			m[v] = 2
		}
	}

	for k, v := range m {
		if v > 1 {
			arr = append(arr, k)
		}
	}

	return arr
}

func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{1, 0, 1}

	fmt.Print(intersection(nums1, nums2))

}
