---
title: "Go Example"
date: 2020-12-14T12:18:37+08:00
draft: false
categories: ['GO']
tags: ['GO']
---

### 1. 编译

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```

要运行这个程序，将这些代码放到 hello-world.go 中并且使用 go run 命令。

```bash
go run hello-world.go
hello world
```

有时候我们想将我们的程序编译成二进制文件。我们可以通过 go build 命来达到目的。
```
go build hello-world.go
ls
```

> hello-world.exe

然后我们可以直接运行这个二进制文件。

```bash
ls
hello-world.exe	hello-world.go
./hello-world
```


### 2. 字符串

```go
package main
import "fmt"

func main() {
    fmt.Println("go" + "lang")
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

    // golang
    // 1+1 = 2
    // 7.0/3.0 = 2.3333333333333335
    // false
    // true
    // fals
}
```

### 3. 变量

1. var 声明 1 个或者多个变量。
2. 你可以申明一次性声明多个变量。
3. Go 将自动推断已经初始化的变量类型。
3. 声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
4. := 语句是申明并初始化变量的简写，例如这个例子中的 var f string = "short"。

```go
package main
import "fmt"
func main() {
    var a string = "initial"
    var b, c int = 1, 2
    var d = true
    var e int
    f := "short"
    fmt.Println(f)

    // initial
    // 1 2
    // true
    // 0
    // short
}
```

### 4. 常量
- Go 支持字符、字符串、布尔和数值 常量 。
- const 用于声明一个常量。
- const 语句可以出现在任何 var 语句可以出现的地方
- 常数表达式可以执行任意精度的运算
- 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化

```go
package main
import "fmt"
import "math"

const s string = "constant"

func main() {
    fmt.Println(s)
    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))

    // constant
    // 6e+11
    // 600000000000
    // -0.28470407323754404
}
```

### 5. For 循环

for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式。

- 单个循环条件
- 经典的初始化/条件/后续形式 for 循环
- 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环

```go
// `for` 是 Go 中唯一的循环结构。这里有 `for` 循环的三个基本使用方式。
package main
import "fmt"

func main() {

	// 最常用的方式，带单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 `for` 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环将一直执行，直到在循环体内使用了 `break` 或者 `return` 来跳出循环。
	for {
		fmt.Println("loop")
		break
	}
}
	
// $ go run for.go
// 1
// 2
// 3
// 7
// 8
// 9
// loop
```

### 6. if/else 分支

```go
// `if` 和 `else` 分支结构在 Go 中当然是直接了当的了。
package main
import "fmt"

func main() {

    // 这里是一个基本的例子。
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // 你可以不要 `else` 只用 `if` 语句。
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // 在条件语句之前可以有一个语句；任何在这里声明的变量
    // 都可以在所有的条件分支中使用。
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// 注意，在 Go 中，你可以不适用圆括号，但是花括号是需
// 要的。

```
> Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句。

### 7. switch 分支结构

```go
// _switch_ ，方便的条件分支语句。
package main
import "fmt"
import "time"

func main() {
	// 一个基本的 `switch`。
	i := 2
	fmt.Print("write ", i, " as ")

	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	// 在一个 `case` 语句中，你可以使用逗号来分隔多个表达式。在这个例子中，我们很好的使用了可选的`default` 分支。
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("it's the weekend")
		default:
			fmt.Println("it's a weekday")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种方式。这里展示了 `case` 表达式是如何使用非常量的。
	t := time.Now()
	switch {
		case t.Hour() < 12:
				fmt.Println("it's before noon")
		default:
				fmt.Println("it's after noon")
	}
}

// todo: type switches
```

### 8. 数组

```go
// 在 Go 中，_数组_ 是一个固定长度的数列。

package main
import "fmt"

func main() {
    // 这里我们创建了一个数组 `a` 来存放刚好 5 个 `int`。
    // 元素的类型和长度都是数组类型的一部分。数组默认是
    // 零值的，对于 `int` 数组来说也就是 `0`。
    var a [5]int
    fmt.Println("emp:", a)

    // 我们可以使用 `array[index] = value` 语法来设置数组
    // 指定位置的值，或者用 `array[index]` 得到值。
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // 使用内置函数 `len` 返回数组的长度
    fmt.Println("len:", len(a))

    // 使用这个语法在一行内初始化一个数组
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 数组的存储类型是单一的，但是你可以组合这些数据
    // 来构造多维的数据结构。
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

```

### 9. 切片

```go
// _Slice_ 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

package main
import "fmt"

func main() {

    // 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。
    // 要创建一个长度不为 0 的空 slice，需要使用内建函数 `make`。
    // 这里我们创建了一个长度为 3 的 `string` 类型的 slice（初始值为零值）。
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // 我们可以和数组一样设置和得到值
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` 返回 slice 的长度
    fmt.Println("len:", len(s))

    // 除了基本操作外，slice 支持比数组更丰富的操作。比如 slice 支持内建函数 `append`，
    // 该函数会返回一个包含了一个或者多个新值的 slice。
    // 注意由于 `append` 可能返回一个新的 slice，我们需要接收其返回值。
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // slice 还可以 `copy`。这里我们创建一个空的和 `s` 有相同长度的 slice——`c`，
    // 然后将 `s` 复制给 `c`。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // slice 支持通过 `slice[low:high]` 语法进行“切片”操作。
    // 例如，右边的操作可以得到一个包含元素 `s[2]`、`s[3]` 和 `s[4]` 的 slice。
    l := s[2:5]
    fmt.Println("sl1:", l)

    // 这个 slice 包含从 `s[0]` 到 `s[5]`（不包含 5）的元素。
    l = s[:5]
    fmt.Println("sl2:", l)

    // 这个 slice 包含从 `s[2]`（包含 2）之后的元素。
    l = s[2:]
    fmt.Println("sl3:", l)

    // 我们可以在一行代码中声明并初始化一个 slice 变量。
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

### 10. Map

```go
// _map_ 是 Go 内建的[关联数据类型](http://zh.wikipedia.org/wiki/关联数组)
// （在一些其他的语言中也被称为 _哈希(hash)_ 或者 _字典(dict)_ ）。

package main
import "fmt"

func main() {

    // 要创建一个空 map，需要使用内建函数 `make`：`make(map[key-type]val-type)`。
    m := make(map[string]int)

    // 使用典型的 `make[key] = val` 语法来设置键值对。
    m["k1"] = 7
    m["k2"] = 13

    // 打印 map。例如，使用 `fmt.Println` 打印一个 map，会输出它所有的键值对。
    fmt.Println("map:", m)

    // 使用 `name[key]` 来获取一个键的值。
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // 内建函数 `len` 可以返回一个 map 的键值对数量。
    fmt.Println("len:", len(m))

    // 内建函数 `delete` 可以从一个 map 中移除键值对。
    delete(m, "k2")
    fmt.Println("map:", m)

    // 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。
    // 这可以用来消除 `键不存在` 和 `键的值为零值` 产生的歧义，
    // 例如 `0` 和 `""`。这里我们不需要值，所以用 _空白标识符(blank identifier)_ _ 将其忽略。
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // 你也可以通过右边的语法在一行代码中声明并初始化一个新的 map。
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

// map: map[k1:7 k2:13]
// v1:  7
// len: 2
// map: map[k1:7]
// prs: false
// map: map[bar:2 foo:1]
```

### 11. Range 遍历

```go
// _range_ 用于迭代各种各样的数据结构。
// 让我们来看看如何在我们已经学过的数据结构上使用 `range`。

package main

import "fmt"

func main() {

    // 这里我们使用 `range` 来对 slice 中的元素求和。
    // 数组也可以用这种方法初始化并赋初值。
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` 在数组和 slice 中提供对每项的索引和值的访问。
    // 上面我们不需要索引，所以我们使用 _空白标识符_ `_` 将其忽略。
    // 实际上，我们有时候是需要这个索引的。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` 在 map 中迭代键值对。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` 也可以只遍历 map 的键。
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` 在字符串中迭代 unicode 码点(code point)。
    // 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}


```

### 12. 函数

```go
// _函数_ 是 Go 的核心。我们将通过一些不同的例子来进行学习它。
package main
import "fmt"

// 这里是一个函数，接受两个 `int` 并且以 `int` 返回它们的和
func plus(a int, b int) int {
    // Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
    return a + b
}

// 当多个连续的参数为同样类型时，
// 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
    return a + b + c
}

// Go 原生支持 _多返回值_。
// 这个特性在 Go 语言中经常用到，例如用来同时返回一个函数的结果和错误信息。
// `(int, int)` 在这个函数中标志着这个函数返回 2 个 `int`。
func vals() (int, int) {
    return 3, 7
}

func main() {
    // 通过 `函数名(参数列表)` 来调用函数，
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    // 这里我们通过 _多赋值_ 操作来使用这两个不同的返回值。
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // 如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 `_`。
    _, c := vals()
    fmt.Println(c)
}

```

### 13. 多返回值

### 14. 可变参数函数

```go

// [_可变参数函数_](https://zh.wikipedia.org/wiki/%E5%8F%AF%E8%AE%8A%E5%8F%83%E6%95%B8%E5%87%BD%E6%95%B8)。
// 在调用时可以传递任意数量的参数。
// 例如，`fmt.Println` 就是一个常见的变参函数。

package main
import "fmt"

// 这个函数接受任意数量的 `int` 作为参数。
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // 变参函数使用常规的调用方式，传入独立的参数。
    sum(1, 2)
    sum(1, 2, 3)

    // 如果你有一个含有多个值的 slice，想把它们作为参数使用，
    // 你需要这样调用 `func(slice...)`。
    nums := []int{1, 2, 3, 4}

    // ...展开符
    sum(nums...)
}
```

### 15. 闭包

```go
// Go 支持[_匿名函数_](http://zh.wikipedia.org/wiki/%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)，
// 并能用其构造 <a href="http://zh.wikipedia.org/wiki/%E9%97%AD%E5%8C%85_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>闭包</em></a>。
// 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

package main
import "fmt"

// `intSeq` 函数返回一个在其函数体内定义的匿名函数。
// 返回的函数使用闭包的方式 _隐藏_ 变量 `i`。
// 返回的函数 _隐藏_ 变量 `i` 以形成闭包。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    // 我们调用 `intSeq` 函数，将返回值（一个函数）赋给 `nextInt`。
    // 这个函数的值包含了自己的值 `i`，这样在每次调用 `nextInt` 时，都会更新 `i` 的值。
    nextInt := intSeq()

    // 通过多次调用 `nextInt` 来看看闭包的效果。
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
    newInts := intSeq()
    fmt.Println(newInts())
}

```


### 16. 递归
```go
// Go 支持 <a href="http://zh.wikipedia.org/wiki/%E9%80%92%E5%BD%92_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>递归</em></a>。
// 这里是一个经典的阶乘示例。

package main

import "fmt"

// `fact` 函数在到达 `fact(0)` 前一直调用自身。
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}

```

### 17. 指针

```go
// Go 支持 <em><a href="http://zh.wikipedia.org/wiki/%E6%8C%87%E6%A8%99_(%E9%9B%BB%E8%85%A6%E7%A7%91%E5%AD%B8)">指针</a></em>，
// 允许在程序中通过 `引用传递` 来传递值和数据结构。

package main
import "fmt"

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较 `指针` 和 `值`。
// `zeroval` 有一个 `int` 型参数，所以使用值传递。
// `zeroval` 将从调用它的那个函数中得到一个实参的拷贝：ival。
func zeroval(ival int) {
    ival = 0
}

// `zeroptr` 有一个和上面不同的参数：`*int`，这意味着它使用了 `int` 指针。
// 紧接着，函数体内的 `*iptr` 会 _解引用_ 这个指针，从它的内存地址得到这个地址当前对应的值。
// 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // 指针也是可以被打印的。
    fmt.Println("pointer:", &i)
}

// initial: 1
// zeroval: 1
// zeroptr: 0
// pointer: 0xc00002c008
```

### 17. 结构体

```go
// Go 的_结构体(struct)_ 是带类型的字段(fields)集合。
// 这在组织数据时非常有用。

package main
import "fmt"

// 这里的 `person` 结构体包含了 `name` 和 `age` 两个字段。
type person struct {
    name string
    age  int
}

func main() {

    // 使用这个语法创建新的结构体元素。
    fmt.Println(person{"Bob", 20})

    // 你可以在初始化一个结构体元素时指定字段名字。
    fmt.Println(person{name: "Alice", age: 30})

    // 省略的字段将被初始化为零值。
    fmt.Println(person{name: "Fred"})

    // `&` 前缀生成一个结构体指针。
    fmt.Println(&person{name: "Ann", age: 40})

    // 使用`.`来访问结构体字段。
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // 也可以对结构体指针使用`.` - 指针会被自动解引用。
    sp := &s
    fmt.Println(sp.age)

    // 结构体是可变(mutable)的。
    sp.age = 51
    fmt.Println(sp.age)
}

```

### 18. 方法

```go
// Go 支持为结构体类型定义_方法(methods)_ 。

package main
import "fmt"

type rect struct {
    width, height int
}

// 这里的 `area` 是一个拥有 `*rect` 类型接收器(receiver)的方法。
func (r *rect) area() int {
    return r.width * r.height
}

// 可以为值类型或者指针类型的接收者定义方法。
// 这是一个值类型接收者的例子。
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // 这里我们调用上面为结构体定义的两个方法。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // 调用方法时，Go 会自动处理值和指针之间的转换。
    // 想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值，
    // 你都可以使用指针来调用方法。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
```


### 19. 接口

```go
// 方法签名的集合叫做：_接口(Interfaces)_。

package main
import (
    "fmt"
    "math"
)

// 这是一个几何体的基本接口。
type geometry interface {
    area() float64
    perim() float64
}

// 在这个例子中，我们将为 `rect` 和 `circle` 实现该接口。
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
// 这里我们为 `rect` 实现了 `geometry` 接口。
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。
// 这儿有一个通用的 `measure` 函数，我们可以通过它来使用所有的 `geometry`。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // 结构体类型 `circle` 和 `rect` 都实现了 `geometry` 接口，
    // 所以我们可以将其实例作为 `measure` 的参数。
    measure(r)
    measure(c)
}

```

### 21. 错误处理

符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。 Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。

```go
package main
import ( "errors" "fmt" )


// 按照惯例，错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口。
func f1(arg int) (int, error) {
    if arg == 42 {
        // errors.New 使用给定的错误信息构造一个基本的 error 值。 
        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}


// 你还可以通过实现 Error() 方法来自定义 error 类型。 这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
    arg  int
    prob string
}

// 你还可以通过实现 Error() 方法来自定义 error 类型。 这里使用自定义错误类型来表示上面例子中的参数错误
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
```

### 22. 协程

协程(goroutine) 是轻量级的执行线程。

```go
package main

import (
    "fmt"
    "time"
)


func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    // 假设我们有一个函数叫做 f(s)。 我们一般会这样 同步地 调用它
    f("direct")

    // 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
    go f("goroutine")

    // 你也可以为匿名函数启动一个协程。
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // 现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成（更好的方法是使用 WaitGroup）。
    time.Sleep(time.Second)
    fmt.Println("done")
}

```

direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

### 23. 通道


通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

```go
package main

import "fmt"

func main() {

    // 使用 make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。
    messages := make(chan string)

    // 使用 channel <- 语法 发送 一个新的值到通道中。 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
    go func() { messages <- "ping" }()

    // 使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
    msg := <-messages
    fmt.Println(msg)
    // 我们运行程序时，通过通道， 成功的将消息 "ping" 从一个协程传送到了另一个协程中。
}
```

$ go run channels.go
ping

默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。

### 24. 通道缓冲


默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。 有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。

```go
package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

### 25. 通道同步

可以使用通道来同步协程之间的执行状态。 这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。 如果需要等待多个协程，WaitGroup 是一个更好的选择。

```go


package main
import (
    "fmt"
    "time"
)
//  我们将要在协程中运行这个函数。 done 通道将被用于通知其他协程这个函数已经完成工作。

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    
    // 发送一个值来通知我们已经完工啦。
    done <- true
}

func main() {
    // 运行一个 worker 协程，并给予用于通知的通道。
    done := make(chan bool, 1)
    go worker(done)
    // 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
    <-done
}

```

### 26. 通道方向

当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。

```go
package main
import "fmt"

// ping 函数定义了一个只能发送数据的（只写）通道。 尝试从这个通道接收数据会是一个编译时错误。

func ping(pings chan<- string, msg string) {
    pings <- msg
}

// pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
$ go run channel-directions.go
```

### 27. 通过选择器

Go 的 选择器（select） 让你可以同时等待多个通道操作。 将协程、通道和选择器结合，是 Go 的一个强大特性。


```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 在这个例子中，我们将从两个通道中选择。
    c1 := make(chan string)
    c2 := make(chan string)

    // 各个通道将在一定时间后接收一个值， 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    
    // 我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

```

// 跟预期的一样，我们首先接收到值 "one"，然后是 "two"。

$ time go run select.go 
received one
received two
注意，程序总共仅运行了两秒左右。因为 1 秒 和 2 秒的 Sleeps 是并发执行的，

real    0m2.245s

### 28. 超时处理

超时 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的。 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。

```golang
package main
import ( "fmt" "time")
func main() {

    // 在这个例子中，假如我们执行一个外部调用， 并在 2 秒后使用通道 c1 返回它的执行结果。
    c1 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()
    
    // 这里是使用 select 实现一个超时操作。 res := <- c1 等待结果，<-Time.After 等待超时（1秒钟）以后发送的值。 由于 select 默认处理第一个已准备好的接收操作， 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。

    select {
        case res := <-c1:
            fmt.Println(res)
        case <-time.After(1 * time.Second):
            fmt.Println("timeout 1")
    }
    // 如果我们允许一个长一点的超时时间：3 秒， 就可以成功的从 c2 接收到值，并且打印出结果。

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
        case res := <-c2:
            fmt.Println(res)
        case <-time.After(3 * time.Second):
            fmt.Println("timeout 2")
    }
}
```
运行这个程序，首先显示运行超时的操作，然后是成功接收的。

$ go run timeouts.go 
timeout 1
result 2


### 29. 非阻塞通道操作

常规的通过通道发送和接收数据是阻塞的。 然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，甚至是非阻塞的多路 select。

```golang
package main
import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)
    
    // 这是一个非阻塞接收的例子。 如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 否则，就直接到 default 分支中。

    select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        default:
            fmt.Println("no message received")
    }
    
    // 一个非阻塞发送的例子，代码结构和上面接收的类似。 msg 不能被发送到 message 通道，因为这是 个无缓冲区通道，并且也没有接收者，因此， default 会执行。

    msg := "hi"
    select {
        case messages <- msg:
            fmt.Println("sent message", msg)
        default:
            fmt.Println("no message sent")
    }
    // 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。 这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。

    select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        case sig := <-signals:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity")
    }
}
```
$ go run non-blocking-channel-operations.go 
no message received
no message sent
no activity


### 30. 通道的关闭

关闭 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。

```golang

package main
import "fmt"

// 在这个例子中，我们将使用一个 jobs 通道，将工作内容， 从 main() 协程传递到一个工作协程中。 当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)
    
    // 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。 当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
    // 使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs。

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
    // 使用前面学到的通道同步方法等待任务结束。

    <-done
}
```
$ go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
根据 关闭通道 的思想，可以引出我们的下一个示例：遍历通道。


### 31. 通道遍历

在前面的例子中， 我们讲过 for 和 range 为基本的数据结构提供了迭代的功能。 我们也可以使用这个语法来遍历的从通道中取值。

```go
package main
import "fmt"
func main() {
    // 我们将遍历在 queue 通道中的两个值。
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    
    // range 迭代从 queue 中得到每个值。 因为我们在前面 close 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。

    for elem := range queue {
        fmt.Println(elem)
    }
}
```

这个例子也让我们看到，一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。

$ go run range-over-channels.go
one
two


### 32. Timer

我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。 Go 内置的 定时器 和 打点器 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。

```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。
    timer1 := time.NewTimer(2 * time.Second)
    // <-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。

    <-timer1.C
    fmt.Println("Timer 1 fired")

    // 如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。 使用定时器的原因之一就是，你可以在定时器触发之前将其取消。 例如这样。
    timer2 := time.NewTimer(time.Second)

    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()

    stop2 := timer2.Stop()

    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // 给 timer2 足够的时间来触发它，以证明它实际上已经停止了。
    time.Sleep(2 * time.Second)
}
```
第一个定时器将在程序开始后大约 2s 触发， 但是第二个定时器还未触发就停止了。

$ go run timers.go
Timer 1 fired
Timer 2 stopped


### 33. Ticker

定时器 是当你想要在未来某一刻执行一次时使用的 - 打点器 则是为你想要以固定的时间间隔重复执行而准备的。 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

```go

package main
import (
    "fmt"
    "time"
)
func main() {
    // 打点器和定时器的机制有点相似：使用一个通道来发送数据。 这里我们使用通道内建的 select，等待每 500ms 到达一次的值。

    ticker := time.NewTicker(500 * time.Millisecond) // ticker 等同于 setinterval, 每500ms 触发一次  
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
    
    // 打点器可以和定时器一样被停止。 打点器一旦停止，将不能再从它的通道中接收到值。 我们将在运行 1600ms 后停止这个打点器。

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

```
当我们运行这个程序时，打点器会在我们停止它前打点 3 次。

$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped


### 参考资料

- [Go by Example](https://gobyexample.com/)
- [Go by Example 中文版](https://gobyexample-cn.github.io/)
- [error-handling-and-go](https://blog.golang.org/error-handling-and-go)

