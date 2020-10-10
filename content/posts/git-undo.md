---
title: "Git Undo"
date: 2020-08-31T15:11:27+08:00
draft: true
---

删除commit提交记录

```shell
  # Checkout
  git checkout --orphan latest_branch;

  # Add all the files
  git add -A;

  # Commit the changes
  git commit -am "Reinitialize";

  # Delete the branch
  git branch -D master;

  # Rename the current branch to master
  git branch -m master;
  
  # Finally, force update your repository
  git push -f origin master;
```

撤销上一次的commit

```shell
  git reset --hard HEAD^ 
  git push --force
```