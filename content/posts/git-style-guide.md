---
title: "Git Style Guide"
date: 2021-03-24T15:41:12+08:00
draft: false
---

## Branches

- 选择简短和具有描述性的名字来命名分支：

```bash
# 好
$ git checkout -b oauth-migration

# 不好，过于模糊
$ git checkout -b login_fix
```

- 来自外部的标识符也适合用作分支的名字，例如来自 Github 的 Issue 序号。

```bash
# GitHub issue #15
$ git checkout -b issue-15
```

- 用破折号分割单词。

- 当不同的人围绕同一个特性开发时，维护整个团队的特性分支与每个人的独立分支是比较方便的做法。使用如下的命名方式：

```bash
git checkout -b feature-a/master # team-wide branch
git checkout -b feature-a/maria # Maria's branch
git checkout -b feature-a/nick # Nick's branch
```

- 合并之后，除非有特殊原因，从上游仓库中删除你的分支。使用如下命令查看已合并的分支：

```bash
git branch --merged | grep -v "\*"
```

- 命名规范

请务必使用前缀，以下为常见前缀，例如：feat/update-user-info

feat(feature):     添加特性的代码更改，比如新增某些功能的代码提交；这个是最常用的提交前缀；
fix:      bug修复
docs:     只是改了文档信息，比如readme.md
style:    不影响代码含义的更改，比如只是增加了空格，格式化了代码，增加了缺少的分号等
refactor: 既不修复bug也不添加特性的代码更改
perf:     改进性能的代码更改
test:     添加缺失的测试或纠正现有的测试的代码提交，比如新增单测的代码提交
build:    影响构建系统或外部依赖的更改(示例范围:gulp, broccoli, npm)
ci:       对CI配置文件和脚本的更改(示例范围:Travis、Circle、BrowserStack、SauceLabs)
chore:    其他的不修改src或test文件夹下的更改
revert:   还原以前的提交的更改

