package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	arr := make([]int, m+n)
	for i+j < m+n {
		if nums1[i] <= nums2[j] && i < m {
			arr[i+j] = nums1[i]
			i += 1
		} else if j < n {
			arr[i+j] = nums2[j]
			j += 1
		}
	}
	nums1 = arr
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
