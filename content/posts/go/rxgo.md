---
title: "Rxgo"
date: 2021-08-11T13:29:03+08:00
draft: false
---

## Reactive Extensions for the Go Language

ReactiveX
ReactiveX, or Rx for short, is an API for programming with Observable streams. This is the official ReactiveX API for the Go language.

ReactiveX is a new, alternative way of asynchronous programming to callbacks, promises, and deferred. It is about processing streams of events or items, with events being any occurrences or changes within the system. A stream of events is called an Observable.

An operator is a function that defines an Observable, how and when it should emit data. The list of operators covered is available here.

### RxGo
The RxGo implementation is based on the concept of pipelines. A pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function.


创建目录并初始化：

```bash
$ mkdir rxgo && cd rxgo
$ go mod init github.com/darjun/go-daily-lib/rxgo
```

安装rxgo库：

```bash
$ go get -u github.com/reactivex/rxgo/v2
```

```go

package main

import (
  "fmt"

  "github.com/reactivex/rxgo/v2"
)

func main() {
  observable := rxgo.Just(1, 2, 3, 4, 5)()
  ch := observable.Observe()
  for item := range ch {
    fmt.Println(item.V)
  }
}
```


### 参考资料

- https://darjun.github.io/2020/10/11/godailylib/rxgo/
- https://play.golang.org/