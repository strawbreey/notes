---
title: "Linux Install Php Imagick"
date: 2020-10-09T14:54:01+08:00
draft: true
---

说明：

php安装目录：/usr/local/php5
php.ini配置文件路径：/usr/local/php5/etc/php.ini
Nginx安装目录：/usr/local/nginx
Nginx网站根目录：/usr/local/nginx/html

1. 安装编译工具
```shell
yum install wget  make gcc gcc-c++ gtk+-devel zlib-devel openssl openssl-devel pcre-devel kernel keyutils  patch perl
```

2. 安装ImageMagick

```shell
cd /usr/local/src  #进入软件包存放目录
wget http://www.imagemagick.org/download/ImageMagick.tar.gz  #下载ImageMagick
tar zxvf ImageMagick.tar.gz  #解压
cd ImageMagick-6.7.9-3   #进入安装目录
./configure --prefix=/usr/local/imagemagick  #配置
make #编译
make install  #安装
export PKG_CONFIG_PATH=/usr/local/imagemagick/lib/pkgconfig/  #设置环境变量
```


3. 安装imagick

```shell
cd /usr/local/src
wget http://pecl.php.net/get/imagick-3.0.1.tgz  #下载imagick
tar zxvf imagick-3.0.1.tgz
cd imagick-3.0.1
/usr/local/php5/bin/phpize #用phpize生成configure配置文件
./configure  --with-php-config=/usr/local/php5/bin/php-config --with-imagick=/usr/local/imagemagick #配置
make #编译
make install #安装
```

备注：在安装过程中出现错误，一般是由于缺少编译工具包导致，可根据提示参照第一步安装相应的工具包即可
安装完成之后，出现下面的界面，记住以下路径，后面会用到
Installing shared extensions:     /usr/local/php5/lib/php/extensions/no-debug-non-zts-20090626/ #imagick模块路径

4. 配置php支持imagick
```
vi /usr/local/php5/etc/php.ini  #编辑配置文件，在最后一行添加以下内容
extension="imagick.so"
```
5. 测试
```shell
vi /usr/local/nginx/html/phpinfo.php   #编辑，输入以下代码

vi /usr/local/nginx/html/imagick.php   #编辑，输入以下代码

service php-fpm restart  #重启php-fpm
service nginx restart  #重启nginx
```

```php
# vi /usr/local/nginx/html/phpinfo.php 
<?php
phpinfo();
?>

# /usr/local/nginx/html/imagick.php 
<?php
  header('Content-type: image/jpeg');
  $image = new Imagick('www.osyunwei.com.jpg');
  $image->thumbnailImage(300, 225);
  echo $image;
?>

```