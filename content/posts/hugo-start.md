---
title: "Hugo Start"
date: 2020-08-31T10:15:27+08:00
draft: false
---

## start

### Install Hugo

```shell
# macOS install 

brew intall hugo
hugo version
```

### Create a New Site

```shell
hugo new site quickStart
```

### Add a Theme
```shell
cd quickstart
git init
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
``` 

### Add Some Content
```shell
hugo new posts/new-content.md
```

### Start the Hugo server
```shell
hugo serve -D

```

### Build Static pages
```shell
hugo -D
```