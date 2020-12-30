---
title: "Composer Install"
date: 2020-12-30T19:31:39+08:00
draft: false
---


1. The openssl extension is required for SSL/TLS protection but is not available. If you can not enable the openssl extension, you can disable this error, at your own risk, by setting the 'disable-tls' option to true.

```shell
composer config -g -- disable-tls true
```

2.  Your configuration does not allow connections to http://repo.packagist.org/packages.json.

```shell
extension=php_openssl.dll
```