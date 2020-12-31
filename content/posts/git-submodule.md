---
title: "Git Submodule"
date: 2020-08-31T10:20:39+08:00
draft: false
---


```bash
echo "Deleting old publication"
rm -rf public
mkdir public
git worktree prune
rm -rf .git/worktrees/public/

echo "Checking out gh-pages branch into public"
git worktree add -B gh-pages public upstream/gh-pages

echo "Removing existing files"
rm -rf public/*

echo "Generating site"
hugo

echo "Updating gh-pages branch"
cd public && git add --all && git commit -m "Publishing to gh-pages (publish.sh)"
```

Git 本身不支持对某一特定文件夹或者子目录的权限控制，但是可以通过子模块的形式间接实现。 

### 子模块的初始化
```shell
# ​通过git submodule add的方式：
git submodule add https://github.com/xxx/example.git        

```
回车后，将在本地的仓库内创建了一个子模块

我们可以是用git status来查看一下状态，如下面： 在这里我们可以看到两个文件：

```shell
git status
```

- 文件.gitmodules .gitmodules文件里面记录着子模块的路径和项目的url。如果你有多个子模块，你都可以在这里查看到具体的信息，同时也是他人获取你子模块信息的地方。

- 文件example

xxxx是你的子模块的路径。当你不在该路径时，git将不会记录其内容。相反的，git把该模块当作一个特别的提交。当你项目提交的时候，例如：

```shell
git commit -am 'added one module'
```
在这里你可以再一次看到子模块当作一个特别的提交（160000模式是git的一个特殊的模式）

小结：直接在要引入的项目目录下中克隆一个子模块：git submodule add [url]

### 克隆一个带有子模块的项目

当你克隆一个带有子模块的项目时，你会发现，子模块（子路径）里面的内容完全为空。
```shell
git clone https://git.code.oa.com/xxx/submodules-test.git
cd submodules-test/
ls -la

cd example/
ls
```
此时，你应该在主项目路径下运行git submodule init初始化，再运行git submodule update命令抓取内容下来。

``` shell
cd ..
git submodule init
# Submodule 'example' (https://git.code.oa.com/xxx/example.git) registered for path 'example'

git submodule update
# Cloning into 'example'...
# remote: Counting objects: 3, done.        
# remote: Total 3 (delta 0), reused 0 (delta 0)        
# Unpacking objects: 100% (3/3), done.
# Submodule path 'example': checked out '1cd699621fb9f29a910e5ede48782bbbfd41a627'

git submodule update
# fatal: reference isn’t a tree: 6c5e70b984a60b3cecd395edd5b48a7575bf58e0
# Unable to checkout '6c5e70b984a60b3cecd395edd5ba7575bf58e0' in submodule path…
```
这个时候，你就能看到你子模块下上次最新提交后的内容。 如果你不想这么麻烦，你可以在克隆项目时，附上--recursive参数，这样git会自动的初始化和抓取子模块的过程， git clone --recursive https://git.code.oa.com/xxx/submodules-test.git。克隆项目的时候需要谨记的是，每一次抓取包含子模块的项目都必须要这么做。 另外还需要注意的是，如果开发者在本地对子模块进行了修改，却没有push上远程仓库上，其他开发者进行根据它非公开的指针拉取的时会出现以下这个错误。


会出现这个原因是因为子模块会从.gitmodules中获取url相应的提交，可是此时由于之前的开发者没有push到相应的url，导致而找不到提交的内容，因此会发生非公开的指针和url不匹配的原因。如果想要避免这样的情况，你需要查看一下谁是最后提交的人(git log -1)，然后告知该开发者进行push。

小结：克隆完一个带有子模块的项目，需要初始化子模块，然后更新子模块（或者克隆项目时带有--recursive参数）。

### 工作在带有子模块的项目中

克隆了一个新的带有子模块的项目在本地仓库，如果此时我们想要修改并提交submodule。这里分两种情况，一种是直接在子项目修改并提交，另一种是其它人修改子模块的git仓库。

第一种情况

需要注意一点，子模块的分支处于游离状态，我们在修改它的时候第一步需要执行检出对应的分支。

```shell

# 首先，我们进入子模块
$ cd example

# 检出master分支
$ git checkout master

# 然后做修改
$ vim test.py

# 最后做提交
$ git add .
$ git commit -m "Get something done"
$ git push origin master

# 如何将submodule的变更在父项目中提交

# 回到父项目中更新
$ cd ..

# 子模块被提交后父项目会检测到，正常提交即可
$ git status

# modified:   example (new commits)

# add的时候注意再最后不要加 / 斜杠，否则会出现很棘手问题(push远端后子模块不带版本号)

$ git add example
$ git commit -m "update submodule"
$ git push
```

第二种情况

当子模块是属于别人维护的时候，当他更新的时候，如果我们需要同步更新依赖的子模块。

```shell
$ git submodule update --remote example
  Submodule path 'example': checked out '826a71f68f64cad5cafb3cafdb0b726d41b8d868'
  
# 子模块更新后父项目会检测到，正常提交即可
$ git status

# modified:   example (new commits)

# add的时候注意再最后不要加 / 斜杠，否则会出现很棘手问题(push远端后子模块不带版本号)

$ git add example
$ git commit -m "update submodule"
$ git push
```

### 子模块的删除
删除submodule，需要借助git submodule deinit命令，具体参考如下：
```shell
# 逆初始化模块，其中{MOD_NAME}为模块目录，执行后可发现模块目录被清空 .git/config 已被重写
# git submodule deinit {MOD_NAME}
$ git submodule deinit example

# 删除子模块文件，同时删除.gitmodules中记录的模块信息
$ git rm example

# 提交更改到代码库，可观察到'.gitmodules'内容发生变更
$ git commit -m "Removed submodule "

# 删除.git的保存
$ rm -rf .git/modules/example
```
    
### 子模块注意事项
1. 子模块可以理解为某仓库的指定版本的内容作为当前仓库的某一子目录
2. 父仓库不直接关注子模块的内容变化，只关注子模块版本号的变化
3. 父仓库的任意一个版本都唯一的确定了子模块的版本
4. 子模块的版本号发生变化将导致父仓库产生变化（diff）父仓库需要提交这一变化来记录相应改动，而这，也将导致父仓库产生一个新的版本
5. 子仓库可以完全不知道父仓库的存在
6. 父仓库可以不存储任何代码，而仅仅是记录一系列子模块的版本号，这些版本号的变更形成了父仓库的历史
7. 可以使用子模块来划分目录的权限

### 更新submodule的坑

submodule 项目和它的父项目本质上是 2 个独立的 git 仓库。只是父项目存储了它依赖的 submodule 项目的版本号信息而已。如果有人更新了submodule，然后更新了父项目中依赖的版本号。 你需要在 git pull 之后，调用 git submodule update 来更新 submodule 信息。这儿的问题在于，如果你 git pull 之后，忘记了调用 git submodule update，那么你极有可能再次把旧的 submodule 依赖信息提交上去。 对于那些习惯使用 git commit -a 的人来说，这种危险会更大一些。所以建议：

git pull 之后，立即执行 git status, 如果发现 submodule 有修改，立即执行 git submodule update
尽量不要使用 git commit -a， git add 命令存在的意义就是让你对加入暂存区的文件做二次确认，而 git commit -a 相当于跳过了这个确认过程。
更复杂一些，如果你的 submodule 又依赖了 submodule，那么很可能你需要在 git pull 和 git submodule update 之后，再分别到每个 submodule 中再执行一次 git submodule update。 这里可以使用 git submodule foreach 命令来实现： git submodule foreach git submodule update