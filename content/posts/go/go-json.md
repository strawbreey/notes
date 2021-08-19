---
title: "Go Json"
date: 2021-08-18T20:01:32+08:00
draft: false
---

JSON（JavaScript 对象表示，JavaScript Object Notation）作为一种轻量级的数据交换格式1，在今天几乎占据了绝大多数的市场份额。虽然与更紧凑的数据交换格式相比，它的序列化和反序列化性能不足，但是 JSON 提供了良好的可读性与易用性，在不追求极致机制性能的情况下，使用 JSON 作为序列化格式是一种非常好的选择。

encoding/json 对外提供标准的 JSON 序列化和反序列化方法，即 encoding/json.Marshal 和 encoding/json.Unmarshal，它们也是包中最常用的两个方法。

序列化和反序列化的开销完全不同，JSON 反序列化的开销是序列化开销的好几倍，相信这背后的原因也非常好理解。Go 语言中的 JSON 序列化过程不需要被序列化的对象预先实现任何接口，它会通过反射获取结构体或者数组中的值并以树形的结构递归地进行编码，标准库也会根据 encoding/json.Unmarshal 中传入的值对 JSON 进行解码。

Go 语言 JSON 标准库编码和解码的过程大量地运用了反射这一特性，你会在本节的后半部分看到大量的反射代码，这一小节就不过多介绍了。我们在这里会简单介绍 JSON 标准库中的接口和标签，这是它为开发者提供的为数不多的影响编解码过程的接口。

接口 #
JSON 标准库中提供了 encoding/json.Marshaler 和 encoding/json.Unmarshaler 两个接口分别可以影响 JSON 的序列化和反序列化结果：
```go
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
```

在 JSON 序列化和反序列化的过程中，它会使用反射判断结构体类型是否实现了上述接口，如果实现了上述接口就会优先使用对应的方法进行编码和解码操作，除了这两个方法之外，Go 语言其实还提供了另外两个用于控制编解码结果的方法，即 encoding.TextMarshaler 和 encoding.TextUnmarshaler：
```go
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}

type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
```

一旦发现 JSON 相关的序列化方法没有被实现，上述两个方法会作为候选方法被 JSON 标准库调用并参与编解码的过程。总的来说，我们可以在任意类型上实现上述这四个方法自定义最终的结果，后面的两个方法的适用范围更广，但是不会被 JSON 标准库优先调用。

### 参考资料

- https://draveness.me/golang/docs/part4-advanced/ch09-stdlib/golang-json/