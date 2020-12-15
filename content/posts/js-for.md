---
title: "Js For"
date: 2020-12-14T23:37:43+08:00
draft: false
---

### 1. for循环

1. for有三个表达式：①声明循环变量；②判断循环条件；③更新循环变量；三个表达式之间，用`;`分割，for循环三个表达式都可以省略，但是两个“;”缺一 不可。

2. for循环的执行特点：先判断再执行，与while相同　　

3. for循环三个表达式都可以有多部分组成，第二部分多个判断条件用&& ||连接，第一三部分用逗号分割

```js
for (var num = 1; num<10; num++) {
    console.log(num); //1,2,3,4,5,6,7,8,9
}
```

### 2. while循环   

while循环()中的表达式，运算结果可以是各种类型，但是最终都会转为真假，转换规则如下。

- Boolean：true为真，false为假；
- String：空字符串为假，所有非空字符串为真；
- Number：0为假，一切非0数字为真；
- null/Undefined/NaN:全为假；
- Object：全为真。

```js
var num = 1;//1、声明循环变量
while (num<10) {//2、判断循环条件;
	console.log(num);//3、执行循环体操作；
	num++;//4、更新循环变量；
}
```

### 3. do-while循环  

while循环特点：先判断后执行；

do-while循环特点：先执行再判断，即使初始条件不成立，do-while循环至少执行一次，也就是说do-while循环比while循环多执行一次。

```js
var num = 10;
            
do {
  console.log(num);//10 9 8 7 6 5 4 3 2 1 0
  num--;
} while(num>=0);
            
console.log(num);//-1
```

### 4. for - in

for - in语句用于对数组或者对象的属性进行循环操作。

for - in循环中的代码每执行一次，就会对数组或者对象的属性进行一次操作。

```js
let obj={'name':'programmer','age':'22','height':'180'};
for (let i in obj) {
	console.log(i,obj[i])
}
```

### 5. for - of

for...of循环可以使用的范围包括数组、Set 和 Map 结构、某些类似数组的对象（比如arguments对象、DOM NodeList 对象）、后文的 Generator 对象，以及字符 串。

JavaScript 原有的for-in循环，只能获得对象的键名，不能直接获取键值。ES6 提供for...of循环，允许遍历获得键值

1. 数组操作：

```js
  var arr = ['a', 'b', 'c', 'd'];

	for (let a in arr) {
		console.log(a); // 0 1 2 3
	}
	
	for (let a of arr) {
		console.log(a); // a b c d
	}
```

2. 类似数组的对象操作：

```js
	// 字符串
	var str = "hello";
	
	for (let s of str) {
		console.log(s); // h e l l o
	}
	
	// DOM NodeList对象
	let paras = document.querySelectorAll("p");
	
	for (let p of paras) {
		p.classList.add("test");
	}
	
	// arguments对象
	function printArgs() {
		for (let x of arguments) {
			console.log(x);
		}
	}
	printArgs('a', 'b');// 'a' 'b'
```
 

###   循环控制语句

1. break：跳出本层循环，继续执行循环后面的语句。如果循环有多层，则break只能跳出一层。
2. continue：跳过本次循环剩余的代码，继续执行下一次循环。
	- 对与for循环，continue之后执行的语句，是循环变量更新语句i++；
	- 对于while、do-while循环，continue之后执行的语句，是循环条件判断；
　
因此，使用这两个循环时，必须将continue放到i++之后使用，否则continue将跳过 i++进入死循环。

    