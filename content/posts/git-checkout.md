---
title: "Git Checkout"
date: 2020-11-20T15:52:25+08:00
draft: false
tags: ['git']
---

checkout是Git最常用的命令之一

checkout在CVS和SVN中都是检出的意思，从版本库检出一个版本，在Git中就不是这么简单了。手册上是这样介绍的： git-checkout - Switch branches or restore working tree files

语法

```bash
git checkout [-q] [-f] [-m] [<branch>]
git checkout [-q] [-f] [-m] --detach [<branch>]
git checkout [-q] [-f] [-m] [--detach] <commit>
git checkout [-q] [-f] [-m] [[-b|-B|--orphan] <new_branch>] [<start_point>]
git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>] [--] <pathspec>…​
git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>] --pathspec-from-file=<file> [--pathspec-file-nul]
git checkout (-p|--patch) [<tree-ish>] [--] [<pathspec>…​]
```

描述



在Git里面，checkout用于切换分支或者恢复工作树的文件。

1. 新建一个hotfix分支

```bash
git checkout -b | -B <new _branch> [<start point>]

git checkout -b hotfix-1.2.1 master
# 这个时候分支是本地分支，并没有提交到服务器上去，如果这个分支已经被创建，这个命令会失败，这个时候，如果想要重置这个分支，需要使用-B参数。
```

2. 查看分支

```bash
git branch -av
```

本地发生了一些修改，但是想放弃这些修改

```bash
git checkout [<tree-ish>] [--] <pathspec>

git checkout 26a2e80 # 26a2e80 是一个commit号，这个命令会把index区域和工作区域的内容都更新
git checkout -- README # README是想恢复的文件名，恢复成index区域里面的内容，为什么要加“--”呢，这个是为了告诉Git，这是一个文件而不是一个分支
Git checkout . # 从index区域恢复所有文件
```

这个命令很灵活，既可以带一个commit号，又可以带着一个路径，tree-ish 可以理解成一个commit号，就是恢复到某一个commit号，index就是暂存区，这里要理解Git的三个区域，如果这个还不明白，那需要单开一篇文章去讲了。


### 参考文献

- [git-checkout](https://git-scm.com/docs/git-checkout)
- [Git checkout with dot](https://stackoverflow.com/questions/14460595/git-checkout-with-dot)