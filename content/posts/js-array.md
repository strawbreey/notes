---
title: "Js Array"
date: 2020-09-28T11:06:27+08:00
draft: false
---

### Array.from()

```js
// Syntax
// arrayLike: An array-like or iterable object to convert to an array.
// mapFn Optional: Map function to call on every element of the array.
// thisArg Optional: Value to use as this when executing mapFn.

// return value:  A new Array instance.
Array.from(arrayLike [, mapFn [, thisArg]])

console.log(Array.from('foo'));
// expected output: Array ["f", "o", "o"]

console.log(Array.from([1, 2, 3], x => x + x));
// expected output: Array [2, 4, 6]
```

### Array.isArray()
```js

// Array.isArray(value)
// The value to be checked.

Array.isArray([1, 2, 3]);  // true
Array.isArray({foo: 123}); // false
Array.isArray('foobar');   // false
Array.isArray(undefined);  // false

// polyfill
if (!Array.isArray) {
  Array.isArray = function(arg) {
    return Object.prototype.toString.call(arg) === '[object Array]';
  };
}

//Examples: 
// return true
Array.isArray([]);
Array.isArray([1]);
Array.isArray(new Array());
Array.isArray(new Array('a', 'b', 'c', 'd'));
Array.isArray(new Array(3));
Array.isArray(Array.prototype); 

// return false
Array.isArray();
Array.isArray({});
Array.isArray(null);
Array.isArray(undefined);
Array.isArray(17);
Array.isArray('Array');
Array.isArray(true);
Array.isArray(false);
Array.isArray(new Uint8Array(32));
Array.isArray({ __proto__: Array.prototype });
```

###  Array.of()

Array.of() 方法创建一个具有可变数量参数的新数组实例，而不考虑参数的数量或类型。

 Array.of() 和 Array 构造函数之间的区别在于处理整数参数：Array.of(7) 创建一个具有单个元素 7 的数组，而 Array(7) 创建一个长度为7的空数组

```js
Array.of(7);       // [7] 
Array.of(1, 2, 3); // [1, 2, 3]

Array(7);          // array of 7 empty slots
Array(1, 2, 3);    // [1, 2, 3]
```
#### Syntax
```js
Array.of(element0[, element1[, ...[, elementN]]])
```

#### polyfill
```js
if (!Array.of) {
  Array.of = function() {
    return Array.prototype.slice.call(arguments);
    // Or 
    let vals = [];
    for(let prop in arguments){
        vals.push(arguments[prop]);
    }
    return vals;
  }
}
```


### concat()

concat() 方法用于合并两个或多个数组。此方法不会更改现有数组，而是返回一个新数组

```js
const array1 = ['a', 'b', 'c'];
const array2 = ['d', 'e', 'f'];
const array3 = array1.concat(array2);
```

### copyWithin()

方法浅复制数组的一部分到同一数组中的另一个位置，并返回它，不会改变原数组的长度。

arr.copyWithin(target[, start[, end]])

target 为基底的索引，复制序列到该位置。如果是负数，target 将从末尾开始计算。
如果 target 大于等于 arr.length，将会不发生拷贝。如果 target 在 start 之后，复制的序列将被修改以符合 arr.length。

start 为基底的索引，开始复制元素的起始位置。如果是负数，start 将从末尾开始计算。
如果 start 被忽略，copyWithin 将会从0开始复制。

end 为基底的索引，开始复制元素的结束位置。copyWithin 将会拷贝到该位置，但不包括 end 这个位置的元素。如果是负数， end 将从末尾开始计算。
如果 end 被忽略，copyWithin 方法将会一直复制至数组结尾（默认为 arr.length）。

 ```js
const array1 = ['a', 'b', 'c', 'd', 'e'];

// copy to index 0 the element at index 3
array1.copyWithin(0, 3, 4)
// expected output: Array ["d", "b", "c", "d", "e"]

// copy to index 1 all elements from index 3 to the end
array1.copyWithin(1, 3)
// expected output: Array ["d", "d", "e", "d", "e"]

[1, 2, 3, 4, 5].copyWithin(-2) 
// [1, 2, 3, 1, 2] 从下标4开始复制数据

[1, 2, 3, 4, 5].copyWithin(0, 3)
// [4, 5, 3, 4, 5] 复制数据到下标零处, 从下标3开始 

[1, 2, 3, 4, 5].copyWithin(0, 3, 4)
// [4, 2, 3, 4, 5]

[1, 2, 3, 4, 5].copyWithin(-2, -3, -1)
// [1, 2, 3, 3, 4]

[].copyWithin.call({length: 5, 3: 1}, 0, 3);
// {0: 1, 3: 1, length: 5}

// ES2015 Typed Arrays are subclasses of Array
var i32a = new Int32Array([1, 2, 3, 4, 5]);

i32a.copyWithin(0, 2);
// Int32Array [3, 4, 5, 4, 5]

// On platforms that are not yet ES2015 compliant: 
[].copyWithin.call(new Int32Array([1, 2, 3, 4, 5]), 0, 3, 4);
// Int32Array [4, 2, 3, 4, 5]
 ```

### entries

方法返回一个新的Array Iterator对象，该对象包含数组中每个索引的键/值对

```js
const array1 = ['a', 'b', 'c'];

const iterator1 = array1.entries();

iterator1.next().value;
// expected output: Array [0, "a"]

iterator1.next().value;
// expected output: Array [1, "b"]
```

### every 

方法测试一个数组内的所有元素是否都能通过某个指定函数的测试

若收到一个空数组，此方法在一切情况下都会返回 true。

callback 用来测试每个元素的函数，它可以接收三个参数：

element 用于测试的当前值。  
index(可选) 用于测试的当前值的索引。  
array(可选) 调用 every 的当前数组。thisArg 执行 callback 时使用的 this 值。 

如果回调函数的每一次返回都为 truthy 值，返回 true ，否则返回 false。

#### 语法
arr.every(callback(element[, index[, array]])[, thisArg])



```js
const isBelowThreshold = (currentValue) => currentValue < 40;

const array1 = [1, 30, 39, 29, 10, 13];

console.log(array1.every(isBelowThreshold));
// expected output: true

array1.every((value) => value < 50); 
// expected output: true

array2 = []
array1.every(); //true 

```

#### Polyfill

```js
if (!Array.prototype.every) {
  Array.prototype.every = function(callbackfn, thisArg) {
    'use strict';
    var T, k;

    if (this == null) {
      throw new TypeError('this is null or not defined');
    }

    // 1. Let O be the result of calling ToObject passing the this 
    //    value as the argument.
    var O = Object(this);

    // 2. Let lenValue be the result of calling the Get internal method
    //    of O with the argument "length".
    // 3. Let len be ToUint32(lenValue).
    var len = O.length >>> 0;

    // 4. If IsCallable(callbackfn) is false, throw a TypeError exception.
    if (typeof callbackfn !== 'function') {
      throw new TypeError();
    }

    // 5. If thisArg was supplied, let T be thisArg; else let T be undefined.
    if (arguments.length > 1) {
      T = thisArg;
    }

    // 6. Let k be 0.
    k = 0;

    // 7. Repeat, while k < len
    while (k < len) {

      var kValue;

      // a. Let Pk be ToString(k).
      //   This is implicit for LHS operands of the in operator
      // b. Let kPresent be the result of calling the HasProperty internal 
      //    method of O with argument Pk.
      //   This step can be combined with c
      // c. If kPresent is true, then
      if (k in O) {

        // i. Let kValue be the result of calling the Get internal method
        //    of O with argument Pk.
        kValue = O[k];

        // ii. Let testResult be the result of calling the Call internal method
        //     of callbackfn with T as the this value and argument list 
        //     containing kValue, k, and O.
        var testResult = callbackfn.call(T, kValue, k, O);

        // iii. If ToBoolean(testResult) is false, return false.
        if (!testResult) {
          return false;
        }
      }
      k++;
    }
    return true;
  };
}
```

###  fill()

方法用一个固定值填充一个数组中从起始索引到终止索引内的全部元素。不包括终止索引。

arr.fill(value[, start[, end]])

value
    用来填充数组元素的值。
start 可选
    起始索引，默认值为0。
end 可选  
    终止索引，默认值为 this.length。

```js
const array1 = [1, 2, 3, 4];

// fill with 0 from position 2 until position 4
array1.fill(0, 2, 4)
// expected output: [1, 2, 0, 0]

// fill with 5 from position 1
array1.fill(5, 1)
// expected output: [1, 5, 5, 5]

array1.fill(6)
// expected output: [6, 6, 6, 6]

```

### filter() 

方法创建一个新数组, 其包含通过所提供函数实现的测试的所有元素。 

filter 为数组中的每个元素调用一次 callback 函数，并利用所有使得 callback 返回 true 或等价于 true 的值的元素创建一个新数组。callback 只会在已经赋值的索引上被调用，对于那些已经被删除或者从未被赋值的索引不会被调用。那些没有通过 callback 测试的元素会被跳过，不会被包含在新数组中。

var newArray = arr.filter(callback(element[, index[, array]])[, thisArg])

callback 用来测试数组的每个元素的函数。返回 true 表示该元素通过测试，保留该元素，false 则不保留。它接受以下三个参数：

element 数组中当前正在处理的元素。

index可选 正在处理的元素在数组中的索引。

array可选 调用了 filter 的数组本身。

thisArg可选
执行 callback 时，用于 this 的值。

返回值
一个新的、由通过测试的元素组成的数组，如果没有任何数组元素通过测试，则返回空数组。

```js
function isBigEnough(element) {
  return element >= 10;
}
var filtered = [12, 5, 8, 130, 44].filter(isBigEnough);
var filtered = [12, 5, 8, 130, 44].filter((element) => element > 40);

// 过滤 JSON 中的无效条目
var arrByID = arr.filter((item) => isNumber(item.id) && item.id !==0);

// 数组中搜索
arr.filter((el) => el.toLowerCase().indexOf(query.toLowerCase()) > -1)
```

#### Polyfill 

filter 被添加到 ECMA-262 标准第 5 版中，因此在某些实现环境中不被支持。

```js
if (!Array.prototype.filter){
  Array.prototype.filter = function(func, thisArg) {
    'use strict';
    if (!((typeof func === 'Function' || typeof func === 'function') && this))
        throw new TypeError();
   
    var len = this.length >>> 0,
        res = new Array(len), // preallocate array
        t = this, c = 0, i = -1;
    if (thisArg === undefined){
      while (++i !== len){
        // checks to see if the key was set
        if (i in this){
          if (func(t[i], i, t)){
            res[c++] = t[i];
          }
        }
      }
    }
    else{
      while (++i !== len){
        // checks to see if the key was set
        if (i in this){
          if (func.call(thisArg, t[i], i, t)){
            res[c++] = t[i];
          }
        }
      }
    }
   
    res.length = c; // shrink down array to proper size
    return res;
  };
}
```

### find()

方法返回数组中满足提供的测试函数的第一个元素的值

find方法对数组中的每一项元素执行一次 callback 函数，直至有一个 callback 返回 true。当找到了这样一个元素后，该方法会立即返回这个元素的值，否则返回 undefined。注意 callback 函数会为数组中的每个索引调用即从 0 到 length - 1，而不仅仅是那些被赋值的索引，这意味着对于稀疏数组来说，该方法的效率要低于那些只遍历有值的索引的方法。

```js
inventory.find(item => item.name == 'apples')); // { name: 'apples', quantity: 5 }


```

### Polyfill

本方法在ECMAScript 6规范中被加入，可能不存在于某些实现中

```js
// https://tc39.github.io/ecma262/#sec-array.prototype.find
if (!Array.prototype.find) {
  Object.defineProperty(Array.prototype, 'find', {
    value: function(predicate) {
     // 1. Let O be ? ToObject(this value).
      if (this == null) {
        throw new TypeError('"this" is null or not defined');
      }

      var o = Object(this);

      // 2. Let len be ? ToLength(? Get(O, "length")).
      var len = o.length >>> 0;

      // 3. If IsCallable(predicate) is false, throw a TypeError exception.
      if (typeof predicate !== 'function') {
        throw new TypeError('predicate must be a function');
      }

      // 4. If thisArg was supplied, let T be thisArg; else let T be undefined.
      var thisArg = arguments[1];

      // 5. Let k be 0.
      var k = 0;

      // 6. Repeat, while k < len
      while (k < len) {
        // a. Let Pk be ! ToString(k).
        // b. Let kValue be ? Get(O, Pk).
        // c. Let testResult be ToBoolean(? Call(predicate, T, « kValue, k, O »)).
        // d. If testResult is true, return kValue.
        var kValue = o[k];
        if (predicate.call(thisArg, kValue, k, o)) {
          return kValue;
        }
        // e. Increase k by 1.
        k++;
      }

      // 7. Return undefined.
      return undefined;
    }
  });
}
```

### findIndex 

方法返回数组中满足提供的测试函数的第一个元素的索引

findIndex方法对数组中的每个数组索引0..length-1（包括）执行一次callback函数，直到找到一个callback函数返回真实值（强制为true）的值。如果找到这样的元素，findIndex会立即返回该元素的索引。如果回调从不返回真值，或者数组的length为0，则findIndex返回-1。 与某些其他数组方法（如Array#some）不同，在稀疏数组中，即使对于数组中不存在的条目的索引也会调用回调函数。

```js
const array1 = [5, 12, 8, 130, 44];
const isLargeNumber = (element) => element > 13;
array1.findIndex(isLargeNumber);
// or 
array1.findIndex((element) => element > 13);
// expected output: 3

```

### flat

方法会按照一个可指定的深度递归遍历数组，并将所有元素与遍历到的子数组中的元素合并为一个新数组返回。

```js
[0, 1, 2, [3, 4]].flat();
// expected output: [0, 1, 2, 3, 4]

[0, 1, 2, [[[3, 4]]]].flat(2);
// expected output: [0, 1, 2, [3, 4]]

[1, 2, [3, 4]].flat() 
// [1, 2, 3, 4]

[1, 2, [3, 4, [5, 6]]].flat();
// [1, 2, 3, 4, [5, 6]]

[1, 2, [3, 4, [5, 6]]].flat(2);
// [1, 2, 3, 4, 5, 6]

//使用 Infinity，可展开任意深度的嵌套数组
[1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]].flat(Infinity);
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
```

### flatMap

方法首先使用映射函数映射每个元素，然后将结果压缩成一个新数组。它与 map 连着深度值为1的 flat 几乎相同，但 flatMap 通常在合并成一种方法的效率稍微高一些。

### forEach

方法对数组的每个元素执行一次给定的函数。

forEach() 方法按升序为数组中含有效值的每一项执行一次 callback 函数，那些已删除或者未初始化的项将被跳过（例如在稀疏数组上）。

注意： 除了抛出异常以外，没有办法中止或跳出 forEach() 循环。如果你需要中止或跳出循环，forEach() 方法不是应当使用的工具。

```js
const array1 = ['a', 'b', 'c'];
array1.forEach(element => console.log(element));
```

### includes

includes() 方法用来判断一个数组是否包含一个指定的值，根据情况，如果包含则返回 true，否则返回false。

从fromIndex 索引处开始查找 valueToFind。如果为负值，则按升序从 array.length + fromIndex 的索引开始搜

```js
[1, 2, 3].includes(2)
// expected output: true

['cat', 'dog', 'bat'].includes('cat')
// expected output: true

['cat', 'dog', 'bat'].includes('at')
// expected output: false

[1, 2, 3].includes(2);     // true
[1, 2, 3].includes(4);     // false
[1, 2, 3].includes(3, 3);  // false
[1, 2, 3].includes(3, -1); // true
[1, 2, NaN].includes(NaN); // true

var arr = ['a', 'b', 'c'];

arr.includes('c', 3);   // false
arr.includes('c', 100); // false

var arr = ['a', 'b', 'c'];

arr.includes('a', -100); // true
arr.includes('b', -100); // true
arr.includes('c', -100); // true
arr.includes('a', -2); // false
```

### Polyfill 
```js
// https://tc39.github.io/ecma262/#sec-array.prototype.includes
if (!Array.prototype.includes) {
  Object.defineProperty(Array.prototype, 'includes', {
    value: function(valueToFind, fromIndex) {

      if (this == null) {
        throw new TypeError('"this" is null or not defined');
      }

      // 1. Let O be ? ToObject(this value).
      var o = Object(this);

      // 2. Let len be ? ToLength(? Get(O, "length")).
      var len = o.length >>> 0;

      // 3. If len is 0, return false.
      if (len === 0) {
        return false;
      }

      // 4. Let n be ? ToInteger(fromIndex).
      //    (If fromIndex is undefined, this step produces the value 0.)
      var n = fromIndex | 0;

      // 5. If n ≥ 0, then
      //  a. Let k be n.
      // 6. Else n < 0,
      //  a. Let k be len + n.
      //  b. If k < 0, let k be 0.
      var k = Math.max(n >= 0 ? n : len - Math.abs(n), 0);

      function sameValueZero(x, y) {
        return x === y || (typeof x === 'number' && typeof y === 'number' && isNaN(x) && isNaN(y));
      }

      // 7. Repeat, while k < len
      while (k < len) {
        // a. Let elementK be the result of ? Get(O, ! ToString(k)).
        // b. If SameValueZero(valueToFind, elementK) is true, return true.
        if (sameValueZero(o[k], valueToFind)) {
          return true;
        }
        // c. Increase k by 1. 
        k++;
      }

      // 8. Return false
      return false;
    }
  });
}
```