---
title: "Micro Front Ends"
date: 2020-11-06T11:26:18+08:00
draft: false
---

### 聊聊微前端

如果把前端拆分，可以分为框架, 组件, 依赖库，配置 (业务) 和 数据

webpack打包的时候，可以把框架, 组件，依赖库打包成公共库

前端路由，业务逻辑 打包成私有库

数据共享 可以 考虑使用 内存/stroge 持久化


-[微前端在美团外卖的实践](https://tech.meituan.com/2020/02/27/meituan-waimai-micro-frontends-practice.html)