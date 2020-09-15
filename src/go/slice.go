package main

import "fmt"

func main(){

	// s := []int{5}
	// s = append(s, 7)
	// s = append(s, 9)
	// x := append(s, 11)
	// y := append(s, 12)
	// fmt.Println(s, x, y)

	s := []int{5}
	fmt.Println(len(s), cap(s), s, &s[0])

	s = append(s, 7) //新建切片并拷贝内容, 按Slice扩容机制，cap(s)翻倍 == 2

	fmt.Println(len(s), cap(s), s, &s[0])

	s = append(s, 9) //新建切片并拷贝内容, 按Slice扩容机制，cap(s)翻倍 == 4
	fmt.Println(len(s), cap(s), s, &s[0])

	x := append(s, 11)
	fmt.Println(len(s), cap(s), s, x, &s[0])

	y := append(s, 12)
	fmt.Println(len(s), cap(s), s, x, y, &s[0]) // s的容量为4, 长度为3, 不会发生扩容。append 12 改变了容量4的地址

	z := append(s, 13)
	fmt.Println(len(s), cap(s), s, x, y, z, &s[0]) // s的容量为4, 长度为3, 增加13, 15 触发发生扩容。append 12 改变了容量4的地址

	// 输出结果是 [5 7 9] [5 7 9 12] [5 7 9 12]

	z[3] = 16
	fmt.Println(len(s), cap(s), s, x, y, z, &s[0])
}