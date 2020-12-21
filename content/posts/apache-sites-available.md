---
title: "Apache Sites Available"
date: 2020-12-18T15:10:19+08:00
draft: false
---

由于php的url需要重写, 需要用到apache的mod_rewrite.so模块，用.htaccess文件去掉index.php优化SEO。

需要修改apache的配置文件，就是这个配置文件啊。跟老的不一样，他的把原来一个httpd.conf 拆分成几个文件，od_rewrite.so模块在mods-available文件夹里，mods-enabled做一个链接指向它。然后再apache2.conf 里 IncludeOptional mods-enabled/*.load，

sites-enabled只是sites-available的一个软连接

这是为了服务器的配置方便，具体怎么方便这里就不细说了。

一般来说是配置sites-available里面的文件。

apace2的配置文件通常位于/etc/apaches下面。

Linux下 Apache的配置文件是 /etc/apache2/apache2.conf，Apache在启动时会自动读取这个文件的配置信息。而其他的一些配置文件，如 httpd.conf等，则是通过Include指令包含进来。

在apache2.conf里有sites-enabled目录，而在 /etc/apache2下还有一个sites-available目录，其实，这里面才是真正的配置文件，而sites- enabled目录存放的只是一些指向这里的文件的符号链接，可以通过ls -al /etc/apache2/sites-enabled/来查看。

所以，如果apache上配置了多个虚拟主机，每个虚拟主机的配置文件都放在 sites-available下，那么对于虚拟主机的停用、启用就非常方便了：当在sites-enabled下建立一个指向某个虚拟主机配置文件的链接时，就启用了它；如果要关闭某个虚拟主机的话，只需删除相应的链接即可，无需去改配置文件，还是很方便的。

