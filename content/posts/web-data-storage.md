---
title: "Web Data Storage"
date: 2020-10-14T22:36:39+08:00
draft: false
---

## Cookie

Cookie 是直接存储在浏览器中的一小串数据。它们是 HTTP 协议的一部分，由 [RFC 6265](https://tools.ietf.org/html/rfc6265) 规范定义。

Cookie 通常是由 Web 服务器使用响应 Set-Cookie HTTP-header 设置的。然后浏览器使用 Cookie HTTP-header 将它们自动添加到（几乎）每个对相同域的请求中

最常见的用处之一就是身份验证：

登录后，服务器在响应中使用 Set-Cookie HTTP-header 来设置具有唯一“会话标识符（session identifier）”的 cookie。
下次如果请求是由相同域发起的，浏览器会使用 Cookie HTTP-header 通过网络发送 cookie。
所以服务器知道是谁发起了请求。

```js
alert(document.cookie)
```

写入 document.cookie

```js
document.cookie = "user=John"; // 只会更新名称为 user 的 cookie
alert(document.cookie); // 展示所有 cookie
```

从技术上讲，cookie 的名称和值可以是任何字符，为了保持有效的格式，它们应该使用内建的 encodeURIComponent 函数对其进行转义：

```js
// 特殊字符（空格），需要编码
let name = "my name";
let value = "John Smith"

// 将 cookie 编码为 my%20name=John%20Smith
document.cookie = encodeURIComponent(name) + '=' + encodeURIComponent(value);

alert(document.cookie); // ...; my%20name=John%20Smith
```

存在一些限制：

    encodeURIComponent 编码后的 name=value 对，大小不能超过 4kb。因此，我们不能在一个 cookie 中保存大的东西。
    每个域的 cookie 总数不得超过 20+ 左右，具体限制取决于浏览器。

domain

可访问 cookie 的域。但是在实际中，有一些限制。我们无法设置任何域。

默认情况下，cookie 只有在设置的域下才能被访问到。所以，如果 cookie 设置在 site.com 下，我们在 other.com 下就无法获取它。


### expires，max-age

默认情况下，如果一个 cookie 没有设置这两个参数中的任何一个，那么在关闭浏览器之后，它就会消失。此类 cookie 被称为 "session cookie”。

为了让 cookie 在浏览器关闭后仍然存在，我们可以设置 expires 或 max-age 选项中的一个。

expires=Tue, 19 Jan 2038 03:14:07 GMT
cookie 的到期日期，那时浏览器会自动删除它。

日期必须完全采用 GMT 时区的这种格式。我们可以使用 date.toUTCString 来获取它。例如，我们可以将 cookie 设置为 1 天后过期。

```js
// 当前时间 +1 天
let date = new Date(Date.now() + 86400e3);
date = date.toUTCString();
document.cookie = "user=John; expires=" + date;
```

max-age=3600
expires 的替代选项，具指明 cookie 的过期时间距离当前时间的秒数。

如果为 0 或负数，则 cookie 会被删除：

// cookie 会在一小时后失效
document.cookie = "user=John; max-age=3600";

// 删除 cookie（让它立即过期）
document.cookie = "user=John; max-age=0";