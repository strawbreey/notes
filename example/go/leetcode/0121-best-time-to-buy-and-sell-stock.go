package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	num := len(prices)
	min := prices[0]
	for i := 0; i < num; i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	fmt.Println(maxProfit(nums))
}
