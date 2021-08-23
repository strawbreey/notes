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
array(  
    key => value,
    ...
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

### 数组 函数

array_change_key_case — 将数组中的所有键名修改为全大写或小写
array_chunk — 将一个数组分割成多个
array_column — 返回输入数组中指定列的值
array_combine — 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
array_count_values — 统计数组中所有的值
array_diff_assoc — 带索引检查计算数组的差集
array_diff_key — 使用键名比较计算数组的差集
array_diff_uassoc — 用用户提供的回调函数做索引检查来计算数组的差集
array_diff_ukey — 用回调函数对键名比较计算数组的差集
array_diff — 计算数组的差集
array_fill_keys — 使用指定的键和值填充数组
array_fill — 用给定的值填充数组
array_filter — 使用回调函数过滤数组的元素
array_flip — 交换数组中的键和值
array_intersect_assoc — 带索引检查计算数组的交集
array_intersect_key — 使用键名比较计算数组的交集
array_intersect_uassoc — 带索引检查计算数组的交集，用回调函数比较索引
array_intersect_ukey — 在键名上使用回调函数来比较计算数组的交集
array_intersect — 计算数组的交集
array_key_exists — 检查数组里是否有指定的键名或索引
array_key_first — 获取指定数组的第一个键值
array_key_last — 获取一个数组的最后一个键值
array_keys — 返回数组中部分的或所有的键名
array_map — 为数组的每个元素应用回调函数
array_merge_recursive — 递归地合并一个或多个数组
array_merge — 合并一个或多个数组
array_multisort — 对多个数组或多维数组进行排序
array_pad — 以指定长度将一个值填充进数组
array_pop — 弹出数组最后一个单元（出栈）
array_product — 计算数组中所有值的乘积
array_push — 将一个或多个单元压入数组的末尾（入栈）
array_rand — 从数组中随机取出一个或多个随机键
array_reduce — 用回调函数迭代地将数组简化为单一的值
array_replace_recursive — 使用传递的数组递归替换第一个数组的元素
array_replace — 使用传递的数组替换第一个数组的元素
array_reverse — 返回单元顺序相反的数组
array_search — 在数组中搜索给定的值，如果成功则返回首个相应的键名
array_shift — 将数组开头的单元移出数组
array_slice — 从数组中取出一段
array_splice — 去掉数组中的某一部分并用其它值取代
array_sum — 对数组中所有值求和
array_udiff_assoc — 带索引检查计算数组的差集，用回调函数比较数据
array_udiff_uassoc — 带索引检查计算数组的差集，用回调函数比较数据和索引
array_udiff — 用回调函数比较数据来计算数组的差集
array_uintersect_assoc — 带索引检查计算数组的交集，用回调函数比较数据
array_uintersect_uassoc — 带索引检查计算数组的交集，用单独的回调函数比较数据和索引
array_uintersect — 计算数组的交集，用回调函数比较数据
array_unique — 移除数组中重复的值
array_unshift — 在数组开头插入一个或多个单元
array_values — 返回数组中所有的值
array_walk_recursive — 对数组中的每个成员递归地应用用户函数
array_walk — 使用用户自定义函数对数组中的每个元素做回调处理
array — 新建一个数组
arsort — 对数组进行降向排序并保持索引关系
asort — 对数组进行升序排序并保持索引关系
compact — 建立一个数组，包括变量名和它们的值
count — 计算数组中的单元数目，或对象中的属性个数
current — 返回数组中的当前值
each — 返回数组中当前的键／值对并将数组指针向前移动一步
end — 将数组的内部指针指向最后一个单元
extract — 从数组中将变量导入到当前的符号表
in_array — 检查数组中是否存在某个值
key_exists — 别名 array_key_exists
key — 从关联数组中取得键名
krsort — 对数组按照键名逆向排序
ksort — 对数组根据键名升序排序
list — 把数组中的值赋给一组变量
natcasesort — 用“自然排序”算法对数组进行不区分大小写字母的排序
natsort — 用“自然排序”算法对数组排序
next — 将数组中的内部指针向前移动一位
pos — current 的别名
prev — 将数组的内部指针倒回一位
range — 根据范围创建数组，包含指定的元素
reset — 将数组的内部指针指向第一个单元
rsort — 对数组降序排序
shuffle — 打乱数组
sizeof — count 的别名
sort — 对数组升序排序
uasort — 使用用户自定义的比较函数对数组中的值进行排序并保持索引关联
uksort — 使用用户自定义的比较函数对数组中的键名进行排序
usort — 使用用户自定义的比较函数对数组中的值进行排序


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