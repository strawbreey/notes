---
title: "Sql Like"
date: 2021-08-23T15:48:32+08:00
draft: false
---


LIKE 操作符用于在 WHERE 子句中搜索列中的指定模式。

例0： 查找含有 特殊字符 `%` 的字段 

```sql
SELECT * FROM Persons
    WHERE City LIKE '%\%%'
```

例1：

现在，我们希望从上面的 "Persons" 表中选取居住在以 "N" 开始的城市里的人：

我们可以使用下面的 SELECT 语句：

```sql
SELECT * FROM Persons
WHERE City LIKE 'N%'
```
提示："%" 可用于定义通配符（模式中缺少的字母）。

例2：

接下来，我们希望从 "Persons" 表中选取居住在以 "g" 结尾的城市里的人：

我们可以使用下面的 SELECT 语句：

```sql
SELECT * FROM Persons
WHERE City LIKE '%g'
```

提示："%" 可用于定义通配符（模式中缺少的字母）。


例子 3
接下来，我们希望从 "Persons" 表中选取居住在包含 "lon" 的城市里的人：

我们可以使用下面的 SELECT 语句：

```sql
SELECT * FROM Persons
WHERE City LIKE '%lon%'
```

例子 4

通过使用 NOT 关键字，我们可以从 "Persons" 表中选取居住在不包含 "lon" 的城市里的人：

我们可以使用下面的 SELECT 语句：
```sql
SELECT * FROM Persons
WHERE City NOT LIKE '%lon%'
```