---
title: "Linux Install Php"
date: 2020-09-01T10:10:41+08:00
draft: true
---

## Install PHP

### install php5.6

```shell
  wget -c http://cn2.php.net/distributions/php-5.6.21.tar.gz
  tar -zxvf php-5.6.21.tar.gz


```


  格式：tar zcvf 压缩后的路径及包名 你要压缩的文件

　　z:gzip压缩

　　c:创建压缩包

　　v:显示打包压缩解压过程

　　f:接着压缩

　　t:查看压缩包内容

　　x:解压

　　X:指定文件列表形式排除不需要打包压缩的文件或目录

　　-exclude:指定排除文件或目录不需要打包压缩的文件或目录（也可以用正则匹配/通配符等）

　　-C:解压到指定目录


例如：将/root/cs/test/下文件压缩到/root/cs/test1文件下命名为1

tar zcf /root/cs/test1/1.tar.gz /root/cs/test/

zcvf :   打包压缩     
xvf: 解压缩

## php 下载 解压 配置 编译
https://www.php.net/distributions/php-7.2.33.tar.gz

wget http://cn2.php.net/distributions/php-7.2.33.tar.gz tar xf php-7.2.33.tar.xz

tar -zxvf php-7.2.33.tar.xz

./configure --prefix=/usr/local/php72 --with-config-file-path=/usr/local/php72/etc --with-mysql=mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd --with-gd --with-iconv --with-zlib --enable-xml --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization  --enable-mbregex --enable-fpm --enable-mbstring --enable-ftp --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-pear --with-gettext --enable-session --with-mcrypt --with-curl

make && make install

make test


### 5.6
./configure --prefix=/usr/local/php56 --with-config-file-path=/usr/local/php56/etc --with-mysql=mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd --with-gd --with-iconv --with-zlib --enable-xml --enable-bcmath --enable-shmop --enable-sysvsem --enable-inline-optimization  --enable-mbregex --enable-fpm --enable-mbstring --enable-ftp --enable-gd-native-ttf --with-openssl --enable-pcntl --enable-sockets --with-xmlrpc --enable-zip --enable-soap --with-pear --with-gettext --enable-session --with-mcrypt --with-curl


## 编辑php-fpm.conf

cd /usr/local/php72/etc
cp php-fpm.conf.default php-fpm.conf
vi php-fpm.conf


/usr/local/php56/bin/phpize  ./configure  --with-php-config=/usr/local/php56/bin/php-config

## quote
[CentOs 同时使用多个php版本](https://www.jianshu.com/p/600ef6e83af1)



https://www.osyunwei.com/archives/5327.html