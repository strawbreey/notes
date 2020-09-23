---
title: "Sql Transaction Processing"
date: 2020-09-21T17:17:43+08:00
draft: true
---

### transaction processing (事务处理)

使用事务处理，要么完全支持，要么完全不执行，来维护数据库的完整性。

```SQL
-- 开始事务
BEGIN TRANSACTION 
COMMIT TRANSACTION
```

#### ROLLBACL 回退/测回

```
DELECT FORM ORDER;
ROLLBACK;
```

##### COMMIT
一般的sql语句都是针对数据库的直接执行, 而事务处理中，提交

