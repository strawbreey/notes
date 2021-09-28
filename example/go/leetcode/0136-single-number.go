package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	obj := make(map[int]int)
	for _, v := range nums {
		obj[v] += 1
	}

	for k, v := range obj {
		if v == 1 {
			return k
		}
	}

	return 0
}

// 答案是使用位运算。对于这道题，可使用异或运算 \oplus⊕。异或运算有以下三个性质。

// 任何数和 00 做异或运算，结果仍然是原来的数，即 a \oplus 0=aa⊕0=a。
// 任何数和其自身做异或运算，结果是 00，即 a \oplus a=0a⊕a=0。
// 异或运算满足交换律和结合律，即 a \oplus b \oplus a=b \oplus a \oplus a=b \oplus (a \oplus a)=b \oplus0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。

func singleNumber1(nums []int) int {
	single := 0
	// 位运算
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Print(singleNumber(nums))
}
