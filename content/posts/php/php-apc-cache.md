---
title: "Php Apc Cache"
date: 2021-08-06T20:07:55+08:00
draft: false
---


### APC缓存简介

APC(Alternative PHP Cache)是一个PHP缓存。它在内存中存储PHP页面并且减少了硬盘的I/O。这对于性能的提升十分明显。你甚至可以在CPU使用率下降50%的情况下提升系统50%的性能。

### windows下安装PHP的APC拓展

1、首先要确定你的Php版本。
2、下载和你的PHP对应的 php_apc.dll
3、官网下载地址：http://pecl.php.net/package/apc
4、把下载php_apc.dll 复制到PHP的ext目录下
5、修改php.ini 在末尾添加 extension=php_apc.dll
6、来配置php.ini文件

    默认的官方APC配置：

    ```bash
    apc.cache_by_default = On
    　　apc.enable_cli = Off
    　　apc.enabled = On
    　　apc.file_update_protection = 2
    　　apc.filters =
    　　apc.gc_ttl = 3600
    　　apc.include_once_override = Off
    　　apc.max_file_size = 1M
    　　apc.num_files_hint = 1000
    　　apc.optimization = Off
    　　apc.report_autofilter = Off
    　　apc.shm_segments = 1
    　　apc.shm_size = 30
    　　apc.slam_defense = 0
    　　apc.stat = On
    　　apc.ttl = 0
    　　apc.user_entries_hint = 100
    　　apc.user_ttl = 0
    　　apc.write_lock = On
    ```

    关于APC完整的参数设置的解释，请查阅：http://www.php.net/apc
    我们可以做如下设置php.ini

    ```bash
    apc.enabled = 1
    apc.shm_segments = 1
    apc.shm_size = 128
    apc.max_file_size = 10M
    apc.stat=1
    ```

    剩余其它的设置将会使用默认值。

7、设置临时目录APC需要一个临时目录来存储文件。它会尝试在windows的临时目录缓存文件，事先请给临时目录写的权限。
8、重新启动apache服务器,查看phpinfo中是否有apc的配置项目,有的话就配置成功了

### APC小结

1、缓存期限: APC的缓存分两部分: 系统缓存和用户数据缓存.

系统缓存 是自动使用的, 是指APC把PHP文件源码的编译结果缓存起来,然后在再次调用时先对比时间标记。如果未过期,则使用缓存代码运行。默认缓存 3600s(一小时). 但是这样仍会浪费大量CPU时间. 因此可以在php.ini中设置system缓存为永不过期(apc.ttl=0).不过如果这样设置,改运php代码后需要restart一下您的web服务器(比如apache…).目前对APC的性能测试一般指的是这一层cache。

2、用户数据缓存 由用户在编写php代码时用apc_store和apc_fetch函数操作读取、写入的. 如果量不大的话我建议可以使用一下.如果量大,我建议使用memcache会更好. 如果要享受APC带来的缓存大文件上传进度的特性, 需要在php.ini中将apc.rfc1867设为1,并且在表单中加一个隐藏域APC_UPLOAD_PROGRESS,这个域的值可以随机生成一个hash,以确何唯一.具体例子请参见前面给出的链接.

3、你不能在同一台服务器上同时安装APC和Zend Optimiser，二者只能选其一。

4、其实所有Windows环境下的php环境都可以使用这个方法。