---
title: "Http ContentType"
date: 2021-08-31T11:44:00+08:00
draft: false
---

## Content-Type

Content-Type 实体头部用于指示资源的 MIME 类型  `media type` 。

在响应中，Content-Type标头告诉客户端实际返回的内容的内容类型。

浏览器会在某些情况下进行MIME查找，并不一定遵循此标题的值; 为了防止这种行为，可以将标题 X-Content-Type-Options 设置为 nosniff。

在请求中 (如POST 或 PUT)，客户端告诉服务器实际发送的数据类型。

### 句法

```bash
Content-Type: text/html; charset=utf-8
Content-Type: multipart/form-data; boundary=something
```

### 指令

- media-type 资源或数据的 MIME type 。
- charset 字符编码标准。
- boundary 对于多部分实体，boundary 是必需的，其包括来自一组字符的1到70个字符，已知通过电子邮件网关是非常健壮的，而不是以空白结尾。它用于封装消息的多个部分的边界。

### 例子

在通过HTML form提交生成的POST请求中，请求头的Content-Type由<form>元素上的enctype属性指定

```html
<form action="/" method="post" enctype="multipart/form-data">
  <input type="text" name="description" value="some text">
  <input type="file" name="myFile">
  <button type="submit">Submit</button>
</form>
```

请求头看起来像这样（在这里省略了一些 headers）：

```bash
POST /foo HTTP/1.1
Content-Length: 68137
Content-Type: multipart/form-data; boundary=---------------------------974767299852498929531610575

---------------------------974767299852498929531610575
Content-Disposition: form-data; name="description"

some text
---------------------------974767299852498929531610575
Content-Disposition: form-data; name="myFile"; filename="foo.txt"
Content-Type: text/plain

(content of the uploaded file foo.txt)
---------------------------974767299852498929531610575
```


## HTTP content-type

Content-Type（内容类型），一般是指网页中存在的 Content-Type，用于定义网络文件的类型和网页的编码，决定浏览器将以什么形式、什么编码读取这个文件，这就是经常看到一些 PHP 网页点击的结果却是下载一个文件或一张图片的原因。

Content-Type 标头告诉客户端实际返回的内容的内容类型。

语法格式：

Content-Type: text/html; charset=utf-8
Content-Type: multipart/form-data; boundary=something

常见的媒体格式类型如下：

```
  text/html ： HTML格式
  text/plain ：纯文本格式
  text/xml ： XML格式
  image/gif ：gif图片格式
  image/jpeg ：jpg图片格式
  image/png：png图片格式

```
以application开头的媒体格式类型：
```
  application/xhtml+xml ：XHTML格式
  application/xml： XML数据格式
  application/atom+xml ：Atom XML聚合格式
  application/json： JSON数据格式
  application/pdf：pdf格式
  application/msword ： Word文档格式
  application/octet-stream ： 二进制流数据（如常见的文件下载）
  application/x-www-form-urlencoded ： <form encType=””>中默认的encType，form表单数据被编码为key/value格式发送到服务器（表单默认的提交数据的格式）
```
另外一种常见的媒体格式是上传文件之时使用的：
```
  multipart/form-data ： 需要在表单中进行文件上传时，就需要使用该格式
``` 


- https://qastack.cn/superuser/1277819/why-does-chrome-sometimes-download-a-pdf-instead-of-opening-it-duplicate