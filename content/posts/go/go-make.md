---
title: "Go Make"
date: 2021-08-11T09:59:15+08:00
draft: false
---

> 当我们想要在 Go 语言中初始化一个结构时，可能会用到两个不同的关键字 — make 和 new。因为它们的功能相似，所以初学者可能会对这两个关键字的作用感到困惑1，但是它们两者能够初始化的变量却有较大的不同。
> - make (分配空间+初始化) 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel;
> - new (分配空间) 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针;

## Go 内置 make 解析

1. 创建一个切片

```go
slice := make([]int, 0, 100)
```


2. 创建一个map：

```go
ages := make(map[string]int) // mapping from strings to ints
ages["alice"] = 31
ages["charlie"] = 34
// => 等同于

ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```

3. 创建一个channel 

```go
ch := make(chan int, 5)
```


相比与复杂的 make 关键字，new 的功能就简单多了，它只能接收类型作为参数然后返回一个指向该类型的指针：

```go
i := new(int)

var v int
i := &v
```

![golang-make-and-new.png](/images/golang-make-and-new.png)

### 总结

- make 关键字的作用是创建切片、哈希表和 Channel 等内置的数据结构，
- new 的作用是为类型申请一片内存空间，并返回指向这片内存的指针。