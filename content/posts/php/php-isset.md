---
title: "Php Isset"
date: 2021-09-26T18:12:40+08:00
draft: false
---

isset(mixed $var, mixed $... = ?): bool

检测变量是否设置，并且不是 null。

如果已经使用 unset() 释放了一个变量之后，它将不再是 isset()。若使用 isset() 测试一个被设置成 null 的变量，将返回 false。同时要注意的是 null 字符（"\0"）并不等同于 PHP 的 null 常量。

如果一次传入多个参数，那么 isset() 只有在全部参数都以被设置时返回 true 计算过程从左至右，中途遇到没有设置的变量时就会立即停止。

```php

$var = '';

// 结果为 TRUE，所以后边的文本将被打印出来。
if (isset($var)) {
    echo "This var is set so I will print.";
}

// 在后边的例子中，我们将使用 var_dump 输出 isset() 的返回值。
// the return value of isset().

$a = "test";
$b = "anothertest";

var_dump(isset($a));     // TRUE
var_dump(isset($a, $b)); // TRUE

unset ($a);

var_dump(isset($a));     // FALSE
var_dump(isset($a, $b)); // FALSE

$foo = NULL;
var_dump(isset($foo));   // FALSE


$a = array ('test' => 1, 'hello' => NULL, 'pie' => array('a' => 'apple'));

var_dump(isset($a['test']));            // TRUE
var_dump(isset($a['foo']));             // FALSE
var_dump(isset($a['hello']));           // FALSE

// 键 'hello' 的值等于 NULL，所以被认为是未置值的。
// 如果想检测 NULL 键值，可以试试下边的方法。 
var_dump(array_key_exists('hello', $a)); // TRUE

// Checking deeper array values
var_dump(isset($a['pie']['a']));        // TRUE
var_dump(isset($a['pie']['b']));        // FALSE
var_dump(isset($a['cake']['a']['b']));  // FALSE

```


### 参考

empty() - 检查一个变量是否为空
__isset()
unset() - 释放给定的变量
defined() - 检查某个名称的常量是否存在
the type comparison tables
array_key_exists() - 检查数组里是否有指定的键名或索引
is_null() - 检测变量是否为 null
错误控制 @ 运算符。