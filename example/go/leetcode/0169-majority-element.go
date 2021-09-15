package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	// m := make(map[int]int) // 声明map
	// num := len(nums)
	// for i := 0; i < num; i++ {
	// 	m[nums[i]]++
	// }

	// fmt.Println(m)

	// for k, v := range m {
	// 	if v > num/2 {
	// 		return k
	// 	}
	// }
	// return 0
	sort.Ints(nums)
	num := nums[len(nums)/2]
	return num
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Print(majorityElement(nums))
}
