---
title: "Php Filter"
date: 2020-09-23T15:56:14+08:00
draft: false
tags: ['php']

---

### 过滤器函数

filter_has_var — 检测是否存在指定类型的变量
filter_id — 返回与某个特定名称的过滤器相关联的id
filter_input_array — 获取一系列外部变量，并且可以通过过滤器处理它们
filter_input — 通过名称获取特定的外部变量，并且可以通过过滤器处理它
filter_list — 返回所支持的过滤器列表
filter_var_array — 获取多个变量并且过滤它们
filter_var — 使用特定的过滤器过滤一个变量


#### filter_has_var

检测是否存在指定类型的变量

```php
filter_has_var ( int $type , string $variable_name ) : bool
```

type: INPUT_GET、 INPUT_POST、 INPUT_COOKIE、 INPUT_SERVER、 INPUT_ENV 里的其中一个。

variable_name 要检查的变量名。


### filter_id  返回与某个特定名称的过滤器相关联的id

filter_id ( string $filtername ) : int


### filter_input_array 获取一系列外部变量，并且可以通过过滤器处理它们

```php
filter_input_array ( int $type [, mixed $definition [, bool $add_empty = true ]] ) : mixed
```