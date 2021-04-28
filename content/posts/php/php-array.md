---
title: "Php Array"
date: 2020-09-23T10:09:47+08:00
draft: false
tags: ['php']

---

An array in PHP is actually an ordered map. A map is a type that associates values to keys. This type is optimized for several different uses; it can be treated as an array, list (vector), hash table (an implementation of a map), dictionary, collection, stack, queue, and probably more. As array values can be other arrays, trees and multidimensional arrays are also possible.

PHP 中的数组实际上是一个有序映射。映射是一种把 values 关联到 keys 的类型。此类型在很多方面做了优化，因此可以把它当成真正的数组，或列表（向量），散列表（是映射的一种实现），字典，集合，栈，队列以及更多可能性。由于数组元素的值也可以是另一个数组，树形结构和多维数组也是允许的。


### 定义数组

可以用 array() 语言结构来新建一个数组。它接受任意数量用逗号分隔的 键（key） => 值（value）对。
```php
array(  key =>  value
     , ...
     )
// 键（key）可是是一个整数 integer 或字符串 string
// 值（value）可以是任意类型的值
```

最后一个数组单元之后的逗号可以省略。通常用于单行数组定义中，例如常用 array(1, 2) 而不是 array(1, 2, )。对多行数组定义通常保留最后一个逗号，这样要添加一个新单元时更方便。

自 5.4 起可以使用短数组定义语法，用 [] 替代 array()。


```php
$array = array(
    "foo" => "bar",
    "bar" => "foo",
);

// 自 PHP 5.4 起
$array = [
    "foo" => "bar",
    "bar" => "foo",
];
```
key 可以是 integer 或者 string。value 可以是任意类型。

此外 `key` 会有如下的强制转换：

- 包含有合法整型值的字符串会被转换为整型。例如键名 "8" 实际会被储存为 8。但是 "08" 则不会强制转换，因为其不是一个合法的十进制数值。
- 浮点数也会被转换为整型，意味着其小数部分会被舍去。例如键名 8.7 实际会被储存为 8。
- 布尔值也会被转换成整型。即键名 true 实际会被储存为 1 而键名 false 会被储存为 0。
- Null 会被转换为空字符串，即键名 null 实际会被储存为 ""。
- 数组和对象不能被用为键名。坚持这么做会导致警告：Illegal offset type。


#### 类型强制与覆盖示例

如果在数组定义中多个单元都使用了同一个键名，则只使用了最后一个，之前的都被覆盖了。

```php
$array = array(
    1    => "a",
    "1"  => "b",
    1.5  => "c",
    true => "d",
);

var_dump($array); 

// array(1) {
//   [1]=>
//   string(1) "d"
// }

```


####  混合 integer 和 string 键名

```php
$array = array(
    "foo" => "bar",
    "bar" => "foo",
    100   => -100,
    -100  => 100,
);
// array(4) {
//   ["foo"]=>
//   string(3) "bar"
//   ["bar"]=>
//   string(3) "foo"
//   [100]=>
//   int(-100)
//   [-100]=>
//   int(100)
// }
```
PHP 数组可以同时含有 integer 和 string 类型的键名，因为 PHP 实际并不区分索引数组和关联数组。

#### 没有键名的索引数组

```php

$array = array("foo", "bar", "hallo", "world");

// array(4) {
//   [0]=>
//   string(3) "foo"
//   [1]=>
//   string(3) "bar"
//   [2]=>
//   string(5) "hallo"
//   [3]=>
//   string(5) "world"
// }
```

#### 仅对部分单元指定键名

```php
$array = array(
         "a",
         "b",
    6 => "c",
         "d",
);

// array(4) {
//   [0]=>
//   string(1) "a"
//   [1]=>
//   string(1) "b"
//   [6]=>
//   string(1) "c"
//   [7]=>
//   string(1) "d"
// }
```

#### 方括号语法访问数组单元

数组单元可以通过 array[key] 语法来访问。

```php
$array = array(
    "foo" => "bar",
    42    => 24,
    "multi" => array(
         "dimensional" => array(
             "array" => "foo"
         )
    )
);

var_dump($array["foo"]); // string(3) "bar"
var_dump($array[42]); // init(24)
var_dump($array["multi"]["dimensional"]["array"]); // string(3) "f66"

```

方括号和花括号可以互换使用来访问数组单元（例如 $array[42] 和 $array{42} 在上例中效果相同）。

#### 数组解引用

```php
function getArray() {
    return array(1, 2, 3);
}

// on PHP 5.4
$secondElement = getArray()[1];

// previously
$tmp = getArray();
$secondElement = $tmp[1];

// or
list(, $secondElement) = getArray();
```

> 试图访问一个未定义的数组键名与访问任何未定义变量一样：会导致 E_NOTICE 级别错误信息，其结果为 NULL。

#### 用方括号的语法新建/修改

可以通过明示地设定其中的值来修改一个已有数组。

```php
$arr[key] = value;
$arr[] = value;
// key 可以是 integer 或 string
// value 可以是任意类型的值`
```

要修改某个值，通过其键名给该单元赋一个新值。要删除某键值对，对其调用 unset() 函数。

```php

$arr = array(5 => 1, 12 => 2);

$arr[] = 56;    // This is the same as $arr[13] = 56;
                // at this point of the script

$arr["x"] = 42; // This adds a new element to
                // the array with key "x"
                
unset($arr[5]); // This removes the element from the array

unset($arr);    // This deletes the whole array

```
> 如上所述，如果给出方括号但没有指定键名，则取当前最大整数索引值，新的键名将是该值加上 1（但是最小为 0）。如果当前还没有整数索引，则键名将为 0。


### 常用Array函数

#### unset 函数

unset() 函数允许删除数组中的某个键。但要注意数组将不会重建索引。如果需要删除后重建索引，可以用 array_values() 函数。

```php
$a = array(1 => 'one', 2 => 'two', 3 => 'three');
unset($a[2]);
/* will produce an array that would have been defined as
   $a = array(1 => 'one', 3 => 'three');
   and NOT
   $a = array(1 => 'one', 2 =>'three');
*/

$b = array_values($a);
// Now $b is array(0 => 'one', 1 =>'three')
```

#### foreach

### 转换为数组

对于任意 integer，float，string，boolean 和 resource 类型，如果将一个值转换为数组，将得到一个仅有一个元素的数组，其下标为 0，该元素即为此标量的值。换句话说，(array)$scalarValue 与 array($scalarValue) 完全一样。

如果一个 object 类型转换为 array，则结果为一个数组，其单元为该对象的属性。键名将为成员变量名，不过有几点例外：整数属性不可访问；私有变量前会加上类名作前缀；保护变量前会加上一个 '*' 做前缀。这些前缀的前后都各有一个 NULL 字符。这会导致一些不可预知的行为：

```php

class A {
    private $A; // This will become '\0A\0A'
}

class B extends A {
    private $A; // This will become '\0B\0A'
    public $AA; // This will become 'AA'
}

var_dump((array) new B());

```


```php
foreach ($colors as &$color) {
    $color = strtoupper($color); // 大写
}
unset($color); /* ensure that following writes to
$color will not modify the last array element */

// Workaround for older versions
foreach ($colors as $key => $color) {
   
```

#### 比较

可以用array_diff() 和 数组运算符来比较数组


#### 在循环中改变单元

```PHP
// PHP 5
foreach ($colors as &$color) {
    $color = strtoupper($color);
}
unset($color); /* ensure that following writes to
$color will not modify the last array element */

// Workaround for older versions
foreach ($colors as $key => $color) {
    $colors[$key] = strtoupper($color);
}

print_r($colors);
```

#### 下标从 1 开始的数组

```php
$firstquarter  = array(1 => 'January', 'February', 'March');
print_r($firstquarter);

// Array 
// (
//     [1] => 'January'
//     [2] => 'February'
//     [3] => 'March'
// )


```

#### 填充数组

```php
$handle = opendir('.');
while (false !== ($file = readdir($handle))) {
    $files[] = $file;
}
closedir($handle); 

```

#### 数组排序

因为数组中的值可以为任意值，也可是另一个数组。这样可以产生递归或多维数组。
s
```php

sort($files);
print_r($files);
```


### 生词
```
ordered map 有序映射
associat 关联
optimized optimized
treated as 当作 像...一样对待
dictionary 字典
collection 集合
multidimensional 多维数组
hash table 散列表
implementation 实现
```