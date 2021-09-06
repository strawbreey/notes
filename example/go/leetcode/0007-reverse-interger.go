package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	y := 0

	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	} else {
		return y
	}
}

func main() {
	// i := 'a'
	// // fmt.Print(reverse(i))
	// s := "hello world"
	// z := "你好中国"

	// fmt.Println(i)
	// fmt.Println(s[:2])
	// fmt.Println(string([]rune(z)[:2]))

	// fmt.Println(len(strconv.Itoa(1234)))
	// fmt.Println(int(^uint32(0) >> 1))

	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.Abs(-120))
	fmt.Println(reverse(123))

	fmt.Println(strconv.Itoa(123))
}
