---
title: "Go Effective"
date: 2020-11-19T14:55:01+08:00
draft: false
---

### Introduction 介绍

Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.
 
- [The Go Programming Language Specification](https://golang.org/ref/spec)  olang 语言 规范

- [A Tour of Go](https://tour.golang.org/welcome/1)

- [How to Write Go Code](https://golang.org/doc/code.html)

### The Go Programming Language Specification （g）


### Formatting 格式化

As an example, there's no need to spend time lining up the comments on the fields of a structure. Gofmt will do that for you. Given the declaration

Go 语言不用担心格式化问题, Gofmt 会帮你统一格式化


```go
type T struct {
    name string // name of the object
    value int // its value
}
```

=> 

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

### Commentary


## Names

Names are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case. It's therefore worth spending a little time talking about naming conventions in Go programs.


### Package names

```
import "bytes"
```

### Getters

Go doesn't provide automatic support for getters and setters. There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so, but it's neither idiomatic nor necessary to put Get into the getter's name. If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner. The use of upper-case names for export provides the hook to discriminate the field from the method. A setter function, if needed, will likely be called SetOwner. Both names read well in practice:

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```


### Interface names

By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

There are a number of such names and it's productive to honor them and the function names they capture. Read, Write, Close, Flush, String and so on have canonical signatures and meanings. To avoid confusion, don't give your method one of those names unless it has the same signature and meaning. Conversely, if your type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call your string-converter method String not ToString.



### word


```
Semicolons 分号
```