---
title: "Js Chaining Methods"
date: 2020-10-26T17:55:46+08:00
draft: false
---

Chaining Methods, also known as Cascading, means repeatedly calling one method after another on an object, in one continuous line of code. Let us look at an example where method chaining can help us avoid repetition.

```js
class Car {
   constructor() {
      this.wheels = 4
      this.doors = 4
      this.topSpeed = 100
      this.feulCapacity = "400 Litres"
   }
   setWheels(w) {
      this.wheels = w;
      return this;
   }
   setDoors(d) {
      this.doors = d;
      return this;
   }
   setTopSpeed(t) {
      this.topSpeed = t;
      return this;
   }
   setFeulCapacity(fc) {
      this.feulCapacity = fc;
      return this;
   }
   displayCarProps() {
      console.log(`Your car has ${this.wheels} wheels,\
      ${this.doors} doors with a top speed of ${this.topSpeed}\
      and feul capacity of ${this.feulCapacity}`)
   }
}

let sportsCar = new Car()
   .setDoors(2)
   .setTopSpeed(250)
   .setFeulCapacity("600 Litres")
   .displayCarProps()
```


### 参考

- [Method Chaining in JavaScript](https://www.tutorialspoint.com/method-chaining-in-javascript#:~:text=Chaining%20Methods%2C%20also%20known%20as,one%20continuous%20line%20of%20code.)
- [实现链式调用](https://github.com/lgwebdream/FE-Interview/issues/22)
- [Understanding Method Chaining In Javascript](https://medium.com/backticks-tildes/understanding-method-chaining-in-javascript-647a9004bd4f)

