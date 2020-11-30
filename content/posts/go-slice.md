---
title: "Go Slice"
date: 2020-09-03T15:42:47+08:00
draft: false
---


Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。



### 定义切片

```go
// 切片不需要说明长度
var identifier []type

// 或使用make()函数来创建切片:
var slice1 []type = make([]type, len)

// 也可以简写为
slice1 := make([]type, len)
```


也可以指定容量，其中capacity为可选参数。
```go

make([]T, length, capacity)
```

### slice init 切片初始化
```go
s := [] int {1,2,3 } 
// 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s := arr[:]
// 初始化切片s,是数组arr的引用

s := arr[startIndex:endIndex] 
// 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

s1 := s[startIndex:endIndex] 
// 通过切片s初始化切片s1

s :=make([]int,len,cap) 
// 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
```

### len() 和 cap()

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少


### 空(nil)切片
```go
// 一个切片在未初始化之前默认为 nil，长度为 0，实例如下：
 var numbers []int
```

### 切片截取

可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：

```go

  x := []int{0,1,2,3,4,5,6,7,8}  
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
  // len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

  /* 打印原始切片 */
  fmt.Println("numbers ==", numbers)
  // numbers == [0 1 2 3 4 5 6 7 8]

  /* 打印子切片从索引1(包含) 到索引4(不包含)*/
  fmt.Println("numbers[1:4] ==", numbers[1:4])
  // numbers[1:4] == [1 2 3]

  /* 默认下限为 0*/
  fmt.Println("numbers[:3] ==", numbers[:3])
  // numbers[:3] == [0 1 2]

  /* 默认上限为 len(s)*/
  fmt.Println("numbers[4:] ==", numbers[4:])
  // numbers[4:] == [4 5 6 7 8]

  x1 := make([]int,0,5)
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x1),cap(x1),x1)
  // len=0 cap=5 slice=[]

  /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
  x2 := numbers[:2]
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x2),cap(x2),x2)
  // len=2 cap=9 slice=[0 1]

  /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
  x3 := numbers[2:5]
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x3),cap(x3),x3)
  // len=3 cap=7 slice=[2 3 4]
```

## append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来

```go
    var x []int
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
    // len=0 cap=0 slice=[]

    /* 允许空切片追加数据 */
    x = append(numbers, 0)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
    // len=1 cap=1 slice=[0]

    /* 向切片添加一个元素 */
    x = append(numbers, 1)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
    // len=2 cap=2 slice=[0 1]


    /* 同时添加多个元素 */
    x = append(numbers, 2,3,4)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
    // len=5 cap=6 slice=[0 1 2 3 4]

    /* 创建切片 numbers1 是之前切片的两倍容量*/
    x1 := make([]int, len(numbers), (cap(numbers))*2)

    /* 拷贝 numbers 的内容到 numbers1 */
    copy(x1,numbers)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x1),cap(x1),x1)

```