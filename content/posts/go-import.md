---
title: "Go Import"
date: 2020-12-14T11:29:30+08:00
draft: false
---

可以在一个 Go语言源文件包声明语句之后，其它非导入声明语句之前，包含零到多个导入包声明语句。

每个导入声明可以单独指定一个导入路径，也可以通过圆括号同时导入多个导入路径。

要引用其他包的标识符，可以使用 import 关键字，导入的包名使用双引号包围，包名是从 GOPATH 开始计算的路径，使用/进行路径分隔。

默认导入的写法

导入有两种基本格式，即单行导入和多行导入，两种导入方法的导入代码效果是一致的。


```go
// 1. 单行导入 单行导入格式如下：
import "fmt"
import "math"

// 2. 多行导入 当多行导入时，包名在 import 中的顺序不影响导入效果，格式如下：
import(
  "fmt"
  "math"
)

```

Go was designed to be a language that encourages good software engineering practices. An important part of high quality software is code reuse – embodied in the principle “Don't Repeat Yourself.”

As we saw in chapter 7 functions are the first layer we turn to allow code reuse. Go also provides another mechanism for code reuse: packages. Nearly every program we've seen so far included this line:

import "fmt"
fmt is the name of a package that includes a variety of functions related to formatting and output to the screen. Bundling code in this way serves 3 purposes:

It reduces the chance of having overlapping names. This keeps our function names short and succinct

It organizes code so that its easier to find code you want to reuse.

It speeds up the compiler by only requiring recompilation of smaller chunks of a program. Although we use the package fmt, we don't have to recompile it every time we change our program.

Creating Packages
Packages only really make sense in the context of a separate program which uses them. Without this separate program we have no way of using the package we create. Let's create an application that will use a package we will write. Create a folder in ~/src/golang-book called chapter11. Inside that folder create a file called main.go which contains this:

package main

import "fmt"
import "golang-book/chapter11/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  fmt.Println(avg)
}
Now create another folder inside of the chapter11 folder called math. Inside of this folder create a file called math.go that contains this:

package math

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}
Using a terminal in the math folder you just created run go install. This will compile the math.go program and create a linkable object file: ~/pkg/os_arch/golang-book/chapter11/math.a. (where os is something like windows and arch is something like amd64)

Now go back to the chapter11 folder and run go run main.go. You should see 2.5. Some things to note:

math is the name of a package that is part of Go's standard distribution, but since Go packages can be hierarchical we are safe to use the same name for our package. (The real math package is just math, ours is golang-book/chapter11/math)

When we import our math library we use its full name (import "golang-book/chapter11/math"), but inside of the math.go file we only use the last part of the name (package math).

We also only use the short name math when we reference functions from our library. If we wanted to use both libraries in the same program Go allows us to use an alias:

import m "golang-book/chapter11/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := m.Average(xs)
  fmt.Println(avg)
}
m is the alias.

You may have noticed that every function in the packages we've seen start with a capital letter. In Go if something starts with a capital letter that means other packages (and programs) are able to see it. If we had named the function average instead of Average our main program would not have been able to see it.

It's a good practice to only expose the parts of our package that we want other packages using and hide everything else. This allows us to freely change those parts later without having to worry about breaking other programs, and it makes our package easier to use.

Package names match the folders they fall in. There are ways around this, but it's a lot easier if you stay within this pattern.

Documentation
Go has the ability to automatically generate documentation for packages we write in a similar way to the standard package documentation. In a terminal run this command:

godoc golang-book/chapter11/math Average
You should see information displayed for the function we just wrote. We can improve this documentation by adding a comment before the function:

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
If you run go install in the math folder, then re-run the godoc command you should see our comment below the function definition. This documentation is also available in web form by running this command:

godoc -http=":6060"
and entering this URL into your browser:

http://localhost:6060/pkg/
You should be able to browse through all of the packages installed on your system.

Problems
Why do we use packages?

What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)

What is a package alias? How do you make one?

We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

How would you document the functions you created in #3?