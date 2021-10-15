---
title: "X Requested With"
date: 2021-10-14T20:16:26+08:00
draft: false
---

```js
// String requestedWith = ((HttpServletRequest) request).getHeader("X-Requested-With");
  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const authReq = req.clone({setHeaders: {
      'X-Requested-With': 'XMLHttpRequest',
      'request-id': Math.random().toString(36).slice(-8) +  '-' + new Date().getTime() + '-'
    }});
    return next.handle(authReq);
  }
```

于是查了一下

如果 requestedWith 为 null，则为同步请求。

如果 requestedWith 为 XMLHttpRequest 则为 Ajax 请求。