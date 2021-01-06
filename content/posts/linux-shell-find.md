---
title: "Linux Shell Find"
date: 2021-01-04T15:45:29+08:00
draft: false
---

```bash
# find ./ -exec rm {} \; 
# find ./ | xargs rm -rf
```

两者都可以把find命令查找到的结果删除，其区别简单的说是前者是把find发现的结果一次性传给exec选项，这样当文件数量较多的时候，就可能会出现“参数太多”之类的错误，相比较而言，后者就可以避免这个错误，因为xargs命令会分批次的处理结果。这样看来，“find ./| xargs rm -rf”是更通用的方法，推荐使用！

rm不接受标准输入，所以不能用find / -name "*.txt" ｜rm

-exec 必须由一个 ; 结束，而因为通常 shell 都会对 ; 进行处理，所以用 \; 防止这种情况。 
{} 可能需要写做 '{}'，也是为了避免被 shell 过滤

```shell
find ./ -type f -exec grep txt {} /dev/null \; 
```

`./` 表示从当前目录找 
-type f，表示只找file，文件类型的，目录和其他字节啥的不要 
-exec 把find到的文件名作为参数传递给后面的命令行，代替{}的部分 
-exec 后便跟的命令行，必须用“ \;”结束

```bash
# find ./ -type f -name "*.txt"|xargs grep "test" -n
# find . -name "*.txt" -exec grep "test" {} \; -print
```