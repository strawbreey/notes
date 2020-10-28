---
title: "Git Undo"
date: 2020-08-31T15:11:27+08:00
draft: true
---

### 删除commit提交记录

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

### 撤销上一次的commit

```shell
  git reset --hard HEAD^ 
  git push --force


# --mixed  意思是：不删除工作空间改动代码，撤销commit，并且撤销git add 。 操作这个为默认参数,git reset --mixed HEAD^ 和 git reset HEAD^ 效果是一样的。
  
# --soft 不删除工作空间改动代码，撤销commit，不撤销git add . 
 
# --hard 删除工作空间改动代码，撤销commit，撤销git add . 

# 注意完成这个操作后，就恢复到了上一次的commit状态。
```
