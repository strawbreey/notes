---
title: "Yii Run Applications"
date: 2020-09-15T15:22:56+08:00
draft: false
---

### 应用结构 

```menu
basic/                  应用根目录
    composer.json       Composer 配置文件, 描述包信息
    config/             包含应用配置及其它配置
        console.php     控制台应用配置信息应用配置信息配置信息
        web.php         Web 应用配置信息配置信息
    commands/           包含控制台命令类
    controllers/        包含控制器类控制器类
    models/             包含模型类
    runtime/            包含
    vendor/             包含已经安装的
    views/              包含视图文件视图文件
    web/                Web 应用根目录
        assets/         包含
        index.php       应用入口文件入口文件
    yii                 Yii 控制台命令执行脚本命令执行脚本执行脚本
```

一般来说，应用中的文件可被分为两类：在 basic/web 下的和在其它目录下的。 前者可以直接通过 HTTP 访问（例如浏览器），后者不能也不应该被直接访问。

一个应用的静态结构

![application-structure](/images/application-structure.png)

每个应用都有一个入口脚本 web/index.php，这是整个应用中唯一可以访问的 PHP 脚本。 入口脚本接受一个 Web 请求并创建应用实例去处理它。 应用在它的组件辅助下解析请求， 并分派请求至 MVC 元素。视图使用小部件 去创建复杂和动态的用户界面。

### 请求生命周期

![request-lifecycle](/images/request-lifecycle.png)

- 用户向入口脚本 web/index.php 发起请求。
- 入口脚本加载应用配置并创建一个应用 实例去处理请求。
- 应用通过请求组件解析请求的 路由。
- 应用创建一个控制器实例去处理请求。
- 控制器创建一个动作实例并针对操作执行过滤器。
- 如果任何一个过滤器返回失败，则动作取消。
- 如果所有过滤器都通过，动作将被执行。
- 动作会加载一个数据模型，或许是来自数据库。
- 动作会渲染一个视图，把数据模型提供给它。
- 渲染结果返回给响应组件。
- 响应组件发送渲染结果给用户浏览器。


## 入口脚本

入口脚本是应用启动流程中的第一环， 一个应用（不管是网页应用还是控制台应用）只有一个入口脚本。 终端用户的请求通过入口脚本实例化应用并将请求转发到应用。

Web 应用的入口脚本必须放在终端用户能够访问的目录下， 通常命名为 index.php， 也可以使用 Web 服务器能定位到的其他名称。

控制台应用的入口脚本一般在应用根目录下命名为 yii（后缀为.php）， 该文件需要有执行权限， 这样用户就能通过命令 ./yii <route> [arguments] [options] 来运行控制台应用。

入口脚本主要完成以下工作：

- 定义全局常量；
- 注册 Composer 自动加载器；
- 包含 Yii 类文件；
- 加载应用配置；
- 创建一个应用实例并配置;
- 调用 yii\base\Application::run() 来处理请求。

```php
<?php

defined('YII_DEBUG') or define('YII_DEBUG', true);
defined('YII_ENV') or define('YII_ENV', 'dev');

// 注册 Composer 自动加载器
require __DIR__ . '/../vendor/autoload.php';

// 包含 Yii 类文件
require __DIR__ . '/../vendor/yiisoft/yii2/Yii.php';

// 加载应用配置
$config = require __DIR__ . '/../config/web.php';

// 创建、配置、运行一个应用
(new yii\web\Application($config))->run();
```

![application-lifecycle](/images/application-lifecycle.png)


当运行 入口脚本 处理请求时， 应用主体会经历以下生命周期:

- 入口脚本加载应用主体配置数组。
- 入口脚本创建一个应用主体实例：
- 调用 preInit() 配置几个高级别应用主体属性， 比如 basePath。
- 注册 error handler 错误处理方法。
- 配置应用主体属性。
- 调用 init() 初始化，该函数会调用 bootstrap() 运行引导启动组件。
- 入口脚本调用 yii\base\Application::run() 运行应用主体:
- 触发 EVENT_BEFORE_REQUEST 事件。
- 处理请求：解析请求 路由 和相关参数； 创建路由指定的模块、控制器和动作对应的类，并运行动作。
- 触发 EVENT_AFTER_REQUEST 事件。
- 发送响应到终端用户。
- 入口脚本接收应用主体传来的退出状态并完成请求的处理。



应用主体是服务定位器， 它部署一组提供各种不同功能的 应用组件 来处理请求。 例如，urlManager组件负责处理网页请求路由到对应的控制器。 db组件提供数据库相关服务等等。