---
title: "Vscode Toggle Word Wrap"
date: 2020-12-07T10:34:32+08:00
draft: false
---

### vscode 自动换行
- View-> Toggle Word Wrap 选项.

- 快捷键 : alt + Z

- code-首选项-设置，找到"editor.wordWrap"属性，该属性默认值为off，改为“on”。vscode的"editor.wordWrapColumn"默认值为80，即每行80字符时自动换行。

- vscode setting.json

```
"editor.wordWrap": "on"
```