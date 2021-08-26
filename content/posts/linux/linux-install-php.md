---
title: "Linux Install Php"
date: 2020-09-01T10:10:41+08:00
draft: false
---

#### linux 常用命令

格式：tar zcvf 压缩后的路径及包名 你要压缩的文件
　z: gzip压缩
  c: 创建压缩包
  v: 显示打包压缩解压过程
  f: 接着压缩
  t: 查看压缩包内容
  x: 解压
  X: 指定文件列表形式排除不需要打包压缩的文件或目录
  -exclude: 指定排除文件或目录不需要打包压缩的文件或目录（也可以用正则匹配/通配符等）
  -C: 解压到指定目录


```bash
# 例如：将`/root/cs/test/`下文件压缩到`/root/cs/test1`文件下命名为1
tar zcf /root/cs/test1/1.tar.gz /root/cs/test/
```

zcvf:   打包压缩     
xvf: 解压缩



### php7.2安装

```bash
# https://www.php.net/distributions/php-7.2.33.tar.gz

# 下载php7.2
wget http://cn2.php.net/distributions/php-7.2.33.tar.gz tar xf php-7.2.33.tar.xz

# 解压
tar -zxvf php-7.2.33.tar.xz

# 配置
./configure --prefix=/usr/local/php72 --with-config-file-path=/usr/local/php72/etc --with-mysql=mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd --with-gd --with-iconv --with-zlib --enable-xml --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization  --enable-mbregex --enable-fpm --enable-mbstring --enable-ftp --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-pear --with-gettext --enable-session --with-mcrypt --with-curl

# 安装
make && make install

# 测试
make test
```

### php405安装

### 安装php5.6

```bash

# 下载 php5.6
wget -c http://cn2.php.net/distributions/php-5.6.21.tar.gz

# 解压
tar -zxvf php-5.6.21.tar.gz

# 编译
./configure --prefix=/usr/local/php56 --with-config-file-path=/usr/local/php56/etc --with-mysql=mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd --with-gd --with-iconv --with-zlib --enable-xml --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization  --enable-mbregex --enable-fpm --enable-mbstring --enable-ftp --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-pear --with-gettext --enable-session --with-mcrypt --with-curl

# 安装
make && make install

# 测试
make test
```

## 配置php-fpm.conf

```bash
# 修改php-fpm.cong
cd /usr/local/php72/etc
cp php-fpm.conf.default php-fpm.conf
vi php-fpm.conf

# 编译
/usr/local/php56/bin/phpize  ./configure  --with-php-config=/usr/local/php56/bin/php-config
```


### linux 


php-5.3.28.tar.bz

yum install php54w.x86_64 php54w-cli.x86_64 php54w-common.x86_64 php54w-gd.x86_64 php54w-ldap.x86_64 php54w-mbstring.x86_64 php54w-mcrypt.x86_64 php54w-mysql.x86_64 php54w-pdo.x86_64


yum -y install php php-fpm php-mysql php-mcrypt php-pdo php-mbstring php-ldap

### yum 安装 php

> centos7.0用yum直接安装apache、php他们的默认版本是 apache2.4 和 php5.4

```bash
# 在安装前 ，更新一下系统
yum update

# 安装一些必备的包
yum -y install gcc
yum -y install gcc-c++
yum -y install make

# 开始安装apache
yum install httpd

# 现在去浏览器中输入的服务器的ip，正常情况你是访问不了的，因为有防火墙默认是没有对80端口开启的，所以现在要去开放防火墙对80端口开放

yum install iptables-services ## 安装iptables防火墙
vi /etc/sysconfig/iptables ##修改配置
-A INPUT -m state –state NEW -m tcp -p tcp –dport 80 -j ACCEPT #允许80端口通过防火墙 1
-A INPUT -m state –state NEW -m tcp -p tcp –dport 3306 -j ACCEPT #允许3306端口通过防火墙 2
# 把1、2两条规则保存到打开的配置文件里面，注意：要放在20端口下面
systemctl restart firewalld.service ##重启防火墙

find / -name httpd.conf ## 假如找不到httpd.conf文件
# 找到ServerName —–改成：ServerName localhost:80
# 重启：
systemctl restart httpd.service
# 设置apache开机启动
systemctl enable httpd.service 


# 安装php, 默认是php5.4
yum -y install php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc

# 安装php7.0

rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm
yum -y install php70w.x86_64 php70w-cli.x86_64 php70w-common.x86_64 php70w-gd.x86_64 php70w-ldap.x86_64 php70w-mbstring.x86_64 php70w-mysql.x86_64 php70w-pdo.x86_64 php70w-pear.noarch php70w-process.x86_64 php70w-xml.x86_64 php70w-xmlrpc.x86_64


# 安装PHP FPM
yum install php55w-fpm
yum install php56w-fpm
yum install php70w-fpm

# 安装mysql
yum -y install mysql-connector-odbc mysql-devel libdbi-dbd-mysql


# 安装php支持redis扩展

# 下载phpredis包
# tar -zxvf 包名
/usr/bin/phpize
./configure –with-php-config=php-config的路径 你可以find / -name php-config 查看路径
make
make && install
```

## 引用
[CentOs 同时使用多个php版本](https://www.jianshu.com/p/600ef6e83af1)