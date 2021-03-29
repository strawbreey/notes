---
title: "Php Time Task"
date: 2021-03-25T14:22:47+08:00
draft: false
---

## PHP 定时任务

1. 固定执行文件脚本

```bash
# PHP 定时任务
* */1 * * * /bin/sh /home/wwwroot/php-script.sh
```

2. php-script.sh 脚本

```bash
#!/bin/sh
php /home/wwwroot/wkr2019.com/Public/cli.php /cli/task/exec >> /home/wwwlogs/sh.log;
php /home/wwwroot/op.cheduo.com/think task >> /home/wwwlogs/sh.log;
```

3. php 文件代码

```php
<?php
namespace Cli\Controller;
use Think\Controller;

class TaskController extends Controller {

    public function _initialize()
    {
        // 判断执行环境
        if(PHP_SAPI != 'cli'){
            exit('deny!');
        }
    }

    function exec(){
        $hour = date('H');  // 小时

        if(intval($hour) == 0) {
            D('YybStatistics')->increase([]);  // 自动进行今日统计
        }
    }
}
```