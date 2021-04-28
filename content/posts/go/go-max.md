---
title: "Go Max"
date: 2020-12-01T13:09:31+08:00
draft: false
---

```go
maxCandies = max(maxCandies, candies[i])
```

go语言math包里面定义了min/max函数，但是是float64类型的，而并没有整数类型的min/max。

go语言的math包里面定义的min/max函数如下：

```go
math.Min(float64, float64) float64
math.Max(float64, float64) float64
```

事实上我们更经常要比较的是两个整数的场景：

```go
math.Min/Max(int, int), or
math.Min/Max(int64, int64)
```

由于float64类型要处理infinity和not-a-number这种值，而他们的处理非常复杂，一般用户没有能力，所有go需要为用户提供系统级别的解决办法。

对于int/int64类型的数据，min/max的实现非常简单直接，用户完全可以自己实现，例如：

```go
func Min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}
```


```go
package main

import (
    "fmt"
    "math"
)

func main() {
    ints := []int{1, 2, 3, 4, 5}
    max := 0

    for _, n := range ints {
        max = int(math.Max(float64(max), float64(n)))
    }

    fmt.Println(max)
}
```

```go
func Min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}

func Max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}
```


### 参考资料

- [Don't abuse math.Max / math.Min](https://mrekucci.blogspot.com/2015/07/dont-abuse-mathmax-mathmin.html)