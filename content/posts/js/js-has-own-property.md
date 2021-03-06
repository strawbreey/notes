---
title: "Js Has Own Property"
date: 2020-12-25T10:08:24+08:00
draft: false
---

## avaScript 判断对象中是否有某属性

判断对象中是否有某属性的常见方式总结，不同的场景要使用不同的方式。

### 1. 点( . )或者方括号( [ ] )
　　通过点或者方括号可以获取对象的属性值，如果对象上不存在该属性，则会返回undefined。当然，这里的“不存在”指的是对象自身和原型链上都不存在，如果原型链有该属性，则会返回原型链上的属性值。

```js
// 创建对象
let test = {name : 'lei'}
// 获取对象的自身的属性
test.name            //"lei"
test["name"]         //"lei"

// 获取不存在的属性
test.age             //undefined

// 获取原型上的属性
test["toString"]     //toString() { [native code] }

// 新增一个值为undefined的属性
test.un = undefined

test.un              //undefined    不能用在属性值存在，但可能为 undefined的场景
```
　
所以，我们可以根据 Obj.x !== undefined 的返回值 来判断Obj是否有x属性。


这种方式很简单方便，局限性就是：不能用在x的属性值存在，但可能为 undefined的场景。 in运算符可以解决这个问题

### 2. in 运算符
　　
MDN 上对in运算符的介绍：如果指定的属性在指定的对象或其原型链中，则in 运算符返回true。

```js
'name' in test        //true
'un' in test             //true
'toString' in test    //true
'age' in test           //false
```

示例中可以看出，值为undefined的属性也可正常判断。

这种方式的局限性就是无法区分自身和原型链上的属性，在只需要判断自身属性是否存在时，这种方式就不适用了。这时需要hasOwnProperty()

 

### 3.hasOwnProperty()

```js
test.hasOwnProperty('name')          // true   自身属性
test.hasOwnProperty('age')           // false  不存在
test.hasOwnProperty('toString')      // false  原型链上属性
```

可以看到，只有自身存在该属性时，才会返回true。适用于只判断自身属性的场景。

### 总结

三种方式各有优缺点，不同的场景使用不同的方式，有时还需要结合使用，比如遍历自身属性的时候，就会把 for ··· in  ···和 hasOwnProperty()结合使用。