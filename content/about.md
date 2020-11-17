---
title: "About"
date: 2020-08-28T19:17:20+08:00
draft: false
---

### 创建文章

创建一个`hugo`页面

```shell
hugo new hugo.md
```

内容是 Markdown 格式的，+++ 之间的内容是 TOML 格式的，根据你的喜好，你可以换成 YAML 格式（使用 --- 标记）或者 JSON 格式。

### 安装皮肤

```bash
  cd themes
  git clone https://github.com/spf13/hyde.git
```

### 运行hugo
```bash
  hugo server --theme=hyde --buildDrafts
```