---
title: "Vim Swap"
date: 2020-10-20T11:26:29+08:00
draft: false
---

### Found a swap file by the name

```bash
E325: ATTENTION
Found a swap file by the name ".git/.MERGE_MSG.swp"
          owned by: xxxxxx   dated: Mon Nov 12 23:17:40 2012
         file name: ~xxxxxx/Desktop/My-ios-App/.git/MERGE_MSG
          modified: YES
         user name: xxxxxx   host name: unknown-b8-8d-12-22-27-72.lan
        process ID: 1639
While opening file ".git/MERGE_MSG"
             dated: Tue Nov 13 14:06:48 2012
      NEWER than swap file!

(1) Another program may be editing the same file.
    If this is the case, be careful not to end up with two
    different instances of the same file when making changes.
    Quit, or continue with caution.

(2) An edit session for this file crashed.
    If this is the case, use ":recover" or "vim -r .git/MERGE_MSG"
    to recover the changes (see ":help recovery").
    If you did this already, delete the swap file ".git/.MERGE_MSG.swp"
    to avoid this message.

Swap file ".git/.MERGE_MSG.swp" already exists!
[O]pen Read-Only, (E)dit anyway, (R)ecover, (D)elete it, (Q)uit, (A)bort:

```


Accepted answer fails to mention how to delete the .swp file.

Hit "D" when the prompt comes up and it will remove it.

In my case, after I hit D it left the latest saved version intact and deleted the .swp which got created because I exited VIM incorrectly

### 参考链接 

- [Found a swap file by the name](https://stackoverflow.com/questions/13361729/found-a-swap-file-by-the-name)

idea.waibao.oa.com/web/file/145