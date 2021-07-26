---
title: "Git Auth"
date: 2020-12-31T15:16:37+08:00
draft: false
---



### 用户信息

```shell
# 设置全局的用户信息
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com

# 设置项目的用户信息
$ git config user.name "John Doe"
$ git config user.email johndoe@example.com

```

### 

如果我们git clone的下载代码的时候是连接的https://而不是git@git (ssh)的形式，当我们操作git pull/push到远程的时候，总是提示我们输入账号和密码才能操作成功，频繁的输入账号和密码会很麻烦。


```bash
# 设置默认账户
git config --global credential.helper store

# 这样保存的密码是明文的，保存在用户目录~的.git-credentials文件中
$ file ~/.git-credentials
$ cat ~/.git-credentials

# 1、设置记住密码（默认15分钟）：
git config --global credential.helper cache

# 2、如果想自己设置时间，可以这样做：
git config credential.helper 'cache --timeout=3600' # 这样就设置一个小时之后失效

# 3、长期存储密码：
git config --global credential.helper store

4、增加远程地址的时候带上密码也是可以的。(推荐)
http://yourname:password@git.oschina.net/name/project.git
```

