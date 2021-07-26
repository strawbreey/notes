---
title: "Yii Db Query"
date: 2021-07-23T10:31:58+08:00
draft: false
---

查询构建器建立在 Database Access Objects 基础之上，可让你创建 程序化的、DBMS无关的SQL语句。相比于原生的SQL语句，查询构建器可以帮你 写出可读性更强的SQL相关的代码，并生成安全性更强的SQL语句。

使用查询构建器通常包含以下两个步骤：

1. 创建一个 yii\db\Query 对象来代表一条 SELECT SQL 语句的不同子句（例如 SELECT, FROM）。
2. 执行 yii\db\Query 的一个查询方法（例如：all()）从数据库当中检索数据。

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

### 创建查询

#### select()

select() 方法用来指定 SQL 语句当中的 SELECT 子句。 你可以像下面的例子一样使用一个数组或者字符串来定义需要查询的字段。当 SQL 语句 是由查询对象生成的时候，被查询的字段名称将会自动的被引号括起来。

```php
$query->select(['id', 'email']);
// 等同于：
$query->select('id, email');

$query->select(['user.id AS user_id', 'email']);
// 等同于：
$query->select('user.id AS user_id, email');
$query->select(['user_id' => 'user.id', 'email']);
```

#### from() 

from() 方法指定了 SQL 语句当中的 FROM 子句

#### where()

where() 方法定义了 SQL 语句当中的 WHERE 子句。 你可以使用如下四种格式来定义 WHERE 条件：

字符串格式，例如：'status=1'
哈希格式，例如： ['status' => 1, 'type' => 2]
操作符格式，例如：['like', 'name', 'test']
对象格式，例如：new LikeCondition('name', 'LIKE', 'test')


#### orderBy()

orderBy() 方法是用来指定 SQL 语句当中的 ORDER BY 子句的。


#### groupBy()

groupBy() 方法是用来指定 SQL 语句当中的 GROUP BY 片断的。


#### having()

having() 方法是用来指定 SQL 语句当中的 HAVING 子句。它带有一个条件， 和 where() 中指定条件的方法一样

#### limit() 和 offset()
limit() 和 offset() 是用来指定 SQL 语句当中 的 LIMIT 和 OFFSET 子句的


#### join()


join() 是用来指定 SQL 语句当中的 JOIN 子句的。

#### union()

union() 方法是用来指定 SQL 语句当中的 UNION 子句的

### 查询方法 

- all()：将返回一个由行组成的数组，每一行是一个由名称和值构成的关联数组（译者注：省略键的数组称为索引数组）。
- one()：返回结果集的第一行。
- column()：返回结果集的第一列。
- scalar()：返回结果集的第一行第一列的标量值。
- exists()：返回一个表示该查询是否包结果集的值。
- count()：返回 COUNT 查询的结果。