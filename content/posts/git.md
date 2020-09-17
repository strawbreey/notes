---
title: "Git"
date: 2020-08-28T19:29:52+08:00
draft: true
---

## create a new repository on the command line
```
git init
git add README.md
git commit -m "first commit"
git branch -M master
git remote add origin https://github.com/strawbreey/notes.git
git push -u origin master
```

## push an existing repository from the command line
```
git remote add origin https://github.com/strawbreey/notes.git
git branch -M master
git push -u origin master
```

## create gh-pages branch

```shell
git checkout --orphan gh-pages
git reset --hard
git commit --allow-empty -m "Initializing gh-pages branch"
git push upstream gh-pages
git checkout master

git remote add origin https://github.com/strawbreey/notes.git

rm -rf public
git worktree add -B gh-pages public upstream/gh-pages

hugo
cd public && git add --all && git commit -m "Publishing to gh-pages" && cd ..

git push upstream gh-pages
git push --set-upstream origin gh-pages

git branch --set-upstream-to=origin/<branch> gh-pagesgit branch --set-upstream-to=origin/<branch> gh-pages
```

### get theme
```
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
git submodule add https://github.com/strawbreey/hugo-theme-hello-friend.git themes/friend
```


## new gitignore 
```
echo "public" >> .gitignore
```

## Set Git Account

```shell
# 查看当前用户
git config user.name
git config user.email

# 设置当前的用户信息
git config user.name "strawbreey" 
git config user.email "2675882608@qq.com"
```


# 删除分支
 
```shell
git branch -d [branch-name]
``` 
 
 
# 删除远程分支
```shell 
git push origin --delete [branch-name]
 
git branch -dr [remote/branch]
```


## gh-pages 分之



```shell
git worktree list # 查看worktree

删除worktree
git worktree prune