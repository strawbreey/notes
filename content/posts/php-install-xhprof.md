---
title: "Php Install Xhprof"
date: 2020-11-30T11:51:49+08:00
draft: false
---


Xhprof是facebook开源出来的一个php轻量级的性能分析工具，跟Xdebug类似，但性能开销更低，还可以用在生产环境中，也可以由程序开关来控制是否进行profile。

### 安装xhprof

```bash
  # 下载xhprof 
  wget https://github.com/longxinH/xhprof/archive/v2.1.0.tar.gz

  # 解压缩包
  tar xf v2.1.0.tar.gz

  cd xhprof-2.1.0/extension

  # 运行phpize
  /data/nmp/php7/bin/phpize

  # 运行configure
  ./configure --with-php-config=/data/nmp/php7/bin/php-config
  
  # 编译安装
  make && make install
```

修改php.ini

```ini
[xhprof]
extension = xhprof.so
xhprof.output_dir = /tmp/xhprof
```

### 配置xhprof


把扩展目录下xhprof_html和xhprof_lib 复制出来，单独存放在一个目录中

```
cp -a xhprof_html /data/wwwroot/xhprof
cp -a xhprof_lib /data/wwwroot/xhprof
```

把要针对的项目，添加到入口文件中

```js
//开启性能分析
xhprof_enable(XHPROF_FLAGS_MEMORY | XHPROF_FLAGS_CPU);
//php中止时运行的函数
register_shutdown_function(function () {
    //停止性能分析
    $xhprof_data = xhprof_disable();
    if (function_exists('fastcgi_finish_request')) {
        fastcgi_finish_request();
    }
    //引入xhprof库文件，路径请自行修改
    $XHPROF_ROOT = realpath(dirname(__FILE__) . '/xhprof');
    include_once $XHPROF_ROOT . "/xhprof_lib/utils/xhprof_lib.php";
    include_once $XHPROF_ROOT . "/xhprof_lib/utils/xhprof_runs.php";
    $xhprof_runs = new XHProfRuns_Default();
    //导出性能分析数据，默认xhprof.output_dir指定的目录
    $run_id = $xhprof_runs->save_run($xhprof_data, 'xhprof');
});
```

数据默认会导出到xhprof.output_dir指定目录, 给以创建目录相关权限

```bash
mkdir /tmp/xhprof
chmod -R 777 /tmp/xhprof
```


### 配置虚拟主机，用来查看分析日志

```bash
server {
    listen       80;
    server_name  xhprof.xxx.com;
    charset utf-8;
  
    root   /data/wwwroot/xhprof/xhprof_html;
    index  index.html index.htm index.php;
 
    location ~ \.php$ {
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

### 安装 可视化工具

```bash
yum install -y libpng
yum install -y graphviz

# 手动安装
https://graphviz.gitlab.io/_pages/Download/Download_source.html

# 编译系统
tar xf graphviz.tar.gz
cd graphviz-2.40.1
./configure
make
make install
```

