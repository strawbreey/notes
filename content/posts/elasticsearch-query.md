---
title: "Elasticsearch Query"
date: 2021-08-09T10:58:42+08:00
draft: false
---

### ElasticSearch的Query String语法基本介绍

1. 根据字段名查询(Field names)
接口(interface)为app

interface: app

接口(interface)为app或live

interface: (app OR live)

精确匹配, 接口是app

interface: “app”

2. 通配符(Fuzziness)
‘?’表示单个字符，’*’表示0个或多个字符

查找视频域名
domain: *.video.sina.com.cn

3. 范围查询(Range)
502次数大于等于50

code502: [50 TO *]

总数大于等于10且小于50

total: [10 TO 50}

4. 布尔运算(Boolean operators）
AND

domain: *.video.sina.com.cn AND interface: app

NOT

NOT interface: app

OR

interface: app OR interface: live

5. 正则表达式(Regular expressions)
匹配200、300、400、500状态码

status:/[2-5]00/

匹配两个域名

domain:/[nl].portal.com/

### Reference
Query String syntax