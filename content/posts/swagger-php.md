---
title: "Swagger Php"
date: 2020-09-27T16:24:23+08:00
draft: false
---

Generate interactive (交互式) OpenAPI documentation for your RESTful API using doctrine annotations (注释).

Swagger 是一个简单但功能强大的 API 表达工具。它具有地球上最大的 API 工具生态系统，数以千计的开发人员，使用几乎所有的现代编程语言，都在支持和使用 Swagger。使用 Swagger 生成 API，我们可以得到交互式文档，自动生成代码的 SDK 以及 API 的发现特性等。

现在，Swagger 已经帮助包括 Apigee, Getty 图像，Intuit, LivingSocial, McKesson, 微软，Morningstar 和 PayPal 等世界知名企业建立起了一套基于 RESTful API 的完美服务系统。

### OpenAPI 规范

OpenAPI 规范是 Linux 基金会的一个项目，试图通过定义一种用来描述 API 格式或 API 定义的语言，来规范 RESTful 服务开发过程。OpenAPI 规范帮助我们描述一个 API 的基本信息，比如：

- 有关该 API 的一般性描述
- 可用路径（/ 资源）
- 在每个路径上的可用操作（获取 / 提交...）
- 每个操作的输入 / 输出格式

目前 V2.0 版本的 OpenAPI 规范（也就是 SwaggerV2.0 规范）已经发布并开源在 github 上。该文档写的非常好，结构清晰，方便随时查阅。关于规范的学习和理解，本文最后还有个彩蛋。

### 安装

在项目中使用 composer 安装 swagger-php 

```php
composer require zircote/swagger-php
```

使用php脚本动态生成文档

```php
<?php
require("vendor/autoload.php");
$openapi = \OpenApi\scan('/path/to/project');
header('Content-Type: application/x-yaml');
echo $openapi->toYaml();
```

使用命令行动态生成脚本

```bash
./vendor/bin/openapi --help

# 为了从任何地方使用cli，请在全局安装swagger-php并将~/.composer/vendor/bin目录添加到您环境中的PATH中。
composer global require zircote/swagger-php
```

使用变量

使用CLI时，您需要使用以下--bootstrap选项将php文件包含在常量中：

constants.php
```php
define("API_HOST", ($env === "production") ? "example.com" : "localhost");
```

```bash
openapi --bootstrap constants.php
```

注解位置

您不应将所有注释都放在一个大的@OA \ OpenApi（）注释块中，而应将它们分散在整个代码库中。swagger-php将扫描您的项目并将所有注释合并为一个@OA \ OpenApi注释。

swagger-php提供的最大好处是文档与实现API的代码非常接近



### 参考链接 

- [如何编写基于 Swagger-PHP 的 API 文档
](https://learnku.com/laravel/t/7430/how-to-write-api-documents-based-on-swagger-php)

- [Swagger-PHP v3.x](https://zircote.github.io/swagger-php/#links)