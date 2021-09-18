---
title: "Http ContentType"
date: 2021-08-31T11:44:00+08:00
draft: false
---

### Content-Type

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

1. Content-Type 在HTML表单中

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