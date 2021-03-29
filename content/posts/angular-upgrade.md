---
title: "Angular Upgrade"
date: 2021-03-26T17:47:55+08:00
draft: false
---

### Angular 升级

1. 升级 angular-cli 到最新你版本
```bash
npm install -g @angular/cli@latest
````

2. 更新项目中angular相关的包。
```bash
ng update --all
```

### 更新前

- 无论在那里使用了 Renderer，现在都使用 Renderer2
- 如果使用了旧版的 HttpModule 和 Http 服务，请切换到 HttpClientModule 和 HttpClient 服务。HttpClient 简化了默认的 ergonomics(不再需要映射(map)到JSON)，并且现在支持类型化的返回值和拦截器。进一步了解 https://angular.io/guide/http
- 一旦所有依赖更新到 RxJS 6，请移除 rxjs-compat.
- 如果使用了 Angular Service worker，使用 files 替代 versionedFiles

### 更新期间
- 在终端中运行 ng update @angular/cli@8 @angular/core@8 更新 CLI 和核心框架到 v8。

- 使用 ::ng-deep 替代 /deep/, 了解有关 Angular 组件样式和 ::ng-deep 的更多信息。/deep/ 和 ::ng-deep 都不推荐使用，在从浏览器和工具(tools)中移除 shadow-piercing descendant 组合器前最好使用 ::ng-deep。

- Angular 现在使用 TypeScript 3.4，阅读更多关于改进类型检查可能产生的错误的信息
- 确保正在使用 Node 10 或更高版本
- CLI 的 build 命令现在可以自动创建一个具有最少 polyfill 的 ES6 构建(build)，和一个适用于较旧浏览器的兼容 ES5 构建，并根据浏览器加载适当的文件。您可以通过在 tsconfig.json 中将 target 设置为 es5 来取消此更改。了解详细
- 使用新版本的 CLI 时，将询问您是否要选择共享 CLI 使用数据。您也可以添加自己的 Google Analytics 帐户。这样，我们就可以优先考虑哪些 CLI 功能，从而做出更好的决策，并衡量改进的影响。了解详细

- 如果您使用 ViewChild 或 ContentChild，则我们将更新解决这些查询的方式，以使开发人员可以更好地控制。现在，您必须指定更改检测应在设置结果之前运行。例如：@ContentChild('foo', {static: false}) foo !: ElementRef;. ng update 会自动更新您的查询，但会因为使您的查询成为 static 而导致兼容性问题。了解详细

- 我们已经从本地 Sass 编译器切换到 JavaScript 编译器。要切换回 native 版本，请将其安装为devDependency：npm install node-sass --save-dev。

- 如果要构建自己的 Schematics，它们可能(potentially)是异步(asynchronous)的。从 8.0 开始，Schematics 都是异步的。

### 更新后

- 对于通过路由器延迟加载的模块，请确保您正在使用动态导入。在 v9 中删除了通过字符串导入。ng update 会自动处理。

- 我们不建议使用 @angular/platform-webworker，因为它与 CLI 不兼容。在 Web Worker 中运行 Angular 的渲染架构无法满足开发人员的需求。您仍然可以将Web worker与Angular一起使用。通过 Web worker 指南了解更多。如果您有需要的用例，请通过 devrel@angular.io 与我们联系！

- 在版本 8 中已弃用了对 Angular 中的 Web 跟踪框架的支持。您应该停止使用所有 wtf*API。要进行性能跟踪，我们建议使用浏览器性能工具。

- 删除 angular.json 中的 es5BrowserSupport，并在 tsconfig.json 中将 target 设置为 es2015。Angular 现在使用您的浏览器列表来确定是否需要 ES5 构建。ng update 将自动迁移此项。



### 升级过程中遇到的警告

```bash
# 升级angular 到 8
$ ng update @angular/cli@8 @angular/core@8
Using package manager: 'yarn'
Collecting installed dependencies...
Found 50 dependencies.
Fetching dependency metadata from registry...
                  Package "@ant-design/icons-angular" has an incompatible peer dependency to "typescript" (requires "~3.1.3", would install "3.5.3").
× Migration failed: Incompatible peer dependencies found.
Peer dependency warnings when installing dependencies means that those dependencies might not work correctly together.
You can use the '--force' option to ignore incompatible peer dependencies and instead address these warnings later.
  See "C:\Users\V_GUAN~1\AppData\Local\Temp\ng-T8prCL\angular-errors.log" for further details.

# 强制升级到8
ng update @angular/cli@8 @angular/core@8 --force

# 升级antd 到8

ng update ng-zorro-antd@8

```

