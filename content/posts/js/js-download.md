---
title: "Js Download"
date: 2021-09-28T16:27:42+08:00
draft: false
---


## 前端下载

1. HTML <a> download Attribute

```html
<a href="/images/myw3schoolsimage.jpg" download>
```

默认情况下会直接下载图片 (jpg, pdf, gif, png 等), 当跨域时, 浏览器会新开一个tab 页面打开

2. 动态创建一个a 标签

```js
  downloadByUrl (url, name = null) {
    const eleLink = document.createElement('a');
    eleLink.download =  name ? name : 'file';
    eleLink.style.display = 'none';
    eleLink.href = url;
    document.body.appendChild(eleLink);
    eleLink.click();
    document.body.removeChild(eleLink);
  }
```

效果如上，当跨域时, 浏览器会新开一个tab 页面打开

3. 使用ajax 获取远程的二进制数据流

```js
  downloadImageByUrl (url, filename = null) {
    this.http.get(url, {
      responseType: "blob",
    }).subscribe(resp=>{
      // console.log(resp)
      var eleLink = document.createElement('a');
      eleLink.download = filename ? filename : '';
      eleLink.style.display = 'none';
      // 字符内容转变成blob地址
      var blob = new Blob([resp]);
      eleLink.href = URL.createObjectURL(blob);
      // 触发点击
      document.body.appendChild(eleLink);
      eleLink.click();
      // 然后移除
      document.body.removeChild(eleLink);
    })
  }
```