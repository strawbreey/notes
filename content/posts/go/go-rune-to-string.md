---
title: "Go Rune to String"
date: 2021-09-24T20:25:31+08:00
draft: false
---

rune int32的别名，几乎在所有方面等同于int32
// 它用来区分字符值和整数值

在 Go 中有几种方式将 rune 类型转换为 string

```go
string('c')
// string(rune('c'))
```
或者借助 strconv 包

```
strconv.QuoteRune('c')
```

这种方式会将单引号也打印出来