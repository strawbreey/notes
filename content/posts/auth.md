---
title: "Auth"
date: 2020-09-15T10:45:56+08:00
draft: false
---

前后端中常用的鉴权方式有四种

- HTTP Basic Authentication
- session-cookie
- Token 验证
- OAuth(开发授权)

### HTTP Basic Authentication

 - 客户端向服务器请求数据，请求的内容可能是一个网页或者是一个ajax异步请求，此时，假设客户端尚未被验证，则客户端提供如下请求至服务器:
 ```
  Get /index.html HTTP/1.0 
  Host:www.google.com
 ```

  - 服务器向客户端发送验证请求代码401,（WWW-Authenticate: Basic realm=”google.com”这句话是关键，如果没有客户端不会弹出用户名和密码输入界面）服务器返回的数据大抵如下：
  ```
  HTTP/1.0 401 Unauthorised 
  Server: SokEvo/1.0 
  WWW-Authenticate: Basic realm=”google.com” 
  Content-Type: text/html 
  Content-Length: xxx
  ```

  - 当符合http1.0或1.1规范的客户端（如IE，FIREFOX）收到401返回值时，将自动弹出一个登录窗口，要求用户输入用户名和密码。

  - 用户输入用户名和密码后，将用户名及密码以BASE64加密方式加密，并将密文放入前一条请求信息中，则客户端发送的第一条请求信息则变成如下内容：
  ```
  Get /index.html HTTP/1.0 
  Host:www.google.com 
  Authorization: Basic d2FuZzp3YW5n
  ```
  - 服务器收到上述请求信息后，将 Authorization 字段后的用户信息取出、解密，将解密后的用户名及密码与用户数据库进行比较验证，如用户名及密码正确，服务器则根据请求，将所请求资源发送给客户端

优点：
- 基本认证的一个优点是基本上所有流行的网页浏览器都支持基本认证。基本认证很少在可公开访问的互联网网站上使用，有时候会在小的私有系统中使用（如路由器网页管理接口）。后来的机制HTTP摘要认证是为替代基本认证而开发的，允许密钥以相对安全的方式在不安全的通道上传输。

- 程序员和系统管理员有时会在可信网络环境中使用基本认证，使用Telnet或其他明文网络协议工具手动地测试Web服务器。这是一个麻烦的过程，但是网络上传输的内容是人可读的，以便进行诊断。

缺点：
- 虽然基本认证非常容易实现，但该方案创建在以下的假设的基础上，即：客户端和服务器主机之间的连接是安全可信的。特别是，如果没有使用SSL/TLS这样的传输层安全的协议，那么以明文传输的密钥和口令很容易被拦截。该方案也同样没有对服务器返回的信息提供保护。

- 现存的浏览器保存认证信息直到标签页或浏览器被关闭，或者用户清除历史记录。HTTP没有为服务器提供一种方法指示客户端丢弃这些被缓存的密钥。这意味着服务器端在用户不关闭浏览器的情况下，并没有一种有效的方法来让用户注销。
。

### session-cookie
Http协议是一个无状态的协议，服务器不会知道到底是哪一台浏览器访问了它，因此需要一个标识用来让服务器区分不同的浏览器。cookie 就是这个管理服务器与客户端之间状态的标识。
cookie 的原理是，浏览器第一次向服务器发送请求时，服务器在 response 头部设置 Set-Cookie 字段，浏览器收到响应就会设置 cookie 并存储，在下一次该浏览器向服务器发送请求时，就会在 request 头部自动带上 Cookie 字段，服务器端收到该 cookie 用以区分不同的浏览器。当然，这个 cookie 与某个用户的对应关系应该在第一次访问时就存在服务器端，这时就需要 session 了。

session 是会话的意思，浏览器第一次访问服务端，服务端就会创建一次会话，在会话中保存标识该浏览器的信息。它与 cookie 的区别就是 session 是缓存在服务端的，cookie 则是缓存在客户端，他们都由服务端生成，为了弥补 Http 协议无状态的缺陷。

服务器在接受客户端首次访问时在服务器端创建seesion，然后保存seesion(我们可以将seesion保存在 内存中，也可以保存在redis中，推荐使用后者)，然后给这个session生成一个唯一的标识字符串,然后在 响应头中种下这个唯一标识字符串。
签名。这一步通过秘钥对sid进行签名处理，避免客户端修改sid。(非必需步骤)
浏览器中收到请求响应的时候会解析响应头，然后将sid保存在本地cookie中，浏览器在下次http请求的请求头中会带上该域名下的cookie信息。
服务器在接受客户端请求时会去解析请求头cookie中的sid，然后根据这个sid去找服务器端保存的该客户端的session，然后判断该请求是否合法。

redis是一个键值服务器，可以专门放session的键值对。

### Token

token 是一个令牌，浏览器第一次访问服务端时会签发一张令牌，之后浏览器每次携带这张令牌访问服务端就会认证该令牌是否有效，只要服务端可以解密该令牌，就说明请求是合法的，令牌中包含的用户信息还可以区分不同身份的用户。一般 token 由用户信息、时间戳和由 hash 算法加密的签名构成。

#### Token认证流程

- 客户端使用用户名跟密码请求登录
- 服务端收到请求，去验证用户名与密码
- 验证成功后，服务端会签发一个 Token，再把这个 Token 发送给客户端
- 客户端收到 Token 以后可以把它存储起来，比如放在 Cookie 里或者Local Storage 里
- 客户端每次向服务端请求资源的时候需要带着服务端签发的 Token
- 服务端收到请求，然后去验证客户端请求里面带着的 Token（request头部添加Authorization），如果验证成功，就向客户端返回请求的数据 ，如果不成功返回401错误码，鉴权失败。


#### Token 和 session 的区别
session-cookie的缺点：
- 认证方式局限于在浏览器中使用，cookie 是浏览器端的机制，如果在app端就无法使用 cookie。
- 为了满足全局一致性，我们最好把 session 存储在 redis 中做持久化，而在分布式环境下，我们可能需要在每个服务器上都备份，占用了大量的存储空间。
- 在不是 Https 协议下使用 cookie ，容易受到 CSRF 跨站点请求伪造攻击。

token的缺点：
- 加密解密消耗使得 token 认证比 session-cookie 更消耗性能。
- token 比 sessionId 大，更占带宽。

两者对比，它们的区别显而易见：
- token 认证不局限于 cookie ，这样就使得这种认证方式可以支持多种客户端，而不仅是浏览器。且不受同源策略的影响。
- 不使用 cookie 就可以规避CSRF攻击。
- token 不需要存储，token 中已包含了用户信息，服务器端变成无状态，服务器端只需要根据定义的规则校验这个 token 是否合法就行。这也使得 token 的可扩展性更强。


基于 token 的解决方案有许多，常用的是JWT，JWT 的原理是，服务器认证以后，生成一个 JSON 对象，这个 JSON 对象肯定不能裸传给用户，那谁都可以篡改这个对象发送请求。因此这个 JSON 对象会被服务器端签名加密后返回给用户，返回的内容就是一张令牌，以后用户每次访问服务器端就带着这张令牌。
这个 JSON 对象可能包含的内容就是用户的信息，用户的身份以及令牌的过期时间。

### OAuth(开放授权)

OAuth（Open Authorization）是一个开放标准，允许用户授权第三方网站访问他们存储在另外的服务提供者上的信息，而不需要将用户名和密码提供给第三方网站或分享他们数据的所有内容，为了保护用户数据的安全和隐私，第三方网站访问用户数据前都需要显式的向用户征求授权。我们常见的提供OAuth认证服务的厂商有支付宝，QQ,微信。
OAuth协议又有1.0和2.0两个版本。相比较1.0版，2.0版整个授权验证流程更简单更安全，也是目前最主要的用户身份验证和授权方式。

#### OAuth认证流程

OAuth就是一种授权机制。数据的所有者告诉系统，同意授权第三方应用进入系统，获取这些数据。系统从而产生一个短期的进入令牌（token），用来代替密码，供第三方应用使用。
OAuth有四种获取令牌的方式，不管哪一种授权方式，第三方应用申请令牌之前，都必须先到系统备案，说明自己的身份，然后会拿到两个身份识别码：客户端 ID（client ID）和客户端密钥（client secret）。这是为了防止令牌被滥用，没有备案过的第三方应用，是不会拿到令牌的。
在前后端分离的情境下，我们常使用授权码方式，指的是第三方应用先申请一个授权码，然后再用该码获取令牌。

![16cd59ea3ff2066b](/images/16cd59ea3ff2066b.png)