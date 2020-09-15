---
title: "Reactive Forms"
date: 2020-09-14T15:33:34+08:00
draft: false
---

响应式表单提供了一种模型驱动的方式来处理表单输入，其中的值会随时间而变化。

### 响应式表单和模板驱动表单的区别

响应式表单和模板驱动表单以不同的方式处理和管理表单数据。每种方法都有各自的优点。

响应式表单提供对底层表单对象模型直接、显式的访问。它们与模板驱动表单相比，更加健壮：它们的可扩展性、可复用性和可测试性都更高。如果表单是你的应用程序的关键部分，或者你已经在使用响应式表单来构建应用，那就使用响应式表单。

模板驱动表单依赖模板中的指令来创建和操作底层的对象模型。它们对于向应用添加一个简单的表单非常有用，比如电子邮件列表注册表单。它们很容易添加到应用中，但在扩展性方面不如响应式表单。如果你有可以只在模板中管理的非常基本的表单需求和逻辑，那么模板驱动表单就很合适。


| Tables           | 响应式                | 模板驱动         |
|       ---        |          ---         |         ---      |
| 建立表单模型      | 显式的，在组件类中创建 | 隐式的，由指令创建 |
| 数据模型          | 结构化和不可变的      | 非结构化和可变的   |
| 可预测性          | 同步                 | 异步              |
| 表单验证          | 函数                 | 指令              |

响应式表单比模板驱动表单更有可伸缩性。它们提供对底层表单 API 的直接访问，以及对表单数据模型的同步访问，从而可以更轻松地创建大型表单。响应式表单需要较少的测试设置，测试时不需要深入理解变更检测，就能正确测试表单更新和验证。

模板驱动表单专注于简单的场景，可复用性没那么高。它们抽象出了底层表单 API，并且只提供对表单数据模型的异步访问。对模板驱动表单的这种抽象也会影响测试。测试程序非常依赖于手动触发变更检测才能正常运行，并且需要进行更多设置工作。


### 常用表单基础类

响应式表单和模板驱动表单都建立在下列基础类之上。

- FormControl 实例用于追踪单个表单控件的值和验证状态
- FormGroup 用于追踪一个表单控件组的值和状态
- FormArray 用于追踪表单控件数组的值和状态
- ControlValueAccessor 用于在 Angular 的 FormControl 实例和原生 DOM 元素之间创建一个桥梁。


```ts
interface ControlValueAccessor {
  writeValue(obj: any): void // Writes a new value to the element.
  registerOnChange(fn: any): void // Registers a callback function that is called when the control's value changes in the UI.
  registerOnTouched(fn: any): void // Registers a callback function that is called by the forms API on initialization to update the form model on blur.
  setDisabledState(isDisabled: boolean)?: void // Function that is called by the forms API when the control status changes to or from 'DISABLED'. Depending on the status, it enables or disables the appropriate DOM element.
}

// writeValue
writeValue(value: any): void {
  this._renderer.setProperty(this._elementRef.nativeElement, 'value', value);
}

// registerOnChange
registerOnChange(fn: (_: any) => void): void {
  this._onChange = fn;
}

//  registerOnTouched
registerOnTouched(fn: any): void {
  this._onTouched = fn;
}

host: {
   '(change)': '_onChange($event.target.value)'
}

// setDisabledState
setDisabledState(isDisabled: boolean): void {
  this._renderer.setProperty(this._elementRef.nativeElement, 'disabled', isDisabled);
}



```

// 实现类
- CheckboxControlValueAccessor
- DefaultValueAccessor
- NumberValueAccessor
- RadioControlValueAccessor
- RangeValueAccessor
- SelectControlValueAccessor
- SelectMultipleControlValueAccessor

DefaultValueAccessor


### Setup in reactive forms

对于响应式表单，你可以直接在组件类中定义表单模型。`[formControl]` 指令会通过内部值访问器来把显式创建的 FormControl 实例与视图中的特定表单元素联系起来。

```ts
import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-reactive-favorite-color',
  template: `
    Favorite Color: <input type="text" [formControl]="favoriteColorControl">
  `
})
export class FavoriteColorComponent {
  favoriteColorControl = new FormControl('');
}
```

 展示了在响应式表单中，表单模型是如何成为事实之源（source of truth）的。它通过输入元素上的 [formControl] 指令，在任何给定的时间点提供表单元素的值和状态。
在响应式表单中直接访问表单模型
![key-diff-reactive-forms.png](/images/key-diff-reactive-forms.png)


### 表单中的数据流

当应用包含一个表单时，Angular 必须让该视图与组件模型保持同步，并让组件模型与视图保持同步。当用户通过视图更改值并进行选择时，新值必须反映在数据模型中。同样，当程序逻辑改变数据模型中的值时，这些值也必须反映到视图中。

在响应式表单中，视图中的每个表单元素都直接链接到一个表单模型（FormControl 实例）。 从视图到模型的修改以及从模型到视图的修改都是同步的，而且不依赖于 UI 的渲染方式。

这个视图到模型的示意图展示了当输入字段的值发生变化时数据是如何从视图开始，经过下列步骤进行流动的。

- 用户在输入框元素中键入了一个值"Blue"。
- 这个输入框元素会发出一个带有最新值的 "input" 事件。
- 这个控件 ControlValueAccessor 会监听输入框元素，并立即把新值传给 FormControl 实例。
- FormControl 实例会通过 valueChanges 这个可观察对象发出这个新值。
- valueChanges 的任何一个订阅者都会收到这个新值。

![dataflow-reactive-forms-vtm](/images/dataflow-reactive-forms-vtm.png)


这个模型到视图的示意图体现了程序中对模型的修改是如何通过下列步骤传播到视图中的。

- favoriteColorControl.setValue() 方法被调用，它会更新这个 FormControl 的值。

- FormControl 实例会通过 valueChanges 这个可观察对象发出新值。

- valueChanges 的任何订阅者都会收到这个新值。

- 该表单输入框元素上的控件值访问器会把控件更新为这个新值

![dataflow-reactive-forms-mtv](dataflow-reactive-forms-mtv.png)


响应式表单通过以不可变的数据结构提供数据模型，来保持数据模型的纯粹性。每当在数据模型上触发更改时，FormControl 实例都会返回一个新的数据模型，而不会更新现有的数据模型。这使你能够通过该控件的可观察对象跟踪对数据模型的唯一更改。这让变更检测更有效率，因为它只需在唯一性更改（译注：也就是对象引用发生变化）时进行更新。由于数据更新遵循响应式模式，因此你可以把它和可观察对象的各种运算符集成起来以转换数据。


### 响应式表单
响应式表单使用显式的、不可变的方式，管理表单在特定的时间点上的状态。对表单状态的每一次变更都会返回一个新的状态，这样可以在变化时维护模型的整体性。响应式表单是围绕 Observable 流构建的，表单的输入和值都是通过这些输入值组成的流来提供的，它可以同步访问。

使用表单

使用表单控件有三个步骤。

- 在你的应用中注册响应式表单模块。该模块声明了一些你要用在响应式表单中的指令。

- 生成一个新的 FormControl 实例，并把它保存在组件中。

- 在模板中注册这个 FormControl。

```ts
// src/app/app.module.ts (excerpt)
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  imports: [
    // other imports ...
    ReactiveFormsModule
  ],
})
export class AppModule { }


// src/app/name-editor/name-editor.component.ts
import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-name-editor',
  templateUrl: './name-editor.component.html',
  styleUrls: ['./name-editor.component.css']
})
export class NameEditorComponent {
  name = new FormControl('');
}

```
```html
<!-- src/app/name-editor/name-editor.component.html -->
<label>
  Name:
  <input type="text" [formControl]="name">
</label>
```

### 更新部分数据模型 (patchValue)

当修改包含多个 FormGroup 实例的值时，你可能只希望更新模型中的一部分，而不是完全替换掉时

有两种更新模型值的方式：

- 使用 setValue() 方法来为单个控件设置新值。 setValue() 方法会严格遵循表单组的结构，并整体性替换控件的值。

- 使用 patchValue() 方法可以用对象中所定义的任何属性为表单模型进行替换。

注: setValue() 方法的严格检查可以帮助你捕获复杂表单嵌套中的错误，而 patchValue() 在遇到那些错误时可能会默默的失败。

patchValue() 方法要针对模型的结构进行更新。patchValue() 只会更新表单模型中所定义的那些属性。


### 响应式表单 API 汇总

| 类           | 说明                | 
|       ---        |          ---         |   
| AbstractControl      |所有三种表单控件类（FormControl、FormGroup 和 FormArray）的抽象基类。它提供了一些公共的行为和属性 |
| FormControl      |管理单体表单控件的值和有效性状态。它对应于 HTML 的表单控件，比如 `<input>` 或 `<select>`。|
| FormGroup      |管理一组 AbstractControl 实例的值和有效性状态。该组的属性中包括了它的子控件。组件中的顶层表单就是 FormGroup。 |
| FormArray      |管理一些 AbstractControl 实例数组的值和有效性状态。 |
| FormBuilder      |一个可注入的服务，提供一些用于提供创建控件实例的工厂方法。 |


| 指令 (Directive)          | 说明                | 
|       ---                 |          ---         |   
| FormControlDirective      | 把一个独立的 FormControl 实例绑定到表单控件元素。|
| FormControlName           | 把一个现有 FormGroup 中的 FormControl 实例根据名字绑定到表单控件元素。|
| FormGroupDirective        | 把一个现有的 FormGroup 实例绑定到 DOM 元素。|
| FormGroupName             | 把一个内嵌的 FormGroup 实例绑定到一个 DOM 元素。 |
| FormArrayName             | 把一个内嵌的 FormArray 实例绑定到一个 DOM 元素。 |
