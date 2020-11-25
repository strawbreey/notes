---
title: "Elasticsearch"
date: 2020-09-14T10:09:20+08:00
draft: false
---

全文搜索属于最常见的需求，开源的 Elasticsearch （以下简称 Elastic）是目前全文搜索引擎的首选。它可以快速地储存、搜索和分析海量数据。维基百科、Stack Overflow、Github 都采用它

Elastic 的底层是开源库 Lucene。但是，你没法直接用 Lucene，必须自己写代码去调用它的接口。Elastic 是 Lucene 的封装，提供了 REST API 的操作接口，开箱即用。

## 全文搜索引擎(elasticsearch)

### install

```shell
# 下载,解压
wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.5.1.zip
unzip elasticsearch-5.5.1.zip
cd elasticsearch-5.5.1/ 

# 启动
./bin/elasticsearch
```

注: Elastic 需要 Java 8 环境。默认情况下，Elastic 只允许本机访问，如果需要远程访问，可以修改 Elastic 安装目录的`config/elasticsearch.yml`文件，去掉`network.host`的注释，将它的值改成`0.0.0.0`，然后重新启动`Elastic`。

### introduction

#### Node 与 Cluster

`Elastic` 本质上是一个分布式数据库，允许多台服务器协同工作，每台服务器可以运行多个`Elastic` 实例。
单个 `Elastic` 实例称为一个节点（node）。一组节点构成一个集群（cluster）。

#### Index

`Elastic` 会索引所有字段，经过处理后写入一个反向索引（Inverted Index）。查找数据的时候，直接查找该索引。

所以，`Elastic` 数据管理的顶层单位就叫做 Index（索引）。它是单个数据库的同义词。每个 Index （即数据库）的名字必须是小写。

```shell
curl -X GET 'http://localhost:9200/_cat/indices?v'
```

#### Document

Index 里面单条的记录称为 Document（文档）。许多条 Document 构成了一个 Index。

```json
// document demo
{
  "user": "张三",
  "title": "工程师",
  "desc": "数据库管理"
}
```

同一个 Index 里面的 Document，不要求有相同的结构（scheme），但是最好保持相同，这样有利于提高搜索效率。

#### Type

Document 可以分组，比如weather这个 Index 里面，可以按城市分组（北京和上海），也可以按气候分组（晴天和雨天）。这种分组就叫做 Type，它是虚拟的逻辑分组，用来过滤 Document。

#### 新建和删除 Index

```shell
# 新建 Index，可以直接向 Elastic 服务器发出 PUT 请求。
curl -X PUT 'localhost:9200/weather'

# 返回
{
  "acknowledged":true,
  "shards_acknowledged":true
}

# 删除这个 Index
curl -X DELETE 'localhost:9200/weather'
```
### 数据操作

#### 新增记录

向指定的 /Index/Type 发送 PUT 请求，就可以在 Index 里面新增一条记录。比如，向/accounts/person发送请求，就可以新增一条人员记录。

```shell
# 新增数据
curl -X PUT 'localhost:9200/accounts/person/1' -d '
{
  "user": "张三",
  "title": "工程师",
  "desc": "数据库管理"
}' 

# 返回JSON对象
{
  "_index":"accounts",
  "_type":"person",
  "_id":"1",
  "_version":1,
  "result":"created",
  "_shards":{"total":2,"successful":1,"failed":0},
  "created":true
}

# 不指定id
$ curl -X POST 'localhost:9200/accounts/person' -d '
{
  "user": "李四",
  "title": "工程师",
  "desc": "系统管理"
}
```
#### 查看记录
```shell
# URL 的参数pretty=true表示以易读的格式返回。
curl 'localhost:9200/accounts/person/1?pretty=true'

# 返回数据
{
  "_index" : "accounts",
  "_type" : "person",
  "_id" : "1",
  "_version" : 1,
  "found" : true,
  "_source" : {
    "user" : "张三",
    "title" : "工程师",
    "desc" : "数据库管理"
  }
}
```

#### 删除记录

```shell
curl -X DELETE 'localhost:9200/accounts/person/1'
```

#### 更新记录
```shell
$ curl -X PUT 'localhost:9200/accounts/person/1' -d '
{
    "user" : "张三",
    "title" : "工程师",
    "desc" : "数据库管理，软件开发"
}' 

# 返回数据
{
  "_index":"accounts",
  "_type":"person",
  "_id":"1",
  "_version":2,
  "result":"updated",
  "_shards":{"total":2,"successful":1,"failed":0},
  "created":false
}
```

#### 数据查询

使用 GET 方法，直接请求/Index/Type/_search，就会返回所有记录。

```shell
$ curl 'localhost:9200/accounts/person/_search'

# 返回
{
  "took":2,
  "timed_out":false,
  "_shards":{"total":5,"successful":5,"failed":0},
  "hits":{
    "total":2,
    "max_score":1.0,
    "hits":[
      {
        "_index":"accounts",
        "_type":"person",
        "_id":"AV3qGfrC6jMbsbXb6k1p",
        "_score":1.0,
        "_source": {
          "user": "李四",
          "title": "工程师",
          "desc": "系统管理"
        }
      },
      {
        "_index":"accounts",
        "_type":"person",
        "_id":"1",
        "_score":1.0,
        "_source": {
          "user" : "张三",
          "title" : "工程师",
          "desc" : "数据库管理，软件开发"
        }
      }
    ]
  }
}

# 上面代码中，返回结果的 took字段表示该操作的耗时（单位为毫秒），timed_out字段表示是否超时，hits字段表示命中的记录，里面子字段的含义如下。
# total：返回记录数，本例是2条。
# max_score：最高的匹配程度，本例是1.0。
# hits：返回的记录组成的数组。
```

### 全文搜索

Elastic 的查询非常特别，使用自己的查询语法，要求 GET 请求带有数据体。

```shell
curl 'localhost:9200/accounts/person/_search'  -d '
{
  "query" : { "match" : { "desc" : "软件" }}
}'

# 返回
{
  "took":3,
  "timed_out":false,
  "_shards":{"total":5,"successful":5,"failed":0},
  "hits":{
    "total":1,
    "max_score":0.28582606,
    "hits":[
      {
        "_index":"accounts",
        "_type":"person",
        "_id":"1",
        "_score":0.28582606,
        "_source": {
          "user" : "张三",
          "title" : "工程师",
          "desc" : "数据库管理，软件开发"
        }
      }
    ]
  }
}

# Elastic 默认一次返回10条结果，可以通过size字段改变这个设置。
$ curl 'localhost:9200/accounts/person/_search'  -d '
{
  "query" : { "match" : { "desc" : "管理" }},
  "size": 1
}'

# 可以通过from字段，指定位移。
$ curl 'localhost:9200/accounts/person/_search'  -d '
{
  "query" : { "match" : { "desc" : "管理" }},
  "from": 1,
  "size": 1
}'

```

### 逻辑运算

如果有多个搜索关键字， Elastic 认为它们是or关系。

```shell
$ curl 'localhost:9200/accounts/person/_search'  -d'
{
  "query" : { "match" : { "desc" : "软件 系统" }}
}'
```

### 面试题

- elasticsearch了解多少，说说你们公司es的集群架构，索引数据大小，分片有多少，以及一些调优手段 

- elasticsearch的倒排索引是什么

- elasticsearch 索引数据多了怎么办，如何调优，部署

- elasticsearch是如何实现master选举的

- 详细描述一下Elasticsearch索引文档的过程

- 详细描述一下Elasticsearch搜索的过程 

- Elasticsearch在部署时，对Linux的设置有哪些优化方法

- lucence内部结构是什么


### 参考链接
- [ElasticSearch 官方手册](https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started.html)
- [ElasticSearch Source Code](https://github.com/elastic/elasticsearch)
- [ElasticSearch 最佳实践](https://www.elastic.co/blog/a-practical-introduction-to-elasticsearch)
- [阮一峰--全文搜索引擎 Elasticsearch 入门教程](https://www.ruanyifeng.com/blog/2017/08/elasticsearch.html)
- [sql](/posts/sql)