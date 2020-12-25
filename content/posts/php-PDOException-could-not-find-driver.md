---
title: "Php PDOException Could Not Find Driver"
date: 2020-12-24T15:24:42+08:00
draft: false
---

新配了个环境，访问数据库发现抱着个错：
```
PDOException “could not find driver”
```
经排查是pdo拓展没有开启。在php.ini里面解开注释即可。

解放注释之后，发现路径保存

修改php路径，之后可行


### 参考资料

- [php-warning-php-startup-unable-to-load-dynamic-library](https://stackoverflow.com/questions/5282264/php-warning-php-startup-unable-to-load-dynamic-library)