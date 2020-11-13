---
title: "Git Remote"
date: 2020-11-13T10:32:42+08:00
draft: true
---

```bash
# 更新远程分支
git remote update origin --prune
git remote update origin -p
```

### 参考

```bash
# git-remote - Manage set of tracked repositories
git remote [-v | --verbose]
git remote add [-t <branch>] [-m <master>] [-f] [--[no-]tags] [--mirror=(fetch|push)] <name> <url>
git remote rename <old> <new>
git remote remove <name>
git remote set-head <name> (-a | --auto | -d | --delete | <branch>)
git remote set-branches [--add] <name> <branch>…​
git remote get-url [--push] [--all] <name>
git remote set-url [--push] <name> <newurl> [<oldurl>]
git remote set-url --add [--push] <name> <newurl>
git remote set-url --delete [--push] <name> <url>
git remote [-v | --verbose] show [-n] <name>…​
git remote prune [-n | --dry-run] <name>…​
git remote [-v | --verbose] update [-p | --prune] [(<group> | <remote>)…​]
```

### 例子

```bash
# 列出已经存在的远程分支y
git remote
# => origin

# 列出详细信息，在每一个名字后面列出其远程url，
git remote -v
# => origin

git branch -r
# => origin/HEAD -> origin/master
# => origin/master

# 添加一个远程仓库
git remote add staging git://git.kernel.org/.../gregkh/staging.git

# 列出已经存在的远程分支y
git remote
# => origin
# => staging

git fetch staging


git switch -c staging staging/master
```