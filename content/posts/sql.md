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

 #### 下划线(_) 通配符： 只能匹配单个字符，而不是多个字符
```sql 
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE '__ inch teddy bear'; -- 检索所有jet开头的词

-- 备注: 与%匹配多个字符不同, _ 总是刚好匹配一个字符
 ```

 #### 方括号([]) 通配符 用来指定一个字符集
```sql  
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE '__ inch teddy bear'; -- 检索所有jet开头的词

-- 备注: 与%匹配多个字符不同, _ 总是刚好匹配一个字符
 ```

 注意: SQL的通配符很有用,但是这种功能需要消耗更长的时间处理
 - 不要过度使用通配符
 - 在使用通配符时，尽量别放在搜素开始处
 - 主要通配符的位置


#### 创建计算字段

```SQL
SELECT vend_name + '(' + vend_countray + ')' FROM Vendors ORDER BY vend_name;
SELECT vend_name || '(' ||  vend_countray || ')' FROM Vendors ORDER BY vend_name;
SELECT RTRIM(vend_name) + '(' + RTRIM(vend_contry) + ')' FROM Vendors ORDER BY vend_name; -- TRIM 去除两遍空格 LTRIM LTRIM 左右空格
SELECT prod_id, quantily, item_price FROM OrderItem WHERE order_num = 2008;
SELECT prod_id, quantily, item_price, quantily * item_price AS expanded_prce FROM OrderItems WHERE order_num = 2008; -- 算术表达式 + - * /
```

### 汇总数据
```sql

-- 聚合函数
 
SELECT AVG(prod_price) AS avg_price FROM Products; -- 平均数

SELECT AVG(prod_price) AS avg_price FROM Products WHERE vend_id = 'DLL01'

SELECT COUNT(*) AS num_cust FROM Custiomers; -- 确定表中的行数 count(*) 

SELECT COUNT(cust_email) AS num_cust FROM Custiomers;

SELECT MAX(prid_price) AS max_price FROM Products;

SELECT MIN(prid_price) AS min_price FROM Products;

SELECT SUM(quantity) AS item_ordered FROM OrderItem;

SELECT SUM(item_price * quantity) AS total_price FROM OrderItem;

-- 聚合不同值

SELECT COUNT(*) As num_items, 
       MIN(prod_price) AS price_min, 
       MAX(prod_price) AS price_max, 
       AVG(prod_price) AS price_avg;
FROM Products;

```

#### 分组数据

数据分组

```SQL
SELECT COUNT(*) AS num_prods FROM Products WHERE vend_id = 'DLL01' num_prods;
-- 创建分组
SELECT vend_id, COUNT(*) AS num_prods FROM Products GROUP BY vend_id;

SELECT cust_id, COUNT(*) AS orders FROM Orders GROUP BY cust_id HAVING COUNT(*) >= 2;

SELECT order_num, COUNT(*) AS items FROM OrderItems GROUP BY order_num HAVING COUNT(*) >=3

SELECT order_num, COUNT(*) AS itmes FROM OrderItems GROUP BY order_num HAVING COUNT(*) >= 3 ORDER BY items, order_num;


```

#### 子查询