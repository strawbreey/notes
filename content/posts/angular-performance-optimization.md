---
title: "Angular Performance Optimization"
date: 2020-09-16T17:39:04+08:00
draft: false
---

### Angualr 性能优化

angular编译的js基本都很大达到了mb级别，加载起来很占用带宽，不过angualr也提供了编译时压缩代码的功能，angular build指定为 --prod即可

```js
// 注释下面这一行可以激活延迟加载策略
export const routing: ModuleWithProviders = RouterModule.forRoot(routes, {
  // preloadingStrategy: PreloadAllModules,
  // useHash: true
});
```

`ng build --prod --build--optimizer` // 编译后进一步压缩文件的大小


### Angular脏检查（Change Detection）机制

当我们遇到了性能问题，首先我们需要明白是什么导致我们的Angular应用变慢了。众所周知，在Angular中使用了双向绑定连接了model和DOM，当我们在component中改变了model的值，DOM中对应的值也会改变。但是Angular是怎么知道它需要在什么时候去更新DOM呢？这就是Angular Change Detection机制需要完成的事。

而通常情况下，Angular Change Detection会由下几类事件触发：

- 所有的浏览器事件（click, mouseover, keyup等）
- Ajax异步请求
- setTimeout()和setInterval()

一旦我们触发了这些事件，Change Detection机制就会开始运行。在这些事件所绑定的方法执行完成之后，Change Detection会根据当前模板中数据绑定的情况来更新DOM。

Angular的性能好坏就和每一次Change Detection周期的执行时间长短有关。理想情况下如果每次Change Detection周期能控制在17ms以内，那么界面将会十分流畅。而在这个周期中，触发事件是周期的开始，这个时间可以忽略不计。浏览器渲染是浏览器的行为，不受Angular控制。所以我们可以优化的就是在Event Handler执行和Chanege Detection检测绑定这两个步骤。

#### 减少Event Handler的运行时间

### CDN加速


### nginx缓存
```shell
events {
  #的最大连接数（包含所有连接数）1024
  worker_connections  1024;  ## Default: 1024
}
 
http{
 
   # 代理缓存配置 nginx根目录指定缓存文件夹 kawa_cachedata名字自己定义跟proxy_cache_path对应上
   proxy_cache_path "./kawa_cachedata"  levels=1:2 keys_zone=kawacache:256m inactive=1d 
   
   max_size=1000g;

   server {
     listen 80;
     location /{
        #使用缓存名称
        proxy_cache kawacache;
        #对以下状态码实现缓存
        proxy_cache_valid 200 206 304 301 302 1d;
        #缓存的key为url地址
        proxy_cache_key $request_uri;
        add_header X-Cache-Status $upstream_cache_status;
        #反向代理地址
        proxy_pass http://127.0.0.1:8080;
      }
   
   }
}
```