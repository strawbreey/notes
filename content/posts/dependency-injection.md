---
title: "Dependency Injection"
date: 2020-09-08T15:19:09+08:00
draft: false
---

依赖注入

In software engineering, dependency injection is a technique in which an object receives other objects that it depends on. These other objects are called dependencies. In the typical "using" relationship the receiving object is called a client and the passed (that is, "injected") object is called a service. The code that passes the service to the client can be many kinds of things and is called the injector. Instead of the client specifying which service it will use, the injector tells the client what service to use. The "injection" refers to the passing of a dependency (a service) into the object (a client) that would use it.

![Dependency Injection](/images/Dependency_Injection_Design_Pattern_UML.jpg)


#### JavaScript 中的依赖注入
```js
class CatController {
    private catService = new CatService();
    getCats() {
        return this.catService.get();
    }
}
const cc = new CatController();

//  => 改成依赖注入的方式
class CatController {
    private catService: ICatService;
    constructor(catService: ICatService){
        this.catService = catService;
    }
}

const cs = new CatService();
const cc = new CatController(cs);

```

#### Inversion of Control (控制反转)
IoC(控制反转: Inversion of Control)能做什么 
　　IoC 不是一种技术，只是一种思想，一个重要的面向对象编程的法则，它能指导我们如何设计出松耦合、更优良的程序。传统应用程序都是由我们在类内部主动创建依赖对象，从而导致类与类之间高耦合，难于测试；有了IoC容器后，把创建和查找依赖对象的控制权交给了容器，由容器进行注入组合对象，所以对象与对象之间是松散耦合，这样也方便测试，利于功能复用，更重要的是使得程序的整个体系结构变得非常灵活。

控制反转（Inversion of Control，英文缩写为IoC）是一个重要的面向对象编程的法则来削减计算机程序的耦合问题，也是轻量级的Spring框架的核心。 控制反转一般分为两种类型，依赖注入（Dependency Injection，简称DI）和依赖查找（Dependency Lookup）。依赖注入应用比较广泛。


依赖注入(DI)和控制反转(IoC)是具体的手段，是OOP理论中依赖倒置原则的体现形式，通过信息隐藏来降低对象之间的耦合，也就是说解决的是面向对象编程中对象耦合的问题。


#### angular 中的依赖注入
[!Dependency injection](https://angular.cn/guide/dependency-injection)





#### 相关书籍
[Dependency injection](https://en.wikipedia.org/wiki/Dependency_injection)