---
title: "Yii Create Index"
date: 2021-08-09T11:26:00+08:00
draft: false
---

### YII 创建索引

```php
    $this->createIndex(
        'thing_pre_supplier_id',
        '\{\{%thing\}\}',
        array('pre_supplier_id')
    );
```

createIndex() public method

Builds and executes a SQL statement for creating a new index.

public void createIndex ( $name, $table, $columns, $unique = false )
- $name	string	
The name of the index. The name will be properly quoted by the method.

- $table	string	
The table that the new index will be created for. The table name will be properly quoted by the method.

- $columns	string|array	
The column(s) that should be included in the index. If there are multiple columns, please separate them by commas or use an array. Each column name will be properly quoted by the method. Quoting will be skipped for column names that include a left parenthesis "(".

- $unique	boolean	
Whether to add UNIQUE constraint on the created index.