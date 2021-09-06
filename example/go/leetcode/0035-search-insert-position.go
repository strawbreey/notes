package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	} else if nums[len(nums)-1] < target {
		return len(nums)
	} else if nums[0] > target {
		return 0
	}

	for k, _ := range nums {
		if nums[k] == target {
			return k
		} else if nums[k+1] > target {
			return k + 1
		}
	}
	return len(nums)
}

func searchInsert1(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func searchInsert2(nums []int, target int) int {
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	} else if nums[len(nums)-1] < target {
		return len(nums)
	} else if nums[0] > target {
		return 0
	}

	for k, _ := range nums {
		if nums[k] == target {
			return k
		} else if nums[k+1] > target {
			return k + 1
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Print(searchInsert(nums, target))
}
