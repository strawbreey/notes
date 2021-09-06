package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(123))
}
