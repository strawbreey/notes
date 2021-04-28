---
title: "Git Auth"
date: 2020-12-31T15:16:37+08:00
draft: false
---

如果我们git clone的下载代码的时候是连接的https://而不是git@git (ssh)的形式，当我们操作git pull/push到远程的时候，总是提示我们输入账号和密码才能操作成功，频繁的输入账号和密码会很麻烦。

解决办法：

git bash进入你的项目目录，输入：

git config --global credential.helper store