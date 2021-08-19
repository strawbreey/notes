---
title: "Go Version"
date: 2021-08-17T15:33:39+08:00
draft: false
---

## 获取go 当前的版本号

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.Version())
}
```