---
title: "Sql Insert"
date: 2020-09-23T16:24:30+08:00
draft: true
---



### SQL INSERT INTO

```sql
INSERT INTO table_name VALUES (value1, value2, value3, ...);

INSERT INTO table_name (column1, column2， column3) VALUES (value1, value2, value3, ...);

```


数据库大批量SQL插入性能优化

1. 一条SQL语句插入多条数据
2. 在事务中进行插入
3. 数据有序插入

从测试结果可以看到，合并数据+事务的方法在较小数据量时，性能提高是很明显的，数据量较大时（1千万以上），性能会急剧下降，这是由于此时数据量超过了innodb_buffer的容量，每次定位索引涉及较多的磁盘读写操作，性能下降较快。而使用合并数据+事务+有序数据的方式在数据量达到千万级以上表现依旧是良好，在数据量较大时，有序数据索引定位较为方便，不需要频繁对磁盘进行读写操作，所以可以维持较高的性能。


[insert-optimization](https://dev.mysql.com/doc/refman/5.7/en/insert-optimization.html)