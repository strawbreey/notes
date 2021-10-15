---
title: "NgTemplateOutlet"
date: 2021-10-13T19:23:00+08:00
draft: false
---

该指令用于基于已有的 `TemplateRef` 对象，插入对应的内嵌视图。在应用 `NgTemplateOutlet` 指令时，我们可以通过 `[ngTemplateOutletContext]` 属性来设置 `EmbeddedViewRef` 的上下文对象。绑定的上下文应该是一个对象，此外可通过 let 语法来声明绑定上下文对象属性名。

> 友情提示：若 let 语法未绑定任何属性名，则上下文对象中 $implicit 属性，对应的值将作为默认值。

NgTemplateOutlet 指令语法

```html
<ng-container *ngTemplateOutlet="templateRefExp; context: contextExp"></ng-container>
```

NgTemplateOutlet 使用示例

```js

@Component({
  selector: 'ng-template-outlet-example',
  template: `
    <ng-container *ngTemplateOutlet="greet"></ng-container>
    <hr>
    <ng-container *ngTemplateOutlet="eng; context: myContext"></ng-container>
    <hr>
    <ng-container *ngTemplateOutlet="svk; context: myContext"></ng-container>
    <hr>
    <ng-template #greet><span>Hello</span></ng-template>
    <ng-template #eng let-name><span>Hello {{name}}!</span></ng-template>
    <ng-template #svk let-person="localSk"><span>Ahoj {{person}}!</span></ng-template>
`
})
class NgTemplateOutletExample {
  myContext = {$implicit: 'World', localSk: 'Svet'};
}
```
### 基础知识

TemplateRef

TemplateRef 实例用于表示模板对象。

ViewContainerRef

ViewContainerRef 实例提供了 createEmbeddedView() 方法，该方法接收 TemplateRef 对象作为参数，并将模板中的内容作为容器 (comment 元素) 的兄弟元素，插入到页面中。

`<ng-template>`

`<ng-template>` 用于定义模板，使用 * 语法糖的结构指令，最终都会转换为 `<ng-template>` 模板指令，模板内的内容如果不进行处理，是不会在页面中显示。

`<ng-container>`

`<ng-container>` 是一个逻辑容器，可用于对节点进行分组，但不作为 DOM 树中的节点，它将被渲染为 HTML中的 comment 元素，它可用于避免添加额外的元素来使用结构指令。


>注意: 使用 let 语法创建模板局部变量，若未设置绑定的值，则默认是上下文对象中 $implicit 属性对应的值。为什么属性名是 $implicit 呢？因为 Angular 不知道用户会如何命名，所以定义了一个默认的属性名。 即 let-name="$implicit" 与 let-name 是等价的。

