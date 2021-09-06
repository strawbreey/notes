---
title: "Ts As"
date: 2021-08-30T17:36:54+08:00
draft: false
---

## 理解Typescript中的as、问号与叹号

### as关键字表示断言

在Typescript中，表示断言有两种方式。一种是尖扩号表示法:

```ts
let someValue: any = "this is a string";
let strLength: number = (<string>someValue).length;

```

另一种使用as关键字：

```ts
let someValue: any = "this is a string";
let strLength: number = (someValue as string).length;
```

唯一区别是，在JSX中，尖扩号与JSX语法冲突，只能使用as关键字。

### 问号(?)用于属性定义

准备工作：以下示例可以建立一个optional.ts的文件，然后用tsc进行编译。tsc可以用npm install -g typescript安装。

```ts
//运行以下命令可以在同目录下生成optional.js，有错误时会报错
tsc optional.ts
tsc --strictNullChecks optional.ts
```

问号表示可选的属性，一般用于属性定义，如，用于接口时：

```ts
interface SquareConfig {
    color?: string;
    width?: number;
}
function createSquare(config: SquareConfig) {
    if (config.clor) {
        console.log(config);
    }   
}
```

可选属性的含义是：使用这个属性时，要么这个属性名不存在，要么必须符合属性的类型定义（官方解释）。比如上述createSquare函数编译时会报error错误：

error TS2551: Property 'clor' does not exist on type 'SquareConfig'.

如果修改createSquare，将config.color的值改为undefined，会怎样？

```ts
interface SquareConfig {
    color?: string;
    width?: number;
}
function createSquare(config: SquareConfig) {
    config.color = undefined;
    if (config.color) {
        console.log(config);
    }   
}
```

这时并没有编译报错！明明config.color定义的是string类型呀？

即便是添加--strictNullChecks进行编译，也不会报错。可见，可选属性所定义的类型，并没有被typescript严格地对待，默认并不检查undefined。需要注意的是，将上述undefined改成null时，普通编译也不报错，--strictNullChecks编译会报如下错误：

error TS2322: Type 'null' is not assignable to type 'string | undefined'

从这句报错中，我们可以得出这样的结论：可选属性等价于一个union类型，union了undefined；不加--strictNullChecks编译时，null可以赋值给undfined类型。也就是说，SquareConfig的定义与下面的代码等价：

```ts
interface SquareConfig {
    color: string|undefined;
    width: number|undefined;
}
```

下面比较一下可选属性与正常属性。再次修改createSquare，将color属性修改为正常属性。

```ts
interface SquareConfig {
    color: string;
    width?: number;
}
function createSquare(config: SquareConfig) {
    config.color = undefined;
    if (config.color) {
        console.log(config);
    }   
}    
```


以--strictNullChecks编译，报错了:

error TS2322: Type 'undefined' is not assignable to type 'string'

这个比较也验证了上述的结论。

### 问号(?)用于属性读取

以下示例都使用--strictNullChecks编译

问号用于属性读取，主要有两个场景：一是读取数组元素（如下面的node[i]），二是读取不确定的类型如any，union，可选类型（如node[i].data）等。如下例，保存为index.ts：
```ts
interface VNodeData {
    class?: string;
}
interface VNode {
    sel?: string;
    data?: VNodeData;
}
function test(node: VNode[]) {
    let i = 0;
    var b = node[i].data.class;
    if(b !== undefined) {
        console.log(1);
    }   
}
```

用tsc --strictNullChecks index.ts，报错：

error TS2532: Object is possibly 'undefined'

下面将一一展示这一行代码var b = node[i].data.class;修改改后的效果。

1. 修改为var b = node[i]?.data.class;，然后编译。报错：

Object is possibly 'undefined'

2. 修改为var b = node[i]?.data?.class;，然后编译。编译通过，查看编译后的对应代码为：

```ts
function test(node) {
    var _a, _b; 
    var i = 0;
    var b = (_b = (_a = node[i]) === null || _a === void 0 ? void 0 : _a.data) === null || _b === void 0 ? void 0 : _b["class"];
    if (b !== undefined) {
        console.log(1);
    }
}

```
摘出来看，var b = node[i]?表示，如果node[i]的值为null或者undefined，则b等于undefined，否则b=node[i]。

3. 修改为var b = (node[i] as VNode).data?.class;，然后编译。编译通过，查看编译后的对应代码为：

```ts
function test(node) {
    var _a;
    var i = 0;
    var b = (_a = node[i].data) === null || _a === void 0 ? void 0 : _a["class"];
    if (b !== undefined) {
        console.log(1);
    }
}
```

此时，使用node[i]时，Typescript编译器将不再对其判断是否为null和undefined。即：var b = node[i] as VNode直接会被编成var b = node[i]。

4、修改为var b = node[i]!.data?.class，然后编译。编译通过，查看编译后的对应代码为：

```ts
function test(node) {
    var _a;
    var i = 0;
    var b = (_a = node[i].data) === null || _a === void 0 ? void 0 : _a["class"];
    if (b !== undefined) {
        console.log(1);
    }
}
```

可见，3和4的编译后代码完全一样，!的作用此时与as是等价的。然而，!只是用来判断null和undefined；as则可用于变更（缩小或者放大都可以）类型检测范围，仅当as后面跟的类型是一个非空类型时，两者才等价。如下例中，不能将as用法改为!。

```ts
interface Cat {
    action: string;
}
interface Dog {
    action: string;
}
type Animal = Cat | Dog;
let action:Animal = {} as Cat;
```
### 结论

1. as和!用于属性的读取，都可以缩小类型检查范围，都做判空用途时是等价的。只是!具体用于告知编译器此值不可能为空值（null和undefined），而as不限于此。

2. ?可用于属性的定义和读取，读取时告诉编译器此值可能为空值（null和undefined），需要做判断。