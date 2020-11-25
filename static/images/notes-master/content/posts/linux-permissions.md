---
title: "Linux Permissions"
date: 2020-08-31T11:17:11+08:00
draft: false
---

### Linux 文件属性

  查看文件

```shell
ls -al # 查看当前文件夹中文件的所有信息

[root@www ~]# ls -al

total 156

drwxr-x---   4    root   root     4096   Sep  8 14:06 .
drwxr-xr-x  23    root   root     4096   Sep  8 14:21 ..
-rw-------   1    root   root     1474   Sep  4 18:27 anaconda-ks.cfg
-rw-------   1    root   root      199   Sep  8 17:14 .bash_history
-rw-r--r--   1    root   root       24   Jan  6  2007 .bash_logout
-rw-r--r--   1    root   root      191   Jan  6  2007 .bash_profile
-rw-r--r--   1    root   root      176   Jan  6  2007 .bashrc
-rw-r--r--   1    root   root      100   Jan  6  2007 .cshrc
drwx------   3    root   root     4096   Sep  5 10:37 .gconf
drwx------   2    root   root     4096   Sep  5 14:09 .gconfd
-rw-r--r--   1    root   root    42304   Sep  4 18:26 install.log
-rw-r--r--   1    root   root     5661   Sep  4 18:25 install.log.syslog

```


ls是『list』的意思，重点在显示文件的文件名与相关属性。而选项『-al』则表示列出所有的文件详细的权限与属性

1. 第一栏代表这个文件的类型与权限(permission)

第一个字符代表这个文件是『目录、文件或链接文件等等』：

  当为[ d ]则是目录，例如上表档名为『.gconf』的那一行；
  当为[ - ]则是文件，例如上表档名为『install.log』那一行；
  若是[ l ]则表示为连结档(link file)；
  若是[ b ]则表示为装置文件里面的可供储存的接口设备(可随机存取装置)；
  若是[ c ]则表示为装置文件里面的串行端口设备，例如键盘、鼠标(一次性读取装置)。

第二栏表示有多少档名连结到此节点(i-node)


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


### 添加软链接

```shell
# /sub/uploads 添加 /uploads 软链接
ln -s  /uploads /sub/uploads

# 修改权限
chmod -R 777 uploads
```

## Quote


- [Linux修改目录权限](https://www.jianshu.com/p/ac4e994a47e2)

- [How do I change permissions for a folder and all of its subfolders and files in one step in Linux?](https://stackoverflow.com/questions/3740152/how-do-i-change-permissions-for-a-folder-and-all-of-its-subfolders-and-files-in)
- [chmod command](https://man.linuxde.net/chmod)