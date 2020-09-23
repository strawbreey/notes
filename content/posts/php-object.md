---
title: "Php Object"
date: 2020-09-23T11:36:50+08:00
draft: true
---

#### 对象初始化

```php
class foo
{
    function do_foo()
    {
        echo "Doing foo."; 
    }
}

$bar = new foo;
$bar->do_foo();

```

#### 转换为对象

如果将一个对象转换成对象，它将不会有任何变化。如果其它任何类型的值被转换成对象，将会创建一个内置类 stdClass 的实例。如果该值为 NULL，则新的实例为空。 array 转换成 object 将使键名成为属性名并具有相对应的值。注意：在这个例子里， 使用 PHP 7.2.0 之前的版本，数字键只能通过迭代访问。

```php

$obj = (object) array('1' => 'foo');
var_dump(isset($obj->{'1'})); // PHP 7.2.0 后输出 'bool(true)'，之前版本会输出 'bool(false)' 
var_dump(key($obj)); // PHP 7.2.0 后输出 'string(1) "1"'，之前版本输出  'int(1)' 
```