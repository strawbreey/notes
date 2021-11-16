---
title: "Angular Providers"
date: 2021-10-29T11:58:17+08:00
draft: false
---

在Angular中有很多方式可以将服务类注册到注入器中:

1. @Injectable 元数据中的providedIn属性
2. @NgModule 元数据中的 providers属性
3. @Component 元数据中的 providers属性


创建一个文件名叫名 `hero.service.ts` 叫 hero 的服务

```js
import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class HeroService {

    constructor() { }

    // 新增加setName方法
    setName(name:string):string{
        return `姓名：${name}`;
    }
}
```

@Injectable 元数据中的providedIn属性

providedIn: 'root' 告诉 Angular在根注入器中注册这个服务,这也是使用CLI生成服务时默认的方式.

这种方式注册, 不需要再@NgModule装饰器中写providers, 而且在代码编译打包时, 可以执行摇树优化, 会移除所有没在应用中使用过的服务。


在 `heroes.component.ts` 中

```js
import { Component, OnInit } from '@angular/core';
import { HeroService } from '../hero.service'

@Component({
    selector: 'app-heroes',
    templateUrl: './heroes.component.html',
    styleUrls: ['./heroes.component.css']
})
export class HeroesComponent implements OnInit {

constructor(private heroService:HeroService) { }

    ngOnInit() {
        this.heroService.setName('张三');
    }
}
```

2. @NgModule 元数据中的 providers属性

改写 hero.service.ts里面的 `@Injectable`，如下

```js
import { Injectable } from '@angular/core';

@Injectable() // 删掉了 {providedIn: 'root'}
export class HeroService {...}
```

 在module.ts 中

```js

@NgModule({
    providers: [
        HeroService,
    ],
})

```
然后就可以在使用拉，使用方法，同1 heroes.component.ts文件

3. @Component 元数据中的 providers属性

`hero.service.ts` 里面的@Injectable，删掉 `{providedIn: 'root'}`, 同2 `hero.service.ts` 文件

改写 `heroes.component.ts`

```js
import { Component, OnInit } from '@angular/core';
import { HeroService } from '../hero.service'

@Component({
    selector: 'app-heroes',
    templateUrl: './heroes.component.html',
    styleUrls: ['./heroes.component.css'],
    providers: [HeroService] // 新增 providers: [HeroService]
})
export class HeroesComponent implements OnInit {

    constructor(
        private heroService:HeroService
    ) { }

    ngOnInit() {
        this.heroService.setName('张三');
    }
}
```
