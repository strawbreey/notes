---
title: "Php Proxy"
date: 2021-04-08T19:50:16+08:00
draft: false
---

composer create-project --prefer-dist yiisoft/yii2-app-basic yii-ccxt
You are running Composer with SSL/TLS protection disabled.
Creating a "yiisoft/yii2-app-basic" project at "./yii-ccxt"


  [Composer\Downloader\TransportException]
  curl error 60 while downloading https://packagist.phpcomposer.com/packages.json: SSL certificate problem: unable to get local issuer certificate


  composer config -g repo.packagist composer https://packagist.phpcomposer.com 

You are running Composer with SSL/TLS protection disabled.

v_guancli@v_guancli-PC0 MINGW64 ~/workspace
$ composer config -g -- disable-tls false
You are now running Composer with SSL/TLS protection enabled.

v_guancli@v_guancli-PC0 MINGW64 ~/workspace
$ composer config -g repo.packagist composer https://packagist.phpcomposer.com 

$ composer create-project --prefer-dist yiisoft/yii2-app-basic yii-ccxt
Creating a "yiisoft/yii2-app-basic" project at "./yii-ccxt"


  [InvalidArgumentException]
  Could not find package yiisoft/yii2-app-basic with stability stable.