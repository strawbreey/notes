---
title: "Web Performance"
date: 2020-11-13T11:07:00+08:00
draft: true
---

Content

- Meke Fewer Http Requests 减少http请求

- Reduce DNS Lookups 减少dns查询

- Avoid Redirects 避免重定向

- Make ajax Cacheable ajax 缓存

- PostLoad Components 

- Preload Components 租价预加载

- Reduce the Number of DOM Elements

- Split Components Across Domains

- Minimize Number of iframes

- Avoid 404s

Server

- Use a Content Delivery Network (CDN)

- Add Expires or Cache-Control Header

- Gzip Components

- Configure ETags

- Flush Buffer Early

- Use GET for Ajax Requests

- Avoid Empty Image src

Cookie

- Reduce Cookie Size
- Use Cookie-Free Domains for Components

Css

- Put Stylesheets at Top
- Avoid CSS Expressions
- Choose `<link>` Over @import
- Avoid Filters

JavaScript

- Put Scripts at Bottom
- Make JavaScript and CSS External
- Minify JavaScript and CSS
- Remove Duplicate Scripts
- Minimize DOM Access
- Develop Smart Event Handlers

Images

- Optimize Images

- Optimize CSS Sprites

- Do Not Scale Images in HTML

- Make favicon.ico Small and Cacheable


Mobile

- Keep Components Under 25 KB

- Pack Components Into a Multipart Document




#### 如何尽可能避免重定向

一般来说, 重定向是无法完全避免的，适当地使用重定向能为网站提供更好的功能。（例如本地化，用户体验等方面）。

但是过多地进行重定向也肯定会给网站性能带来显著的影响。那么，有哪些方法可以作为我们改善这一点的参考呢

1. 在定义链接地址的href属性的时候，尽量使用最完整的、直接的地址。例如

  - 使用 www.cnblogs.com 而不是 cnblogs.com
  - 使用 cn.bing.com 而不是 bing.com
  - 使用 www.google.com.hk 而不是 google.com
  - 使用 www.mysite.com/products/ 而不是 www.mysite.com/products

2. 在使用Response.Redirect的时候，设置第二个参数为false

  - 考虑是否可用Server.Execute代替
  - 考虑Respone.RedirectPermanent

3. 如果涉及到从测试环境到生产环境的迁移，建议通过DNS中的CNAME的机制来定义别名，而不是强制地重定向来实现