---
title: "Go Strings Api"
date: 2020-10-12T11:50:11+08:00
draft: false
---

Package strings implements simple functions to manipulate UTF-8 encoded strings.

func Compare

```go
func Compare (a, b string) int

// Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.


package main

import (
	"fmt"
	"strings"
)

func main() {
  strings.Compare("a", "b")  // -1
  strings.Compare("a", "a")  // 0
  strings.Compare("b", "a")  // 1
}
```

### func Containes

func Containes (s, substr string) bool

Contains reports whether substr is within s.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
  strings.Contains("seafood", "foo") // true
  strings.Contains("seafood", "bar") // false
  strings.Contains("seafood", "") // true
  strings.Contains("", "") // true
}
```

### func ContainsAny

func ContainsAny (s, chars string) bool

ntainsAny reports whether any Unicode code points in chars are within s.

```go
strings.ContainsAny("team", "i") //false
strings.ContainsAny("fail", "ui") // true
strings.ContainsAny("ure", "ui") // true
strings.ContainsAny("failure", "ui") //true
strings.ContainsAny("foo", "")  //false
strings.ContainsAny("", "") //false
```

### func ContainsRune  

Rune 符文

rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.

func ContainsRune(s string, r rune) bool

```go
strings.ContainsRune("aardvark", 97) //true
strings.ContainsRune("timeout", 97) //false
```