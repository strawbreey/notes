---
title: "Git Common Commands"
date: 2020-08-31T15:39:29+08:00
draft: false
tags: ['git']
---


### Git不需要输入用户名密码的两种方式
1. 采用ssh连接方式: 将自己主机的id_ras.pub 加到git 的ssh Keys 中

2. 采用https的连接方式:

```
git config --global credential.helper store
```


### 从某一个commit开始创建本地分支 

1、git log 查看提交
2、通过checkout 跟上commitId 即可创建制定commit之前的本地分支 
git checkout commitId -b branchName 

