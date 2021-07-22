---
title: "Git Revert"
date: 2021-07-22T11:00:33+08:00
draft: false
---


### git revert

git revert 撤销 某次操作，此次操作之前和之后的commit和history都会保留，并且把这次撤销
作为一次最新的提交

```bash
    * git revert HEAD                 # 撤销前一次 commit
    * git revert HEAD^                # 撤销前前一次 commit
    * git revert commit               #（比如：fa042ce57ebbe5bb9c8db709f719cec2c58ee7ff）撤销指定的版本，撤销也会作为一次提交进行保存。
```
git revert是提交一个新的版本，将需要revert的版本的内容再反向修改回去，