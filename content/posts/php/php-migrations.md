---
title: "Php Migrations"
date: 2021-07-29T10:35:14+08:00
draft: false
---

在开发和维护一个数据库驱动的应用程序时， 数据库的结构会像代码一样不断`演变`。 例如，在开发应用程序的过程中，会增加一张新表且必须得加进来； 在应用程序被部署到生产环境后，需要建立一个索引来提高查询的性能等等。 因为一个数据库结构发生改变的时候源代码也经常会需要做出改变， Yii 提供了一个 数据库迁移 功能，该功能可以记录数据库的变化， 以便使数据库和源代码一起受版本控制。