---
title: "Php Cache"
date: 2021-08-17T14:44:28+08:00
draft: false
---

### PHP四大加速缓存器`opcache`，`apc`，`xcache`，`eAccelerator`与php解析的初步理解

php缓存器大概有这个几种： opcache、APC、xcache、eAccelerator 

加速器的原理是将编译后的源码缓存起来，当下次执行相同代码时，可以省去词法语法分析等步骤，提高php的执行效率

下面是这些加速的配置方式

### 1、opcache（官方推出的产品还是好用的）

opcache 通过将 PHP 脚本预编译的字节码存储到共享内存中来提升 PHP 的性能， 存储预编译字节码的好处就是 省去了每次加载和解析 PHP 脚本的开销。

PHP 5.5.0 及后续版本中已经绑定了 opcache 扩展。 对于 PHP 5.2，5.3 和 5.4 版本可以使用PECL扩展中的 opcache 库。

windows下的php扩展下载地址：http://windows.php.net/downloads/pecl/releases/

配置如下：

打开php.ini文件，找到[opcache]

```ini
; dll地址
extension=php_opcache.dll
; 开关打开
opcache.enable=1
; 开启CLI
opcache.enable_cli=1
; 可用内存, 酌情而定, 单位为：Mb
opcache.memory_consumption=128
; Zend Optimizer + 暂存池中字符串的占内存总量.(单位:MB)
opcache.interned_strings_buffer=8
; 对多缓存文件限制, 命中率不到 100% 的话, 可以试着提高这个值
opcache.max_accelerated_files=10000
; opcache 会在一定时间内去检查文件的修改时间, 这里设置检查的时间周期, 默认为 2, (单位:秒)
opcache.revalidate_freq=1
; 打开快速关闭, 打开这个在PHP Request Shutdown的时候回收内存的速度会提高
opcache.fast_shutdown=1
```

### 二、APC

Alternative PHP Cache (APC) 是一个开放自由的PHP opcode缓存。它的目标是提供一个自由、 开放，和健全的框架用于缓存和优化PHP的中间代码。

- 下载apc扩展dll，选择你对应的PHP版本: http://windows.php.net/downloads/pecl/releases/apc/

把下载的php_apc.dll放入php的ext扩展目录下。
打开php.ini文件，

配置如下：

```ini
[apc]
extension=php_apc.dll
apc.enabled=1
; 共享内存块的数目
apc.shm_segments=1
; 共享内存块的大小(单位:MB)
apc.shm_size=64
; 优化级别，更高的值则使用更主动的优化
apc.optimization=1
; 源文件的数目，不确定设置为0
apc.num_files_hint=0
; 缓存条目在缓冲区中允许逗留的秒数
apc.ttl=7200
; 针对每个用户缓存条目在缓冲区中允许逗留的秒数
apc.user_ttl=7200
; 缓存条目在垃圾回收表中能够存在的秒数
apc.gc_ttl=7200
; 文件写锁
apc.write_lock=on
```

### 三、xcache（国产）

xcache是一个开源的 opcode 缓存器/优化器。

- 下载xcache，选择对应PHP版本的xcache http://xcache.lighttpd.net/pub/Releases/

把php_xcache.dll放到php的ext目录下。

php.ini配置如下：
```ini
[xcache-common]
extension = php_xcache.dll
[xcache.admin]
xcache.admin.enable_auth = On
xcache.admin.user = "admin"
xcache.admin.pass = "md5后你的密码"
[xcache]
; 选择底层内存共享实现方案
xcache.shm_scheme = "mmap"
xcache.size = 128M
; 设置为CPU数
xcache.count = 1
; 只是个参考值
xcache.slots = 8K
; 缓存时间
xcache.ttl = 1200
; 垃圾回收的时间间隔
xcache.gc_interval = 1200
; 同上，针对变量缓存设置
xcache.var_size = 4M
xcache.var_count = 1
xcache.var_slots = 8K
xcache.var_ttl = 1200

; 变量最大缓存时间
xcache.var_maxttl = 7200
xcache.var_gc_interval = 1200
xcache.var_namespace_mode = 0
xcache.var_namespace = ""
xcache.readonly_protection = Off
; 对于win系统，这里不是文件路径
xcache.mmap_path = "xcache"
xcache.coredump_directory = ""
xcache.coredump_type = 0
xcache.disable_on_crash = Off
xcache.experimental = Off
xcache.cacher = On
xcache.stat = On
xcache.optimizer = Off
[xcache.coverager]
xcache.coverager = Off
xcache.coverager_autostart = On
xcache.coveragedump_directory = ""
```