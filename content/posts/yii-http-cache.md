---
title: "Yii Http Cache"
date: 2020-11-11T18:25:40+08:00
draft: false
---

通过配置 `yii\filters\HttpCache` 过滤器，控制器操作渲染的内容就能缓存在客户端。 

HttpCache 过滤器仅对 GET 和 HEAD 请求生效， 它能为这些请求设置三种与缓存有关的 HTTP 头。

### Last-Modified 头

Last-Modified 头使用时间戳标明页面自上次客户端缓存后是否被修改过。

通过配置 yii\filters\HttpCache::$lastModified 属性向客户端发送 Last-Modified 头。 该属性的值应该为 PHP callable 类型，返回的是页面修改时的 Unix 时间戳。 该 callable 的参数和返回值应该如下：

```php
public function behaviors()
{
    return [
        [
            'class' => 'yii\filters\HttpCache',
            'only' => ['index'],
            'lastModified' => function ($action, $params) {
                $q = new \yii\db\Query();
                return $q->from('post')->max('updated_at');
            },
        ],
    ];
}
```

上述代码表明 HTTP 缓存只在 index 操作时启用。 它会基于页面最后修改时间生成一个 Last-Modified HTTP 头。 当浏览器第一次访问 index 页时，服务器将会生成页面并发送至客户端浏览器。 之后客户端浏览器在页面没被修改期间访问该页， 服务器将不会重新生成页面，浏览器会使用之前客户端缓存下来的内容。 因此服务端渲染和内容传输都将省去。

### ETag 头

“Entity Tag”（实体标签，简称 ETag）使用一个哈希值表示页面内容。如果页面被修改过， 哈希值也会随之改变。通过对比客户端的哈希值和服务器端生成的哈希值， 浏览器就能判断页面是否被修改过，进而决定是否应该重新传输内容。

通过配置 yii\filters\HttpCache::$etagSeed 属性向客户端发送 ETag 头。 该属性的值应该为 PHP callable 类型，返回的是一段种子字符用来生成 ETag 哈希值。 该 callable 的参数和返回值应该如下：

```php
public function behaviors()
{
    return [
        [
            'class' => 'yii\filters\HttpCache',
            'only' => ['view'],
            'etagSeed' => function ($action, $params) {
                $post = $this->findModel(\Yii::$app->request->get('id'));
                return serialize([$post->title, $post->content]);
            },
        ],
    ];
}
```

上述代码表明 HTTP 缓存只在 view 操作时启用。 它会基于用户请求的标题和内容生成一个 ETag HTTP 头。 当浏览器第一次访问 view 页时，服务器将会生成页面并发送至客户端浏览器。 之后客户端浏览器标题和内容没被修改在期间访问该页，服务器将不会重新生成页面， 浏览器会使用之前客户端缓存下来的内容。 因此服务端渲染和内容传输都将省去。

### Cache-Control 头

Cache-Control 头指定了页面的常规缓存策略。 可以通过配置 yii\filters\HttpCache::$cacheControlHeader 属性发送相应的头信息。默认发送以下头：

Cache-Control: public, max-age=3600
