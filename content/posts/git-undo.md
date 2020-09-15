---
title: "Git Undo"
date: 2020-08-31T15:11:27+08:00
draft: true
---

## undo 
```shell
  #Clone your git repo
  git clone https://github.com/lestatzhang/lestatzhang.github.io.git;

  #Entre your local repo
  cd lestatzhang.github.io;

  #Checkout
  git checkout --orphan latest_branch;

  #Add all the files
  git add -A;

  #Commit the changes
  git commit -am "Reinitialize";

  #Delete the branch
  git branch -D master;

  #Rename the current branch to master
  git branch -m master;
  
  #Finally, force update your repository
  git push -f origin master;
```
## 

```
  git reset --hard HEAD^ 
  git push --force
```