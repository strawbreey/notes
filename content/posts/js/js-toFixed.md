---
title: "Js ToFixed"
date: 2021-07-26T11:17:13+08:00
draft: false
showonlyimage: true
image: "img/portfolio/business-card-26.jpg"
---


toFixed() 方法使用定点表示法来格式化一个数值。

```js
function financial(x) {
  return Number.parseFloat(x).toFixed(2);
}

console.log(financial(123.456));
// expected output: "123.46"

console.log(financial(0.004));
// expected output: "0.00"

console.log(financial('1.23e+5'));
// expected output: "123000.00"

```

### numObj.toFixed(digits)

- digits  小数点后数字的个数；介于 0 到 20 （包括）之间，实现环境可能支持更大范围。如果忽略该参数，则默认为 0。