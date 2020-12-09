---
title: "Php Openssl"
date: 2020-12-07T16:40:49+08:00
draft: false
---

本扩展使用 » OpenSSL 库来对称/非对称加解密，以及 PBKDF2、 PKCS7、 PKCS12、 X509 和其他加密操作。除此之外还提供了 TLS 流的实现。

要使用 Open SSL 函数，你必须安装» OpenSSL 库。 PHP 5 要求OpenSSL >= 0.9.6。 然而PHP5后续版本会有些编译上的问题，所以最好是使用OpenSSL >= 0.9.8(PHP7要求的最低版本) 。 其他版本(PHP >= 7.1.0) 要求 OpenSSL >= 1.0.1.

要使用 PHP 的 OpenSSL 支持，你应该使用--with-openssl[=DIR]参数来编译PHP.

OpenSSL 库还在运行时对正常操作有额外的要求。最明显的是，OpenSSL需要访问随机或伪随机数生成器; 在大多数 Unix 和类 Unix 平台上(包括linux),意味着它必须要访问 /dev/urandom 或者 /dev/random 设备。

在OpenSSL模块中有三种资源类型。第一种是一个 pkey(公钥或私钥)标识符，第二种是一个X509证书标识符，第三种是 CSR (证书签名请求) 标识符。