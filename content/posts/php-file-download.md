---
title: "File Download"
date: 2021-04-22T15:51:07+08:00
draft: false
---

## 文件下载

```js
  this.http.get('web/file/downloads',  { params: {
      id: id
    },
    responseType: 'blob',
    observe: 'response'
  }).subscribe(resp => {
    if (resp.headers.get('Content-Disposition').split(';')[1].split('=')[1]) {
      // 新建a标签
      const eleLink = document.createElement('a');
      // 获取文件名
      eleLink.download = resp.headers.get('Content-Disposition').split(';')[1].split('=')[1];
      eleLink.style.display = 'none';
      // 字符内容转变成blob地址
      eleLink.href = URL.createObjectURL(resp.body);
      // 触发点击
      document.body.appendChild(eleLink);
      eleLink.click();
      // 然后移除
      document.body.removeChild(eleLink);
    }
  });
```

3. 非同源跨域下载


## php 输出blob 

1. file_get_contents
```php
  file_get_contents(url)
```

2. fgets
```php
  $fp = fopen($file_path, 'rb');
  while (!feof($fp)) {
      $buffer = fgets($fp, 16384);
      echo $buffer;
  }
  fclose($fp);
```

3. fread

```php
  $handle = fopen($file_path, "rb");
  $contents = fread($handle, filesize($file_path));
  fclose($handle);
```