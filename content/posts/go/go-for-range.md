---
title: "Go for Range"
date: 2020-11-30T15:46:43+08:00
draft: false
---


range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned


```go
package main

import "fmt"

func main() {

    // 声明数组
    nums := []int{2, 3, 4}

    // 声明遍历
    sum := 0
    
    // for i, num := range nums i, 
    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}

    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```


### 参考资料

- [Go by Example: Range](https://gobyexample.com/range)