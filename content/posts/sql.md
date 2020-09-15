---
title: "Sql"
date: 2020-09-10T16:47:25+08:00
draft: false
---

### 数据库
数据库是一以某种方式存储的数据集合

使用crashcourse

```sql
use crashcourse;
```

显示数据库
```sql
show DATABASES; // 返回可用数据库的列表
```

显示数据库内表的列表
```shell
show TABLES; // 显示当前数据库内的表

show status  // 显示数据库的状态信息

show create database
show create table #显示创建特定数据库/表
show GRANTS // 显示
```


### 检索数据(SELECT) 
```sql
SELECT prod_name FROM products;  -- 检索一列数据
SELECT prod_id, prod_name, prod_price FROM products; -- 检索多列数据
SELECT * FROM products; -- 检索所有数据
SELECT vend_id FROM products; -- 检索vend—id 列
SELECT DISTINCT vend_id FROM products; -- 检索vend_id列, 并只返回不同的值
SELECT prod_name FROM products LIMIT 5; -- 检索prod_name, 并只返回5条数据
SELECT prod_name FROM products LIMIT 5, 5; -- 检索prod_name, 并只返回5条数据
SELECT product.prod_name FROM crashcourse.products; -- = SELECT prod_name FROM products;
```

### 排序数据 (ORDER By)
```sql
SELECT prod_name FROM products;

SELECt prod_name FROM products ORDER BY prod_name; -- 对prod_name以字母顺序排序

SELECT prod_name, prod_price, prod_name FROM products ORDER BY prod_price, prod_name; -- 按多列排序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price, prod_name DESC; --按指定方向排序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price DESC, prod_name; -- DESC/ASC 降序/升序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price DESC LIMIT 1; -- 限制一个
```

### 过滤数据1
```sql
-- where 操作符
SELECT prod_name, prod_price, prod_id FROM  products WHERE prod_price = 2.5  -- 只返回值为2.5的值

SELECT prod_name, prod_price FROM products WHERE prod_name = 'fuses' --mysql 执行匹配时默认不区分大小写

SELECT prod_name, prod_price FROM products WHERE prod_price < 10

SELECT prod_name, prod_price FROm products WHERE vend_id <> 1003; -- 不是供应商1003制作的

SELECT prod_name. prod_price FROM producst WHERE vend_id != 1003 -- 同上

-- BETWEEN

SELECT prod_name, prod_price FROM prod_price BETWEEN 5 AND 10; --范围值检查

SELECT prod_name, prod_price FROM prod_price WHERE prod_price IS NULL; -- 返回没有价格的所有产品
SELECT cust_id FROM WHERE cust_email IS NULL;  
```

### 过滤数据2
```sql
SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1003 AND prod_price <= 10; -- AND

SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1002 OR vend_id = 1003 -- OR

SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1002 OR vend_id = 1003 AND prod_price >= 10; -- AND 的优先级大于 OR

SELECT prod_id, prod_price, prod_name FROM products WHERE (vend_id = 1002 OR vend_id = 1003) AND prod_price >= 10;

SELECT prod_name, prod_price FROM products WHERE vend_id IN (1002, 1003) ORDER BY prod_name; -- IN 操作符指定条件范围

SELECT prod_name, prod_price FROM products WHERE vend_id = 1002 OR vend_id = 1000 ORDER BY prod_name; -- in 操作符一般比OR 操作符执行的更快

SELECT prod_name, prod_price FROM products WHERE vend_id NOT IN (1002, 1003) ORDER BY prod_name; -- NOT 否定他之后的所有条件
```

### 通配符过滤

#### 百分号(%) 通配符： 表示任何字符出现任意次数
```sql 
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE 'jet%'; -- 检索所有jet开头的词
SELECT prod_id, prod_name FROm prodUcts WHERE prod_name LIKE '%jet%'; -- 检索所有含有jet的词
SELECT prod_id, prod_name FROm prodUcts WHERE prod_name LIKE '%jet'; -- 检索所有含有jet结尾的词

 ```
