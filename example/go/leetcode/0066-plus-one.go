package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] = digits[n-1] + 1

	for i := n - 1; i >= 0; i-- {
		if digits[i] > 9 {
			digits[i] = digits[i] - 10
			if i > 0 {
				digits[i-1] = digits[i-1] + 1
			} else {
				s := []int{1}
				s = append(s, digits...)
				return s
			}

		}
	}

	return digits
}

func main() {
	nums := []int{1, 1, 2}

	fmt.Println(nums[:1])
	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	fmt.Println(plusOne(nums))
}
