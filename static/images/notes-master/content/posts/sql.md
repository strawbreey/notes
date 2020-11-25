---
title: "Sql"
date: 2020-09-10T16:47:25+08:00
draft: false
---

### 数据库

数据库是一以某种方式存储的数据集合

选择数据库 crashcourse

```sql
use crashcourse;
```

查看数据库信息

```sql
show DATABASES; // 返回可用数据库的列表， 数据库、表、列、用户、权限等的信息
```

显示数据库内表的列表
```shell
show TABLES; // 显示当前数据库内的表

show status  // 显示服务器状态信息

show create database / show create table #显示创建特定数据库/表
show GRANTS // 显示
```


### 4. 查询数据(SELECT) 

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

### 5.排序数据 (ORDER By)

```sql
SELECT prod_name FROM products;

SELECt prod_name FROM products ORDER BY prod_name; -- 对prod_name以字母顺序排序

SELECT prod_name, prod_price, prod_name FROM products ORDER BY prod_price, prod_name; -- 按多列排序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price, prod_name DESC; --按指定方向排序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price DESC, prod_name; -- DESC/ASC 降序/升序

SELECT prod_id, prod_price, prod_name FROM products ORDER BY prod_price DESC LIMIT 1; -- 限制一个
```

### 6.过滤数据

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

### 7. 过滤数据

```sql
SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1003 AND prod_price <= 10; -- AND

SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1002 OR vend_id = 1003 -- OR

SELECT prod_id, prod_price, prod_name FROM products WHERE vend_id = 1002 OR vend_id = 1003 AND prod_price >= 10; -- AND 的优先级大于 OR

SELECT prod_id, prod_price, prod_name FROM products WHERE (vend_id = 1002 OR vend_id = 1003) AND prod_price >= 10;

SELECT prod_name, prod_price FROM products WHERE vend_id IN (1002, 1003) ORDER BY prod_name; -- IN 操作符指定条件范围

SELECT prod_name, prod_price FROM products WHERE vend_id = 1002 OR vend_id = 1000 ORDER BY prod_name; -- in 操作符一般比OR 操作符执行的更快

SELECT prod_name, prod_price FROM products WHERE vend_id NOT IN (1002, 1003) ORDER BY prod_name; -- NOT 否定他之后的所有条件
```

### 8. 通配符过滤

百分号(%) 通配符： 表示任何字符出现任意次数

```sql 
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE 'jet%'; -- 检索所有jet开头的词

SELECT prod_id, prod_name FROm prodUcts WHERE prod_name LIKE '%jet%'; -- 检索所有含有jet的词

SELECT prod_id, prod_name FROm prodUcts WHERE prod_name LIKE '%jet'; -- 检索所有含有jet结尾的词

 ```

下划线(_) 通配符： 只能匹配单个字符，而不是多个字符

```sql 
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE '__ inch teddy bear'; -- 检索所有jet开头的词

-- 备注: 与%匹配多个字符不同, _ 总是刚好匹配一个字符
```

方括号([]) 通配符 用来指定一个字符集

```sql  
SELECT prod_id, prod_name FROM products WHERE prod_name LIKE '__ inch teddy bear'; -- 检索所有jet开头的词

-- 备注: 与%匹配多个字符不同, _ 总是刚好匹配一个字符
```

 注意: SQL的通配符很有用,但是这种功能需要消耗更长的时间处理
 - 不要过度使用通配符
 - 在使用通配符时，尽量别放在搜素开始处
 - 主要通配符的位置


### 10. 创建计算字段

```SQL
SELECT vend_name + '(' + vend_countray + ')' FROM Vendors ORDER BY vend_name;
SELECT vend_name || '(' ||  vend_countray || ')' FROM Vendors ORDER BY vend_name;
SELECT RTRIM(vend_name) + '(' + RTRIM(vend_contry) + ')' FROM Vendors ORDER BY vend_name; -- TRIM 去除两遍空格 LTRIM LTRIM 左右空格
SELECT prod_id, quantily, item_price FROM OrderItem WHERE order_num = 2008;
SELECT prod_id, quantily, item_price, quantily * item_price AS expanded_prce FROM OrderItems WHERE order_num = 2008; -- 算术表达式 + - * /
```

### 12. 汇总数据

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

### 13 分组数据


```sql
SELECT COUNT(*) AS num_prods FROM Products WHERE vend_id = 'DLL01' num_prods;

SELECT vend_id, COUNT(*) AS num_prods FROM Products GROUP BY vend_id; -- 创建分组

SELECT cust_id, COUNT(*) AS orders FROM Orders GROUP BY cust_id HAVING COUNT(*) >= 2;

SELECT order_num, COUNT(*) AS items FROM OrderItems GROUP BY order_num HAVING COUNT(*) >=3

SELECT order_num, COUNT(*) AS itmes FROM OrderItems GROUP BY order_num HAVING COUNT(*) >= 3 ORDER BY items, order_num;

```

### 14. 子查询

MySQL 4.1引入了对子查询的支持，所以要想使用本章描述的SQL，必须使用MySQL 4.1或更高级的版本。

问题:  现在，假如需要列出订购物品TNT2的所有客户，应该怎样检索？下面列出具体的步骤。

- (1) 检索包含物品TNT2的所有订单的编号。
- (2) 检索具有前一步骤列出的订单编号的所有客户的ID。
- (3) 检索前一步骤返回的所有客户ID的客户信息。


```sql
select order_num from orderitems where prod_id = 'TNT2'

select coust_id from orders order_num IN (20005, 20007)

-- => 合并成为子查询

select cust_id from orders where order_num in (
       select order_num from orderitems where prod_id = 'TNT2'
)

-- 多重子查询
select cust_name, cust_contact from customers where cust in (
       select cust_id from orders where order_num in (
              select order_num from orderitems where prod_id = 'TNT2'
       )
)

```

### 15. 联结表

SQL最强大的功能之一就是能在数据检索查询的执行中联结（join）表。联结是利用SQL的SELECT能执行的最重要的操作，很好地理解联结及其语法是学习SQL的一个极为重要的组成部分。

```sql
-- 创建联结

select vend_name, prod_name, prod_price 
       from vendors, products
       where vendors.vend_id = products.vend_id
       order by vend_name, prodname


-- 在联结两个表时，你实际上做的是将第一个表中的每一行与第二个表中的每一行配对。WHERE子句作为过滤条件，它只包含那些匹配给定条件（这里是联结条件）的行。没有WHERE子句，第一个表中的每个行将与第二个表中的每个行配对，而不管它们逻辑上是否可以配在一起。

-- 由没有联结条件的表关系返回的结果为笛卡儿积。检索出的行的数目将是第一个表中的行数乘以第二个表中的行数。

select vend_name, prod_name, prod_price 
       from vendors, products
       order by vend_name, prod_name


-- => 等同于 内部联结

select vend_name, prod_name, prod_price 
from venfors inner join products 
on vendors.vend_id = products.vend_id

--  ANSI SQL规范首选INNER JOIN语法

-- 联结多个表
select prod_name, vend_name, prod_price, quantity
from orderitems, products, vendors,
where products.vend_id = vendors.vend_id
     and orderitems.prod_id = products.prod_id
     and order_num = 20005

-- 性能考虑： MySQL在运行时关联指定的每个表以处理联结。这种处理可能是非常耗费资源的，因此应该仔细，不要联结不必要的表。联结的表越多，性能下降越厉害。


```

### 16. 高级联结

```sql
-- 使用别名
select cust_name, cust_contact
from customers as c, order as o, orderitmes as oi
where c.cust_id = o.cust_id
       and oi.order_num = o.order_num
       and prod_id = 'TNT2'

-- 自联结
select prod_id, prod_name 
from products
where vend_id = (
       select vend_id from products 
       where prod_id = 'DTNTR'
)

-- 联结

select p1.prod_id, p1.prod_name
from products as p1, products as p2
where p1.vend_id = p2.vend_id
    and p2.prod_id = 'DTNTR'

-- 自联结通常作为外部语句用来替代从相同表中检索数据时使用的子查询语句。虽然最终的结果是相同的，但有时候处理联结远比处理子查询快得多

-- 自然联结
-- 事实上，迄今为止我们建立的每个内部联结都是自然联结，很可能我们永远都不会用到不是自然联结

--  内联结是做交集，外连接是做并集
-- 外部联结
select customers.cust_id, order.order_num, 
from customers inner join order
on customer.cust_id = orders.cust_id

-- 为了检索所有客户，包括那些没有订单的客户，可如下进行：
select customer.cust_id, order.order_num
from customers left outer join orders
on cunstomers.cust_id = order.cust_id
```

在总结关于联结的这两章前，有必要汇总一下关于联结及其使用的某些要点。

❑ 注意所使用的联结类型。一般我们使用内部联结，但使用外部联结也是有效的。

❑ 保证使用正确的联结条件，否则将返回不正确的数据。

❑ 应该总是提供联结条件，否则会得出笛卡儿积。

❑ 在一个联结中可以包含多个表，甚至对于每个联结可以采用不同的联结类型。虽然这样做是合法的，一般也很有用，但应该在一起测试它们前，分别测试每个联结。这将使故障排除更为简单

### 17. 组合查询

创建UNION涉及编写多条SELECT语句
```sql
select vend_id, prod_id, prod_price 
from products
where prod_price <=5

select vend_id. prod_id, prod_price
from products
where vend_id in (1001, 1002)

-- => 组合
select vend_id, prod_id, prod_price
from products
where prod_price <=5
union
select vend_id, prod_id, prod_price
from products
where vend_id in (1001, 1002)

-- => or
select vend_id, prod_id, prod_price
from products
where prod_price <=5 or vend_id_in (1001, 1002)

-- 在这个简单的例子中，使用UNION可能比使用WHERE子句更为复杂。但对于更复杂的过滤条件，或者从多个表（而不是单个表）中检索数据的情形，使用UNION可能会使处理更简单。

-- 返回全部的结果
select vend_id, prod_id, prod_price
from products
where prod_price <=5
union all
select vend_id, prod_id, prod_price
from products
where vend_id in (1001, 1002)

-- 在用UNION组合查询时，只能使用一条ORDER BY子句，它必须出现在最后一条SELECT语句之后。对于结果集，不存在用一种方式排序一部分，而又用另一种方式排序另一部分的情况，因此不允许使用多条ORDER BY子句。
select vend_id, prod_id, prod_price
from products
where prod_price <=5
union all
select vend_id, prod_id, prod_price
from products
where vend_id in (1001, 1002)
order by vend_id, prod_price
```


### 18.全文本搜索

MySQL支持几种基本的数据库引擎。并非所有的引擎都支持本书所描述的全文本搜索。两个最常使用的引擎为MyISAM和InnoDB，前者支持全文本搜索，而后者不支持。

性能考虑: 

❑ 性能——通配符和正则表达式匹配通常要求MySQL尝试匹配表中所有行（而且这些搜索极少使用表索引）。因此，由于被搜索行数不断增加，这些搜索可能非常耗时。

❑ 明确控制——使用通配符和正则表达式匹配，很难（而且并不总是能）明确地控制匹配什么和不匹配什么。例如，指定一个词必须匹配，一个词必须不匹配，而一个词仅在第一个词确实匹配的情况下才可以匹配或者才可以不匹配。

❑ 智能化的结果——虽然基于通配符和正则表达式的搜索提供了非常灵活的搜索，但它们都不能提供一种智能化的选择结果的方法。例如，一个特殊词的搜索将会返回包含该词的所有行，而不区分包含单个匹配的行和包含多个匹配的行（按照可能是更好的匹配来排列它们）。类似，一个特殊词的搜索将不会找出不包含该词但包含其他相关词的行。

### 19. 插入数据

```sql
insert info customers values(
       null,
       'pep',
       '100',
       'Los',
       'CA',
       '90046',
       'USA',
       'NULL',
       'NULL'
)

-- 以上语法很简单，但并不安全，应该尽量避免使用

insert info customers (
       cust_name,
       cust_address,
       cust_city,
       cust_state,
       cust_zip,
       cust_country,
       cust_contact,
       cust_email
) values(
       null,
       'pep',
       '100',
       'Los',
       'CA',
       '90046',
       'USA',
       'NULL',
       'NULL'
)

-- 性能考虑：数据库经常被多个客户访问，对处理什么请求以及用什么次序处理进行管理是MySQL的任务。INSERT操作可能很耗时（特别是有很多索引需要更新时），而且它可能降低等待处理的SELECT语句的性能。

insert low_priority info customers (
       cust_name,
       cust_address,
       cust_city,
       cust_state,
       cust_zip,
       cust_country,
       cust_contact,
       cust_email
) values(
       null,
       'pep',
       '100',
       'Los',
       'CA',
       '90046',
       'USA',
       'NULL',
       'NULL'
)
-- low_priority 降低 insert 语句的优先级

-- 插入多个行
-- 可以使用多条INSERT语句，甚至一次提交它们，每条语句用一个分号结束
-- 此技术可以提高数据库处理的性能，因为MySQL用单条INSERT语句处理多个插入比使用多条INSERT语句快。
```