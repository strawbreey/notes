---
title: "Angular TrackBy"
date: 2020-10-27T17:40:15+08:00
draft: true
---

trackBy: trackByFn

[ngForTrackBy]="trackByFn"

```js
  trackByFn(index, item) {
    return item.id ? item.id : index; // or item.id
  }
```


### 参考文献

- [Using Pure Pipes To Generate NgFor TrackBy Identity Functions In Angular 7.2.7](https://www.bennadel.com/blog/3579-using-pure-pipes-to-generate-ngfor-trackby-identity-functions-in-angular-7-2-7.htm)