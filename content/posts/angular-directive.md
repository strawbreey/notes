---
title: "Angular Directive"
date: 2020-09-14T16:31:10+08:00
description: "Angular Directive"
draft: false
---

### 概述
angular 中有三种指令类型
- 组件 拥有模板的指令
- 结构型指令 通过添加和删除DOM元素改变DOM布局的指令 如: `ngFor`, `Ngif`  
- 属性型指令 改变元素,组件外观或者行为的指令 ,如果内置的NgStyle


#### 创建的一个属性型指令
```shell
ng generate directive heightlight
```

```ts
// src/app/highlight.directive.ts
import { Directive, ElementRef } from '@angular/core';

// 导入的 Directive 符号提供了 Angular 的 @Directive 装饰器。

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {
    constructor(el: ElementRef) {
       el.nativeElement.style.backgroundColor = 'yellow';
    }
}

```

```html
<p appHighlight>Highlight me!</p>
```

当然，你可以通过标准的 JavaScript 方式手动给宿主 DOM 元素附加一个事件监听器。 但这种方法至少有三个问题：
```
必须正确的书写事件监听器。

当指令被销毁的时候，必须拆卸事件监听器，否则会导致内存泄露。

必须直接和 DOM API 打交道，应该避免这样做。
```


#### 使用 @Input 数据绑定向指令传递值

先从 @angular/core 中导入 Input。

```ts
import { Directive, ElementRef, HostListener, Input } from '@angular/core';

@Input() highlightColor: string;
```

```html
<p appHighlight highlightColor="yellow">Highlighted in yellow</p>
<p appHighlight [highlightColor]="'orange'">Highlighted in orange</p>
```

最终的指令类
```js
import { Directive, ElementRef, HostListener, Input } from '@angular/core';
// 装饰器
@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {

  constructor(private el: ElementRef) { }

  @Input() defaultColor: string;
  
  // 绑定@Input属性 
  @Input('appHighlight') highlightColor: string; // 绑定到 @Input 别名 highlightColor

  @HostListener('mouseenter') onMouseEnter() {   // 监听鼠标
     this.highlight(this.highlightColor || this.defaultColor || 'red');
  }

  @HostListener('mouseleave') onMouseLeave() {
    this.highlight(null);
  }

  // 设置高亮
  private highlight(color: string) {    
    this.el.nativeElement.style.backgroundColor = color;
  }
}
```

#### 附录：为什么要加@Input？
@Input 装饰器都告诉 Angular，该属性是公共的，并且能被父组件绑定。 如果没有 @Input，Angular 就会拒绝绑定到该属性。

模板 HTML 绑定到组件的属性，而且从来没有用过 @Input。 差异何在？

差异在于信任度不同。 Angular 把组件的模板看做从属于该组件的。 组件和它的模板默认会相互信任。 这也就是意味着，组件自己的模板可以绑定到组件的任意属性，无论是否使用了 @Input 装饰器。

但组件或指令不应该盲目的信任其它组件或指令。 因此组件或指令的属性默认是不能被绑定的。 从 Angular 绑定机制的角度来看，它们是私有的，而当添加了 @Input 时，Angular 绑定机制才会把它们当成公共的。 只有这样，它们才能被其它组件或属性绑定。


你可以根据属性名在绑定中出现的位置来判定是否要加 @Input。

- 当它出现在等号右侧的模板表达式中时，它属于模板所在的组件，不需要 @Input 装饰器。

- 当它出现在等号左边的方括号（[ ]）中时，该属性属于其它组件或指令，它必须带有 @Input 装饰器。


### 结构型指令

结构型指令的职责是 HTML 布局。 它们塑造或重塑 DOM 的结构，比如添加、移除或维护这些元素。

三个常用的内置结构型指令 —— NgIf、NgFor和NgSwitch...

你将看到指令同时具有两种拼写形式大驼峰 UpperCamelCase 和小驼峰 lowerCamelCase，比如你已经看过的 NgIf 和 ngIf。 这里的原因在于，NgIf 引用的是指令的类名，而 ngIf 引用的是指令的属性名*

指令的类名拼写成大驼峰形式（NgIf），而它的属性名则拼写成小驼峰形式（ngIf）。 本章会在谈论指令的属性和工作原理时引用指令的类名，在描述如何在 HTML 模板中把该指令应用到元素时，引用指令的属性名。


#### NgIf 分析
NgIf 是一个很好的结构型指令案例：它接受一个布尔值，并据此让一整块 DOM 树出现或消失。ngIf 指令并不是使用 CSS 来隐藏元素的。它会把这些元素从 DOM 中物理删除。

指令也可以通过把它的 display 风格设置为 none 而隐藏不需要的段落。


#### *ngFor 内幕
Angular 会把 `*ngFor` 用同样的方式把星号（*）语法的 template属性转换成 <ng-template>元素。
```html
<div *ngFor="let hero of heroes; let i=index; let odd=odd; trackBy: trackById" [class.odd]="odd">
  ({{i}}) {{hero.name}}
</div>

<ng-template ngFor let-hero [ngForOf]="heroes" let-i="index" let-odd="odd" [ngForTrackBy]="trackById">
  <div [class.odd]="odd">({{i}}) {{hero.name}}</div>
</ng-template>
```

### 微语法

#### 每个宿主元素上只能有一个结构型指

有时你会希望只有当特定的条件为真时才重复渲染一个 HTML 块。 你可能试过把 *ngFor 和 *ngIf 放在同一个宿主元素上，但 Angular 不允许。这是因为你在一个元素上只能放一个结构型指令。

原因很简单。结构型指令可能会对宿主元素及其子元素做很复杂的事。当两个指令放在同一个元素上时，谁先谁后？NgIf 优先还是 NgFor 优先？NgIf 可以取消 NgFor 的效果吗？ 如果要这样做，Angular 应该如何把这种能力泛化，以取消其它结构型指令的效果呢？

对这些问题，没有办法简单回答。而禁止多个结构型指令则可以简单地解决这个问题。 这种情况下有一个简单的解决方案：把 *ngIf 放在一个"容器"元素上，再包装进 *ngFor 元素。 这个元素可以使用ng-container，以免引入一个新的 HTML 层级。


#### NgSwitch 内幕

Angular 的 NgSwitch 实际上是一组相互合作的指令：NgSwitch、NgSwitchCase 和 NgSwitchDefault。

```html
<div [ngSwitch]="hero?.emotion">
  <app-happy-hero    *ngSwitchCase="'happy'"    [hero]="hero"></app-happy-hero>
  <app-sad-hero      *ngSwitchCase="'sad'"      [hero]="hero"></app-sad-hero>
  <app-confused-hero *ngSwitchCase="'confused'" [hero]="hero"></app-confused-hero>
  <app-unknown-hero  *ngSwitchDefault           [hero]="hero"></app-unknown-hero>
</div>
```

NgSwitch 本身不是结构型指令，而是一个属性型指令，它控制其它两个 switch 指令的行为。 这也就是为什么你要写成 [ngSwitch] 而不是 *ngSwitch 的原因。

NgSwitchCase 和 NgSwitchDefault 都是结构型指令。 因此你要使用星号（*）前缀来把它们附着到元素上。 NgSwitchCase 会在它的值匹配上选项值的时候显示它的宿主元素。 NgSwitchDefault 则会当没有兄弟 NgSwitchCase 匹配上时显示它的宿主元素。

像其它的结构型指令一样，NgSwitchCase 和 NgSwitchDefault 也可以解开语法糖，变成 <ng-template> 的形式。

```html
<div [ngSwitch]="hero?.emotion">
  <ng-template [ngSwitchCase]="'happy'">
    <app-happy-hero [hero]="hero"></app-happy-hero>
  </ng-template>
  <ng-template [ngSwitchCase]="'sad'">
    <app-sad-hero [hero]="hero"></app-sad-hero>
  </ng-template>
  <ng-template [ngSwitchCase]="'confused'">
    <app-confused-hero [hero]="hero"></app-confused-hero>
  </ng-template >
  <ng-template ngSwitchDefault>
    <app-unknown-hero [hero]="hero"></app-unknown-hero>
  </ng-template>
</div>
```

#### 优先使用星号（*）语法
星号（*）语法比不带语法糖的形式更加清晰。 如果找不到单一的元素来应用该指令，可以使用`<ng-container>`作为该指令的容器。


### <ng-template>元素

ng-template>是一个 Angular 元素，用来渲染 HTML。 它永远不会直接显示出来。 事实上，在渲染视图之前，Angular 会把 <ng-template> 及其内容替换为一个注释。


### 写一个结构型指令

写一个名叫 UnlessDirective 的结构型指令，它是 NgIf 的反义词。 NgIf 在条件为 true 的时候显示模板内容，而 UnlessDirective 则会在条件为 false 时显示模板内容。

```html
<p *appUnless="condition">Show this sentence unless the condition is true.</p>
```

创建指令很像创建组件。

- 导入 Directive 装饰器（而不再是 Component）。

- 导入符号 Input、TemplateRef 和 ViewContainerRef，你在任何结构型指令中都会需要它们。

- 给指令类添加装饰器。

- 设置 CSS 属性选择器，以便在模板中标识出这个指令该应用于哪个元素。

```js
import { Directive, Input, TemplateRef, ViewContainerRef } from '@angular/core';

/**
 * Add the template content to the DOM unless the condition is true.
 */
@Directive({ selector: '[appUnless]'})
export class UnlessDirective {
  private hasView = false;

  constructor(
    private templateRef: TemplateRef<any>,
    private viewContainer: ViewContainerRef
  ) { }

  @Input() set appUnless(condition: boolean) {
    if (!condition && !this.hasView) {
      this.viewContainer.createEmbeddedView(this.templateRef);
      this.hasView = true;
    } else if (condition && this.hasView) {
      this.viewContainer.clear();
      this.hasView = false;
    }
  }
}
```

```html
<p *appUnless="condition" class="unless a">
  (A) This paragraph is displayed because the condition is false.
</p>

<p *appUnless="!condition" class="unless b">
  (B) Although the condition is true,
  this paragraph is displayed because appUnless is set to false.
</p>

```