---
title: "Hot Module Replacement"
date: 2020-09-25T10:14:13+08:00
draft: false
---

### 浏览器刷新页面的方法


#### Location.reload

```js
// 语法
location.reload([forcedReload])

// demo
window.location.reload(true);
```
`forcedReload` 可选 该参数要求为 布尔 类型，当取值为 true 时，将强制浏览器从服务器重新获取当前页面资源，而不是从浏览器的缓存中读取，如果取值为 false 或不传该参数时，浏览器则可能会从缓存中读取当前页面。


无缓存刷新页面但页面引用的资源还是可能使用缓存，大多数浏览器可以通过设置在打开开发者工具时禁用缓存实现无缓存需求）


#### Location.replace

```js
// 语法
object.replace(url);

// demo
document.location.replace('https://developer.mozilla.org/en-US/docs/Web/API/Location/reload');
```

Location.replace() 方法以给定的URL来替换当前的资源。 与assign() 方法 不同的是，调用 replace() 方法后，当前页面不会保存到会话历史中（session History），这样，用户点击回退按钮时，将不会再跳转到该页面。

因违反安全规则导致的赋值失败，浏览器将会抛出类型为 SECURITY_ERROR 的 DOMException 异常。当调用该方法的脚本所属的源与拥有 Location 对象所属源不同时，通常情况会发生这种异常,此时通常该脚本是存在不同的域下。



#### Location.assign

```js
// 语法
location.assign(url);

// demo
document.location.assign('https://developer.mozilla.org/zh-CN/docs/Web/API/Location/reload');
```

Location.assign  方法会触发窗口加载并显示指定的URL的内容。 如果由于安全原因无法执行跳转，那么会抛出一个 SECURITY_ERROR 类型的 DOMException。当调用此方法的脚本来源和页面的 Location 对象中定义的来源隶属于不同域的时候，就会抛出上述错误


#### JavaScript刷新页面的几种方法

```js
  history.go(0)
  location.reload()
  location=location
  location.assign(location)
  document.execCommand('Refresh')
  window.navigate(location)
  location.replace(location)
  document.URL=location.href
```

#### 自动刷新页面的方法
```html
<!-- 页面自动刷新 -->
<meta http-equiv="refresh" content="20">

<!-- 页面自动跳转 -->
<meta http-equiv="refresh" content="20;url=http://www.baidu.com">

<!-- js自动刷新 -->
<script type="text/javascript">
function myrefresh()
{
   window.location.reload();
}
setTimeout('myrefresh()',1000); //指定1秒刷新一次
</script>
```

### 浏览器更热新

Hot Module Replacement，简称HMR，无需完全刷新整个页面的同时，更新模块。HMR的好处，在日常开发工作中体会颇深：节省宝贵的开发时间、提升开发体验。

- 页面刷新，不保留页面状态，就是简单粗暴，直接window.location.reload()。
- 基于WDS ([Webpack-dev-server](https://github.com/webpack/webpack-dev-server))的模块热替换，只需要局部刷新页面上发生变化的模块，同时可以保留当前的页面状态，比如复选框的选中状态、输入框的输入等。

#### Webpack-dev-server 实现原理

修改代码 -> webpack-dev-middleware (监听本地代码变化) -> Webpack-dev-server(webpack 变化通知) -> websoket通知浏览器 -> 浏览器接收通知 -> 使用reload 变化?

首先是建立起浏览器端和服务器端之间的通信，浏览器会接收服务器端推送的消息，如果需要热更新，浏览器发起http请求去服务器端获取打包好的资源解析并局部刷新页面。

#### webpack-dev-server 在三个框架中使用情况

angular -> angular-cli -> webpack -> webpack-dev-server
react -> angular-cli -> webpack -> webpack-dev-server


[angular webpack-dev-server](https://github.com/angular/angular/blob/3e57ca1d98ab6d263762a6e9a57512e50a1fb27e/aio/content/guide/deployment.md)

### 参考链接 
- [Location.reload()](https://developer.mozilla.org/zh-CN/docs/Web/API/Location/reload)
- [Location.replace()](https://developer.mozilla.org/zh-CN/docs/Web/API/Location/replace)
- [轻松理解webpack热更新原理](https://juejin.im/post/6844904008432222215)