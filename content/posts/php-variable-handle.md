---
title: "Php Variable Handle"
date: 2020-09-23T15:36:50+08:00
draft: false
---

#### boolval — 获取变量的布尔值

```php
echo '0:        '.(boolval(0) ? 'true' : 'false')."\n";
echo '42:       '.(boolval(42) ? 'true' : 'false')."\n";
echo '0.0:      '.(boolval(0.0) ? 'true' : 'false')."\n";
echo '4.2:      '.(boolval(4.2) ? 'true' : 'false')."\n";
echo '"":       '.(boolval("") ? 'true' : 'false')."\n";
echo '"string": '.(boolval("string") ? 'true' : 'false')."\n";
echo '"0":      '.(boolval("0") ? 'true' : 'false')."\n";
echo '"1":      '.(boolval("1") ? 'true' : 'false')."\n";
echo '[1, 2]:   '.(boolval([1, 2]) ? 'true' : 'false')."\n";
echo '[]:       '.(boolval([]) ? 'true' : 'false')."\n";
echo 'stdClass: '.(boolval(new stdClass) ? 'true' : 'false')."\n";
```
#### debug_zval_dump — Dumps a string representation of an internal zend value to output
#### doubleval — floatval 的别名
#### empty — 检查一个变量是否为空

Note:

在 PHP 5.5 之前，empty() 仅支持变量；任何其他东西将会导致一个解析错误。换言之，下列代码不会生效： empty(trim($name))。 作为替代，应该使用trim($name) == false.

没有警告会产生，哪怕变量并不存在。 这意味着 empty() 本质上与 !isset($var) || $var == false 等价。

#### floatval — 获取变量的浮点值 
```php
$var = '122.34343The';
$float_value_of_var = floatval ($var);
print $float_value_of_var; // 打印出 122.34343 类型转换
```
#### get_defined_vars — 返回由所有已定义变量所组成的数组
```php
$b = array(1,1,2,3,5,8);

$arr = get_defined_vars();

// 打印 $b
print_r($arr["b"]);

// 打印 PHP 解释程序的路径（如果 PHP 作为 CGI 使用的话）
// 例如：/usr/local/bin/php
echo $arr["_"];

// 打印命令行参数（如果有的话）
print_r($arr["argv"]);

// 打印所有服务器变量
print_r($arr["_SERVER"]);

// 打印变量数组的所有可用键值
print_r(array_keys(get_defined_vars()));
```
#### get_resource_type — 返回资源（resource）类型
#### gettype — 获取变量的类型
#### import_request_variables — 将 GET／POST／Cookie 变量导入到全局作用域中
#### intval — 获取变量的整数值
#### is_array — 检测变量是否是数组
#### is_bool — 检测变量是否是布尔型
#### is_callable — 检测参数是否为合法的可调用结构
#### is_countable — Verify that the contents of a variable is a countable value
#### is_double — is_float 的别名
#### is_float — 检测变量是否是浮点型
#### is_int — 检测变量是否是整数
#### is_integer — is_int 的别名
#### is_iterable — Verify that the contents of a variable is an iterable value
#### is_long — is_int 的别名
#### is_null — 检测变量是否为 NULL
#### is_numeric — 检测变量是否为数字或数字字符串
#### is_object — 检测变量是否是一个对象
#### is_real — is_float 的别名
#### is_resource — 检测变量是否为资源类型
#### is_scalar — 检测变量是否是一个标量
#### is_string — 检测变量是否是字符串
#### isset — 检测变量是否已设置并且非 NULL
#### print_r — 以易于理解的格式打印变量。
#### serialize — 产生一个可存储的值的表示
#### settype — 设置变量的类型
#### strval — 获取变量的字符串值
#### unserialize — 从已存储的表示中创建 PHP 的值
#### unset — 释放给定的变量
#### var_dump — 打印变量的相关信息
#### var_export — 输出或返回一个变量的字符串表示