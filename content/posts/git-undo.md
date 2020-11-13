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
  git reset --mixed HEAD^ 
  git reset --soft HEAD^ 

  git push --force


# --mixed  意思是：不删除工作空间改动代码，撤销commit，并且撤销git add 。 操作这个为默认参数,git reset --mixed HEAD^ 和 git reset HEAD^ 效果是一样的。
  
# --soft 不删除工作空间改动代码，撤销commit，不撤销git add . 
 
# --hard 删除工作空间改动代码，撤销commit，撤销git add . 

# 注意完成这个操作后，就恢复到了上一次的commit状态。
```


常用git stash命令：

- git stash save "save message"  : 执行存储时，添加备注，方便查找，只有git stash 也要可以的，但查找时不方便识别。

- git stash list  ：查看stash了哪些存储

- git stash show ：显示做了哪些改动，默认show第一个存储,如果要显示其他存贮，后面加stash@{$num}，比如第二个 git stash show stash@{1}

- git stash show -p : 显示第一个存储的改动，如果想显示其他存存储，命令：git stash show  stash@{$num}  -p ，比如第二个：git stash show  stash@{1}  -p

- git stash apply :应用某个存储,但不会把存储从存储列表中删除，默认使用第一个存储,即stash@{0}，如果要使用其他个，git stash apply stash@{$num} ， 比如第二个：git stash apply stash@{1} 

- git stash pop ：命令恢复之前缓存的工作目录，将缓存堆栈中的对应stash删除，并将对应修改应用到当前的工作目录下,默认为第一个stash,即stash@{0}，如果要应用并删除其他stash，命令：git stash pop stash@{$num} ，比如应用并删除第二个：git stash pop stash@{1}

- git stash drop stash@{$num} ：丢弃stash@{$num}存储，从列表中删除这个存储

- git stash clear ：删除所有缓存的stash