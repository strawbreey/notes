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