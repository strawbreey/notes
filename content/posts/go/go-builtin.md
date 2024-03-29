---
title: "Go Builtin"
date: 2020-10-12T15:32:11+08:00
draft: false
categories: ['GO']
tags: ['GO']
---

Package builtin provides documentation for Go's predeclared identifiers (预声明的标识符). The items documented here are not actually in package builtin but their descriptions here allow godoc to present documentation for the language's special identifiers (特殊标识符).

### Package files

[builtin.go](https://github.com/golang/go/blob/master/src/builtin/builtin.go)

### Constants

true and false are the two untyped boolean values.

```go
const (
    true  = 0 == 0 // Untyped bool.
    false = 0 != 0 // Untyped bool.
)
```

iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.


### Variables

nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.

```go
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
```

### func append

The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:

```go
func append(slice []Type, elems ...Type) []Type

slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)

// As a special case, it is legal to append a string to a byte slice, like this:
slice = append([]byte("hello "), "world"...)
```

### func cap

func cap (v Type) init

The cap built-in function returns the capacity of v, according to its type:

```go
Array: the number of elements in v (same as len(v)).
Pointer to array: the number of elements in *v (same as len(v)).
Slice: the maximum length the slice can reach when resliced;
if v is nil, cap(v) is zero.
Channel: the channel buffer capacity, in units of elements;
if v is nil, cap(v) is zero.
```

### func close 

func close(c chan<- Type)

The close built-in function closes a channel, which must be either bidirectional(双向的) or send-only(仅发送). It should be executed only by the sender, never the receiver, and has the effect of shutting down the channel after the last sent value is received. After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element. The form

```go
x, ok := <-c
```

will also set ok to false for a closed channel.


### func complex

func complex(r, i FloatType) ComplexType

The complex built-in function constructs a complex value from two floating-point values. The real and imaginary parts must be of the same size, either float32 or float64 (or assignable to them), and the return value will be the corresponding complex type (complex64 for float32, complex128 for float64).

### func copy

func copy(dst, src []Type) int

The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).

### func delete

func delete(m map[Type]Type1, key Type)

The delete built-in function deletes the element with the specified(指定的) key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.

### func imga

func imag(c ComplexType) FloatTypes

The imag built-in function returns the imaginary part of the complex number c. The return value will be floating point type corresponding to the type of c.

### func len
func len(v Type) int
The len built-in function returns the length of v, according to its type:

```go
Array: the number of elements in v.
Pointer to array: the number of elements in *v (even if v is nil).
Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
String: the number of bytes in v.
Channel: the number of elements queued (unread) in the channel buffer;
         if v is nil, len(v) is zero.
```

### func make 

func make(t Type, size ...IntegerType) Type

The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:

```go
Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length. For example, make([]int, 0, 10) allocates an underlying array
of size 10 and returns a slice of length 0 and capacity 10 that is
backed by this underlying array.
Map: An empty map is allocated with enough space to hold the
specified number of elements. The size may be omitted, in which case
a small starting size is allocated.
Channel: The channel is buffer is initialized with the specified
buffer capacity. If zero, or the size is omitted, the channel is
unbuffered.
```

### func new

func new(Type) *Type
The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.


### func panic 

func panic(v interface{})
The panic(恐慌) built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated with a non-zero exit code. This termination sequence is called panicking and can be controlled by the built-in function recover.

### func print

func print(args ...Type)

The print built-in function formats its arguments in an implementation-specific (具体实施) way and writes the result to standard error. Print is useful for bootstrapping (自举) and debugging( 调试); it is not guaranteed (保证的) to stay in the language.


### func println

func println(args ...Type)

The println built-in function formats its arguments in an implementation-specific way and writes the result to standard error. Spaces are always added between arguments and a newline is appended. Println is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.

### func real

  func real(c ComplexType) FloatType

The real built-in function returns the real part of the complex number c. The return value will be floating point type corresponding to the type of c.

### func recover
  
  func recover() interface{}

The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether the goroutine is panicking.


type ComplexType
ComplexType is here for the purposes of documentation only. It is a stand-in for either complex type: complex64 or complex128.

type ComplexType complex64
type FloatType
FloatType is here for the purposes of documentation only. It is a stand-in for either float type: float32 or float64.

type FloatType float32
type IntegerType
IntegerType is here for the purposes of documentation only. It is a stand-in for any integer type: int, uint, int8 etc.

type IntegerType int
type Type
Type is here for the purposes of documentation only. It is a stand-in for any Go type, but represents the same type for any given function invocation.

type Type int
type Type1
Type1 is here for the purposes of documentation only. It is a stand-in for any Go type, but represents the same type for any given function invocation.

type Type1 int
type bool
bool is the set of boolean values, true and false.

type bool bool
type byte
byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.

type byte = uint8
type complex128
complex128 is the set of all complex numbers with float64 real and imaginary parts.

type complex128 complex128
type complex64
complex64 is the set of all complex numbers with float32 real and imaginary parts.

type complex64 complex64
type error
The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error.

type error interface {
    Error() string
}
type float32
float32 is the set of all IEEE-754 32-bit floating-point numbers.

type float32 float32
type float64
float64 is the set of all IEEE-754 64-bit floating-point numbers.

type float64 float64
type int
int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.

type int int
type int16
int16 is the set of all signed 16-bit integers. Range: -32768 through 32767.

type int16 int16
type int32
int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.

type int32 int32
type int64
int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807.

type int64 int64
type int8
int8 is the set of all signed 8-bit integers. Range: -128 through 127.

type int8 int8
type rune
rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.

type rune = int32
type string
string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable.

type string string
type uint
uint is an unsigned integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, uint32.

type uint uint
type uint16
uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.

type uint16 uint16
type uint32
uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.

type uint32 uint32
type uint64
uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615.

type uint64 uint64
type uint8
uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.

type uint8 uint8
type uintptr
uintptr is an integer type that is large enough to hold the bit pattern of any pointer.

type uintptr uintptr