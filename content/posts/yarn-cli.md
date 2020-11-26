---
title: "Yarn Cli"
date: 2020-09-18T10:11:16+08:00
draft: false
---

yarn is a package manager for your code. It allows you to use and share code with other developers from around the world. Yarn does this quickly, securely, and reliably so you don't ever have to worry.

Yarn allows you to use other developers' solutions to different problems, making it easier for you to develop your software. If you have problems, you can report issues or contribute back, and when the problem is fixed, you can use Yarn to keep it all up to date.


### Installation

- Global install

```shell
# 全局安装
npm install -g yarn
```

- Per-project install

```shell
cd ~/path/to/project

# 设置yarn 的版本
yarn set version berry

# Updating to the latest versions
yarn set version latest

# Installing the latest build fresh from master
yarn set version from sources
# or
yarn set version from sources --branch 121
```

### Usage

```shell
# Accessing the list of commands 
yarn help

# Starting a new project
yarn init

# Installing all the dependencies
yarn # or
yarn install

# upgrading a dependency
yarn up [package]
yarn up [package]@[version]
yarn up [package]@[tag]

# Removing a dependency
yarn remove [package]

# Upgrading Yarn itself
yarn set version latest
yarn set version from sources

```

### yarn CLI
yarn add 

```shell
# usage
yarn add [--json] [-E,--exact] [-T,--tilde] [-C,--caret] [-D,--dev] [-P,--peer] [-O,--optional] [--prefer-dev] [-i,--interactive] [--cached] ...

# example
yarn add lodash
yarn add lodash@1.2.3
yarn add lodash@https://github.com/lodash/lodash
yarn add lodash@lodash/lodash
yarn add lodash-es@lodash/lodash#es
```

yarn bin

```shell
# usage
yarn bin [-v,--verbose] [--json] [name]

# example
yarn bin
yarn bin eslint

```
yarn cache clean

```shell
# usage
yarn cache clean [--mirror] [--all]

# example
yarn cache clean
yarn cache clean --mirror
```
yarn config get

```shell
# usage
yarn config get [--json] [--no-redacted] <name>

# example
yarn config get yarnPath
yarn config get packageExtensions
yarn config get 'npmScopes["my-company"].npmRegistryServer'
yarn config get npmAuthToken --no-redacted
yarn config get packageExtensions --json
```

yarn config set
```shell
# usage
yarn config set [--json] [-H,--home] <name> <value>

# example
yarn config set initScope myScope
yarn config set initScope --json \"myScope\"
yarn config set unsafeHttpWhitelist --json '["*.example.com", "example.com"]'
yarn config set packageExtensions --json '{ "@babel/parser@*": { "dependencies": { "@babel/types": "*" } } }'
```

yarn config

```shell
# usage
yarn config [-v,--verbose] [--why] [--json]

# example
yarn config set initScope myScope
yarn config set initScope --json \"myScope\"
yarn config set unsafeHttpWhitelist --json '["*.example.com", "example.com"]'
yarn config set packageExtensions --json '{ "@babel/parser@*": { "dependencies": { "@babel/types": "*" } } }'
```

### Words
```
reliably  英 [rɪ'laɪəbli]   美 [rɪ'laɪəbli] 
adv. 可靠地
```


### 参考链接 

- [yarn cli](https://yarnpkg.com/cli/add)
- [There appears to be trouble with your network connection. Retrying](https://stackoverflow.com/questions/51508364/yarn-there-appears-to-be-trouble-with-your-network-connection-retrying)