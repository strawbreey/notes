---
title: "Angular Schematic"
date: 2020-09-18T19:40:57+08:00
draft: false
---

### 原理图 (Schematic)

原理图是一个基于模板的支持复杂逻辑的代码生成器。它是一组通过生成代码或修改代码来转换软件项目的指令。原理图会打包成集合（collection）并用 npm 安装。

原理图的集合可以作为一个强大的工具，以创建、修改和维护任何软件项目，特别是当要自定义 Angular 项目以满足你自己组织的特定需求时。例如，你可以借助原理图来用预定义的模板或布局生成常用的 UI 模式或特定的组件。你也可以使用原理图来强制执行架构规则和约定，让你的项目保持一致性和互操作性。


#### Angular CLI 中的原理图

原理图是 Angular 生态系统的一部分。Angular CLI 使用原理图对 Web 应用项目进行转换。 你可以修改这些原理图，并定义新的原理图，比如更新代码以修复依赖中的重大变更，或者把新的配置项或框架添加到现有的项目中。

```shell
# 生成原理图
ng generate my-schematic-collection:my-schematic-name

# or

ng generate my-schematic-name --collection collection-name
```

#### 配置 CLI 的原理图

编写库的原理图  (Developing schematics for libraries)

为一名库开发人员，你可以创建自己的自定义原理图集合，以便把你的库与 Angular CLI 集成在一起。

添加（Add）原理图允许开发人员使用 ng add 在 Angular 工作空间中安装你的库。

生成（Generation）原理图可以告诉 ng generate 子命令如何修改项目、添加配置和脚本，以及为库中定义的工件提供脚手架。

更新（Update）原理图可以告诉 ng update 命令，如何更新库的依赖，并在发布新版本时调整其中的重大变更。

#### 生成（Generation）原理图


生成器原理图是 ng generate 的操作指令。 那些已经有文档的子命令会使用默认的 Angular 生成器原理图，但你可以在子命令中指定另一个原理图来生成你的库中定义的那些工件。

```
ng generate @angular/material:table 
```

#### 更新原理图

```
ng update
```

如果你给这个命令指定一组要更新的库（或 --all 标志），它就会更新这些库、这些库的对等依赖，以及对等依赖的对等依赖。

> 如果存在不一致（例如，如果在某个简单的 semver 范围内无法匹配对等依赖），那么该命令会生成一个错误，并且不会更改工作空间中的任何内容。

我们建议你不要强制更新所有的依赖项，而应该首先尝试更新特定的依赖项。

如果你创建的新版本的库引入了潜在的重大更改，你可以提供一个更新原理图，让 ng update 命令能够自动解决所更新项目中的任何重大修改。

```shell
ng update @angular/material
```

### 创作原理图

#### 概念
原理图的公共 API 定义了表达其基本概念的类。

- 虚拟文件系统用 Tree（树）表示。Tree 数据结构包含一个基础状态 base（一组已经存在的文件）和一个 暂存区 staging（需要应用到 base 的更改列表）。在进行修改的过程中，你并没有真正改变它的 base，而是把那些修改添加到了暂存区。

- Rule（规则）对象定义了一个函数，它接受 Tree，进行转换，并返回一个新的 Tree。原理图的主文件 index.ts 定义了一组实现原理图逻辑的规则。

- 转换由 Action（动作）表示。有四种动作类型：Create、Rename、Overwrite 和 Delete。

- 每个原理图都在一个上下文中运行，上下文由一个 SchematicContext 对象表示。

#### 定义规则和动作

你使用 Schematics CLI 创建一个新的空白原理图时，它所生成的入口函数就是一个规则工厂。RuleFactory 对象定义了一个用于创建 Rule 的高阶函数。

```ts
import { Rule, SchematicContext, Tree } from '@angular-devkit/schematics';

// You don't have to export the function as default.
// You can also have more than one rule factory per file.
export function helloWorld(_options: any): Rule {
 return (tree: Tree, _context: SchematicContext) => {
   return tree;
 };
}
```
你的这些规则可以通过调用外部工具和实现逻辑来修改你的项目。比如，你需要一个规则来定义如何将原理图中的模板合并到宿主项目中。

```ts
import {
  JsonAstObject,
  JsonObject,
  JsonValue,
  Path,
  normalize,
  parseJsonAst,
  strings,
} from '@angular-devkit/core';
```

