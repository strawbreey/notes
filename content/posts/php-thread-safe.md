---
title: "Php Thread Safe"
date: 2020-12-24T14:53:04+08:00
draft: false
---

非线程安全的 PHP 主要用于 IIS，因为 IIS 采用 FastCGI 方式调用 PHP，是自带多线程相关的冲突处理代码的，和 Apache、Nginx 上的运行方式不一样。但 [线程安全] 和 [非线程安全] 仅限于在 windows 环境下纠结。Linux 中很少用到多线程模型，官方也只提供了一个版本。
