---
title: "Go Package"
date: 2020-12-03T13:04:44+08:00
draft: false
---

Go语言有超过100个的标准包（译注：可以用 `go list std | wc -l`命令查看标准包的具体数目），标准库为大多数的程序提供了必要的基础构件。在Go的社区，有很多成熟的包被设计、共享、重用和改进，目前互联网上已经发布了非常多的Go语言开源包，它们可以通过 `https://pkg.go.dev/` 检索

任何包系统设计的目的都是为了简化大型程序的设计和维护工作，通过将一组相关的特性放进一个独立的单元以便于理解和更新，在每个单元更新的同时保持和程序中其它单元的相对独立性。这种模块化的特性允许每个包可以被其它的不同项目共享和重用，在项目范围内、甚至全球范围统一的分发和复用。

当我们修改了一个源文件，我们必须`重新编译`该源文件对应的包和所有依赖该包的其他包。即使是从头构建，Go语言编译器的编译速度也明显快于其它编译语言。Go语言的闪电般的编译速度主要得益于三个语言特性。

- 第一点，所有导入的包必须在每个文件的开头显式声明，这样的话编译器就没有必要读取和分析整个源文件来判断包的依赖关系。
- 第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，每个包可以被独立编译，而且很可能是被并发编译。
- 第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的依赖关系。因此，在编译一个包的时候，编译器只需要读取每个直接导入包的目标文件，而不需要遍历所有依赖的的文件


### 导入路径

```go
import (
    "fmt"
    "math/rand"
    "encoding/json"
    "golang.org/x/net/html"
    "github.com/go-sql-driver/mysql"
)

// 包声明
```

通常来说，默认的包名就是包导入路径名的最后一段，因此即使两个包的导入路径不同，它们依然可能有一个相同的包名。例如，math/rand包和crypto/rand包的包名都是rand。

关于默认包名一般采用导入路径名的最后一段的约定也有三种例外情况。

- 第一个例外，包对应一个可执行程序，也就是main包，这时候main包本身的导入路径是无关紧要的。名字为main的包是给go build（§10.7.3）构建命令一个信息，这个包编译完之后必须调用连接器生成一个可执行程序。

- 第二个例外，包所在的目录中可能有一些文件名是以_test.go为后缀的Go源文件（译注：前面必须有其它的字符，因为以_或.开头的源文件会被构建工具忽略），并且这些源文件声明的包名也是以_test为后缀名的。这种目录可以包含两种包：一种是普通包，另一种则是测试的外部扩展包。所有以_test为后缀包名的测试外部扩展包都由go test命令独立编译，普通包和测试的外部扩展包是相互独立的。测试的外部扩展包一般用来避免测试代码中的循环导入依赖，具体细节我们将在11.2.4节中介绍。

- 第三个例外，一些依赖版本号的管理工具会在导入路径后追加版本号信息，例如“gopkg.in/yaml.v2”。这种情况下包的名字并不包含版本号后缀，而是yaml。

### 导入声明
```go
import "fmt"
import "os"

// => 等价于
import (
    "fmt"
    "os"
)

import (
    "fmt"
    "html/template"
    "os"

    "golang.org/x/net/html"
    "golang.org/x/net/ipv4"
)

// 如果我们想同时导入两个有着名字相同的包，例如math/rand包和crypto/rand包，那么导入声明必须至少为一个同名包指定一个新的包名以避免冲突。这叫做导入包的重命名。

import (
    "crypto/rand"
    mrand "math/rand" // alternative name mrand avoids conflict
)
```

### 包的匿名导入

如果只是导入一个包而并不使用导入的包将会导致一个编译错误。但是有时候我们只是想利用导入包而产生的副作用：它会计算包级变量的初始化表达式和执行导入包的init初始化函数（§2.6.2）。这时候我们需要抑制“unused import”编译错误，我们可以用下划线_来重命名导入的包。像往常一样，下划线_为空白标识符，并不能被访问。

```go
import _ "image/png" // register PNG decoder
```

### 包和命名

当创建一个包，一般要用短小的包名，但也不能太短导致难以理解。标准库中最常用的包有bufio、bytes、flag、fmt、http、io、json、os、sort、sync和time等包。

尽可能让命名有描述性且无歧义。例如，类似imageutil或ioutilis的工具包命名已经足够简洁了，就无须再命名为util了。要尽量避免包名使用可能被经常用于局部变量的名字，这样可能导致用户重命名导入包，例如前面看到的path包。

### 工具

```bash
#  下载 golint 包, 使用命令go get可以下载一个单一的包或者用...下载整个子目录里面的每个包
go get github.com/golang/lint/golint
# $GOPATH/bin/golint gopl.io/ch2/popcount
# src/gopl.io/ch2/popcount/main.go:1:1:
#   package comment should be of the form "Package popcount ..."


# 构建包
# go build命令编译命令行参数指定的每个包
cd $GOPATH/src/gopl.io/ch1/helloworld
go build

go build gopl.io/ch1/helloworld

cd $GOPATH
go build ./src/gopl.io/ch1/helloworld

# 但不能
cd $GOPATH
go build src/gopl.io/ch1/helloworld

go install命令和go build命令很相似，但是它会保存每个包的编译成果，而不是将它们都丢弃。被编译的包会被保存到`$GOPATH/pkg`目录下，目录路径和 src目录路径对应，可执行程序被保存到`$GOPATH/bin`目录。（很多用户会将`$GOPATH/bin`添加到可执行程序的搜索列表中。）还有，`go install`命令和`go build`命令都不会重新编译没有发生变化的包，这可以使后续构建更快捷。为了方便编译依赖的包，`go build -i`命令将安装每个目标所依赖的包。


# 包文档
go doc time
go doc time.Since
go doc time.Duration.Seconds

# godoc
godoc -http :8000

# 查询包
go list github.com/go-sql-driver/mysql

# go list命令的参数还可以用"..."表示匹配任意的包的导入路径。我们可以用它来列出工作区中的所有包
go list ... 

# 特定子目录下的所有包
go list gopl.io/ch3/...

# 某个主题相关的所有包:
go list ...xml...

# go list命令还可以获取每个包完整的元信息，而不仅仅只是导入路径，这些元信息可以以不同格式提供给用户。其中-json命令行参数表示用JSON格式打印每个包的元信息。
go list -json hash

# 命令行参数-f则允许用户使用text/template包（§4.6）的模板语言定义输出文本的格式
go list -f '{{join .Deps " "}}' strconv

go list -f "{{join .Deps \" \"}}" strconv

go list -f "{{.ImportPath}} -> {{join .Imports \" \"}}" compress/...
```