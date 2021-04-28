---
title: "Git Fast Forward"
date: 2020-12-07T14:32:23+08:00
draft: false
---

When you try to merge one commit with a commit that can be reached by following the first commit’s history, Git simplifies things by moving the pointer forward because there is no divergent work to merge together – this is called a “fast-forward.”

In another way,

If Master has not diverged, instead of creating a new commit, git will just point master to the latest commit of the feature branch. This is a “fast forward.”

There won't be any "merge commit" in fast-forwarding merge.

### 参考资料

- [What is git fast-forwarding](https://stackoverflow.com/questions/29673869/what-is-git-fast-forwarding)

- [What is the difference between `git merge` and `git merge --no-ff`?](https://stackoverflow.com/questions/9069061/what-is-the-difference-between-git-merge-and-git-merge-no-ff)