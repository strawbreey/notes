---
title: "Go Huge Size of Compile"
date: 2021-09-16T11:04:56+08:00
draft: false
---

编译了一个hello world Go程序，该程序在linux机器上生成了本机可执行文件。但是我很惊讶地看到简单的Hello world Go程序的大小为1.9MB！

为什么Go中如此简单的程序的可执行文件如此庞大？



1. 如果我在Linux AMD64机器（Go 1.9）上构建以下文件：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}

```

编译

```bash
$ go build ./deme.go
```

我得到的二进制文件大小约为2 Mb。

这样做的原因是因为我们使用的“ fmt”包很大，但是二进制文件也没有被剥离，这意味着符号表仍然存在。如果我们改为指示编译器剥离二进制文件，它将变得更小：


2. 如果我们重写程序以使用内置函数print而不是fmt.Println，如下所示：

```go
package main

func main() {
    print("Hello World!\n")
}

```

```bash
$ go build ./deme.go
```

我得到的二进制文件大小约为1200 kb。

3. 如果我们改为指示编译器剥离二进制文件，它将变得更小

```bash
go build  -ldflags=-w .\test-fmt.go 
```

我们最终得到了一个甚至更小的二进制文件, 这是我们能得到的最小尺寸，而无需借助UPX打包之类的技巧，因此Go运行时的开销约为700 Kb。

### 参考资料

- https://golang.org/doc/faq#Why_is_my_trivial_program_such_a_large_binary