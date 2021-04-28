---
title: "Linux Shell"
date: 2020-10-19T23:12:33+08:00
draft: false
---

Shell 是一个用 C 语言编写的程序，它是用户使用 Linux 的桥梁。Shell 既是一种命令语言，又是一种程序设计语言。

Shell 是指一种应用程序，这个应用程序提供了一个界面，用户通过这个界面访问操作系统内核的服务。

Ken Thompson 的 sh 是第一种 Unix Shell，Windows Explorer 是一个典型的图形界面 Shell。

### Shell 脚本（shell script）
Shell 脚本（shell script），是一种为 shell 编写的脚本程序。

业界所说的 shell 通常都是指 shell 脚本，但读者朋友要知道，shell 和 shell script 是两个不同的概念。

由于习惯的原因，简洁起见，本文出现的 "shell编程" 都是指 shell 脚本编程，不是指开发 shell 自身。

###
```bash
#!/bin/bash #! 是一个约定的标记，它告诉系统这个脚本需要什么解释器来执行，即使用哪一种 Shell。
echo "Hello World !"
```

运行 Shell 脚本有两种方法：

1. 作为可执行程序
```bash
./test.sh  #执行脚本
```
注意，一定要写成 ./test.sh，而不是 test.sh，运行其它二进制的程序也一样，直接写 test.sh，linux 系统会去 PATH 里寻找有没有叫 test.sh 的，而只有 /bin, /sbin, /usr/bin，/usr/sbin 等在 PATH 里，你的当前目录通常不在 PATH 里，所以写成 test.sh 是会找不到命令的，要用 ./test.sh 告诉系统说，就在当前目录找。

2. 作为解释器参数

这种运行方式是，直接运行解释器，其参数就是 shell 脚本的文件名，如：
```bash
/bin/sh test.sh
/bin/php test.php
```
这种方式运行的脚本，不需要在第一行指定解释器信息，写了也没用。

### 变量

变量名和等号之间不能有空格，这可能和你熟悉的所有编程语言都不一样。同时，变量名的命名须遵循如下规则：

  - 命名只能使用英文字母，数字和下划线，首个字符不能以数字开头。
  - 中间不能有空格，可以使用下划线（_）。
  - 不能使用标点符号。
  - 不能使用bash里的关键字（可用help命令查看保留关键字）。
```shell
# 有效的 Shell 变量名
RUNOOB
LD_LIBRARY_PATH
_var
var2
```



```bash

```


```bash
#!/bin/bash

```


```bash
# 使用变量
your_name="qinjx"
echo $your_name
echo ${your_name}

# 只读变量
myUrl="https://www.google.com"
readonly myUrl
myUrl="https://www.runoob.com"

# 删除变量
unset variable_name

# 变量类型
运行shell时，会同时存在三种变量：

1. 局部变量 局部变量在脚本或命令中定义，仅在当前shell实例中有效，其他shell启动的程序不能访问局部变量。
2. 环境变量 所有的程序，包括shell启动的程序，都能访问环境变量，有些程序需要环境变量来保证其正常运行。必要的时候shell脚本也可以定义环境变量。
3. shell变量 shell变量是由shell程序设置的特殊变量。shell变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了shell的正常运行
```

### Shell 字符串

单引号
```bash
str='this is a string'
```
单引号字符串的限制：

- 单引号里的任何字符都会原样输出，单引号字符串中的变量是无效的；
- 单引号字串中不能出现单独一个的单引号（对单引号使用转义符后也不行），但可成对出现，作为字符串拼接使用。

双引号
```bash
  your_name='runoob'
  str="Hello, I know you are \"$your_name\"! \n"
  echo -e $str
```

双引号的优点：

- 双引号里可以有变量
- 双引号里可以出现转义字符

### Shell 数组

bash支持一维数组（不支持多维数组），并且没有限定数组的大小。