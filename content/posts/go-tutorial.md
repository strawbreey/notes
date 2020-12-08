---
title: "Go Tutorial"
date: 2020-12-07T23:51:46+08:00
draft: false
---

### 标识符

25 个关键字

```
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
```

36个预定义标识符

```
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr
```

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

程序中可能会使用到这些标点符号：.、,、;、: 和 …。

### 字符串连接

```go
 fmt.Println("Google" + "Runoob")
```

### 数据类型

1.	布尔型
布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
2.	数字类型
整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。

数字类型: 
    uint8 无符号 8 位整型 (0 到 255)
    uint16 无符号 16 位整型 (0 到 65535)
    uint32 无符号 32 位整型 (0 到 4294967295)
    uint64 无符号 64 位整型 (0 到 18446744073709551615)
    int8 有符号 8 位整型 (-128 到 127)
    int16 有符号 16 位整型 (-32768 到 32767)
    int32 有符号 32 位整型 (-2147483648 到 2147483647)
    int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

浮点型:
	float32 IEEE-754 32位浮点型数
	float64 IEEE-754 64位浮点型数
	complex64 32 位实数和虚数
	complex128 64 位实数和虚数

其他数字类型:

    byte 类似 uint8
    rune 类似 int32
    uint 32 或 64 位
    int 与 uint 一样大小
    uintptr 无符号整型，用于存放一个指针

3.	字符串类型:
字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
4.	派生类型:
    (a) 指针类型（Pointer）
    (b) 数组类型
    (c) 结构化类型(struct)
    (d) Channel 类型
    (e) 函数类型
    (f) 切片类型
    (g) 接口类型（interface）
    (h) Map 类型


### 变量

变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用 var 关键字：

```go
var identifier type

// 可以一次声明多个变量：
var identifier1, identifier2 type

// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
var v_name v_type

v_name = value

// 第二种, 根据值自行判定变量类型。
var v_name = value

var d = true
fmt.Println(d)

// 第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
v_name := value

var intVal int 

intVal :=1 // 这时候会产生编译错误

intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句


// 多变量声明
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)

```

数值类型（包括complex64/128）为 0

布尔类型为 false

字符串为 ""（空字符串）

以下几种类型为 nil：


### 常量

### todo