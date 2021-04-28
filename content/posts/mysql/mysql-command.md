---
title: "Mysql Command"
date: 2020-09-01T16:19:22+08:00
draft: false
---

```shell
Can't connect to local MySQL server through socket '/var/lib/mysql/mysql.sock' 
```


# 导出数据
mysqldump -u root -p ieg_waibao_test > test_db.sql;

## 登录mysql
mysql -u root -p


mysql -h 服务器ip地址 -P 3306 -u root -p