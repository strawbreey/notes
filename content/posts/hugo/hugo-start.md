---
title: "Hugo Start"
date: 2020-08-31T10:15:27+08:00
draft: false
---

## start

### Install Hugo

```bash
# macOS install 
brew intall hugo
hugo version
```

### Create a New Site

```bash
hugo new site quickStart
```

### Add a Theme
```bash
cd quickstart
git init
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
``` 

### Add Some Content
```bash
hugo new posts/new-content.md
```

### Start the Hugo server
```bash
hugo serve -D

```

### Build Static pages
```bash
hugo -D
```