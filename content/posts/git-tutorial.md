---
title: "Git Tutorial"
date: 2020-09-24T19:29:58+08:00
draft: false
---

### 本地版本控制系统
许多人习惯用复制整个项目目录的方式来保存不同的版本，或许还会改名加上备份时间以示区别。 这么做唯一的好处就是简单，但是特别容易犯错。 有时候会混淆所在的工作目录，一不小心会写错文件或者覆盖意想外的文件。

其中最流行的一种叫做 RCS，现今许多计算机系统上都还看得到它的踪影。 RCS 的工作原理是在硬盘上保存补丁集（补丁是指文件修订前后的变化）；通过应用所有的补丁，可以重新计算出各个版本的文件内容。


### Git 配置

Git 自带一个 git config 的工具来帮助设置控制 Git 外观和行为的配置变量。 这些变量存储在三个不同的位置：

- /etc/gitconfig 文件: 包含系统上每一个用户及他们仓库的通用配置。 如果在执行 git config 时带上 --system 选项，那么它就会读写该文件中的配置变量。 （由于它是系统配置文件，因此你需要管理员或超级用户权限来修改它。）

- ~/.gitconfig 或 ~/.config/git/config 文件：只针对当前用户。 你可以传递 --global 选项让 Git 读写此文件，这会对你系统上 所有 的仓库生效。

- 当前使用仓库的 Git 目录中的 config 文件（即 .git/config）：针对该仓库。 你可以传递 --local 选项让 Git 强制读写此文件，虽然默认情况下用的就是它。。 （当然，你需要进入某个 Git 仓库中才能让该选项生效。）

每一个级别会覆盖上一级别的配置，所以 .git/config 的配置变量会覆盖 /etc/gitconfig 中的配置变量。

查看git所有配置以及他所在的文件

```shell
git config --list --show-origin
```

### 用户信息

```shell
# 设置全局的用户信息
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com

# 设置项目的用户信息
$ git config user.name "John Doe"
$ git config user.email johndoe@example.com

```

### 检查配置信息
```shell
# 检查git配置信息
git config --list

# git config <key>： 来检查 Git 的某一项配置
git config user.name

# 检查git配置信息和其配置地址
git config --list --show-origin
```

### 获取Git仓库

```shell
# 在已存在目录中初始化仓库
# 该命令将创建一个名为 .git 的子目录，这个子目录含有你初始化的 Git 仓库中所有的必须文件，这些文件是 Git 仓库的骨干。
git init

$ git add *.c
$ git add LICENSE
$ git commit -m 'initial project version'

# 克隆现有的仓库
git clone https://github.com/libgit2/libgit2

# 克隆仓库并重命名
git clone https://github.com/libgit2/libgit2 mylibgit
```


## git 状态变化周期

![lifecycle](/images/lifecycle.png)

```shell
# 检查当前文件状态
git status 

# 添加跟踪文件
git add README

# 简述
git status -s
git status -short

# 忽略文件
.gitignone

```

文件 .gitignore 的格式规范如下：

- 所有空行或者以 # 开头的行都会被 Git 忽略。

- 可以使用标准的 glob 模式匹配，它会递归地应用在整个工作区中。

- 匹配模式可以以（/）开头防止递归。

- 匹配模式可以以（/）结尾指定目录。

- 要忽略指定模式以外的文件或目录，可以在模式前加上叹号（!）取反。

```shell
# 忽略所有的 .a 文件
*.a

# 但跟踪所有的 lib.a，即便你在前面忽略了 .a 文件
!lib.a

# 只忽略当前目录下的 TODO 文件，而不忽略 subdir/TODO
/TODO

# 忽略任何目录下名为 build 的文件夹
build/

# 忽略 doc/notes.txt，但不忽略 doc/server/arch.txt
doc/*.txt

# 忽略 doc/ 目录及其所有子目录下的 .pdf 文件
doc/**/*.pdf
```

GitHub 有一个十分详细的针对数十种项目及语言的 .gitignore 文件列表， 你可以在 https://github.com/github/gitignore

```shell
# 要查看尚未暂存的文件更新了哪些部分
git diff
```

#### 查看提交历史
```shell
git log

# 查看最近2次提交补丁
git log -p -2

# 查看简略提示信息
git log --stat
# 
git log --pretty=oneline
```

#### 撤销操作

```shell
# 提交完了才发现漏掉了几个文件没有添加，或者提交信息写错了。 此时，可以运行带有 --amend 选项的提交命令来重新提交
git commit --amend

git status
git reset HEAD CONTRIBUTING.md

git checkout -- CONTRIBUTING.md
git status

# 撤销 HEAD^的意思是上一个版本 也可以写成HEAD~1
git reset --soft HEAD^ 
# 等同
git reset --mixed HEAD^

# 删除工作空间改动代码，撤销commit，撤销git add .
git reset --hard HEAD^
```

记住，在 Git 中任何 已提交 的东西几乎总是可以恢复的。 甚至那些被删除的分支中的提交或使用 --amend 选项覆盖的提交也可以恢复 （阅读 数据恢复 了解数据恢复）。 然而，任何你未提交的东西丢失后很可能再也找不到了。

### 远程仓库使用

```shell
git clone https://github.com/schacon/ticgit

cd ticgit

# 查看你已经配置的远程仓库服务器
git remote

# 需要读写远程仓库使用的 Git 保存的简写与其对应的 URL
git remote -v

# 添加远程仓库
git remote
git remote add pb https://github.com/paulboone/ticgit
git remote -v

# 现在你可以在命令行中使用字符串 pb 来代替整个 URL
git fetch pb

#  必须注意 git fetch 命令只会将数据下载到你的本地仓库——它并不会自动合并或修改你当前的工作。

# 推送到远程仓库, 当你想要将 master 分支推送到 origin 服务器时
git push origin master

# 只有当你有所克隆服务器的写入权限，并且之前没有人推送过时，这条命令才能生效。 当你和其他人在同一时间克隆，他们先推送到上游然后你再推送到上游，你的推送就会毫无疑问地被拒绝。 你必须先抓取他们的工作并将其合并进你的工作后才能推送

# 查看某个远程仓库, 可以使用 git remote show <remote> 命令

git remote show origin # 

# 远程仓库的重命名与移除
git remote rename pb paul # 修改一个远程仓库的简写名
git remote remove paul # 移除一个远程仓库
```

### 打标签

像其他版本控制系统（VCS）一样，Git 可以给仓库历史中的某一个提交打上标签，以示重要。 比较有代表性的是人们会使用这个功能来标记发布结点（ v1.0 、 v2.0 等等）

#### 列出标签
```shell
# 列出标签
git tag

# 按照特定的模式查找标签
git tag -l "v1.8.5*"
```

#### 创建标签

git 支持两种标签：轻量标签（lightweight）与附注标签（annotated）

轻量标签很像一个不会改变的分支——它只是某个特定提交的引用。

附注标签是存储在 Git 数据库中的一个完整对象， 它们是可以被校验的，其中包含打标签者的名字、电子邮件地址、日期时间， 此外还有一个标签信息，并且可以使用 GNU Privacy Guard （GPG）签名并验证。

```shell
# 轻量标签
git tag v1.4-lw

# 创建附注标签
git tag -a v1.4 -m "my version 1.4" # -m 选项指定了一条将会存储在标签中的信息。 如果没有为附注标签指定一条信息，Git 会启动编辑器要求你输入信息。

git show # 看到标签信息和与之对应的提交信息

# 对过去的提交打标签
git tag -a v1.2 9fceb02

# 共享标签
git push origin v1.5

# 一次性推送很多标签
git push origin --tags

# 删除标签
git tag -d <tagname>

# 删除远程标签 将冒号前面的空值推送到远程标签名，从而高效地删除它。
git push origin :refs/tags/v1.4-lw
# or
git push origin --delete <tagname>

# 从标签中迁出分支
git checkout -b version2 v2.0.0
```

### 别名

```shell
# 设置别名
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```

## Git 分支

```shell
# 追踪系统中的 #53 问题
git checkout -b iss53
# or
git branch iss53
git checkout iss53

# 备份当前的工作区的内容，从最近的一次提交中读取相关内容，让工作区保证和上次提交的内容一致。
git stash

git checkout master

# 切换一个紧急分支, 并修复bug
git checkout -b hotfix

# 切回master
git checkout master

# 合并hotfix分支
git merge hotfix

# 删除hotfix分支
git branch -d hotfix

# 返回iss53, 修复bug
git checkout iss53
git checkout master
git merge iss53
git branch -d iss53

# 查看所有分支
git branch -a

# 切换远程分支
# git checkout -b 本地分支名 origin/远程分支名
git checkout -b test origin/test


# 将本地分支推送到远程
# git push <远程主机名> <本地分支名>:<远程分支名>
git push -u origin dev:release/caigou_v1.0

```


#### 分支管理
```shell
# 查看分支
git branch

# 查看每一个分支的最后一次提交
git branch -v 

# --merged 与 --no-merged 这两个有用的选项可以过滤这个列表中已经合并或尚未合并到当前分支的分支。
git branch --merged
git branch --no-merged

# 当删除其他还未合并的分支时, 会失败
```

### 远程分支

> 远程仓库名字 “origin” 与分支名字 “master” 一样，在 Git 中并没有任何特别的含义一样。 同时 “master” 是当你运行 git init 时默认的起始分支名字，原因仅仅是它的广泛使用， “origin” 是当你运行 git clone 时默认的远程仓库名字。 如果你运行 git clone -o booyah，那么你默认的远程分支名字将会是 booyah/master。

远程引用是对远程仓库的引用（指针），包括分支、标签等等。 你可以通过 git ls-remote <remote> 来显式地获得远程引用的完整列表， 或者通过 git remote show <remote> 获得远程分支的更多信息。 然而，一个更常见的做法是利用远程跟踪分支。

#### 变基 vs. 合并
至此，你已在实战中学习了变基和合并的用法，你一定会想问，到底哪种方式更好。 在回答这个问题之前，让我们退后一步，想讨论一下提交历史到底意味着什么。

有一种观点认为，仓库的提交历史即是 记录实际发生过什么。 它是针对历史的文档，本身就有价值，不能乱改。 从这个角度看来，改变提交历史是一种亵渎，你使用 谎言 掩盖了实际发生过的事情。 如果由合并产生的提交历史是一团糟怎么办？ 既然事实就是如此，那么这些痕迹就应该被保留下来，让后人能够查阅。

另一种观点则正好相反，他们认为提交历史是 项目过程中发生的事。 没人会出版一本书的第一版草稿，软件维护手册也是需要反复修订才能方便使用。 持这一观点的人会使用 rebase 及 filter-branch 等工具来编写故事，怎么方便后来的读者就怎么写。

## Git 协议

Git 可以使用四种不同的协议来传输资料：本地协议（Local），HTTP 协议，SSH（Secure Shell）协议及 Git 协议。

### 在服务器上搭建Git

### 生成ssh公钥
许多 Git 服务器都使用 SSH 公钥进行认证。 为了向 Git 服务器提供 SSH 公钥，如果某系统用户尚未拥有密钥，必须事先为其生成一份。 这个过程在所有操作系统上都是相似的。 首先，你需要确认自己是否已经拥有密钥。 默认情况下，用户的 SSH 密钥存储在其 ~/.ssh 目录下。 进入该目录并列出其中内容，你便可以快速确认自己是否已拥有密钥：

cd ~/.shh



text