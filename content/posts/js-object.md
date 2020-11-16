---
title: "Js Object"
date: 2020-10-30T15:50:58+08:00
draft: false
---

### 移除对象中的空值

```js
let obj

for ( var key in obj ){
	if ( data[key] === '' ){
		delete data[key]
	}
}
```