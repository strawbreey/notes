---
title: "Php Syntax"
date: 2021-05-06T14:57:21+08:00
draft: false
---

## php 基本语法

1. php 标记

  当解析一个文件时，PHP 会寻找起始和结束标记，也就是 <?php 和 ?>，这告诉 PHP 开始和停止解析二者之间的代码

  ```php
    <?php echo 'if you want to serve PHP code in XHTML or XML documents,
      use these tags'; ?>
  ```
2. 数据类型

  PHP 支持 10 种原始数据类型, 
  
  四种标量类型：

  - bool（布尔型）

  这是最简单的类型。boolean 表达了真值，可以为 true 或 false。
  
  - int（整型）

  int 是集合 ℤ = {..., -2, -1, 0, 1, 2, ...} 中的某个数。

  - float（浮点型，也称作 double)

  浮点型（也叫浮点数 float，双精度数 double 或实数 real）可以用以下任一语法定义：
  
  - string（字符串）

  四种复合类型：

  - array（数组）

  PHP 中的数组实际上是一个有序映射。映射是一种把 values 关联到 keys 的类型。此类型在很多方面做了优化，因此可以把它当成真正的数组，或列表（向量），散列表（是映射的一种实现），字典，集合，栈，队列以及更多可能性。由于数组元素的值也可以是另一个数组，树形结构和多维数组也是允许的。
  
  - object（对象）
  
  - callable（可调用）
  
  - iterable（可迭代）
  Iterable是 PHP 7.1 中引入的一个伪类型。它接受任何 array 或实现了 Traversable 接口的对象。这些类型都能用 foreach 迭代， 也可以和 生成器 里的 yield from 一起使用。

  最后是两种特殊类型：

  - resource（资源）
  
  资源 resource 是一种特殊变量，保存了到外部资源的一个引用。资源是通过专门的函数来建立和使用的。所有这些函数及其相应资源类型见附录。

  - NULL（无类型）
  
  特殊的 null 值表示一个变量没有值。NULL 类型唯一可能的值就是 null。

  可能还会读到一些关于“双精度（double）”类型的参考。实际上 double 和 float 是相同的，由于一些历史的原因，这两个名称同时存在。


  注意:  如果想查看某个表达式的值和类型，用 var_dump() 函数。如果只是想得到一个易读懂的类型的表达方式用于调试，用 gettype() 函数。要检验某个类型，不要用 gettype()，而用 is_type 函数。

  3. 类型转化

  PHP 在变量定义中不需要（或不支持）明确的类型定义；变量类型是根据使用该变量的上下文所决定的。也就是说，如果把一个 string 值赋给变量 $var，$var 就成了一个 string。如果又把一个int 赋给 $var，那它就成了一个int。

  PHP 的自动类型转换的一个例子是乘法运算符“*”。如果任何一个操作数是float， 则所有的操作数都被当成float，结果也是float。 否则操作数会被解释为int，结果也是int。 注意这并没有改变这些操作数本身的类型； 改变的仅是这些操作数如何被求值以及表达式本身的类型。

  ```php
  $foo = "1";  // $foo 是字符串 (ASCII 49)
  $foo *= 2;   // $foo 现在是一个整数 (2)
  $foo = $foo * 1.3;  // $foo 现在是一个浮点数 (2.6)
  $foo = 5 * "10 Little Piggies"; // $foo 是整数 (50)
  $foo = 5 * "10 Small Pigs";     // $foo 是整数 (50)
  ```

  4. 强制类型转化

  PHP 中的类型强制转换和 C 中的非常像：在要转换的变量之前加上用括号括起来的目标类型。

  ```php
  (int), (integer) - 转换为整形 int
  (bool), (boolean) - 转换为布尔类型 bool
  (float), (double), (real) - 转换为浮点型 float
  (string) - 转换为字符串 string
  (array) - 转换为数组 array
  (object) - 转换为对象 object
  (unset) - 转换为 NULL
  ```

  5. 常量: 常量是一个简单值的标识符，在脚本执行期间该值不能改变。常量默认为大小写敏感。传统上常量标识符总是大写的。

  ```php
  define("FOO",     "something");
  define("FOO2",    "something else");
  define("FOO_BAR", "something more");  
  ```

  6. 魔术常量: 魔术常量，不是常量, 但不能改变

  __LINE__	    文件中的当前行号。
  __FILE__	    文件的完整路径和文件名。如果用在被包含文件中，则返回被包含的文件名。
  __DIR__	      文件所在的目录。如果用在被包括文件中，则返回被包括的文件所在的目录。它等价于 dirname(__FILE__)。除非是根目录，否则目录中名不包括末尾的斜杠。
  __FUNCTION__	当前函数的名称。匿名函数则为 {closure}。
  __CLASS__	    当前类的名称。类名包括其被声明的作用区域（例如 Foo\Bar）。注意自 PHP 5.4 起 __CLASS__ 对 trait 也起作用。当用在 trait 方法中时，__CLASS__ 是调用 trait 方法的类的名字。
  __TRAIT__	    Trait 的名字。Trait 名包括其被声明的作用区域（例如 Foo\Bar）。
  __METHOD__  	类的方法名。
  __NAMESPACE__	当前命名空间的名称。

7. 流程控制

- if if 结构是很多语言包括 PHP 在内最重要的特性之一，它允许按照条件执行代码片段。PHP 的 if 结构和 C 语言相似

- else  经常需要在满足某个条件时执行一条语句，而在不满足该条件时执行其它语句，这正是 else 的功能

- elseif 和此名称暗示的一样，是 if 和 else 的组合。和 else 一样，它延伸了 if 语句，可以在原来的 if 表达式值为 false 时执行不同语句

- while

- do-while

- for

- foreach

8. 变量： PHP 中的变量用一个美元符号后面跟变量名来表示。变量名是区分大小写的。

9. PHP 的命令行模式

  - 让php运行指定文件: php script.php / php -f script.php

  - 在命令行直接运行 PHP 代码。 

  ```shell
    php -r 'print_r(get_defined_constants());'
  ```

  - 内置web Server 
  
    PHP 5.4.0起， CLI SAPI 提供了一个内置的Web服务器。
    这个内置的Web服务器主要用于本地开发使用，不可用于线上产品环境。




