---
title: "Go Nil"
date: 2020-12-07T13:21:39+08:00
draft: false
---

nil is a frequently (频繁的) used and important predeclared identifier (预定义标识符) in Go. It is the literal representation (代表) of zero values of many kinds of types. Many new Go programmers with experiences of some other popular languages may view nil as the counterpart (相对应) of null (or NULL) in other languages. This is partly right, but there are many differences between nil in Go and null (or NULL) in other languages.

The remaining (剩下) of this article will list all kinds of facts and details related to nil (nil 相关的事实和细节).

### nil Is a Predeclared Identifier in Go

We can use nil without declaring it.

### nil Can Represent Zero Values of Many Types

In Go, nil can represent zero values of the following kinds of types:

- pointer types (including type-unsafe ones).
- map types.
- slice types.
- function types.
- channel types.
- interface types.

### Predeclared nil Has Not a Default Type

Each of other (其他每个) predeclared identifiers in Go has a default type. For example,
- the default types of true and false are both bool type.
- the default type of iota is int.

But the predeclared nil has not a default type, though it has many possible types. In fact, the predeclared nil is the only untyped value who has not a default type in Go. There must be sufficient information for compiler to deduce the type of a nil from context.

Example:

```go
package main

func main() {
	// There must be sufficient information for
	// compiler to deduce the type of a nil value.
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// These lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// This following line doesn't compile.
	var _ = nil
}
```

### Predeclared nil Is Not a Keyword in Go

The predeclared nil can be shadowed.

Example:

```go
package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil) // 123

	// The following line fails to compile,
	// for nil represents an int value now
	// in this scope.
	var _ map[string]int = nil
}
```

(BTW, null and NULL in many other languages are also not keywords.)

### The Sizes of Nil Values With Types of Different Kinds May Be Different

The memory layouts of all values of a type are always the same. nil values of the type are not exceptions (assume the zero values of the type can be represented as nil). The size of a nil value is always the same as the sizes of non-nil values whose types are the same as the nil value. But nil values of different kinds of types may have different sizes.

// => todo see

Example:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *struct{} = nil
	fmt.Println( unsafe.Sizeof( p ) ) // 8

	var s []int = nil
	fmt.Println( unsafe.Sizeof( s ) ) // 24

	var m map[int]bool = nil
	fmt.Println( unsafe.Sizeof( m ) ) // 8

	var c chan string = nil
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	var f func() = nil
	fmt.Println( unsafe.Sizeof( f ) ) // 8

	var i interface{} = nil
	fmt.Println( unsafe.Sizeof( i ) ) // 16
}
```

The sizes are compiler and architecture dependent. The above printed results are for 64-bit architectures and the standard Go compiler. For 32-bit architectures, the printed sizes will be half.

For the standard Go compiler, the sizes of two values of different types of the same kind whose zero values can be represented as the predeclared nil are always the same. For example, the sizes of all values of all different slice types are the same.

### Two Nil Values of Two Different Types May Be Not Comparable

For example, the two comparisons in the following example both fail to compile. The reason is, in each comparison, neither operand can be implicitly converted to the type of the other.

```go
// Compilation failure reason: mismatched types.
var _ = (*int)(nil) == (*bool)(nil)         // error
var _ = (chan int)(nil) == (chan bool)(nil) // error
```

Please read comparison rules in Go to get which two values can be compared with each other. Typed nil values are not exceptions of the comparison rules.

The code lines in the following example all compile okay.
```go
type IntPtr *int
// The underlying of type IntPtr is *int.
var _ = IntPtr(nil) == (*int)(nil)

// Every type in Go implements interface{} type.
var _ = (interface{})(nil) == (*int)(nil)

// Values of a directional channel type can be
// converted to the bidirectional channel type
// which has the same element type.
var _ = (chan int)(nil) == (chan<- int)(nil)
var _ = (chan int)(nil) == (<-chan int)(nil)
```

### Two Nil Values of the Same Type May Be Not Comparable

In Go, map, slice and function types don't support comparison. Comparing two values, including nil values, of an incomparable types is illegal. The following comparisons fail to compile.

```go
var _ = ([]int)(nil) == ([]int)(nil)
var _ = (map[string]int)(nil) == (map[string]int)(nil)
var _ = (func())(nil) == (func())(nil)
```
But any values of the above mentioned incomparable types can be compared with the bare nil identifier.
// The following lines compile okay.
```go
var _ = ([]int)(nil) == nil
var _ = (map[string]int)(nil) == nil
var _ = (func())(nil) == nil
```

### Two Nil Values May Be Not Equal

If one of the two compared nil values is an interface value and the other is not, assume they are comparable, then the comparison result is always false. The reason is the non-interface value will be converted to the type of the interface value before making the comparison. The converted interface value has a concrete dynamic type but the other interface value has not. That is why the comparison result is always false.

Example:

```go
fmt.Println( (interface{})(nil) == (*int)(nil) ) // false
```
### Retrieving Elements From Nil Maps Will Not Panic

Retrieving element from a nil map value will always return a zero element value.

For example:
```go
fmt.Println( (map[string]int)(nil)["key"] ) // 0
fmt.Println( (map[int]bool)(nil)[123] )     // false
fmt.Println( (map[int]*int64)(nil)[123] )   // <nil>
```
### It Is Legal to Range Over Nil Channels, Maps, Slices, and Array Pointers

The number of loop steps by iterate nil maps and slices is zero.

The number of loop steps by iterate a nil array pointer is the length of its corresponding array type. (However, if the length of the corresponding array type is not zero, and the second iteration is neither ignored nor omitted, the iteration will panic at run time.)

Ranging over a nil channel will block the current goroutine for ever.

For example, the following code will print 0, 1, 2, 3 and 4, then block for ever. Hello, world and Bye will never be printed.

```go
for range []int(nil) {
	fmt.Println("Hello")
}

for range map[string]string(nil) {
	fmt.Println("world")
}

for i := range (*[5]int)(nil) {
	fmt.Println(i)
}

for range chan bool(nil) { // block here
	fmt.Println("Bye")
}
```

### Invoking Methods Through Non-Interface Nil Receiver Arguments Will Not Panic
Example:

```go
package main

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}

func (s Slice) Modify(i int, x bool) {
	s[i] = x // panic if s is nil
}

func (p *Slice) DoNothing() {
}

func (p *Slice) Append(x bool) {
	*p = append(*p, x) // panic if p is nil
}

func main() {
	// The following selectors will not cause panics.
	_ = ((Slice)(nil)).Length
	_ = ((Slice)(nil)).Modify
	_ = ((*Slice)(nil)).DoNothing
	_ = ((*Slice)(nil)).Append

	// The following two lines will also not panic.
	_ = ((Slice)(nil)).Length()
	((*Slice)(nil)).DoNothing()

	// The following two lines will panic. But panics
	// will not be triggered at the time of invoking
	// the methods. They will be triggered on
	// dereferencing nil pointers in the method bodies.
	/*
	((Slice)(nil)).Modify(0, true)
	((*Slice)(nil)).Append(true)
	*/
}
```

In fact, the above implementation of the Append method is not perfect, it should be modified as the following one.

```go
func (p *Slice) Append(x bool) {
	if p == nil {
		*p = []bool{x}
		return
	}
	*p = append(*p, x)
}
```

*new(T) Results a Nil T Value if the Zero Value of Type T Is Represented With the Predeclared nil Identifier
Example:

```go
package main

import "fmt"

func main() {
	fmt.Println(*new(*int) == nil)         // true
	fmt.Println(*new([]int) == nil)        // true
	fmt.Println(*new(map[int]bool) == nil) // true
	fmt.Println(*new(chan string) == nil)  // true
	fmt.Println(*new(func()) == nil)       // true
	fmt.Println(*new(interface{}) == nil)  // true
}
```
### Summary

In Go, for simplicity and convenience, nil is designed as an identifier which can be used to represent the zero values of some kinds of types. It is not a single value. It can represent many values with different memory layouts.


### 参考资料

- [nils in Go](https://go101.org/article/nil.html)