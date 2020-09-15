---
title: "Linux Permissions"
date: 2020-08-31T11:17:11+08:00
draft: false
---

## look auth

r = 读取权限 (4)
w = 写入权限 (2)
x = 执行权限 (1)

1-3位数字代表文件所有者的权限，4-6位数字代表同组用户的权限，7-9数字代表其他用户的权限。

444 r--r--r--
600 rw-------
644 rw-r--r--
666 rw-rw-rw-
700 rwx------
744 rwxr--r--
755 rwxr-xr-x
777 rwxrwxrwx

## Quote


- [Linux修改目录权限](https://www.jianshu.com/p/ac4e994a47e2)

- [How do I change permissions for a folder and all of its subfolders and files in one step in Linux?](https://stackoverflow.com/questions/3740152/how-do-i-change-permissions-for-a-folder-and-all-of-its-subfolders-and-files-in)
- [chmod command](https://man.linuxde.net/chmod)