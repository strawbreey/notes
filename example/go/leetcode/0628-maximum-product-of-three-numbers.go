package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)

	if nums[0]*nums[1]*nums[len(nums)-1] > nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3] {
		return nums[0] * nums[1] * nums[len(nums)-1]
	} else {
		return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	}

}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Print(maximumProduct(nums))
}
