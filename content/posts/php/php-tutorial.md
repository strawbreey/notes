---
title: "Php tutorial"
date: 2021-08-23T15:42:26+08:00
draft: false
---

# PHP 入门教程

## php 变量

PHP 变量规则：

- 变量以 $ 符号开始，后面跟着变量的名称
- 变量名必须以字母或者下划线字符开始
- 变量名只能包含字母、数字以及下划线（A-z、0-9 和 \_ ）
- 变量名不能包含空格
- 变量名是区分大小写的（$y 和 $Y 是两个不同的变量）

> php 是一门弱类型语言, 会根据变量的值，自动把变量转换为正确的数据类型。

## php 变量作用域

PHP 有四种不同的变量作用域：

- local (局部)
- global (全局)： 用于函数内访问全局变量
- static
- parameter

```php
$x=5; // 全局变量

function myTest()
{
    $y=10; // 局部变量
    echo "<p>测试函数内变量:<p>";
    echo "变量 x 为: $x";
    echo "<br>";
    echo "变量 y 为: $y";
}

myTest();

echo "<p>测试函数外变量:<p>";
echo "变量 x 为: $x";
echo "<br>";
echo "变量 y 为: $y"

$x=5;
$y=10;


//  global 关键字用于函数内访问全局变量。在函数内调用函数外定义的全局变量，我们需要在函数中的变量前加上 global 关键字：
function myTest()
{
  // 声明当前变量为全局变量
  global $x,$y;
  $y=$x+$y;
}

myTest();
echo $y; // 输出 15

// PHP 将所有全局变量存储在一个名为 $GLOBALS[index] 的数组中。 index 保存变量的名称。这个数组可以在函数内部访问，也可以直接用来更新全局变量。

function myTest()
{
  // 访问全局变量
  $GLOBALS['y']=$GLOBALS['x']+$GLOBALS['y'];
}

myTest();
echo $y;

// Static 作用域, 当一个函数完成时，它的所有变量通常都会被删除。然而，有时候您希望某个局部变量不要被删除。

function myTest()
{
    static $x=0;
    echo $x;
    $x++;
    echo PHP_EOL;    // 换行符
}

myTest(); // => 0
myTest(); // => 1
myTest(); // => 2

// 参数作用域: 参数是通过调用代码将值传递给函数的局部变量。

function myTest($x)
{
    echo $x;
}
myTest(5);
```

## echo 和 print

echo 和 print 区别:

- echo - 可以输出一个或多个字符串
- print - 只允许输出一个字符串，返回值总为 1

> 提示：echo 输出的速度比 print 快， echo 没有返回值，print 有返回值 1。

## PHP EOF(heredoc) 使用说明

PHP EOF(heredoc)是一种在命令行 shell（如 sh、csh、ksh、bash、PowerShell 和 zsh）和程序语言（像 Perl、PHP、Python 和 Ruby）里定义一个字符串的方法。

使用概述：

1. 必须后接分号，否则编译通不过。
2. EOF 可以用任意其它字符代替，只需保证结束标识与开始标识一致。
3. 结束标识必须顶格独自占一行(即必须从行首开始，前后不能衔接任何空白和字符)。
4. 开始标识可以不带引号或带单双引号，不带引号与带双引号效果一致，解释内嵌的变量和转义符号，带单引号则不解释内嵌的变量和转义符号。
5. 当内容需要内嵌引号（单引号或双引号）时，不需要加转义符，本身对单双引号转义，此处相当与 q 和 qq 的用法。

```php
echo <<<EOF
        <h1>我的第一个标题</h1>
        <p>我的第一个段落。</p>
EOF;
// 结束需要独立一行且前后不能空格

```

> 注意：
>
> 1.以 <<<EOF 开始标记开始，以 EOF 结束标记结束，结束标记必须顶头写，不能有> 缩进和空格，且在结束标记末尾要有分号 。 2.开始标记和结束标记相同，比如常用大写的 EOT、EOD、EOF 来表示，但是不只限于那几个(也可以用：JSON、HTML 等)，只要保证开始标记和结束标记不在正文中出现即可。 3.位于开始标记和结束标记之间的变量可以被正常解析，但是函数则不可以。在 heredoc 中，变量不需要用连接符 . 或 , 来拼接

## PHP 数据类型

PHP 变量存储不同的类型的数据，不同的数据类型可以做不一样的事情。

PHP 支持以下几种数据类型:

- String（字符串）
- Integer（整型）
- Float（浮点型）
- Boolean（布尔型）
- Array（数组）
- Object（对象）
- NULL（空值）
- Resource（资源类型）

```php
// PHP 资源类型 resource 是一种特殊变量，保存了到外部资源的一个引用。

// 使用 get_resource_type() 函数可以返回资源（resource）类型：

// get_resource_type(resource $handle): string

$c = mysql_connect();
echo get_resource_type($c)."\n";
// 打印：mysql link

$fp = fopen("foo","w");
echo get_resource_type($fp)."\n";
// 打印：file

$doc = new_xmldoc("1.0");
echo get_resource_type($doc->doc)."\n";
// 打印：domxml document

```

## PHP 类型比较

虽然 PHP 是弱类型语言，但也需要明白变量类型及它们的意义，因为我们经常需要对 PHP 变量进行比较，包含松散和严格比较。

松散比较：使用两个等号 == 比较，只比较值，不比较类型。
严格比较：用三个等号 === 比较，除了比较值，也比较类型。
例如，"42" 是一个字符串而 42 是一个整数。FALSE 是一个布尔值而 "FALSE" 是一个字符串。

## PHP 常量

设置 PHP 常量

设置常量，使用 define() 函数，函数语法如下：

`bool define ( string $name , mixed $value [, bool $case_insensitive = false ] )`

该函数有三个参数:

- name：必选参数，常量名称，即标志符。
- value：必选参数，常量的值。
- case_insensitive ：可选参数，如果设置为 TRUE，该常量则大小写不敏感，默认是大小写敏感的。

```php
// 区分大小写的常量名
define("GREETING", "欢迎访问 Runoob.com");
echo GREETING;    // 输出 "欢迎访问 Runoob.com"
echo greeting;   // 输出 "greeting"，但是有警告信息，表示该常量未定义

// 不区分大小写的常量名
define("GREETING", "欢迎访问 Runoob.com", true);
echo greeting;  // 输出 "欢迎访问 Runoob.com"

// 常量是全局的，常量在定义后，默认是全局变量，可以在整个运行的脚本的任何地方使用。
define("GREETING", "欢迎访问 Runoob.com");

function myTest() {
  echo GREETING;
}

myTest();    // 输出 "欢迎访问 Runoob.com"
```


## PHP 字符串变量

### PHP 并置运算符

在 PHP 中，只有一个字符串运算符(.) 用于把两个字符串值连接起来。

```php
$txt1="Hello world!";
$txt2="What a nice day!";
echo $txt1 . " " . $txt2;
```

### PHP strlen() 函数

strlen() 函数返回字符串的长度（字节数）。

下面的实例返回字符串 "Hello world!" 的长度：

```php
echo strlen("Hello world!");
```

### PHP strpos() 函数
strpos() 函数用于在字符串内查找一个字符或一段指定的文本。

如果在字符串中找到匹配，该函数会返回第一个匹配的字符位置。如果未找到匹配，则返回 FALSE。

```php
echo strpos("Hello world!","world");
```

- [字符串函数](https://www.php.net/manual/zh/ref.strings.php)


### PHP 运算符


## PHP 数组


在 PHP 中，有三种类型的数组：

- 数值数组 - 带有数字 ID 键的数组
- 关联数组 - 带有指定的键的数组，每个键关联一个值
- 多维数组 - 包含一个或多个数组的数组

PHP 数值数组

这里有两种创建数值数组的方法：

```php
// 自动分配 ID 键（ID 键总是从 0 开始）：

$cars=array("Volvo","BMW","Toyota");

// 人工分配 ID 键：
$cars[0]="Volvo";
$cars[1]="BMW";
$cars[2]="Toyota";
```

获取数组的长度 - count() 函数

count() 函数用于返回数组的长度（元素的数量）：

```php
$cars=array("Volvo","BMW","Toyota");
echo count($cars);
```

遍历数值数组

```php
$cars = array("Volvo","BMW","Toyota");
$arrLength = count($cars);
 
for($x = 0; $x < $arrLength; $x++)
{
  echo $cars[$x];
  echo "<br>";
}
```

### PHP 关联数组
关联数组是使用您分配给数组的指定的键的数组。

```php
$age=array("Peter"=>"35","Ben"=>"37","Joe"=>"43");

// or:
$age['Peter']="35";
$age['Ben']="37";
$age['Joe']="43";


// 实例
$age=array("Peter"=>"35","Ben"=>"37","Joe"=>"43");
echo "Peter is " . $age['Peter'] . " years old.";

// 遍历关联数组
$age=array("Peter"=>"35","Ben"=>"37","Joe"=>"43");
 
foreach($age as $x=>$x_value)
{
    echo "Key=" . $x . ", Value=" . $x_value;
    echo "<br>";
}
```

### 多维数组

多维数组是包含一个或多个数组的数组。

```php
// 二维数组语法格式：
array (
    array (elements...),
    array (elements...),
    ...
)
```

以上数组的元素会自动分配键值，从 0 开始：

```php
$cars = array
(
    array("Volvo",100,96),
    array("BMW",60,59),
    array("Toyota",110,100)
);


$sites = array(
  "runoob"=>array(
      "菜鸟教程",
      "http://www.runoob.com"
  ),
  "google"=>array(
      "Google 搜索",
      "http://www.google.com"
  ),
  "taobao"=>array(
      "淘宝",
      "http://www.taobao.com"
  )
);
print("<pre>"); // 格式化输出数组
print_r($sites);
print("</pre>");

echo $sites['runoob'][0] . '地址为：' . $sites['runoob'][1];

```

三维数组

```php
// 三维数组是在二维数组的基础上再嵌套一层数组，格式如下：

array (
  array (
    array (elements...),
    array (elements...),
    ...
  ),
  array (
    array (elements...),
    array (elements...),
    ...
  ),
  ...
)
```

// 创建三维数组
$myarray = array(
    array(
        array(1, 2),
        array(3, 4),
    ),
    array(
        array(5, 6),
        array(7, 8),
    ),
);
     
完整的 PHP Array 参考手册


## [todo](https://www.runoob.com/php/php-string.html)

