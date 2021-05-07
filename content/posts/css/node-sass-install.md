---
title: "Node Sass Install"
date: 2020-11-12T17:43:31+08:00
draft: false
---

```bash
  gyp info it worked if it ends with ok
  gyp info using node-gyp@5.1.0
  gyp info using node@14.4.0 | win32 | x64
  gyp info find Python using Python version 2.7.16 found at "C:\python2\python.exe"
  gyp http GET https://nodejs.org/download/release/v14.4.0/node-v14.4.0-headers.tar.gz
  gyp WARN install got an error, rolling back install
  gyp ERR! configure error
```

![npm-install-node-sass-error](/images/npm-install-node-sass-error.png)


我们可以看到，安装node-sass有几个步骤：

- 校验本地node_modules中是否已安装node-sass，版本是否一致;

- 如未安装或版本不符，从npm源安装node-sass本体;

- 检测全局缓存和本地中是否有binding.node,如有即跳过安装;

- 没有binding.node则从github下载该二进制文件并将其缓存到全局;

- 假如binding.node下载失败，则尝试本地编译出该文件;

- 将版本信息写到package-lock.json;


### 检查 python 版本

同时安装python2 和 python3 的时候，python3 回默认配置环境遍历。但python2不会

所以python2需要手动加入环境变量中

- 右键点击计算机/此电脑 - 高级系统设置 - 环境变量 - 编辑
- 找到安装python2的默认选择的目录：C:\Python27和C:\Python27\Scripts，复制这两个个目录。
- 点击 编辑环境变量 窗口的 新建，将上一步复制的两个步骤分别添加到环境变量。
- 回到python2的根目录 C:\Python27，将python.exe 改为 python2.exe，然后在cmd 命令窗口输入python2就可以看到安装的python2版本了。


### 检查 yarn 的镜像源

```bash
  # 1、查看一下当前源
  yarn config get registry

  # 2、切换为淘宝源
  yarn config set registry https://registry.npm.taobao.org

  # 3、或者切换为自带的
  yarn config set registry https://registry.yarnpkg.com
```

### 相关依赖

- [node-sass](https://github.com/sass/node-sass)

- [node-gyp](https://github.com/nodejs/node-gyp)

- [What are node.js bindings?](https://stackoverflow.com/questions/20382396/what-are-node-js-bindings)

- [C++ BINDING WITH NODE.JS](https://pravinchavan.wordpress.com/2013/11/08/c-binding-with-node-js/)