---
title: "Php Array Helper"
date: 2020-09-24T15:45:24+08:00
draft: false
---


### 获取值 getValue()

用原生PHP从一个对象、数组、或者包含这两者的一个复杂数据结构中获取数据是非常繁琐的。 你首先得使用 isset 检查 key 是否存在, 然后如果存在你就获取它，如果不存在， 则提供一个默认返回值：

```php
class User
{
    public $name = 'Alex';
}

$array = [
    'foo' => [
        'bar' => new User(),
    ]
];

$value = isset($array['foo']['bar']->name) ? $array['foo']['bar']->name : null;

// or

$value = ArrayHelper::getValue($array, 'foo.bar.name');


// 方法的第一个参数是我们从哪里获取值。第二个参数指定了如何获取数据， 它可以是下述几种类型中的一个：

// 数组键名或者欲从中取值的对象的属性名称；
// 以点号分割的数组键名或者对象属性名称组成的字符串，上例中使用的参数类型就是该类型；
// 返回一个值的回调函数。

$fullName = ArrayHelper::getValue($user, function ($user, $defaultValue) {
    return $user->firstName . ' ' . $user->lastName;
});

// 第三个可选的参数如果没有给定值，则默认为 null，如下例所示：
$username = ArrayHelper::getValue($comment, 'user.username', 'Unknown');
```


### 设置值 setValue()

```php
$array = [
    'key' => [
        'in' => ['k' => 'value']
    ]
];

ArrayHelper::setValue($array, 'key.in', ['arr' => 'val']);
// 在 `$array` 中写入值的路径可以被指定为一个数组
ArrayHelper::setValue($array, ['key', 'in'], ['arr' => 'val']);

//  =>  结果
[
    'key' => [
        'in' => ['arr' => 'val']
    ]
]

// 如果路径包含一个不存在的键，它将被创建

// 如果 `$array['key']['in']['arr0']` 不为空，则该值将被添加到数组中
ArrayHelper::setValue($array, 'key.in.arr0.arr1', 'val');

// 如果你想完全覆盖值 `$array['key']['in']['arr0']`
ArrayHelper::setValue($array, 'key.in.arr0', ['arr1' => 'val']);

// =>
[
    'key' => [
        'in' => [
            'k' => 'value',
            'arr0' => ['arr1' => 'val']
        ]
    ]
]
```

### 删除值 remove()

如果你想获得一个值，然后立即从数组中删除它，你可以使用 remove 方法：

```php
$array = ['type' => 'A', 'options' => [1, 2]];
$type = ArrayHelper::remove($array, 'type');
```

### 检查键名是否存在 keyExists

ArrayHelper::keyExists 工作原理和 array_key_exists 差不多，除了 它还可支持大小写不敏感的键名比较

```php
$data1 = [
    'userName' => 'Alex',
];

$data2 = [
    'username' => 'Carsten',
];

if (!ArrayHelper::keyExists('username', $data1, false) || !ArrayHelper::keyExists('username', $data2, false)) {
    echo "Please provide username.";
}
```


### 检索列  getColumn()

通常你要从多行数据或者多个对象构成的数组中获取某列的值，一个普通的例子是获取 id 值列表。

```php
$data = [
    ['id' => '123', 'data' => 'abc'],
    ['id' => '345', 'data' => 'def'],
];

$ids = array_column($data, 'id');
// or
$ids = ArrayHelper::getColumn($data, 'id');

// 如果需要额外的转换或者取值的方法比较复杂， 第二参数可以指定一个匿名函数：
$result = ArrayHelper::getColumn($array, function ($element) {
    return $element['id'];
});
```

### 重建数组索引 index()
```php
$array = [
    ['id' => '123', 'data' => 'abc', 'device' => 'laptop'],
    ['id' => '345', 'data' => 'def', 'device' => 'tablet'],
    ['id' => '345', 'data' => 'hgi', 'device' => 'smartphone'],
];
$result = ArrayHelper::index($array, 'id');

// =>
[
    '123' => ['id' => '123', 'data' => 'abc', 'device' => 'laptop'],
    '345' => ['id' => '345', 'data' => 'hgi', 'device' => 'smartphone']
    // 原始数组的第二个元素由于相同的 ID 而被最后一个元素覆盖
]

// 第二个参数换成匿名函数
$result = ArrayHelper::index($array, function ($element) {
    return $element['id'];
});
```

### 建立哈希表 map()

为了从一个多维数组或者一个对象数组中建立一个映射表（键值对），你可以使用 map 方法。$from 和 $to 参数分别指定了欲构建的映射表的键名和属性名。 根据需要，你可以按照一个分组字段 $group 将映射表进行分组，例如

```php
$array = [
    ['id' => '123', 'name' => 'aaa', 'class' => 'x'],
    ['id' => '124', 'name' => 'bbb', 'class' => 'x'],
    ['id' => '345', 'name' => 'ccc', 'class' => 'y'],
];

$result = ArrayHelper::map($array, 'id', 'name');

// =>
// [
//     '123' => 'aaa',
//     '124' => 'bbb',
//     '345' => 'ccc',
// ]

$result = ArrayHelper::map($array, 'id', 'name', 'class');
```

### 多维排序

multisort 方法可用来对嵌套数组或者对象数组进行排序，可按一到多个键名排序，比如

```php
$data = [
    ['age' => 30, 'name' => 'Alexander'],
    ['age' => 30, 'name' => 'Brian'],
    ['age' => 19, 'name' => 'Barney'],
];
ArrayHelper::multisort($data, ['age', 'name'], [SORT_ASC, SORT_DESC]);

// => 
[
    ['age' => 19, 'name' => 'Barney'],
    ['age' => 30, 'name' => 'Brian'],
    ['age' => 30, 'name' => 'Alexander'],
];

// 第二个参数指定排序的键名，如果是单键名的话可以是字符串，如果是多键名则是一个数组， 或者是如下例所示的一个匿名函数：
ArrayHelper::multisort($data, function($item) {
    return isset($item['age']) ? ['age', 'name'] : 'name';
});
// 第三个参数表示增降顺序。单键排序时，它可以是 SORT_ASC 或者 SORT_DESC 之一。如果是按多个键名排序，你可以用一个数组为 各个键指定不同的顺序。

// 最后一个参数（译者注：第四个参数）是PHP的排序标识（sort flag），可使用的值和调用 PHP sort() 函数时传递的值一样。


```

### 检测数组类型

想知道一个数组是索引数组还是联合数组很方便，这有个例子：

```php
// 不指定键名的数组
$indexed = ['Qiang', 'Paul'];
echo ArrayHelper::isIndexed($indexed);

// 所有键名都是字符串
$associative = ['framework' => 'Yii', 'version' => '2.0'];
echo ArrayHelper::isAssociative($associative);
```

### HTML 编码和解码值
```html
$encoded = ArrayHelper::htmlEncode($data);
$decoded = ArrayHelper::htmlDecode($data);
```

### 合并数组（Merging Arrays）

```php
$array1 = [
    'name' => 'Yii',
    'version' => '1.1',
    'ids' => [
        1,
    ],
    'validDomains' => [
        'example.com',
        'www.example.com',
    ],
    'emails' => [
        'admin' => 'admin@example.com',
        'dev' => 'dev@example.com',
    ],
];

$array2 = [
    'version' => '2.0',
    'ids' => [
        2,
    ],
    'validDomains' => new \yii\helpers\ReplaceArrayValue([
        'yiiframework.com',
        'www.yiiframework.com',
    ]),
    'emails' => [
        'dev' => new \yii\helpers\UnsetArrayValue(),
    ],
];

$result = ArrayHelper::merge($array1, $array2);

// => 

[
    'name' => 'Yii',
    'version' => '2.0',
    'ids' => [
        1,
        2,
    ],
    'validDomains' => [
        'yiiframework.com',
        'www.yiiframework.com',
    ],
    'emails' => [
        'admin' => 'admin@example.com',
    ],
]
```

### 对象转换为数组

```php
$posts = Post::find()->limit(10)->all();
$data = ArrayHelper::toArray($posts, [
    'app\models\Post' => [
        'id',
        'title',
        // the key name in array result => property name
        'createTime' => 'created_at',
        // the key name in array result => anonymous function
        'length' => function ($post) {
            return strlen($post->content);
        },
    ],
]);
```

### 测试阵列

```php
/ true
ArrayHelper::isIn('a', ['a']);
// true
ArrayHelper::isIn('a', new(ArrayObject['a']));

// true 
ArrayHelper::isSubset(new(ArrayObject['a', 'c']), new(ArrayObject['a', 'b', 'c'])

```


### Test

1. 二维数组去重
```php
array_values(ArrayHelper::index($array, key))
```

2. 数组合并并去重
```php
  array_values(ArrayHelper::index(array_merge($a, $b), key))
```

3. 数组交集

```php
// [1,2,3] [2,3,4] => [2,3]
array_values(array_intersect_key(ArrayHelper::index($a, key), ArrayHelper::index($b, key)))
```

4. 数组并集

```php
// [1,2,3] [2,3,4] => [1,2,3,4]
array_merge($a, $b)
```

5. 数组差集

```php
// [1,2,3,4] [3,4,5,6] => [1,2]
array_values(array_diff_key(ArrayHelper::index($a, key), ArrayHelper::index($b, key)))

// [1,2,3,4] [3,4,5,6] => [5.6]
array_values(array_diff_key(ArrayHelper::index($b, key), ArrayHelper::index($a, key)))
```

6. 二维数组排序

```

```

## 参考文档

- [github yii](https://github.com/yiisoft/yii2/blob/master/framework/helpers/BaseArrayHelper.php)
- [helper-array](https://www.yiiframework.com/doc/guide/2.0/zh-cn/helper-array)