---
title: "Php Iterable"
date: 2021-07-23T13:15:50+08:00
draft: false
---

Iterable是 PHP 7.1 中引入的一个`伪类型`。它接受任何 array 或实现了 Traversable 接口的对象。这些类型都能用 foreach 迭代， 也可以和 生成器 里的 yield from 一起使用。

可迭代对象可以用作参数类型，表示函数需要一组值， 但是不会关心值集的形式，因为它将与 foreach 一起使用。如果一个值不是数组或 Traversable 的实例，则会抛出一个 TypeError。

```php
function foo(iterable $iterable) {
    foreach ($iterable as $value) {
        // ...
    } 
}
```