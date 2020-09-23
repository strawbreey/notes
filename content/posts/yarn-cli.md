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
npm install -g yarn
```

- Per-project install

```shell
cd ~/path/to/project

yarn policies set-version berry
yarn set version berry
```

- Updating to the latest versions

```
yarn set version latest
```

- Installing the latest build fresh from master

```shell
yarn set version from sources
# or
yarn set version from sources --branch 121
```

### Usage

Accessing the list of commands 
```shell
yarn help
```

Starting a new project
```shell
yarn init
```

Installing all the dependencies
```shell
yarn
# or
yarn install
```

Upgrading a dependency
```shell
yarn up [package]
yarn up [package]@[version]
yarn up [package]@[tag]
```

Removing a dependency
```shell
yarn remove [package]
```

Upgrading Yarn itself
```shell
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