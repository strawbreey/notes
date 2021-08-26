---
title: "Js Random"
date: 2021-08-19T16:49:59+08:00
draft: false
---

第一种：使用randomString，e表示长度，默认32位

```js
function randomString(e) {  
  e = e || 32;
  var t = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678",
  a = t.length,
  n = "";
  for (i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
  return n
}
alert(randomString(6));
```

第二种：生成随机数

```js
// 本例子代表生成100000-999999的随机数
function GetRandomNum(Min,Max)
{
var Range = Max - Min;
var Rand = Math.random();
return(Min + Math.round(Rand * Range));
}
var num = GetRandomNum(10000,999999);
alert(num);
```

第三种：对定义的数组字符集进行随机选取

```js
var str = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'];
function generateMixed(n) {
   var res = "";
   for(var i = 0; i < n ; i ++) {
     var id = Math.ceil(Math.random()*35);
     res += str[id];
   }
   return res;
}
alert(generateMixed(6));
```

第四种：生成随机数转成36进制，再截取部分

```js
//Math.random() 生成随机数字, eg: 0.123456
//.toString(36) 转化成36进制 : "0.4fzyo82mvyr"
//.slice(-8); 截取最后八位 : "yo82mvyr"
var str = Math.random().toString(36).slice(-6);
alert(str);
```

第五种：对字符串集合随机排列，随机输出指定的长度

```js
function randomString(length) {
  var str = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
  var result = '';
  for (var i = length; i > 0; --i) 
    result += str[Math.floor(Math.random() * str.length)];
  return result;
}
alert(randomString(6));
````