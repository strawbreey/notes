---
title: "Js Special Characters"
date: 2021-08-19T16:48:00+08:00
draft: false
---


### 删除字符串中的特殊字符

```js
const pattern=/[`~!@#$^&*()=|{}':;',\\\[\]<>\/?~！@#￥……&*（）——|{}【】'；：""'。，、？\s]/g;
const filename = item.file.name.replace(pattern,"");
````