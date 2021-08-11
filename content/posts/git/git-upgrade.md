---
title: "Git Upgrade"
date: 2021-08-10T20:50:32+08:00
draft: false
---

## Git版本升级


### 一、安装依赖包：

1. 下载安装 libiconv-1.14.tar.gz

```bash
wget http://ftp.gnu.org/pub/gnu/libiconv/libiconv-1.14.tar.gz
tar zxvf libiconv-1.14.tar.gz
cd libiconv-1.14
./configure --prefix=/usr/local/libiconv
make && make install
```

2. yum安装依赖包

```bash
yum -y  install curl-devel expat-devel gettext-devel openssl-devel zlib-devel asciidoc gcc perl-ExtUtils-MakeMaker tcl xmlto
# 注：asciidoc如果无法通过yum安装，请下载包，编译安装

wget --no-check-certificate https://jaist.dl.sourceforge.net/project/asciidoc/asciidoc/8.6.9/asciidoc-8.6.9.zip
unzip asciidoc-8.6.9.zip
cd asciidoc-8.6.9
./configure
make && make install
```

### 二、安装 新版git包

1. 卸载git旧版本

```bash
yum remove git
```

2. 编译按照 git 源码

```bash
wget https://github.com/git/git/archive/v2.2.1.tar.gz
tar zxvf v2.2.1.tar.gz
cd git-2.2.1
make configure
./configure --prefix=/usr/local/git --with-iconv=/usr/local/libiconv
make all doc
make install install-doc install-html
 echo "export PATH=$PATH:/usr/local/git/bin:/usr/local/git/libexec/git-core" >> /etc/bashrc
```

### 三、查看新版本

```bash
git version
# git version 2.2.1
```