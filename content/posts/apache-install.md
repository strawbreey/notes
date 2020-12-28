---
title: "Apache Install"
date: 2020-12-28T12:12:45+08:00
draft: false
---

CentOS 7下Apache的安装

CentOS下使用yum安装Apache极为方便，只需要在终端键入以下命令即可

```bash
# 1.安装Apache
sudo apachectl configtest

# 2.设置服务器开机自动启动Apache
systemctl enable httpd.service

# 若要验证是否自动启动可在重启服务器后在终端键入以下命令来检测Apache是否已经启动
systemctl is-enabled httpd.service

# 如果看到了enable这样的响应，则表示Apache已经启动成功
# 3.手动启动Apache
systemctl start httpd.service 在浏览器中输入IP地址即可验证是否启动成功

# 4.手动重启Apache
systemctl restart httpd.service

# 5.手动停止Apache
systemctl stop httpd.service

# 6.安装目录介绍
Apache默认将网站的根目录指向/var/www/html
默认的主配置文件/etc/httpd/conf/httpd.conf
配置存储在的/etc/httpd/conf.d/目录
```