---
title: "Git Subtree"
date: 2021-11-16T15:59:34+08:00
draft: false
---

## Git 迁移子目录

```bash
# 这就是那个大仓库 big-project
$ git clone git@github.com:tom/big-project.git
$ cd big-project
 
# 把所有 `codes-eiyo` 目录下的相关提交整理为一个新的分支 eiyo
$ git subtree split -P codes-eiyo -b eiyo
 
# 另建一个新目录并初始化为 git 仓库
$ mkdir ../eiyo
$ cd ../eiyo
$ git init
 
# 拉取旧仓库的 eiyo 分支到当前的 master 分支
$ git pull ../big-project eiyo
 
# 清理无用日志
$ git gc --aggressive --prune=now
 
# 添加到远程仓库
$ git remote add origin {url}
 
# 提交
$ git push origin master
```