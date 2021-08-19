---
title: "Go Mod Init"
date: 2021-08-09T15:06:55+08:00
draft: false
---

### go Mod Help
```bash
$ go help mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache             # 下载依赖。参数是非必写的，path 是包的路径，version 是包的版本。
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory  # 初始化模块，会在项目根目录下生成 go.mod 文件。
        tidy        add missing and remove unused modules       # 根据 go.mod 文件来处理依赖关系。
        vendor      make vendored copy of dependencies          # 将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令 go build -mod=vendor 编译
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

### get start
```go
go mod init url
```