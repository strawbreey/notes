---
title: "Mysql Leetcode"
date: 2020-12-08T19:30:48+08:00
draft: false
---

1. 编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。

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

select * from cinema where id%2>0 and description <> 'boring' order by rating desc

select * from cinema where id % 2 != 0 and description != 'boring' order by rating desc

```



5. 表1: Person

+-------------+---------+
| 列名         | 类型     |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId 是上表主键
表2: Address

+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId 是上表主键

编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：
 
```
FirstName, LastName, City, State
```


```sql
select FirstName, LastName, City, State from Person left join Address on Person.PersonId = Address.PersonId

select A.FirstName, A.LastName, B.City, B.State from Person A left join Address B
on A.PersonId = B.PersonId;

select A.FirstName, A.LastName, B.City, B.State from Person A left join (
select distinct PersonId, City, State from Address) B
on A.PersonId = B.PersonId;

-- 因为city和state字段是对addressId的解释字段，如果city和state都相同，那addressId也应该相同，这就有主键不能重复相违背了，所以说Address表除主键外的字段是不会也不能出现重复数据的，除非表的前期设计是有问题的

```

考察sql外连接的用法。
外连接包含三个方向：

- 左外连接，left (outer) join 结果表中除了匹配行外，还包括左表有而右表中不匹配的行，对于这样的行，右表选择列置为null。

- 右外连接，right (outer) join 结果表中除了匹配行外，还包括右表有而左表中不匹配的行，对于这样的行，左表选择列置为null。

- 全连接， full (outer) join 完整外部联接返回左表和右表中的所有行。当某行在另一个表中没有匹配行时，则另一个表的选择列表列包含空值。如果表之间有匹配行，则整个结果集行包含基表的数据值。

- 内连接，(inner) join 内联接使用比较运算符根据每个表共有的列的值匹配两个表中的行。例如，检索 students和courses表中学生标识号相同的所有行。

- 交叉连接，cross join，交叉连接，实际上就是将两个表进行笛卡尔积运算，结果表的行数等于两表行数之积


6. Employee 表包含所有员工，他们的经理也属于员工。每个员工都有一个 Id，此外还有一列对应员工的经理的 Id。

+----+-------+--------+-----------+
| Id | Name  | Salary | ManagerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | NULL      |
| 4  | Max   | 90000  | NULL      |
+----+-------+--------+-----------+
给定 Employee 表，编写一个 SQL 查询，该查询可以获取收入超过他们经理的员工的姓名。在上面的表格中，Joe 是唯一一个收入超过他的经理的员工。

+----------+
| Employee |
+----------+
| Joe      |
+----------+

```sql
SELECT
    *
FROM
    Employee AS a,
    Employee AS b
WHERE
    a.ManagerId = b.Id AND a.Salary > b.Salary


SELECT
     a.NAME AS Employee
FROM Employee AS a JOIN Employee AS b
     ON a.ManagerId = b.Id
     AND a.Salary > b.Salary

select Name Employee
from Employee E1
where Salary > (select Salary
                from Employee 
                where Id = E1.ManagerId);


select Name as `Employee` from Employee a where exists (select 1 from Employee b where a.Salary>b.Salary and b.id=a.ManagerId)

```

7. 某网站包含两个表，Customers 表和 Orders 表。编写一个 SQL 查询，找出所有从不订购任何东西的客户。

Customers 表：

+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Orders 表：

+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
例如给定上述表格，你的查询应返回：

+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+

```sql

select Name as Customers
from Name
left Join Orders 
on Orders.CustomerId = Customers.Id

select customers.name as 'Customers'
from customers
where customers.id not in
(
    select customerid from orders
);

select a.Name as Customers
from Customers as a
left join Orders as b
on a.Id=b.CustomerId
where b.CustomerId is null;

SELECT c.Name AS Customers
FROM Customers AS c
LEFT OUTER JOIN Orders AS o
ON c.Id = o.CustomerId
WHERE o.Id IS NULL;

SELECT Name AS Customers
FROM Customers
WHERE Id NOT IN(
	SELECT DISTINCT CustomerId FROM Orders
);

select c.Name as Customers
from Customers c
left join Orders d on c.Id = d.CustomerId
where d.CustomerId is null

```

8. 编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id 是这个表的主键。
例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+

```sql
DELETE p1 FROM Person p1, Person p2 WHERE p1.Email = p2.Email AND p1.Id > p2.Id

-- 自连接。其实delete方法和select的写法很相似。在本题代码中，delete的作用是删掉p1表里的一些记录（不是删掉整张表），要删哪些记录，就要看where里的条件了。

```

9. 部门表 Department：

+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| revenue       | int     |
| month         | varchar |
+---------------+---------+
(id, month) 是表的联合主键。
这个表格有关于每个部门每月收入的信息。
月份（month）可以取下列值 ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"]。
 

编写一个 SQL 查询来重新格式化表，使得新的表中有一个部门 id 列和一些对应 每个月 的收入（revenue）列。

查询结果格式如下面的示例所示：

Department 表：
+------+---------+-------+
| id   | revenue | month |
+------+---------+-------+
| 1    | 8000    | Jan   |
| 2    | 9000    | Jan   |
| 3    | 10000   | Feb   |
| 1    | 7000    | Feb   |
| 1    | 6000    | Mar   |
+------+---------+-------+

查询得到的结果表：
+------+-------------+-------------+-------------+-----+-------------+
| id   | Jan_Revenue | Feb_Revenue | Mar_Revenue | ... | Dec_Revenue |
+------+-------------+-------------+-------------+-----+-------------+
| 1    | 8000        | 7000        | 6000        | ... | null        |
| 2    | 9000        | null        | null        | ... | null        |
| 3    | null        | 10000       | null        | ... | null        |
+------+-------------+-------------+-------------+-----+-------------+

注意，结果表有 13 列 (1个部门 id 列 + 12个月份的收入列)。

```sql
select w1.Id
from Weather w1 join Weather w2
on datediff(w1.RecordDate, w2.RecordDate) = 1 and w1.Temperature > w2.Temperature;

select w1.Id
from Weather w1 , Weather w2
where datediff(w1.RecordDate, w2.RecordDate) = 1 and w1.Temperature > w2.Temperature;
```

10. 有一个courses 表 ，有: student (学生) 和 class (课程)。

请列出所有超过或等于5名学生的课。

例如，表：

+---------+------------+
| student | class      |
+---------+------------+
| A       | Math       |
| B       | English    |
| C       | Math       |
| D       | Biology    |
| E       | Math       |
| F       | Computer   |
| G       | Math       |
| H       | Math       |
| I       | Math       |
+---------+------------+
应该输出:

+---------+
| class   |
+---------+
| Math    |
+---------+

```sql
SELECT class FROM
    (SELECT
        class, COUNT(DISTINCT student) AS num
    FROM
        courses
    GROUP BY class) AS temp_table
WHERE
    num >= 5


SELECT class FROM courses GROUP BY class HAVING COUNT(DISTINCT student) >= 5

```