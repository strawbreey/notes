---
title: "Redis Tutorial"
date: 2020-09-21T10:29:08+08:00
draft: false
---

### Install

```shell
# download redis
wget http://download.redis.io/releases/redis-2.8.17.tar.gz

# Unzip  z:gzip压缩 c:创建压缩包 v:显示打包压缩解压过程 f:接着压缩 t:查看压缩包内容 x:解压 X:指定文件列表形式排除不需要打包压缩的文件或目录
tar xzf redis-2.8.17.tar.gz # xvfz
cd redis-2.8.17
# make
make

# server
cd src
./redis-server

./redis-server ../redis.conf
```

#### Ubuntu 安装

```shell
# install redis
apt-get update
apt-get install redis-server

# start redis
redis-server

# test redis
redis-cli

# 
```


### Redis Data Type 

Redis 支持5中数据类型， string(字符串), hash(哈希), list(列表), set(列表) 和 zset(有序集合)

String 

```shell
SET runoob "菜鸟教程"
GET runoob #菜鸟教程  一个键最大能存储512MB
```

Hash(哈希)
Redis hash 是一个键值对集合
Redis hash 是一个 string 类型的field 和 value 的映射表，hash特别适合用于存储对象
```
HMSET data filed1 "hello" filed2 "world"
HGET data filed1
```

List(列表)
Redis 列表是简单的字符串列表, 按照插入顺序排序，你可以添加一个元素到头部(左边)或者尾部(右边)
```shell
lpush data redis
lpush data mogodb
lpush data rabitmq
lrange data 0 10
```
1) "rabitmq"
2) "mongodb"
3) "redis"

### Set (集合)
Redis 的 Set 是 string 类型的无序集合。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。

```shell
sadd data redis
sadd data mongodb
sadd data rabitmq
smembers data

```
1) "redis"
2) "rabitmq"
3) "mongodb"

zset 
Reset 和set 一样的string类型的元素集合,且不允许重复的成员
```shell
zadd key score member
zadd data 0 redis
zadd data 0 mogodb
zadd runoob 0 rabitmq

ZRANGEBYSCORE runoob 0 1000
```
1) "mongodb"
2) "rabitmq"
3) "redis"

### Redis command

语法

```shell
# 打开终端并输入命令 redis-cli，该命令会连接本地的 redis 服务
redis-cli

# 执行 PING 命令，该命令用于检测 redis 服务是否启动。
PING

# 在远程服务上执行命令
redis-cli -h host -p port -a password

```

### Redis

语法
```shell
#  COMMAND KEY_NAME
SET key redis
GET key  # redis


SET key value # 设定指定key的值
GET key # 获取指定key的值
GETSET key value end # 返回字符串值得子字符
GETSET key value # 将给定的key的值设为value， 并返回key的旧值
GETBIT key offset # 对key所存储的字符串值，或者指定偏移量上的位
MGET key[key2...] 获取所有给定key的值
SETBIT key offset value # 对key所存储的字符串值，设置或者清除指定偏移量上的位

```

### Redis 列表(List)
Redis 列表是简单的字符串列表，按照插入顺序，可以添加元素到列表的头部或者尾部。一个列表最多可以包含 232 - 1 个元素 (4294967295, 每个列表超过40亿个元素)
```shell
LPUSH runoobkey redis
LPUSH runoobkey mongodb
LPUSH runoobkey mysql
LRANGE runoobkey 0 10

1) "mysql"
2) "mongodb"
3) "redis"
```

### Redis 事务
Redis 事务可以一次执行多个命令， 并且带有以下三个重要的保证：

- 批量操作在发送 EXEC 命令前被放入队列缓存。
- 收到 EXEC 命令后进入事务执行，事务中任意命令执行失败，其余的命令依然被执行。
- 在事务执行过程，其他客户端提交的命令请求不会插入到事务执行命令序列中。

一个事务从开始到执行会经历以下三个阶段：

- 开始事务。
- 命令入队。
- 执行事务。

```shell
MULTI
SET book-name "Mastering C++ in 21 days" QUEUED
GET book-name QUEUED
SADD tag "C++" "Programming" "Mastering Series"
SMEMBERS tag QUEUED
EXEC

1) OK
2) "Mastering C++ in 21 days"
3) (integer) 3
4) 1) "Mastering Series"
   2) "C++"
   3) "Programming"
```

#### Redis 事务命令
- DISCRAD 取消事务, 放弃执行事务块内的所有命令
- EXEC 执行所有事务块内的命令
- MULTI 标记一个事务块的开始
- UNMATCH 取消WATCH命令对所有key的监事
- WATCH key[key...] 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。