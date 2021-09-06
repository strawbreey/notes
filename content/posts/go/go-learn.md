---
title: "Go Learn"
date: 2021-09-01T17:29:43+08:00
draft: false
---

### 声明变量

```go
// 声明常量
const s string = "constant"

// 声明变量
var s string = "string" // 声明 字符串(string)
var b, c int = 1, 2     // 声明 数量 number
var d = true            // 声明 boolean 类型
var d int               // 声明变量，默认零值
f := "string"           // 简写声明, 类型自动推导

// 声明数组
var a [5]int            // 声明数组, 默认零值
b := [5]int{1, 2, 3, 4, 5} // 简写声明数组

// 声明切片
s := make([]string, 3)  // 声明切片
s = append(s, "d")      // 切片加值
c := make([]string, len(s)) 
copy(c, s)              // 把 s 复制给 c
l := s[2:5]             // 通过 slice[low:high] 语法进行“切片”操作

str := []string{'a', 'b', 'c'} // 切片简短声明
sort.Strings(strs)      // 字符串排序 [a, b, c]f

ints := []int{3, 2, 1}
sort.Ints(ints)         // 数字排序 [1, 2, 3]

// 自定义排序
type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

fruits := []string{"peach", "banana", "kiwi"}
sort.Sort(ByLength(fruits))

// 声明map
m := make(map[string]int) // 声明map
m["k1"] = 7               // 赋值
len(m)                    // 获取 m 的 键值对数目
delete(m, "k2")           // 删除m 的 k2 键值对
_, prs := m["k2"]         // prs 判断 k2 是否在 m 中
n := map[string]int{"foo": 1, "bar": 2} // 简短声明

s := "截取中文"
res := []rune(s)
string(res[:2])

i := '123'
s := "hello world"
z := "你好中国"


fmt.Println(s[:2])                  // he
fmt.Println(string([]rune(z)[:2]))  // 你好

// 引用和指针
i := 1
&i                          //获取内存地址

address := &i
*address = 1               // 指针 修改内存地址所对应的值

// 通道
messages := make(chan string) // 创建一个通道类型

```

### 命名规则

Go有一套简单的规则，适用于变量、函数和类型的名称：

- 名称必须以字母开头，并且可以有任意数量的额外的字母和数字。
  
- 如果变量、函数或类型的名称以大写字母开头，则认为它是导出的，可以从当前包之外的包访问它。（这就是为什么fmt.Println中的P是大写的：这样它就可以在main包或任何其他包中使用。）如果变量/函数/类型的名称是以小写字母开头的，则认为该名称是未导出的，只能在当前包中使用。
  
- Go有一套简单的规则，适用于变量、函数和类型的名称：·名称必须以字母开头，并且可以有任意数量的额外的字母和数字。·如果变量、函数或类型的名称以大写字母开头，则认为它是导出的，可以从当前包之外的包访问它。（这就是为什么fmt.Println中的P是大写的：这样它就可以在main包或任何其他包中使用。）如果变量/函数/类型的名称是以小写字母开头的，则认为该名称是未导出的，只能在当前包中使用。
  
- Go有一套简单的规则，适用于变量、函数和类型的名称：·名称必须以字母开头，并且可以有任意数量的额外的字母和数字。·如果变量、函数或类型的名称以大写字母开头，则认为它是导出的，可以从当前包之外的包访问它。（这就是为什么fmt.Println中的P是大写的：这样它就可以在main包或任何其他包中使用。）如果变量/函数/类型的名称是以小写字母开头的，则认为该名称是未导出的，只能在当前包中使用。

### 字符串
```go
var s string = "constant"  // 声明常量

// 字符串常用函数集合
s.Contains("test", "es")
s.Count("test", "t")
s.HasPrefix("test", "te")
s.HasSuffix("test", "st")
s.Index("test", "e")
s.Join([]string{"a", "b"}, "-")
s.Repeat("a", 5)
s.Replace("foo", "o", "0", -1)
s.Replace("foo", "o", "0", 1)
s.Split("a-b-c-d-e", "-")
s.ToLower("TEST")
s.ToUpper("test")
len("hello")
"hello"[1]

// 转义字符
\a             匹配响铃符    （相当于 \x07）
                注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
                可以使用 \x08 表示退格符。
\f             匹配换页符    （相当于 \x0C）
\t             匹配横向制表符（相当于 \x09）
\n             匹配换行符    （相当于 \x0A）
\r             匹配回车符    （相当于 \x0D）
\v             匹配纵向制表符（相当于 \x0B）
\123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
\x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
\x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
\Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

\\             匹配字符 \
\^             匹配字符 ^
\$             匹配字符 $
\.             匹配字符 .
\*             匹配字符 *
\+             匹配字符 +
\?             匹配字符 ?
\{             匹配字符 {
\}             匹配字符 }
\(             匹配字符 (
\)             匹配字符 )
\[             匹配字符 [
\]             匹配字符 ]
\|             匹配字符 |


// 类型转换

// 整形到字符串：
var i int = 5 
var s string 
s = strconv.Itoa(i)
s = FormatInt(int64(i), 10)

// 字符串到整形
var s string = "5"  
var i int  
i, err = strconv.Atoi(s)
i, err = ParseInt(s, 10, 0)

// 字符串到float(32 / 64)
var s string = 5  
var f float32  
f, err = ParseFloat(s, 32)

// 格式化	
p := point{1, 2}
fmt.Printf("%v\n", p)

// 如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
fmt.Printf("%+v\n", p)

// %#v 根据 Go 语法输出值，即会产生该值的源码片段。
fmt.Printf("%#v\n", p)

// 需要打印值的类型，使用 %T。
fmt.Printf("%T\n", p)

// 格式化布尔值很简单。
fmt.Printf("%t\n", true)

// 格式化整型数有多种方式，使用 %d 进行标准的十进制格式化。
fmt.Printf("%d\n", 123)

// 这个输出二进制表示形式。
fmt.Printf("%b\n", 14)

// 输出给定整数的对应字符。
fmt.Printf("%c\n", 33)

// %x 提供了十六进制编码。
fmt.Printf("%x\n", 456)

// 同样的，也为浮点型提供了多种格式化选项。 使用 %f 进行最基本的十进制格式化。
fmt.Printf("%f\n", 78.9)

// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学记数法表示形式。
fmt.Printf("%e\n", 123400000.0)
fmt.Printf("%E\n", 123400000.0)

// 使用 %s 进行基本的字符串输出。
fmt.Printf("%s\n", "\"string\"")

// 像 Go 源代码中那样带有双引号的输出，使用 %q。
fmt.Printf("%q\n", "\"string\"")

// 和上面的整型数一样，%x 输出使用 base-16 编码的字符串， 每个字节使用 2 个字符表示。
fmt.Printf("%x\n", "hex this")

// 要输出一个指针的值，使用 %p。
fmt.Printf("%p\n", &p)

// 格式化数字时，您经常会希望控制输出结果的宽度和精度。 要指定整数的宽度，请在动词 “%” 之后使用数字。 默认情况下，结果会右对齐并用空格填充。
fmt.Printf("|%6d|%6d|\n", 12, 345)

// 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

// 要左对齐，使用 - 标志。
fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

// 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。 这是基本的宽度右对齐方法。
fmt.Printf("|%6s|%6s|\n", "foo", "b")

// 要左对齐，和数字一样，使用 - 标志。
fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

// 到目前为止，我们已经看过 Printf 了， 它通过 os.Stdout 输出格式化的字符串。 Sprintf 则格式化并返回一个字符串而没有任何输出。
s := fmt.Sprintf("a %s", "string")
fmt.Println(s)

// 你可以使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout。
fmt.Fprintf(os.Stderr, "an %s\n", "error")

```

### 数字

```go
// 判断类型溢出
if temp := int32(res); (temp*10)/10 != temp {
    return 0
}

const UINT_MIN uint = 0             // 无符号整型uint 最小值
const UINT_MAX uint = ^uint(0)      // 无符号最大值

const INT_MAX = int(^uint(0) >> 1)  // 有符号 最大值
const INT_MIN = ^INT_MAX            // 有符号 最小值
const INT_MAX = int32(^uint32(0) >> 1) // 最大值

^uint32(0)                          // 获取类型最大值

// math 常量
MaxInt    = 1<<(intSize-1) - 1
MinInt    = -1 << (intSize - 1)
MaxInt8   = 1<<7 - 1
MinInt8   = -1 << 7
MaxInt16  = 1<<15 - 1
MinInt16  = -1 << 15
MaxInt32  = 1<<31 - 1
MinInt32  = -1 << 31
MaxInt64  = 1<<63 - 1
MinInt64  = -1 << 63
MaxUint   = 1<<intSize - 1
MaxUint8  = 1<<8 - 1
MaxUint16 = 1<<16 - 1
MaxUint32 = 1<<32 - 1
MaxUint64 = 1<<64 - 1

// func
Abs
Acos
Acosh
Asin
Asinh
Atan
Atan2
Atanh
Cbrt
Ceil
Copysign
Cos
Cosh
Dim
Exp
Exp2
Expm1
Floor
Log
Log10
Log2
Mod
Modf
Pow
Pow10
Round
RoundToEven
Sin
Sincos
Sinh
Sqrt
Tan
Tanh
Trunc
```

### 切片

```go
// 声明一个切片, 一个切片在未初始化之前默认为 nil，长度为 0
var identifier []type
// 使用 make() 函数来创建切片:
var slice1 []type = make([]type, len)
// 简写
slice1 := make([]type, len)
// 指定容量，其中 capacity 为可选参数
make([]T, length, capacity)

s :=[] int {1,2,3 }

// 切片长度
len(s)

// 切片容量
cap(s)

// 切片截取
numbers := []int{0,1,2,3,4,5,6,7,8} 

// 打印子切片从索引1(包含) 到索引4(不包含)
numbers[1:4]

// 追加切片
numbers = append(numbers, 0)

// 将numbers 拷贝到 numbers1
copy(numbers1, numbers)

```

### 判断

```go

// if/else
if num == 100 {
    fmt.Println("is 100")
} else if num > 90 {
    fmt.Println(" > 90")
} else {
    fmt.Println(" < 90")
}

// switch/case
switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
}

```

### 循环

```go
// for 循环
for i: = 0; i < 9; i++ {
    fmt.Println(i)
} 

// range 遍历 数组
for _, num := range nums {
    sum += num
}

// range 遍历map
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
    fmt.Printf(k, v)
}

// range 字符串中迭代 unicode 编码
for i, c := range "go" {
    fmt.Printf(i, c)
}

```

### 函数 结构体 接口 方法
```go

// function
func plus (a int, b int) int {
    return a + b
}

// 多返回值
func plus (a int, b int) (int, int) {
    return a, b
}

// 变参函数
func sum (nums ...int) {
    for _, num := range nums {
        fmt.Println(num)
    }
}

sum(1, 2)
sum(1, 2)

// 递归
func fact (n int) int {
    if n == 0 {
        return 1
    }
    return n * fact (n - 1)
}

// 结构体
type person struct {
    name string
    age int
}

func (p *person) name() string {
    return p.name
}

func (p *person) age() string {
    return p.age
}

r := rect{width: 10, height: 5}

fmt.Println("area: ", r.area())
fmt.Println("perim:", r.perim())

// 接口 todo

```


### 处理错误

```go
// 返回错误
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
    if arg == 42 {
        // 返回错误信息
        return -1, errors.New("can't work with 42")
    }
    // 返回错误值为 nil 代表没有错误
    return arg + 3, nil
}

// panic 意味着有些出乎意料的错误发生
_, err := os.Create("/tmp/file")
if err != nil {
    panic(err)
}


```

### 协程通道

```go

// 协程
	
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

go f("goroutine")

// 匿名函数
go func(msg string) {
    fmt.Println(msg)
}("going")

// 通道
messages := make(chan string)

go func () {
    messages <- "ping"
}()

msg := <- messages

fmt.Println(msg)

// 缓存通道
messages := make(chan string, 2) // 设置通道缓冲值为2

// 通道选择器
c1 := make(chan string)
c2 := make(chan string)
go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
}()
go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
}()
for i := 0; i < 2; i++ {
    select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
    }
}

// 超时处理

// 关闭通道
jobs := make(chan int, 5)
close(jobs)

// 便利通道
queue := make (chan string, 2)
queue <- "one"
queue <- "two"

// 关闭通道
close(queue)

for elem := range queue {
    fmt.Println(elem)
}
```

### 包

``` go
// 声明包
package apis

// 引入包 单行引入
import "github.com/gin-gonic/gin"

// 多行引入
import (
    "github.com/gin-gonic/gin"   
    "github.com/gin-gonic/gin"
    // 使用别名
    mrand "math/rand"
    // 如果只希望导入包，而不使用任何包内的结构和类型，也不调用包内的任何函数时，可以使用匿名导入包
    _ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
)


```

### 时间 (time)

```go
// 定时器
// 获取当前时间
now := time.now()
// 遵循 RFC3339 进行格式化的基本例子
now.Format(time.RFC3339)

time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

// 分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
secs := now.Unix()
nanos := now.UnixNano()

// 将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
time.Unix(secs, 0)
time.Unix(0, nanos)

// 构建一个时间
then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

then.Year()
then.Month()
then.Day()
then.Hour()
then.Minute()
then.Second()
then.Nanosecond()
then.Location()

// Weekday 输出星期一到星期日
then.Weekday()

// 这些方法用来比较两个时间，分别测试一下是否为之前、之后或者是同一时刻，精确到秒。
then.Before(now)
then.After(now)
then.Equal(now)

// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
diff := now.Sub(then)

// 我们可以用各种单位来表示时间段的长度。
diff.Hours()
diff.Minutes()
diff.Seconds()
diff.Nanoseconds()

// 你可以使用 Add 将时间后移一个时间段，或者使用一个 - 来将时间前移一个时间段。
then.Add(diff)
then.Add(-diff)

// 2s 后发送
timer =:= timer.NewTimer(time.Second * 2)

// 直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞
<-timer.c

// 关闭定时器
stop2 := timer2.Stop()

// 没500ms 发送一次
ticker := time.NewTicker(time.Millisecond * 500)
go func() {
    for t := range ticker.C {
        fmt.Println("Tick at", t)
    }
}()


time.Sleep(time.Millisecond * 1600)

// 关闭打点器
ticker.Stop()
fmt.Println("Ticker stopped")

// defer 用来确保一个函数调用在程序执行结束前执行
f := createFile("/tmp/defer.txt")
defer closeFile(f)
writeFile(f)

```

### 组合函数

```go

// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}
// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true。
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true。

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true。

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成。

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到。

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

var strs = []string{"peach", "apple", "pear", "plum"}
Index(strs, "pear")
Include(strs, "grape")
Any(strs, func(v string) bool {
    return strings.HasPrefix(v, "p")
})
All(strs, func(v string) bool {
    return strings.HasPrefix(v, "p")
})
Filter(strs, func(v string) bool {
    return strings.Contains(v, "e")
})
Map(strs, strings.ToUpper)

```

### 正则表达式

```go

// 测试一个字符串是否符合一个正则表达式。
match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

// 通过 Compile 优化 Regexp 结构体。
r, _ := regexp.Compile("p([a-z]+)ch")

// 该结构体有很多方法。这是一个类似于我们前面看到的匹配测试。
r.MatchString("peach")

// 查找匹配的字符串。
r.FindString("peach punch")

// 这个也是查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容。
r.FindStringIndex("peach punch")

// Submatch 返回完全匹配和局部匹配的字符串。 例如，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
r.FindStringSubmatch("peach punch")

// 返回完全匹配和局部匹配位置的索引。
r.FindStringSubmatchIndex("peach punch")

// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。例如查找匹配表达式全部的项。
r.FindAllString("peach punch pinch", -1)

// All 同样可以对应到上面的所有函数。
r.FindAllStringSubmatchIndex("peach punch pinch", -1)

// 这些函数接收一个正整数作为第二个参数，来限制匹配次数。
r.FindAllString("peach punch pinch", 2)

// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。 我们也可以将 String 从函数命中去掉，并提供 []byte 的参数。
r.Match([]byte("peach"))

// 创建正则表达式的全局变量时，可以使用 Compile 的变体 MustCompile 。 MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全。
r = regexp.MustCompile("p([a-z]+)ch")

// regexp 包也可以用来将子字符串替换为为其它值。
r.ReplaceAllString("a peach", "<fruit>")

// Func 变体允许您使用给定的函数来转换匹配的文本。
in := []byte("a peach")
out := r.ReplaceAllFunc(in, bytes.ToUpper)
fmt.Println(string(out))
```

### JSON
```go
// JSON
bolB, _ := json.Marshal(true)
intB, _ := json.Marshal(1)
fltB, _ := json.Marshal(2.34)
strB, _ := json.Marshal("gopher")

// 这是一些切片和 map 编码成 JSON 数组和对象的例子。
slcD := []string{"apple", "peach", "pear"}
slcB, _ := json.Marshal(slcD)

mapD := map[string]int{"apple": 5, "lettuce": 7}
mapB, _ := json.Marshal(mapD)

// JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。

res1D := &response1{
    Page:   1,
    Fruits: []string{"apple", "peach", "pear"}
}
res1B, _ := json.Marshal(res1D)

// 你可以给结构字段声明标签来自定义编码的 JSON 数据的键名。 上面 Response2 的定义，就是这种标签的一个例子。

res2D := &response2{
    Page:   1,
    Fruits: []string{"apple", "peach", "pear"}
}
res2B, _ := json.Marshal(res2D)

// 现在来看看将 JSON 数据解码为 Go 值的过程。 这是一个普通数据结构的解码例子。
byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

// 我们需要提供一个 JSON 包可以存放解码数据的变量。 这里的 map[string]interface{} 是一个键为 string，值为任意值的 map。

var dat map[string]interface{}
// 这是实际的解码，以及相关错误的检查。

if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
}
fmt.Println(dat)

// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。 例如，这里我们将 num 的值转换成 float64 类型。
num := dat["num"].(float64)
fmt.Println(num)

// 访问嵌套的值需要一系列的转化。
strs := dat["strs"].([]interface{})
str1 := strs[0].(string)
fmt.Println(str1)

// 我们还可以将 JSON 解码为自定义数据类型。 这样做的好处是，可以为我们的程序增加附加的类型安全性， 并在访问解码后的数据时不需要类型断言。
str := `{"page": 1, "fruits": ["apple", "peach"]}`
res := response2{}
json.Unmarshal([]byte(str), &res)
fmt.Println(res)
fmt.Println(res.Fruits[0])

// 在上面例子的标准输出上， 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。 当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
enc := json.NewEncoder(os.Stdout)
d := map[string]int{"apple": 5, "lettuce": 7}
enc.Encode(d)

// 随机数, 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。 要产生不同的数字序列，需要给定一个不同的种子。 注意，对于想要加密的随机数，使用此方法并不安全， 应该使用 crypto/rand。
// rand.Intn 返回一个随机的整数 n，且 0 <= n < 100。
rand.Intn(100)

// rand.Float64 返回一个64位浮点数 f，且 0.0 <= f < 1.0。
rand.Float64()

// 生成随机数种子
s1 := rand.NewSource(time.Now().UnixNano())
r1 := rand.New(s1)

// 数字解析
f, _ := strconv.ParseFloat("1.234", 64)

// 在使用 ParseInt 解析整型数时， 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 64 表示返回的整型数是以 64 位存储的。
i, _ := strconv.ParseInt("123", 0, 64)

// ParseInt 会自动识别出字符串是十六进制数。
d, _ := strconv.ParseInt("0x1c8", 0, 64)

// ParseUint 也是可用的。
u, _ := strconv.ParseUint("789", 0, 64)

// Atoi 是一个基础的 10 进制整型数转换函数。
k, _ := strconv.Atoi("135")

// 在输入错误时，解析函数会返回一个错误。
_, e := strconv.Atoi("wat")

```


### url 解析

```go
s := "postgres://user:pass@host.com:5432/path?k=v#f"
u, err := url.Parse(s)
if err != nil {
    panic(err)
}

// 直接访问 scheme。
u.Scheme

// User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
u.User
u.User.Username()
p, _ := u.User.Password()
fmt.Println(p)

// Host 包含主机名和端口号（如果存在）。使用 SplitHostPort 提取它们。
u.Host
host, port, _ := net.SplitHostPort(u.Host)

// 这里我们提取路径和 # 后面的查询片段信息。
u.Path
u.Fragment

// 要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。 你也可以将查询参数解析为一个 map。已解析的查询参数 map 以查询字符串为键， 已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]。
u.RawQuery
m, _ := url.ParseQuery(u.RawQuery)

```

### SHA1 哈希
```go
s := "sha1 this string"

// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
h := sha1.New()

// 写入要处理的字节。如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组。
h.Write([]byte(s))

// Sum 得到最终的散列值的字符切片。Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
bs := h.Sum(nil) // cf23df2207d99a74fbe169e3eba035e633b65d94


```

### base64 

```go
data := "abc123!?$*&()'-=@~"

// Go 同时支持标准 base64 以及 URL 兼容 base64。 这是使用标准编码器进行编码的方法。 编码器需要一个 []byte，因此我们将 string 转换为该类型。
sEnc := b64.StdEncoding.EncodeToString([]byte(data))

// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，你就需要进行错误检查了。
sDec, _ := b64.StdEncoding.DecodeString(sEnc)

// 使用 URL base64 格式进行编解码。
uEnc := b64.URLEncoding.EncodeToString([]byte(data))
uDec, _ := b64.URLEncoding.DecodeString(uEnc)
```

### 文件

```go
// 1. 读文件
// 将文件内容读取到内存中。
dat, err := ioutil.ReadFile("/tmp/dat")

// Open 打开一个文件，以获取一个 os.File 值。
f, err := os.Open("/tmp/dat")

// 从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
b1 := make([]byte, 5)
n1, err := f.Read(b1)

// 你也可以 Seek 到一个文件中已知的位置，并从这个位置开始读取。
o2, err := f.Seek(6, 0)
b2 := make([]byte, 2)
n2, err := f.Read(b2)

o3, err := f.Seek(6, 0)
b3 := make([]byte, 2)
n3, err := io.ReadAtLeast(f, b3, 2)

// 没有内建的倒带，但是 Seek(0, 0) 实现了这一功能。
_, err = f.Seek(0, 0)

r4 := bufio.NewReader(f)
b4, err := r4.Peek(5)

f.Close()

// 2. 写文件
d1 := []byte("hello\ngo\n")
err := ioutil.WriteFile("/tmp/dat1", d1, 0644)

f, err := os.Create("/tmp/dat2")

// 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
defer f.Close()

// 您可以按期望的那样 Write 字节切片。

d2 := []byte{115, 111, 109, 101, 10}
n2, err := f.Write(d2)

// WriteString 也是可用的。
n3, err := f.WriteString("writes\n")

// 调用 Sync 将缓冲区的数据写入硬盘。
f.Sync()

// 与我们前面看到的带缓冲的 Reader 一样，bufio 还提供了的带缓冲的 Writer。
w := bufio.NewWriter(f)
n4, err := w.WriteString("buffered\n")

// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
w.Flush()
```

### HTTP 客户端

```go
resp, err := http.Get("http://gobyexample.com")
if err != nil {
    panic(err)
}
defer resp.Body.Close()

// 打印 response body 前面 5 行的内容。
scanner := bufio.NewScanner(resp.Body)
for i := 0; scanner.Scan() && i < 5; i++ {
    scanner.Text()
}
if err := scanner.Err(); err != nil {
    panic(err)
}
```

### 环境变量

```go
os.Setenv("FOO", "1")
os.Getenv("FOO")  // 1
os.Getenv("BAR")  // ''

// 使用 os.Environ 来列出所有环境变量键值对。 这个函数会返回一个 KEY=value 形式的字符串切片。 你可以使用 strings.SplitN 来得到键和值。这里我们打印所有的键。
for _, e := range os.Environ() {
    pair := strings.SplitN(e, "=", 2)
}
```

### 系统退出

```go
// 当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。
defer fmt.Println("!")

// 退出并且退出状态为 3。
os.Exit(3)
```