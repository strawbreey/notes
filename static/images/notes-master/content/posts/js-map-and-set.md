---
title: "Js Map and Set"
date: 2020-10-23T14:41:10+08:00
draft: false
---


### Map

Map 是一个带键的数据项的集合，就像一个 Object 一样。 但是它们最大的差别是 Map 允许任何类型的键（key）

它的方法和属性如下：

- new Map() —— 创建 map。
- map.set(key, value) —— 根据键存储值。
- map.get(key) —— 根据键来返回值，如果 map 中不存在对应的 key，则返回 undefined。
- map.has(key) —— 如果 key 存在则返回 true，否则返回 false。
- map.delete(key) —— 删除指定键的值。
- map.clear() —— 清空 map。
- map.size —— 返回当前元素个数。

```js

let map = new Map();

map.set('1', 'str1');   // 字符串键
map.set(1, 'num1');     // 数字键
map.set(true, 'bool1'); // 布尔值键

// 还记得普通的 Object 吗? 它会将键转化为字符串
// Map 则会保留键的类型，所以下面这两个结果不同：
alert( map.get(1)   ); // 'num1'
alert( map.get('1') ); // 'str1'

alert( map.size ); // 3

```

  map[key] 不是使用 Map 的正确方式
  虽然 map[key] 也有效，例如我们可以设置 map[key] = 2，这样会将 map 视为 JavaScript 的 plain object，因此它暗含了所有相应的限制（没有对象键等）。

  所以我们应该使用 map 方法：set 和 get 等。

#### Map 还可以使用对象作为键

```js
let john = { name: "John" };

// 存储每个用户的来访次数
let visitsCountMap = new Map();

// john 是 Map 中的键
visitsCountMap.set(john, 123);

alert( visitsCountMap.get(john) ); // 123
```

使用对象作为键是 Map 最值得注意和重要的功能之一。对于字符串键，Object（普通对象）也能正常使用，但对于对象键则不行。

```js
let john = { name: "John" };

let visitsCountObj = {}; // 尝试使用对象

visitsCountObj[john] = 123; // 尝试将 john 对象作为键

// 是写成了这样!
alert( visitsCountObj["[object Object]"] ); // 123
```

链式调用

```js
map.set('1', 'str1')
  .set(1, 'num1')
  .set(true, 'bool1');

```

### Map 迭代

如果要在 map 里使用循环，可以使用以下三个方法：

map.keys() —— 遍历并返回所有的键（returns an iterable for keys），
map.values() —— 遍历并返回所有的值（returns an iterable for values），
map.entries() —— 遍历并返回所有的实体（returns an iterable for entries）[key, value]，for..of 在默认情况下使用的就是这个。

```js
let recipeMap = new Map([
  ['cucumber', 500],
  ['tomatoes', 350],
  ['onion',    50]
]);

// 遍历所有的键（vegetables）
for (let vegetable of recipeMap.keys()) {
  alert(vegetable); // cucumber, tomatoes, onion
}

// 遍历所有的值（amounts）
for (let amount of recipeMap.values()) {
  alert(amount); // 500, 350, 50
}

// 遍历所有的实体 [key, value]
for (let entry of recipeMap) { // 与 recipeMap.entries() 相同
  alert(entry); // cucumber,500 (and so on)
}
```

迭代的顺序与插入值的顺序相同。与普通的 Object 不同，Map 保留了此顺序。

### Object.entries：从对象创建 Map

当创建一个 Map 后，我们可以传入一个带有键值对的数组（或其它可迭代对象）来进行初始化，如下所示

```js
// 键值对 [key, value] 数组
let map = new Map([
  ['1',  'str1'],
  [1,    'num1'],
  [true, 'bool1']
]);

alert( map.get('1') ); // str1
```

如果我们想从一个已有的普通对象（plain object）来创建一个 Map，那么我们可以使用内建方法 Object.entries(obj)，该返回对象的键/值对数组，该数组格式完全按照 Map 所需的格式。

所以可以像下面这样从一个对象创建一个 Map：

```js
let obj = {
  name: "John",
  age: 30
};

let map = new Map(Object.entries(obj));

alert( map.get('name') ); // John


```

这里，Object.entries 返回键/值对数组：[ ["name","John"], ["age", 30] ]。这就是 Map 所需要的格式。

### Object.fromEntries：从 Map 创建对象

```js
let prices = Object.fromEntries([
  ['banana', 1],
  ['orange', 2],
  ['meat', 4]
]);

// 现在 prices = { banana: 1, orange: 2, meat: 4 }

alert(prices.orange); // 2
```

## Set

Set 是一个特殊的类型集合 —— “值的集合”（没有键），它的每一个值只能出现一次。

它的主要方法如下：

- new Set(iterable) —— 创建一个 set，如果提供了一个 iterable 对象（通常是数组），将会从数组里面复制值到 set 中。
- set.add(value) —— 添加一个值，返回 set 本身
- set.delete(value) —— 删除值，如果 value 在这个方法调用的时候存在则返回 true ，否则返回 false。
- set.has(value) —— 如果 value 在 set 中，返回 true，否则返回 false。
- set.clear() —— 清空 set。
- set.size —— 返回元素个数。

```js
let set = new Set();

let john = { name: "John" };
let pete = { name: "Pete" };
let mary = { name: "Mary" };

// visits，一些访客来访好几次
set.add(john);
set.add(pete);
set.add(mary);
set.add(john);
set.add(mary);

// set 只保留不重复的值
alert( set.size ); // 3

for (let user of set) {
  alert(user.name); // John（然后 Pete 和 Mary）
}
```

Set 的替代方法可以是一个用户数组，用 arr.find 在每次插入值时检查是否重复。但是这样性能会很差，因为这个方法会遍历整个数组来检查每个元素。Set 内部对唯一性检查进行了更好的优化。

### Set 迭代（iteration）

```js
let set = new Set(["oranges", "apples", "bananas"]);

for (let value of set) {
  alert(value);
}

// 与 forEach 相同：
set.forEach((value, valueAgain, set) => {
  alert(value);
});
```

Map 中用于迭代的方法在 Set 中也同样支持：

- set.keys() —— 遍历并返回所有的值（returns an iterable object for values），
- set.values() —— 与 set.keys() 作用相同，这是为了兼容 Map，
- set.entries() —— 遍历并返回所有的实体（returns an iterable object for entries）[value, value]，它的存在也是为了兼容 Map。

### 测试题

1. 数组去重

```js
function unique(arr) {
  return Array.from(new Set(arr));
}
```


### 总结

Map —— 是一个带键的数据项的集合。

方法和属性如下：

  - new Map([iterable]) —— 创建 map，可选择带有 [key,value] 对的 iterable（例如数组）来进行初始化。
  - map.set(key, value) —— 根据键存储值。
  - map.get(key) —— 根据键来返回值，如果 map 中不存在对应的 key，则返回 undefined。
  - map.has(key) —— 如果 key 存在则返回 true，否则返回 false。
  - map.delete(key) —— 删除指定键的值。
  - map.clear() —— 清空 map 。
  - map.size —— 返回当前元素个数。

与普通对象 Object 的不同点：

  - 任何键、对象都可以作为键。
  - 有其他的便捷方法，如 size 属性。
  - Set —— 是一组唯一值的集合。

方法和属性：

  - new Set([iterable]) —— 创建 set，可选择带有 iterable（例如数组）来进行初始化。
  - set.add(value) —— 添加一个值（如果 value 存在则不做任何修改），返回 set 本身。
  - set.delete(value) —— 删除值，如果 value 在这个方法调用的时候存在则返回 true ，否则返回 false。
  - set.has(value) —— 如果 value 在 set 中，返回 true，否则返回 false。
  - set.clear() —— 清空 set。
  - set.size —— 元素的个数。
  
  在 Map 和 Set 中迭代总是按照值插入的顺序进行的，所以我们不能说这些集合是无序的，但是我们不能对元素进行重新排序，也不能直接按其编号来获取元素。


### 参考资料

- [Map and Set（映射和集合）](https://zh.javascript.info/map-set)
- [MDN Map](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map)