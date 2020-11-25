---
title: "Js Ajax"
date: 2020-11-12T11:38:37+08:00
draft: false
---

Asynchronous JavaScript + XML（异步JavaScript和XML）, 其本身不是一种新技术，而是一个在 2005年被Jesse James Garrett提出的新术语，用来描述一种使用现有技术集合的‘新’方法，包括: HTML 或 XHTML,  CSS, JavaScript, DOM, XML, XSLT, 以及最重要的 [XMLHttpRequest](https://developer.mozilla.org/zh-CN/docs/Web/API/XMLHttpRequest)。当使用结合了这些技术的AJAX模型以后， 网页应用能够快速地将增量更新呈现在用户界面上，而不需要重载（刷新）整个页面。这使得程序能够更快地回应用户的操作。

简单点说，就是使用 XMLHttpRequest 对象与服务器通信。 它可以使用JSON，XML，HTML和text文本等格式发送和接收数据。AJAX最吸引人的就是它的“异步”特性，也就是说它可以在不重新刷新页面的情况下与服务器通信，交换数据，或更新页面。

### XMLHTTP 的发展 历史

- XMLHTTP最初是由微软公司发明的，在Internet Explorer 5.0 (1999年3月18日)中用作ActiveX对象，可通过JavaScript、VBScript或其它浏览器支持的脚本语言访问。

- Mozilla的开发人员后来在Mozilla 1.0中实现了一个兼容的版本。

- 苹果电脑公司在Safari 1.2中开始支持XMLHTTP，

- Opera从8.0版开始也宣布支持XMLHTTP。

### Ajax 基础知识


```js

// Step 1
// 发送http请求
// Old compatibility code, no longer needed.
if (window.XMLHttpRequest) { // Mozilla, Safari, IE7+ ...
    httpRequest = new XMLHttpRequest();
} else if (window.ActiveXObject) { // IE 6 and older
    httpRequest = new ActiveXObject("Microsoft.XMLHTTP");
}

// 告诉XMLHttp请求对象是由哪一个函数响应
httpRequest.onreadystatechange = function(){
    // Process the server response here.
};


// 声明当你接到响应后要做什么，你要发送一个实际的请求

httpRequest.open('GET', 'http://www.example.org/some.file', true);
httpRequest.send();

// 如果你使用 POST 数据，那就需要设置请求的MIME类型。比如，在调用 send() 方法获取表单数据前要有下面这个：
httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

// Step2
// 函数要检查请求的状态。如果状态的值是 XMLHttpRequest.DONE 

// 全部readyState状态值都在 XMLHTTPRequest.readyState，如下也是：

// 0 (未初始化) or (请求还未初始化)
// 1 (正在加载) or (已建立服务器链接)
// 2 (加载成功) or (请求已接受)
// 3 (交互) or (正在处理请求)
// 4 (完成) or (请求已完成并且响应已准备好)

if (httpRequest.readyState === XMLHttpRequest.DONE) {
    // Everything is good, the response was received.
} else {
    // Not ready yet.
}

// 接下来，点击HTTP响应的 response code。 可能的响应码都已经列在 W3C这个列表里。在下面的例子中，我们通过检查响应码 200 OK 判断AJAX有没有成功。

if (httpRequest.status === 200) {
    // Perfect!
} else {
    // There was a problem with the request.
    // For example, the response may have a 404 (Not Found)
    // or 500 (Internal Server Error) response code.
}


  // Step3
  var httpRequest;
  // 监听ajaxButton按钮的点击事件
  document.getElementById("ajaxButton").addEventListener('click', makeRequest);

  function makeRequest() {
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      alert('Giving up :( Cannot create an XMLHTTP instance');
      return false;
    }

    httpRequest.onreadystatechange = function alertContents() {
      if (httpRequest.readyState === XMLHttpRequest.DONE) {
        if (httpRequest.status === 200) {
          // 返回响应的数据
          alert(httpRequest.responseText);
        } else {
          alert('There was a problem with the request.');
        }
      }
    };

    httpRequest.open('GET', 'test.html');
    httpRequest.send();
  }

  function
```


### Fetch Api

Fetch API 提供一个获取资源的接口

### server-sent 
s
在过去，一个网页必须发送请求到服务器来获取新的数据，也就是说，网页必须主动向服务器请求数据。有了server-sent events之后，服务器可以向网页推送消息，使得服务器可以随时向网页传送数据。这些发送过来的消息可以看作是带有数据的事件

### 参考资料

- [Ajax](https://developer.mozilla.org/zh-CN/docs/Web/Guide/AJAX)

- [xhr 规范](https://xhr.spec.whatwg.org/)