---
title: "Mysql Transaction"
date: 2020-11-09T15:49:56+08:00
draft: true
---

事务(Transaction) 是数据区别于文件系统重要的特性之一

innoDB 的 事务完全符合ACID 的特性

- 原子性(atomicity)

- 一致性(consistency)

- 隔离性(isolation)

- 持久性(durability)


事务可由一条非常简单的SQL语句组成，也可以由一组复杂的SQL语句组成。事务是访问并更新数据库中各种数据项的一个程序执行单元。在事务中的操作，要么都做修改，要么都不做，这就是事务的目的。


事务分类

- 扁平事务

- 带有保存点的扁平事务

- 链事务

- 嵌套事务

- 分布式事务