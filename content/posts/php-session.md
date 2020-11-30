---
title: "Php Session"
date: 2020-10-12T10:56:18+08:00
draft: false
tags: ['php']

---

会话机制(Session）在PHP 中用于保持用户连续访问Web应用时的相关数据，有助于创建高度定制化的程序、增加站点的吸引力。

会话支持在 PHP 中是在并发访问时由一个方法来保存某些数据.从而使你能够构建更多的定制程序 从而提高你的 web 网站的吸引力.

一个访问者访问你的 web 网站将被分配一个唯一的 id, 就是所谓的会话 id. 这个 id 可以存储在用户端的一个 cookie 中，也可以通过 URL 进行传递

会话支持允许你将请求中的数据保存在超全局数组$_SESSION中. 当一个访问者访问你的网站，PHP 将自动检查(如果 session.auto_start 被设置为 1）或者在你要求下检查(明确通过 session_start() 或者隐式通过 session_register()) 当前会话 id 是否是先前发送的请求创建. 如果是这种情况， 那么先前保存的环境将被重建.

```php

<?php
    // Get the private context
    session_name('Private');
    session_start();
    $private_id = session_id();
    $b = $_SESSION['pr_key'];
    session_write_close();
   
    // Get the global context
    session_name('Global');
    session_id('TEST');
    session_start();
   
    $a = $_SESSION['key'];
    session_write_close();

    // Work & modify the global & private context (be ware of changing the global context!)
?>
```

### 基本用法

通过为每个独立用户分配唯一的会话 ID，可以实现针对不同用户分别存储数据的功能。 会话通常被用来在多个页面请求之间保存及共享信息。 一般来说，会话 ID 通过 cookie 的方式发送到浏览器，并且在服务器端也是通过会话 ID 来取回会话中的数据。 如果请求中不包含会话 ID 信息，那么 PHP 就会创建一个新的会话，并为新创建的会话分配新的 ID。

会话的工作流程很简单。当开始一个会话时，PHP 会尝试从请求中查找会话 ID （通常通过会话 cookie）， 如果请求中不包含会话 ID 信息，PHP 就会创建一个新的会话。 会话开始之后，PHP 就会将会话中的数据设置到 $_SESSION 变量中。 当 PHP 停止的时候，它会自动读取 $_SESSION 中的内容，并将其进行序列化， 然后发送给会话保存管理器来进行保存。


```php
session_start();
if (!isset($_SESSION['count'])) {
  $_SESSION['count'] = 0;
} else {
  $_SESSION['count']++;
}
```


### Session 函数

session_abort — Discard session array changes and finish session
session_cache_expire — 返回当前缓存的到期时间
session_cache_limiter — 读取/设置缓存限制器
session_commit — session_write_close 的别名
session_create_id — Create new session id
session_decode — 解码会话数据
session_destroy — 销毁一个会话中的全部数据
session_encode — 将当前会话数据编码为一个字符串
session_gc — Perform session data garbage collection
session_get_cookie_params — 获取会话 cookie 参数
session_id — 获取/设置当前会话 ID
session_is_registered — 检查变量是否在会话中已经注册
session_module_name — 获取/设置会话模块名称
session_name — 读取/设置会话名称
session_regenerate_id — 使用新生成的会话 ID 更新现有会话 ID
session_register_shutdown — 关闭会话
session_register — Register one or more global variables with the current session
session_reset — Re-initialize session array with original values
session_save_path — 读取/设置当前会话的保存路径
session_set_cookie_params — 设置会话 cookie 参数
session_set_save_handler — 设置用户自定义会话存储函数
session_start — 启动新会话或者重用现有会话
session_status — 返回当前会话状态
session_unregister — Unregister a global variable from the current session
session_unset — 释放所有的会话变量
session_write_close — Write session data and end session