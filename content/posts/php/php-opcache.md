---
title: "Php Opcache"
date: 2021-08-06T20:11:19+08:00
draft: false
---

OPcache 通过将 PHP 脚本预编译的字节码存储到共享内存中来提升 PHP 的性能， 存储预编译字节码的好处就是 省去了每次加载和解析 PHP 脚本的开销。

PHP 5.5.0 及后续版本中已经绑定了 OPcache 扩展。 对于 PHP 5.2，5.3 和 5.4 版本可以使用 » PECL 扩展中的 OPcache 库。

### 参考资料

- https://www.php.net/manual/zh/intro.opcache.php
- [PHP Opcache工作原理](https://zhuanlan.zhihu.com/p/75869838)