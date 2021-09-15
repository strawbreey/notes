package main

import (
	"fmt"
	"sort"
)

func mapToSlice(m map[int]string) []string {
	s := make([]string, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func thirdMax(nums []int) int {
	// s := []int{4, 2, 3, 1}
	// sort.Ints(nums)

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	s := make([]int, 0, len(m))
	for k, _ := range m {
		s = append(s, k)
	}

	sort.Ints(s)

	if len(s) < 3 {
		return s[len(s)-1]
	} else {
		return s[len(s)-3]
	}

}

func main() {
	nums1 := []int{3, 1}
	fmt.Print(thirdMax(nums1))
}
