---
title: "Go Interview"
date: 2020-09-03T16:18:11+08:00
draft: false
---

```Go
func main(){
    s := []int{5}

    fmt.Println(len(s), cap(s))

    s = append(s, 7)

    fmt.Println(len(s), cap(s))

    s = append(s, 9)
    fmt.Println(len(s), cap(s))

    x := append(s, 11)
    fmt.Println(len(s), cap(s))

    y := append(s, 12)
    fmt.Println(len(s), cap(s))

    fmt.Println(s, x, y)
    // 输出结果 [5 7 9] [5 7 9 12] [5 7 9 12]

   	s := []int{5, 7, 9}
    x := append(s, 11)
    y := append(s, 12)
    fmt.Println(s,x,y)

    // 输出结果 [5 7 9] [5 7 9 11] [5 7 9 12]

    s := []int{5} // 初始化切片为[5], cap(s) = 1 len (s) = 1
    
    fmt.Println(len(s), cap(s), s)

    s = append(s, 7) //新建切片并拷贝内容, 按Slice扩容机制，cap(s)翻倍 == 2

    fmt.Println(len(s), cap(s), s)

    s = append(s, 9) //新建切片并拷贝内容, 按Slice扩容机制，cap(s)翻倍 == 4
    fmt.Println(len(s), cap(s), s)

    x := append(s, 11)
    fmt.Println(len(s), cap(s), s, x)

    y := append(s, 12)
    fmt.Println(len(s), cap(s), s, x, y) // s的容量为4, 长度为3, 不会发生扩容。append 12 改变了容量4的地址

    z := append(s, 13, 15)
    fmt.Println(len(s), cap(s), s, x, y, z) // s的容量为4, 长度为3, 增加13, 15 触发发生扩容。append 12 改变了容量4的地址

    // 输出结果是 [5 7 9] [5 7 9 12] [5 7 9 12]
}
```

