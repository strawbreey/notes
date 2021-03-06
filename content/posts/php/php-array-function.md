---
title: "Php Array Function"
date: 2020-09-23T11:41:22+08:00
draft: false
tags: ['php']

---

####  array_change_key_case: 将数组所有键名全大写或者小写

array_change_key_case ( array $array [, int $case = CASE_LOWER ] ) : array

// 参数  array 需要操作的数组, case 可以在这里用两个常量，CASE_UPPER 或 CASE_LOWER(默认值)。
// 返回值 返回一个键全是小写或者全是大写的数组；如果输入值（array）不是一个数组，那么返回FALSE

```php
$input_array = array("FirSt" => 1, "SecOnd" => 4);

array_change_key_case($input_array, CASE_LOWER);

// Array
// (
//     [first] => 1
//     [second] => 4
// )
```

#### array_chunk 将一个数组分割成多个

array_chunk ( array $array , int $size [, bool $preserve_keys = false ] ) : array

将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。

- 参数

array: 需要操作的数组 size: 每个数组的单元数目 preserve_keys: 设为 TRUE，可以使 PHP 保留输入数组中原来的键名。如果你指定了 FALSE，那每个结果数组将用从零开始的新数字索引。默认值是 FALSE。

- 返回值
得到的数组是一个多维数组中的单元，其索引从零开始，每一维包含了 size 个元素。

```php
$input_array = array('a', 'b', 'c', 'd', 'e');
print_r(array_chunk($input_array, 2)); // [[a,b],[c,d],[e]]
print_r(array_chunk($input_array, 2, true)); 
// Array
// (
//     [0] => Array
//         (
//             [0] => a
//             [1] => b
//         )

//     [1] => Array
//         (
//             [2] => c
//             [3] => d
//         )

//     [2] => Array
//         (
//             [4] => e
//         )

// )

```

#### arrya_column 返回数组中指定的一列

array_column ( array $input , mixed $column_key [, mixed $index_key = null ] ) : array

array_column() 返回input数组中键值为column_key的列， 如果指定了可选参数index_key，那么input数组中的这一列的值将作为返回数组中对应值的键。

input：需要取出数组列的多维数组。 如果提供的是包含一组对象的数组，只有 public 属性会被直接取出。 为了也能取出 private 和 protected 属性，类必须实现 __get() 和 __isset() 魔术方法。

column_key： 需要返回值的列，它可以是索引数组的列索引，或者是关联数组的列的键，也可以是属性名。 也可以是NULL，此时将返回整个数组（配合index_key参数来重置数组键的时候，非常管用）

index_key: 作为返回数组的索引/键的列，它可以是该列的整数索引，或者字符串键值。

```php

$records = array(
    array(
        'id' => 2135,
        'first_name' => 'John',
        'last_name' => 'Doe',
    ),
    array(
        'id' => 3245,
        'first_name' => 'Sally',
        'last_name' => 'Smith',
    ),
    array(
        'id' => 5342,
        'first_name' => 'Jane',
        'last_name' => 'Jones',
    ),
    array(
        'id' => 5623,
        'first_name' => 'Peter',
        'last_name' => 'Doe',
    )
);
 
// 从记录中取出first names列 
$first_names = array_column($records, 'first_name');


// 从结果集中总取出last names列，用相应的id作为键值
$last_names = array_column($records, 'last_name', 'id');



// username 列是从对象获取 public 的 "username" 属性
class User
{
    public $username;

    public function __construct(string $username)
    {
        $this->username = $username;
    }
}

$users = [
    new User('user 1'),
    new User('user 2'),
    new User('user 3'),
];

array_column($users, 'username')

// 获取 username 列，从对象通过魔术方法 __get() 获取 private 的 "username" 属性。

class Person
{
    private $name;

    public function __construct(string $name)
    {
        $this->name = $name;
    }

    public function __get($prop)
    {
        return $this->$prop;
    }

    public function __isset($prop) : bool
    {
        return isset($this->$prop);
    }
}

$people = [
    new Person('Fred'),
    new Person('Jane'),
    new Person('John'),
];

print_r(array_column($people, 'name'));

```

#### array_combine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值

array_combine ( array $keys , array $values ) : array

返回一个 array，用来自 keys 数组的值作为键名，来自 values 数组的值作为相应的值。

参数：
keys:将被作为新数组的键。非法的值将会被转换为字符串类型（string） values: 将被作为 Array 的值

返回值: 返回合并的 array，如果两个数组的单元数不同则返回 FALSE。

```php
$a = array('green', 'red', 'yellow');
$b = array('avocado', 'apple', 'banana');
$c = array_combine($a, $b);

// Array
// (
//     [green]  => avocado
//     [red]    => apple
//     [yellow] => banana
// )
```

#### array_count_values 统计数组中所有的值

array_count_values ( array $array ) : array

array_count_values() 返回一个数组： 数组的键是 array 里单元的值； 数组的值是 array 单元的值出现的次数。

```php
$array = array(1, "hello", 1, "world", "hello");

array_count_values($array)

// Array
// (
//     [1] => 2
//     [hello] => 2
//     [world] => 1
// )
```

#### array_diff 计算数组的差集

array_diff ( array $array1 , array $array2 [, array $... ] ) : array

对比 array1 和其他一个或者多个数组，返回在 array1 中但是不在其他 array 里的值。

```php
// array_diff() 例子
$array1 = array("a" => "green", "red", "blue", "red");
$array2 = array("b" => "green", "yellow", "red");
$result = array_diff($array1, $array2);

// Array
// (
//     [1] => blue
// )

// array_diff() example with non-matching types
// 数组无法转换成字符串时会产生 Notice 警告
$source = [1, 2, 3, 4];
$filter = [3, 4, [5], 6];
$result = array_diff($source, $filter);













































// 而这个就可以，因为对象可以转换成字符串
class S {
  private $v;

  public function __construct(string $v) {
    $this->v = $v;
  }

  public function __toString() {
    return $this->v;
  }
}

$source = [new S('a'), new S('b'), new S('c')];
$filter = [new S('b'), new S('c'), new S('d')];

$result = array_diff($source, $filter);
```

注意本函数只检查了多维数组中的一维。当然可以用 array_diff($array1[0], $array2[0]); 检查更深的维度。


array_fill_keys — 使用指定的键和值填充数组
array_fill — 用给定的值填充数组
array_filter — 用回调函数过滤数组中的单元
array_flip — 交换数组中的键和值
array_intersect_assoc — 带索引检查计算数组的交集
array_intersect_key — 使用键名比较计算数组的交集
array_intersect_uassoc — 带索引检查计算数组的交集，用回调函数比较索引
array_intersect_ukey — 用回调函数比较键名来计算数组的交集
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
array_rand — 从数组中随机取出一个或多个单元
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
arsort — 对数组进行逆向排序并保持索引关系
asort — 对数组进行排序并保持索引关系
compact — 建立一个数组，包括变量名和它们的值
count — 计算数组中的单元数目，或对象中的属性个数
current — 返回数组中的当前单元
each — 返回数组中当前的键／值对并将数组指针向前移动一步
end — 将数组的内部指针指向最后一个单元
extract — 从数组中将变量导入到当前的符号表
in_array — 检查数组中是否存在某个值
key_exists — 别名 array_key_exists
key — 从关联数组中取得键名
krsort — 对数组按照键名逆向排序
ksort — 对数组按照键名排序
list — 把数组中的值赋给一组变量
natcasesort — 用“自然排序”算法对数组进行不区分大小写字母的排序
natsort — 用“自然排序”算法对数组排序
next — 将数组中的内部指针向前移动一位
pos — current 的别名
prev — 将数组的内部指针倒回一位
range — 根据范围创建数组，包含指定的元素
reset — 将数组的内部指针指向第一个单元
rsort — 对数组逆向排序
shuffle — 打乱数组
sizeof — count 的别名
sort — 对数组排序
uasort — 使用用户自定义的比较函数对数组中的值进行排序并保持索引关联
uksort — 使用用户自定义的比较函数对数组中的键名进行排序
usort — 使用用户自定义的比较函数对数组中的值进行排序