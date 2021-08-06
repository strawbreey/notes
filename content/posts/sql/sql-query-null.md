---
title: "Sql Query Null"
date: 2021-08-06T15:10:34+08:00
draft: false
---

### sql查询字段为null的值

判断使用 is null 就可以啦

```sql
-- 查询数据为空
SELECT * FROM BANK where BANK_CODE is null

-- 查询数据不为空
SELECT * FROM BANK where BANK_CODE is not null
```