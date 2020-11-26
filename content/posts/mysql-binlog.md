---
title: "Mysql Binlog"
date: 2020-11-24T10:43:56+08:00
draft: false
---

binlog是Mysql sever层维护的一种二进制日志，与innodb引擎中的redo/undo log是完全不同的日志；其主要是用来记录对mysql数据更新或潜在发生更新的SQL语句，并以"事务"的形式保存在磁盘中；

作用主要有：

- 复制：MySQL Replication在Master端开启binlog，Master把它的二进制日志传- 递给slaves并回放来达到master-slave数据一致的目的
- 数据恢复：通过mysqlbinlog工具恢复数据
- 增量备份

### binlog管理

- 开启binlogmy.cnf配置中设置：log_bin="存放binlog路径目录"

- reset master 清空binlog日志文件


```bash
# binlog信息查询binlog开启后，可以在配置文件中查看其位置信息，也可以在myslq命令行中查看：
show variables like '%log_bin%';
+---------------------------------+-------------------------------------+
| Variable_name                   | Value                               |
+---------------------------------+-------------------------------------+
| log_bin                         | ON                                  |
| log_bin_basename                | /var/lib/mysql/3306/mysql-bin       |
| log_bin_index                   | /var/lib/mysql/3306/mysql-bin.index |
| log_bin_trust_function_creators | OFF                                 |
| log_bin_use_v1_row_events       | OFF                                 |
| sql_log_bin                     | ON                                  |
+---------------------------------+-------------------------------------+

# binlog文件开启binlog后，会在数据目录（默认）生产host-bin.n（具体binlog信息）文件及host-bin.index索引文件（记录binlog文件列表）。当binlog日志写满(binlog大小max_binlog_size，默认1G),或者数据库重启才会生产新文件，但是也可通过手工进行切换让其重新生成新的文件（flush logs）；另外，如果正使用大的事务，由于一个事务不能横跨两个文件，因此也可能在binlog文件未满的情况下刷新文件

# 查看binlog文件列表,
mysql> show binary logs; 
+------------------+-----------+
| Log_name         | File_size |
+------------------+-----------+
| mysql-bin.000001 |       177 |
| mysql-bin.000002 |       177 |
| mysql-bin.000003 |  10343266 |
| mysql-bin.000004 |  10485660 |
| mysql-bin.000005 |     53177 |
| mysql-bin.000006 |      2177 |
| mysql-bin.000007 |      1383 |
+------------------+-----------+

# 查看binlog的状态：show master status可查看当前二进制日志文件的状态信息，显示正在写入的二进制文件，及当前position
  mysql> show master status;
 +------------------+----------+--------------+------------------+-------------------+
 | File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
 +------------------+----------+--------------+------------------+-------------------+
 | mysql-bin.000007 |      120 |              |                  |                   |
 +------------------+----------+--------------+------------------+-------------------+
```


默认情况下binlog日志是二进制格式，无法直接查看。可使用两种方式进行查看：

- mysqlbinlog: /usr/bin/mysqlbinlog  mysql-bin.000007

    - mysqlbinlog是mysql官方提供的一个binlog查看工具，
    - 也可使用–read-from-remote-server从远程服务器读取二进制日志，
    - 还可使用--start-position --stop-position、--start-time= --stop-time精确解析binlog日志

- 直命令行解析

    SHOW BINLOG EVENTS
        [IN 'log_name'] //要查询的binlog文件名
        [FROM pos]  
        [LIMIT [offset,] row_count]  


### binlog格式


### 参考链接 
- [Binary Log Versions](https://dev.mysql.com/doc/internals/en/binary-log-versions.html)
- [What is binlog in mysql?](https://stackoverflow.com/questions/1366184/what-is-binlog-in-mysql)