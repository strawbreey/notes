---
title: "Yii Db"
date: 2020-11-11T19:09:39+08:00
draft: false
---

Yii 包含了一个建立在 PHP PDO 之上的数据访问层 (DAO： Database Access Objects)。DAO为不同的数据库提供了一套统一的API。 其中 ActiveRecord 提供了数据库与模型(MVC 中的 M,Model) 的交互，QueryBuilder 用于创建动态的查询语句。 DAO提供了简单高效的SQL查询，可以用在与数据库交互的各个地方.

使用 Yii DAO 时，你主要需要处理纯 SQL 语句和 PHP 数组。因此，这是访问数据库最高效的方法。 然而，因为不同数据库之间的 SQL 语法往往是不同的， 使用 Yii DAO 也意味着你必须花些额外的工夫来创建一个”数据库无关“的应用。

### 连接数据库

```php
// 连接数据库
$db = new yii\db\Connection([
    'dsn' => 'mysql:host=localhost;dbname=example',
    'username' => 'root',
    'password' => '',
    'charset' => 'utf8',
]);

return [
    'components' => [
        'db' => [
            'class' => 'yii\db\Connection',
            'dsn' => 'mysql:host=localhost;dbname=example',
            'username' => 'root',
            'password' => '',
            'charset' => 'utf8',
        ],
    ],
];
```

### SQL 查询

```php
// 返回多行. 每行都是列名和值的关联数组.
// 如果该查询没有结果则返回空数组
$posts = Yii::$app->db->createCommand('SELECT * FROM post')
            ->queryAll();

// 返回一行 (第一行)
// 如果该查询没有结果则返回 false
$post = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=1')
           ->queryOne();

// 返回一列 (第一列)
// 如果该查询没有结果则返回空数组
$titles = Yii::$app->db->createCommand('SELECT title FROM post')
             ->queryColumn();

// 返回一个标量值
// 如果该查询没有结果则返回 false
$count = Yii::$app->db->createCommand('SELECT COUNT(*) FROM post')
             ->queryScalar();
```

当使用带参数的 SQL 来创建数据库命令时， 你几乎总是应该使用绑定参数的方法来防止 SQL 注入攻击

bindValue()：绑定一个参数值
bindValues()：在一次调用中绑定多个参数值
bindParam()：与 bindValue() 相似，但是也支持绑定参数引用

```php
$post = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=:id AND status=:status')
           ->bindValue(':id', $_GET['id'])
           ->bindValue(':status', 1)
           ->queryOne();

$params = [':id' => $_GET['id'], ':status' => 1];

$post = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=:id AND status=:status')
           ->bindValues($params)
           ->queryOne();
           
$post = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=:id AND status=:status', $params)
           ->queryOne();
```

绑定参数是通过 预处理语句 实现的。 除了防止 SQL 注入攻击，它也可以通过一次预处理 SQL 语句， 使用不同参数多次执行，来提升性能。例如：

```php
$command = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=:id');

$post1 = $command->bindValue(':id', 1)->queryOne();
$post2 = $command->bindValue(':id', 2)->queryOne();
```

因为 bindParam() 支持通过引用来绑定参数， 上述代码也可以像下面这样写：

```php
$command = Yii::$app->db->createCommand('SELECT * FROM post WHERE id=:id')
              ->bindParam(':id', $id);

$id = 1;
$post1 = $command->queryOne();

$id = 2;
$post2 = $command->queryOne();
```

执行非查询语句

上面部分中介绍的 queryXyz() 方法都处理的是从数据库返回数据的查询语句。对于那些不取回数据的语句， 你应该调用的是 yii\db\Command::execute() 方法。例如，

```php
Yii::$app->db->createCommand('UPDATE post SET status=1 WHERE id=1')
   ->execute();
```

对于 INSERT, UPDATE 和 DELETE 语句，不再需要写纯SQL语句了， 你可以直接调用 insert()、update()、delete()， 来构建相应的 SQL 语句。这些方法将正确地引用表和列名称以及绑定参数值。例如,

```php
Yii::$app->db->createCommand()->insert('user', [
    'name' => 'Sam',
    'age' => 30,
])->execute();

// UPDATE (table name, column values, condition)
Yii::$app->db->createCommand()->update('user', ['status' => 1], 'age > 30')->execute();

// DELETE (table name, condition)
Yii::$app->db->createCommand()->delete('user', 'status = 0')->execute();
```

你也可以调用 batchInsert() 来一次插入多行， 这比一次插入一行要高效得多：

```php

// table name, column names, column values
Yii::$app->db->createCommand()->batchInsert('user', ['name', 'age'], [
    ['Tom', 30],
    ['Jane', 20],
    ['Linda', 25],
])->execute();
```

另一个有用的方法是 upsert()。Upsert 是一种原子操作，如果它们不存在（匹配唯一约束），则将行插入到数据库表中， 或者在它们执行时更新它们：

```php
Yii::$app->db->createCommand()->upsert('pages', [
    'name' => 'Front page',
    'url' => 'http://example.com/', // url is unique
    'visits' => 0,
], [
    'visits' => new \yii\db\Expression('visits + 1'),
], $params)->execute();

```

执行事务

```php
Yii::$app->db->transaction(function($db) {
    $db->createCommand($sql1)->execute();
    $db->createCommand($sql2)->execute();
    // ... executing other SQL statements ...
});

// 上述代码等价于下面的代码，但是下面的代码给予了你对于错误处理代码的更多掌控：

$db = Yii::$app->db;
$transaction = $db->beginTransaction();
try {
    $db->createCommand($sql1)->execute();
    $db->createCommand($sql2)->execute();
    // ... executing other SQL statements ...
    
    $transaction->commit();
} catch(\Exception $e) {
    $transaction->rollBack();
    throw $e;
} catch(\Throwable $e) {
    $transaction->rollBack();
    throw $e;
}
```

嵌套事务

```php
Yii::$app->db->transaction(function ($db) {
    // outer transaction
    
    $db->transaction(function ($db) {
        // inner transaction
    });
});

// 上述代码等价于下面的代码，但是下面的代码给予了你对于错误处理代码的更多掌控：


$db = Yii::$app->db;
$outerTransaction = $db->beginTransaction();
try {
    $db->createCommand($sql1)->execute();

    $innerTransaction = $db->beginTransaction();
    try {
        $db->createCommand($sql2)->execute();
        $innerTransaction->commit();
    } catch (\Exception $e) {
        $innerTransaction->rollBack();
        throw $e;
    } catch (\Throwable $e) {
        $innerTransaction->rollBack();
        throw $e;
    }

    $outerTransaction->commit();
} catch (\Exception $e) {
    $outerTransaction->rollBack();
    throw $e;
} catch (\Throwable $e) {
    $outerTransaction->rollBack();
    throw $e;
}

```

操纵数据库模式

Yii DAO 提供了一套完整的方法来让你操纵数据库模式， 如创建表、从表中删除一列，等等。这些方法罗列如下：

- createTable()：创建一张表
- renameTable()：重命名一张表
- dropTable()：删除一张表
- truncateTable()：删除一张表中的所有行
- addColumn()：增加一列
- renameColumn()：重命名一列
- dropColumn()：删除一列
- alterColumn()：修改一列
- addPrimaryKey()：增加主键
- dropPrimaryKey()：删除主键
- addForeignKey()：增加一个外键
- dropForeignKey()：删除一个外键
- createIndex()：增加一个索引
- dropIndex()：删除一个索引
这些方法可以如下地使用：

```php
Yii::$app->db->createCommand()->createTable('post', [
    'id' => 'pk',
    'title' => 'string',
    'text' => 'text',
]);
```
查询构建器

查询构建器建立在 Database Access Objects 基础之上，可让你创建 程序化的、DBMS无关的SQL语句。相比于原生的SQL语句，查询构建器可以帮你 写出可读性更强的SQL相关的代码，并生成安全性更强的SQL语句。

使用查询构建器通常包含以下两个步骤：

- 创建一个 yii\db\Query 对象来代表一条 SELECT SQL 语句的不同子句（例如 SELECT, FROM）。
- 执行 yii\db\Query 的一个查询方法（例如：all()）从数据库当中检索数据。

```php
$rows = (new \yii\db\Query())
    ->select(['id', 'email'])
    ->from('user')
    ->where(['last_name' => 'Smith'])
    ->limit(10)
    ->all();
```

=> 

```sql
SELECT `id`, `email` 
FROM `user`
WHERE `last_name` = :last_name
LIMIT 10

```

创建查询

select() 

```php
$query->select(['id', 'email']);
// 等同于：
$query->select('id, email');

$query->select(['user.id AS user_id', 'email']);
// 等同于：
$query->select('user.id AS user_id, email');

$query->select(['user_id' => 'user.id', 'email']);

$query->select(["CONCAT(first_name, ' ', last_name) AS full_name", 'email']); 

$subQuery = (new Query())->select('COUNT(*)')->from('user');

// SELECT `id`, (SELECT COUNT(*) FROM `user`) AS `count` FROM `post`
$query = (new Query())->select(['id', 'count' => $subQuery])->from('post');

$query->select('user_id')->distinct();

$query->select(['id', 'username'])
    ->addSelect(['email']);
```

from()

```php
$query->from('user');

$query->from(['public.user u', 'public.post p']);
// 等同于：
$query->from('public.user u, public.post p');

$query->from(['u' => 'public.user', 'p' => 'public.post']);

$subQuery = (new Query())->select('id')->from('user')->where('status=1');
// SELECT * FROM (SELECT `id` FROM `user` WHERE status=1) u 
$query->from(['u' => $subQuery]);

```