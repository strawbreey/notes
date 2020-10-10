---
title: "Linux Soft Link"
date: 2020-10-09T19:49:01+08:00
draft: true
---

#### 创建软链接

ln  -s  [源文件或目录]  [目标文件或目录]

例如：

```shell
# 当前路径创建test 引向/var/www/test 文件夹 
ln –s  /var/www/test  test

# 创建/var/test 引向/var/www/test 文件夹 
ln –s  /var/www/test   /var/test 
```


#### 删除软链接

和删除普通的文件是一眼的，删除都是使用rm来进行操作

 rm –rf 软链接名称（请注意不要在后面加”/”，rm –rf 后面加不加”/” 的区别，可自行去百度下啊）

例如：

```shell
# 删除test
rm –rf test

```

#### 修改软链接

ln –snf  [新的源文件或目录]  [目标文件或目录]

这将会修改原有的链接地址为新的地址

例如：

```shell
# 创建一个软链接

ln –s  /var/www/test   /var/test

# 修改指向的新路径

ln –snf  /var/www/test1   /var/test

```

常用的参数：

-b 删除，覆盖以前建立的链接

-d 允许超级用户制作目录的硬链接

-f 强制执行

-i 交互模式，文件存在则提示用户是否覆盖

-n 把符号链接视为一般目录

-s 软链接(符号链接)

-v 显示详细的处理过程