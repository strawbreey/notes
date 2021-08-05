---
title: "How to Know That SQL Is Slow"
date: 2021-07-28T20:21:17+08:00
draft: false
---

## 如何知道一条sql很慢

看索引的类型，普通索引允许被索引的数据列包含重复的值, 这时候等同于全表扫描。如果确认索引值唯一，建议加上limit 1 .