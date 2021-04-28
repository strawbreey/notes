---
title: "Go String to Number"
date: 2020-12-01T23:20:00+08:00
draft: false
---

在golang中，用字符串与整型有两种方法

- 使用rune（int32位的别名）来转换

- golang中stroncv包的函数来转换

```go
a := "123456"
b,error := strconv.Atoi(a)

c: = 1234
d := strconv.Itoa(c)   //数字变成字符串
```