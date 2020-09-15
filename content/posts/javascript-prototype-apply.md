---
title: "Javascript Prototype Apply"
date: 2020-09-15T14:39:37+08:00
draft: true
---

The apply() method calls a function with a given this value, and arguments provided as an array (or an array-like object).

apply() 方法调用一个具有给定this值的函数，以及以一个数组（或类数组对象）的形式提供的参数。

```js
const numbers = [5, 6, 2, 3, 7];

const max = Math.max.apply(null, numbers); // expected output: 7

const min = Math.min.apply(null, numbers); // expected output: 2
```

### Syntax

func.apply(thisArg, [argsArray])

thisArg
- The value of this provided for the call to func.

- Note that this may not be the actual value seen by the method: if the method is a function in non-strict mode code, null and undefined will be replaced with the global object, and primitive values will be boxed. This argument is required.

argsArray Optional
- An array-like object, specifying the arguments with which func should be called, or null or undefined if no arguments should be provided to the function.

- Starting with ECMAScript 5 these arguments can be a generic array-like object instead of an array. See below for browser compatibility information.

Return value

The result of calling the function with the specified this value and arguments.


Description


### Examples

#### Using apply to append an array to another

```js
const array = ['a', 'b'];
const elements = [0, 1, 2];
array.push.apply(array, elements); // ["a", "b", 0, 1, 2]
array.push(...element);  // ["a", "b", 0, 1, 2]
array.push(element);  // ["a", "b", [0, 1, 2]]

```

#### Using apply and built-in functions

```js
// min|max number in an array
const numbers = [5, 6, 2, 3, 7];

// using Math.min|Math.max apply
let max = Math.max.apply(null, numbers); 
// This about equal to Math.max(numbers[0], ...)
// or Math.max(5, 6, ...)

let min = Math.min.apply(null, numbers);

// vs. simple loop based algorithm
max = -Infinity, min = +Infinity;

for (let i = 0; i < numbers.length; i++) {
  if (numbers[i] > max) {
    max = numbers[i];
  }
  if (numbers[i] < min) {
    min = numbers[i];
  }
}
```

#### Using apply to chain constructors
```js
Function.prototype.construct = function(aArgs) {
  let oNew = Object.create(this.prototype);
  this.apply(oNew, aArgs);
  return oNew;
};
```

```js
function MyConstructor() {
  for (let nProp = 0; nProp < arguments.length; nProp++) {
    this['property' + nProp] = arguments[nProp];
  }
}

let myArray = [4, 'Hello world!', false];
let myInstance = MyConstructor.construct(myArray);

console.log(myInstance.property1);                // logs 'Hello world!'
console.log(myInstance instanceof MyConstructor); // logs 'true'
console.log(myInstance.constructor);
```