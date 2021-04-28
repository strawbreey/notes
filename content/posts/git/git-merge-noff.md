---
title: "Git Merge Noff"
date: 2021-03-24T15:21:51+08:00
draft: false
---

git 合并分支默认是fast forword(快速合并的模式)

缺点是，不能显示历史信息，不知道那些分支曾经合并过。

no-ff (no-fast-forward) ，这种方式会在合并的同时生成一个新的commit, 这样可以看出分支信息

```bash
git merge --no-ff feature-x
```