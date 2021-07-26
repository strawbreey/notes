---
title: "Yii Request"
date: 2021-07-22T17:08:33+08:00
draft: false
---

要获取请求参数，你可以调用 request 组件的 get() 方法和 post() 方法。 他们分别返回 $_GET 和 $_POST 的值。例如，

```php
    $request = Yii::$app->request;

    $get = $request->get(); 
    // 等价于: $get = $_GET;

    $id = $request->get('id');   
    // 等价于: $id = isset($_GET['id']) ? $_GET['id'] : null;

    $id = $request->get('id', 1);   
    // 等价于: $id = isset($_GET['id']) ? $_GET['id'] : 1;

    $post = $request->post(); 
    // 等价于: $post = $_POST;

    $name = $request->post('name');   
    // 等价于: $name = isset($_POST['name']) ? $_POST['name'] : null;

    $name = $request->post('name', '');   
    // 等价于: $name = isset($_POST['name']) ? $_POST['name'] : ''
```

> 信息： 建议你像上面那样通过 request 组件来获取请求参数，而不是 直接访问 $_GET 和 $_POST。 这使你更容易编写测试用例，因为你可以伪造数据来创建一个模拟请求组件。


```php
$request = Yii::$app->request;

// 返回所有参数
$params = $request->bodyParams;

// 返回参数 "id"
$param = $request->getBodyParam('id');
```