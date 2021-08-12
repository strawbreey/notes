---
title: "Sql Join"
date: 2021-08-12T15:10:46+08:00
draft: false
---

### SQL JOIN 

SQL join 用于根据两个或多个表中的列之间的关系，从这些表中查询数据


```sql
SELECT Persons.LastName, Persons.FirstName, Orders.OrderNo
FROM Persons, Orders
WHERE Persons.Id_P = Orders.Id_P 
```

### INNER JOIN

在表中存在至少一个匹配时，INNER JOIN 关键字返回行。

```sql
SELECT Persons.LastName, Persons.FirstName, Orders.OrderNo
FROM Persons
INNER JOIN Orders
ON Persons.Id_P = Orders.Id_P
ORDER BY Persons.LastName
```

### SQL LEFT JOIN 关键字

LEFT JOIN 关键字会从左表 (table_name1) 那里返回所有的行，即使在右表 (table_name2) 中没有匹配的行。

```sql
SELECT column_name(s)
FROM table_name1
LEFT JOIN table_name2 
ON table_name1.column_name=table_name2.column_name
```

### SQL RIGHT JOIN 关键字

RIGHT JOIN 关键字会右表 (table_name2) 那里返回所有的行，即使在左表 (table_name1) 中没有匹配的行。

```sql
SELECT column_name(s)
FROM table_name1
RIGHT JOIN table_name2 
ON table_name1.column_name=table_name2.column_name
```

### SQL FULL JOIN 关键字

只要其中某个表存在匹配，FULL JOIN 关键字就会返回行。

```sql
SELECT column_name(s)
FROM table_name1
FULL JOIN table_name2 
ON table_name1.column_name=table_name2.column_name
```