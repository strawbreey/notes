---
title: "Vim No Write Since Last Change"
date: 2020-12-15T12:29:06+08:00
draft: false
---

### 故障现象

使用vim修改文件, :q或者:wq退出失败 报错，系统提示如下：
```
E37: No write since last change (add ! to override)
```

1. 使用命令:w!强制存盘即可，在vim模式下，键入以下命令：

```bash
:w！
```

2. 这种情况下，多半是没有权限造成的，没有write的权限。直接退出即可

退出方法：ctrl+z


