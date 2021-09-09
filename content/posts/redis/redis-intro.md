---
title: "Redis Intro"
date: 2021-09-09T10:11:52+08:00
draft: false
---

Redis 是完全开源的，遵守 BSD 协议，是一个高性能的 key-value 数据库。

特点:

- Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
- Redis不仅仅支持简单的key-value类型的数据，同时还提供list，set，zset，hash等数据结构的存储。
- Redis支持数据的备份，即master-slave模式的数据备份。

优势:

- 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s 。
- 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。
- 原子 – Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来。
- 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。

### 数据类型

Redis支持五种数据类型：`string（字符串）`，`hash（哈希）`，`list（列表）`，`set（集合`）及`zset(sorted set：有序集合)`。

### String（字符串）

string 是 redis 最基本的类型，你可以理解成与 Memcached 一模一样的类型，一个 key 对应一个 value。

string 类型是二进制安全的。意思是 redis 的 string 可以包含任何数据。比如jpg图片或者序列化的对象。

string 类型是 Redis 最基本的数据类型，string 类型的值最大能存储 512MB。

实列
```bash
SET runoob "菜鸟教程"
GET runoob  # "菜鸟教程"



# 1 SET key value 设置指定 key 的值

# 2	GET key 获取指定 key 的值。

# 3	GETRANGE key start end 返回 key 中字符串值的子字符

# 4	GETSET key value 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。

# 5	GETBIT key offset 对 key 所储存的字符串值，获取指定偏移量上的位(bit)。

# 6	MGET key1 [key2..] 获取所有(一个或多个)给定 key 的值。

# 7	SETBIT key offset value 对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。

# 8	SETEX key seconds value 将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。

# 9	SETNX key value 只有在 key 不存在时设置 key 的值。

# 10 SETRANGE key offset value 用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始。

# 11 STRLEN key 返回 key 所储存的字符串值的长度。

# 12 MSET key value [key value ...] 同时设置一个或多个 key-value 对。

# 13 MSETNX key value [key value ...] 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。

# 14 PSETEX key milliseconds value 这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。

# 15 INCR key 将 key 中储存的数字值增一。

# 16 INCRBY key increment 将 key 所储存的值加上给定的增量值（increment） 。

# 17 INCRBYFLOAT key increment 将 key 所储存的值加上给定的浮点增量值（increment） 。

# 18 DECR key 将 key 中储存的数字值减一。

# 19 DECRBY key decrement key 所储存的值减去给定的减量值（decrement） 。

# 20 APPEND key value 如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。

```

### Hash（哈希）

Redis hash 是一个键值(key=>value)对集合。

Redis hash 是一个 string 类型的 field 和 value 的映射表，hash 特别适合用于存储对象。

每个 hash 可以存储 2^32 -1 键值对（40多亿）

使用场景: 存取，读取，修改用户属性

```bash
HMSET runoob field1 "Hello" field2 "World"
"OK"
HGET runoob field1
"Hello"
HGET runoob field2
"World"

# 1	HDEL key field1 [field2] 删除一个或多个哈希表字段

# 2	HEXISTS key field 查看哈希表 key 中，指定的字段是否存在。

# 3	HGET key field 获取存储在哈希表中指定字段的值。

# 4	HGETALL key 获取在哈希表中指定 key 的所有字段和值

# 5	HINCRBY key field increment 为哈希表 key 中的指定字段的整数值加上增量 increment 。

# 6	HINCRBYFLOAT key field increment 为哈希表 key 中的指定字段的浮点数值加上增量 increment 。

# 7	HKEYS key 获取所有哈希表中的字段

# 8	HLEN key 获取哈希表中字段的数量

# 9	HMGET key field1 [field2] 获取所有给定字段的值

# 10	HMSET key field1 value1 [field2 value2 ] 同时将多个 field-value (域-值)对设置到哈希表 key 中。

# 11	HSET key field value 将哈希表 key 中的字段 field 的值设为 value 。

# 12	HSETNX key field value 只有在字段 field 不存在时，设置哈希表字段的值。

# 13	HVALS key 获取哈希表中所有值。

# 14	HSCAN key cursor [MATCH pattern] [COUNT count] 迭代哈希表中的键值对。
```






### List（列表）

Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。

列表最多可存储 232 - 1 元素 (4294967295, 每个列表可存储40多亿)。

使用场景: 最新消息，排行等功能，消息队列

```bash
lpush runoob redis          # (integer) 1
lpush runoob mongodb        # (integer) 2
lpush runoob rabbitmq       # (integer) 3
lrange runoob 0 10          # 1) "rabbitmq" 2) "mongodb" 3) "redis"

# 1	BLPOP key1 [key2 ] timeout 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

# 2	BRPOP key1 [key2 ] timeout 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

# 3	BRPOPLPUSH source destination timeout 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

# 4	LINDEX key index 通过索引获取列表中的元素

# 5	LINSERT key BEFORE|AFTER pivot value 在列表的元素前或者后插入元素

# 6	LLEN key 获取列表长度

# 7	LPOP key 移出并获取列表的第一个元素

# 8	LPUSH key value1 [value2] 将一个或多个值插入到列表头部

# 9	LPUSHX key value 将一个值插入到已存在的列表头部

# 10	LRANGE key start stop 获取列表指定范围内的元素

# 11	LREM key count value 除列表元素

# 12	LSET key index value 通过索引设置列表元素的值

# 13	LTRIM key start stop 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。

# 14	RPOP key 移除列表的最后一个元素，返回值为移除的元素。

# 15	RPOPLPUSH source destination 移除列表的最后一个元素，并将该元素添加到另一个列表并返回

# 16	RPUSH key value1 [value2] 在列表中添加一个或多个值

# 17	RPUSHX key value 为已存在的列表添加值
```




### Set（集合）

Redis 的 Set 是 string 类型的无序集合。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。

集合中最大的成员数为 232 - 1(4294967295, 每个集合可存储40多亿个成员)。

使用场景: 共同好友，利用唯一性统计访问网站的所有独立ip，好友推荐

```bash
sadd runoob redis       # (integer) 1
sadd runoob mongodb     # (integer) 1
sadd runoob rabbitmq    # (integer) 1
sadd runoob rabbitmq    # (integer) 0
smembers runoob         # 1) "redis" 2) "rabbitmq" 3) "mongodb"
```

注意：以上实例中 rabbitmq 添加了两次，但根据集合内元素的唯一性，第二次插入的元素将被忽略。





### zset(sorted set：有序集合)

Redis zset 和 set 一样也是string类型元素的集合, 且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

zset的成员是唯一的,但分数(score)却可以重复。

使用场景： 1、排行榜 2、带权重的消息队列

```bash

# zadd 命令
# 添加元素到集合，元素在集合中存在则更新对应score

# zadd key score member 

zadd runoob 0 redis         # (integer) 1
zadd runoob 0 mongodb       # (integer) 1
zadd runoob 0 rabbitmq      # (integer) 1
zadd runoob 0 rabbitmq      # (integer) 0
ZRANGEBYSCORE runoob 0 1000     # 1) "mongodb" 2) "rabbitmq" 3) "redis"
```

