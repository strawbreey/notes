---
title: "Swagger Php Start"
date: 2021-03-17T11:07:00+08:00
draft: false
---

## 安装yii 

- 使用composer
- 下载归档文件

### 安装composer

https://getcomposer.org/doc/

- linux / Mac OS X

```
curl -sS https://getcomposer.org/installer | php
mv composer.phar /usr/local/bin/composer
```

- Window 

[Composer-Setup.exe](https://getcomposer.org/Composer-Setup.exe)

 composer 常用命令
```bash
# 更新composer
composer self-update
```

### composer 设置代理

1. 修改composer 的全局配置文件

```
composer config -g repo.packagist composer https://packagist.phpcomposer.com
```

2.  修改当前项目的 composer.json 配置文件

```
composer config repo.packagist composer https://packagist.phpcomposer.com
```

3. 手动添加代理

```json
"repositories": {
    "packagist": {
        "type": "composer",
        "url": "https://packagist.phpcomposer.com"
    }
}
```

4. 解除镜象

composer config -g --unset repos.packagist

> [Packagist / Composer中国全量镜像](https://pkg.phpcomposer.com/)

安装yii

```
composer create-project --prefer-dist yiisoft/yii2-app-basic basic
```





