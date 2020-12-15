---
title: "Go Example"
date: 2020-12-14T12:18:37+08:00
draft: false
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

hello-world
然后我们可以直接运行这个二进制文件。

```bash
ls
hello-world	hello-world.go
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

### 10 Map

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

#todo

### 参考资料

- [Go by Example](https://gobyexample.com/)
- [Go by Example 中文版](https://gobyexample-cn.github.io/)

