---
title: "Git Abort Merge"
date: 2020-12-03T17:42:58+08:00
draft: false
---

### 撤销合并

```bash
git reset --hard HEAD

git pull --strategy=theirs remote_branch

git fetch origin
git reset --hard origin

git reset --merge

git merge --abort #等同于git reset --merge何时MERGE_HEAD存在。

git stash 
git stash pop 

# If you have changes you don't want to commit before starting a merge, just git stash them before the merge and git stash pop after finishing the merge or aborting it.

git revert
git reset
```


### 参考资料

- [I ran into a merge conflict. How can I abort the merge?](https://stackoverflow.com/questions/101752/i-ran-into-a-merge-conflict-how-can-i-abort-the-merge)
