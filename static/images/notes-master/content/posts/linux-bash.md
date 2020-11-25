---
title: "Linux Bash"
date: 2020-10-19T23:08:26+08:00
draft: false
---


### bash

管理整个计算机硬件的其实是操作系统的核心 (kernel)，这个核心是需要被保护的！ 所以我们一般使用者就只能透过 shell 来跟核心沟通，以让核心达到我们所想要达到的工作。 那么系统有多少 shell 可用呢？为什么我们要使用 bash 啊？底下分别来谈一谈喔！


### bash 和 shell
shell是运行在终端中的文本互动程序，bash（GNU Bourne-Again Shell）是最常用的一种shell。是当前大多数Linux发行版的默认Shell。

Shell相当于是一个翻译，把我们在计算机上的操作或我们的命令，翻译为计算机可识别的二进制命令，传递给内核，以便调用计算机硬件执行相关的操作；同时，计算机执行完命令后，再通过Shell翻译成自然语言，呈现在我们面前。

其他的shell还有：sh、bash、ksh、rsh、csh等。Ubuntu系统常用的是bash，Bio-linux系统是基于ubuntu定制的，但是却使用了zsh。

查看当前系统中shell的类型？

```shell
echo $SHELL
```

### shell命令

shell命令可以分为以下三类：

- 内建函数(built-in function)：shell自带的功能
- 可执行文件(executable file)：保存在shell之外的脚本，提供了额外的功能。
- 别名(alias)：给某个命令的简称

### 

- [认识与学习 BASH](http://cn.linux.vbird.org/linux_basic/0320bash_1.php)
