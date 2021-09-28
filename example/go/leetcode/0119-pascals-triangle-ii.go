package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	ans := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans[rowIndex]
}

func main() {
	n := 5
	fmt.Println(getRow(n))
}
