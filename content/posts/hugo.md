---
title: "Hugo"
date: 2020-08-28T19:18:18+08:00
draft: true
---

## 创建文章
创建一个`hugo`页面
```shell
hugo new hugo.md
```

```md

```

内容是 Markdown 格式的，+++ 之间的内容是 TOML 格式的，根据你的喜好，你可以换成 YAML 格式（使用 --- 标记）或者 JSON 格式。


## 安装皮肤
```shell
cd themes
git clone https://github.com/spf13/hyde.git
```

## 运行hugo
```shell
  hugo server --theme=hyde --buildDrafts
```