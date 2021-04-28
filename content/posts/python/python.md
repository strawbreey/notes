---
title: "Python"
date: 2021-04-27T17:29:21+08:00
draft: false
---

## Python 

### 简介

Python 是一种解释型语言： 这意味着开发过程中没有了编译这个环节。类似于PHP和Perl语言。

Python 是交互式语言： 这意味着，您可以在一个 Python 提示符 >>> 后直接执行代码。

Python 是面向对象语言: 这意味着Python支持面向对象的风格或代码封装在对象的编程技术。

Python 是初学者的语言：Python 对初级程序员而言，是一种伟大的语言，它支持广泛的应用程序开发，从简单的文字处理到 WWW 浏览器再到游戏。

### 环境搭建

1. 下载 Python官网：https://www.python.org/


2. Unix & Linux 平台安装 Python:

以下为在 Unix & Linux 平台上安装 Python 的简单步骤：

  - 打开 WEB 浏览器访问https://www.python.org/downloads/source/
  - 选择适用 于Unix/Linux 的源码压缩包。
  - 下载及解压压缩包。
  - 如果你需要自定义一些选项修改Modules/Setup
  - 执行 ./configure 脚本
  - make
  - make install

3. Window 平台安装 Python:
  - 打开 WEB 浏览器访问https://www.python.org/downloads/windows/
  - 在下载列表中选择Window平台安装包，包格式为：python-XYZ.msi 文件 ， XYZ 为你要安装的版本号。
  - 下载后，双击下载包，进入 Python 安装

4. MAC 平台安装 Python:
 - MAC 系统一般都自带有 Python2.x版本的环境, 也可以去官网下载最新板
 

### 运行Python

1. 交互式解释器: 命令行窗口进入 Python，并在交互式解释器中开始编写 Python

2. 命令行脚本： 

```shell
# Unix/Linux
python script.py 
```

[stackblitz](https://stackblitz.com/edit/angular-ewfqoc?file=src/app/app.component.ts)

3. 集成开发环境 PyCharm

### Python 部署

### 执行python程序


```python
#! /usr/bin/python
print('Hello world!')
```

关于脚本第一行的 #!/usr/bin/python 的解释，相信很多不熟悉 Linux 系统的同学需要普及这个知识，脚本语言的第一行，只对 Linux/Unix 用户适用，用来指定本脚本用什么解释器来执行。

有这句的，加上执行权限后，可以直接用 ./ 执行，不然会出错，因为找不到 python 解释器。

#!/usr/bin/python 是告诉操作系统执行这个脚本的时候，调用 /usr/bin 下的 python 解释器。

#!/usr/bin/env python 这种用法是为了防止操作系统用户没有将 python 装在默认的 /usr/bin 路径里。当系统看到这一行的时候，首先会到 env 设置里查找 python 的安装路径，再调用对应路径下的解释器程序完成操作。

#!/usr/bin/python 相当于写死了 python 路径。

#!/usr/bin/env python 会去环境设置寻找 python 目录，可以增强代码的可移植性，推荐这种写法。


分成两种情况：

（1）如果调用 python 脚本时，使用:

python script.py 

#!/usr/bin/python 被忽略，等同于注释

（2）如果调用python脚本时，使用:

./script.py 

#!/usr/bin/python 指定解释器的路径

PS：shell 脚本中在第一行也有类似的声明。