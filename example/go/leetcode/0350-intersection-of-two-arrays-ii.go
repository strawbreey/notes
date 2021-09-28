package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		m1[v] += 1
	}

	for _, v := range nums2 {
		m2[v] += 1
	}

	arr := make([]int, 0)

	for k, _ := range m1 {
		print(k)
		print(m1[k])
		println(m2[k])

		if m1[k] >= m2[k] {
			for i := 0; i < m2[k]; i++ {
				arr = append(arr, k)
			}
		} else {
			for i := 0; i < m1[k]; i++ {
				arr = append(arr, k)
			}
		}
	}
	return arr
}

func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{1, 0, 1}

	fmt.Print(intersect(nums1, nums2))
}
