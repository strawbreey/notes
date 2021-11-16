---
title: "Cors"
date: 2020-09-27T11:05:18+08:00
draft: false
---




跨源资源共享 (CORS: Cross-origin resource sharing) （或通俗地译为跨域资源共享）是一种机制，该机制使用附加的 HTTP 头来告诉浏览器，准许运行在一个源上的Web应用访问位于另一不同源选定的资源。 当一个Web应用发起一个与自身所在源（域，协议和端口）不同的HTTP请求时，它发起的即跨源HTTP请求。

出于安全性，浏览器限制脚本内发起的跨源HTTP请求。 例如，`XMLHttpRequest` 和 `Fetch API` 遵循同源策略。 这意味着使用这些API的Web应用程序只能从加载应用程序的同一个域请求HTTP资源，除非响应报文包含了正确CORS响应头。

跨源域资源共享（ CORS ）机制允许 Web 应用服务器进行跨源访问控制，从而使跨源数据传输得以安全进行。现代浏览器支持在 API 容器中（例如 XMLHttpRequest 或 Fetch ）使用 CORS，以降低跨源 HTTP 请求所带来的风险。

整个CORS通信过程，都是浏览器自动完成，不需要用户参与。对于开发者来说，CORS通信与同源的AJAX通信没有差别，代码完全一样。浏览器一旦发现AJAX请求跨源，就会自动添加一些附加的头信息，有时还会多出一次附加的请求，但用户不会有感觉。

因此，实现CORS通信的关键是服务器。只要服务器实现了CORS接口，就可以跨源通信。

浏览器将CORS请求分成两类：简单请求（simple request）和非简单请求（not-so-simple request）

### CORS headers


Access-Control-Allow-Origin: 域白名单，决定哪些域可以访问资源;

Access-Control-Allow-Credentials: 决定是否响应 withCredentials 为 true即带Cookies的请求;

Access-Control-Allow-Headers 预请求的响应头，决定哪些HTTP头可以用在真正的请求头中;

Access-Control-Allow-Methods 预请求的响应头，决定正式请求可以使用哪些方法（GET, POST, PUT ...）;

Access-Control-Expose-Headers 哪些头可以作为响应头;

Access-Control-Max-Age 预请求的返回可以被缓存的时长;

Access-Control-Request-Headers 预请求的请求头，告知服务器正式请求中哪些请求头会被使用;

Origin 预请求的请求头，告知服务器请求的来源;

当真正一个跨域请求发起时，浏览器会决定这个请求是简单请求还是复杂请求，简单请求则直接发起，复杂请求则需要发起一个预请求来询问服务器是否允许发送该请求。

### 简单请求

（1) 请求方法是以下三种方法之一：

- HEAD
- GET
- POST

（2）HTTP的头信息不超出以下几种字段：

- Accept
- Accept-Language
- Content-Language
- Last-Event-ID
- Content-Type：application/x-www-form-urlencoded、multipart/form-data、text/plain

这是为了兼容表单（form），因为历史上表单一直可以发出跨域请求。AJAX 的跨域设计就是，只要表单可以发，AJAX 就可以直接发。

对于简单请求，浏览器直接发出CORS请求。具体来说，就是在头信息之中，增加一个Origin字段。

```
  GET /cors HTTP/1.1
  Origin: http://api.bob.com
  Host: api.alice.com
  Accept-Language: en-US
  Connection: keep-alive
  User-Agent: Mozilla/5.0...
```

上面的头信息中，Origin字段用来说明，本次请求来自哪个源（协议 + 域名 + 端口）。服务器根据这个值，决定是否同意这次请求。

如果Origin指定的源，不在许可范围内，服务器会返回一个正常的HTTP回应。浏览器发现，这个回应的头信息没有包含Access-Control-Allow-Origin字段（详见下文），就知道出错了，从而抛出一个错误，被XMLHttpRequest的onerror回调函数捕获。注意，这种错误无法通过状态码识别，因为HTTP回应的状态码有可能是200

```
  Access-Control-Allow-Origin: http://api.bob.com
  Access-Control-Allow-Credentials: true
  Access-Control-Expose-Headers: FooBar
  Content-Type: text/html; charset=utf-8
```

### 复杂请求

没有满足简单请求的条件外的即为复杂请求，复杂请求会触发预检机制，
客户端会先发送一个OPTIONS方法的请求得到 `Access-Control-Requeset-Method` 等响应头判断是否能够成功发起请求，否则放弃发送正式请求，由于OPTIONS方法的请求不会改变服务器资源，所以预请求被认定是安全的。

如果预检请求返回了重定向，多大数浏览器是不支持对预检请求的重定向的，可以先发送一个简单请求，获取到重定向的URL，然后再发送正式请求。

