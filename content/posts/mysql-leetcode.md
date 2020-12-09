---
title: "Mysql Leetcode"
date: 2020-12-08T19:30:48+08:00
draft: false
---

### 编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。

```
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
```

```sql
select Email from Person group by Email having count(Email) > 1;
-- 前面提到聚合函数（count），where字句无法与聚合函数一起使用。因为where子句的运行顺序排在第二，运行到where时，表还没有被分组。

select Email from
(
  select Email, count(Email) as num from Person group by Email
) as statistic
where num > 1
```

备注: 
1. where 后不能跟聚合函数，因为where执行顺序大于聚合函数。
2. where 子句的作用是在对查询结果进行分组前，将不符合where条件的行去掉，即在分组之前过滤数据，条件中不能包含聚组函数，使用where条件显示特定的行。
3. having 子句的作用是筛选满足条件的组，即在分组之后过滤数据，条件中经常包含聚组函数，使用having 条件显示特定的组，也可以使用多个分组标准进行分组。


2. 给定一个 salary 表，如下所示，有 m = 男性 和 f = 女性 的值。交换所有的 f 和 m 值（例如，将所有 f 值更改为 m，反之亦然）。要求只使用一个更新（Update）语句，并且没有中间的临时表。

```sql
update salary 
set
   sex = case sex 
      when 'm' then 'f'
               else 'm' 
end;

update salary set sex = char(ascii('m') + ascii('f') - ascii(sex));

update salary set sex = replace("fm", sex, "")

update salary set sex = IF(sex = 'm', 'f', 'm')
```

3. 如果一个国家的面积超过 300 万平方公里，或者人口超过 2500 万，那么这个国家就是大国家。编写一个 SQL 查询，输出表中所有大国家的名称、人口和面积。

```sql
select name, population, area from World where area > 3000000 || population > 25000000


SELECT name, population, area FROM world WHERE area > 3000000 OR population > 25000000


SELECT name, population, area FROM world WHERE area > 3000000
UNION
SELECT name, population, area FROM world WHERE population > 25000000

```

4. 某城市开了一家新的电影院，吸引了很多人过来看电影。该电影院特别注意用户体验，专门有个 LED显示板做电影推荐，上面公布着影评和相关电影描述。

作为该电影院的信息部主管，您需要编写一个 SQL查询，找出所有影片描述为非 boring (不无聊) 的并且 id 为奇数 的影片，结果请按等级 rating 排列。

```sql
select id, movie, description, rating from cinema where description != 'boring' && id % 2 != 0 order by rating desc

select * from cinema where mod(id, 2) = 1 and description != 'boring' order by rating desc

select * from cinema where id%2!=0 and description not like 'boring' order by rating desc;

select * from cinema where description !='boring' and id%2 =1 order by rating desc


```
