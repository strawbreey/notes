---
title: "Angular Platform"
date: 2021-01-20T19:51:27+08:00
draft: false
---

```ts
// 查看平台
import { Platform } from '@angular/cdk/platform';

constructor(
  private platform: Platform,
) {
  if (this.platform.isBrowser) {
    console.log('判断当前平台是浏览器');
  }
}
```

