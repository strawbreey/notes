---
title: "Go Rune"
date: 2021-09-03T18:02:26+08:00
draft: false
---

在编译代码时 出现了以下错误

Go报错：more than one character in rune literal

原因是用错了引号的表达方式。

```go
package main

import(
    "fmt"
)

func main(){
    fmt.Println('hello world')
}
```


在go语法中，双引号是常用的来表达字符串，如果你使用了单引号，编译器会提示出错

invalid character literal (more than one character)

单引号只能包含一个字符，例如’b’ ,程序会输出98表示字符b的ascii码。

如果非要使用单引号输出必须使用string函数转换

fmt.Println(string('b') )

