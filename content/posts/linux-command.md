---
title: "Linux Command"
date: 2020-10-13T10:59:36+08:00
draft: false
---

## Linux Command

1. 查看磁盘空间大小 

```shell
df -hl # G
```

第一列Filesystem，磁盘分区

第二列Size，磁盘分区的大小

第三列Used，已使用的空间

第四列Avail，可用的空间

第五列Use%，已使用的百分比

第六列Mounted on，挂载点

 
参数说明

  -a或--all：包含全部的文件系统；

  --block-size=<区块大小>：以指定的区块大小来显示区块数目；

  -h或--human-readable：以可读性较高的方式来显示信息;

  -H或--si：与-h参数相同，但在计算时是以1000 Bytes为换算单位而非1024 Bytes；

  -i或--inodes：显示inode的信息；

  -k或--kilobytes：指定区块大小为1024字节；

  -l或--local：仅显示本地端的文件系统;

  -m或--megabytes：指定区块大小为1048576字节；

  --no-sync：在取得磁盘使用信息前，不要执行sync指令，此为预设值；

  -P或--portability：使用POSIX的输出格式；

  --sync：在取得磁盘使用信息前，先执行sync指令；

  -t<文件系统类型>或--type=<文件系统类型>：仅显示指定文件系统类型的磁盘信息；

  -T或--print-type：显示文件系统的类型；

  -x<文件系统类型>或--exclude-type=<文件系统类型>：不要显示指定文件系统类型的磁盘信息；

  --help：显示帮助；

  --version：显示版本信息。

2. 查看文件和目录大小

```shell
# 查看目录的总大小
du -sh /data

# 或者进到/data目录后直接执行：
du -sh

# 如果要看/data目录下各个子目录的大小，不包括子目录
du -h

# 如果要看/data目录下各个子目录的大小，包括子目录的子目录
du -h *

# 如果要看/data目录下各个子目录和文件的大小，需要使用-a参数：
du -ah

```

参数说明

-a或-all 显示目录中个别文件的大小。

-b或-bytes 显示目录或文件大小时，以byte为单位。

-c或--total 除了显示个别目录或文件的大小外，同时也显示所有目录或文件的总和。

-k或--kilobytes 以KB(1024bytes)为单位输出。

-m或--megabytes 以MB为单位输出。

-s或--summarize 仅显示总计，只列出最后加总的值。

-h或--human-readable 以K，M，G为单位，提高信息的可读性。

-x或--one-file-xystem 以一开始处理时的文件系统为准，若遇上其它不同的文件系统目录则略过。

-L<符号链接>或--dereference<符号链接> 显示选项中所指定符号链接的源文件大小。

-S或--separate-dirs 显示个别目录的大小时，并不含其子目录的大小。

-X<文件>或--exclude-from=<文件> 在<文件>指定目录或文件。

--exclude=<目录或文件> 略过指定的目录或文件。

-D或--dereference-args 显示指定符号链接的源文件大小。

-H或--si 与-h参数相同，但是K，M，G是以1000为换算单位。

-l或--count-links 重复计算硬件链接的文件。

3. 排序命令

sort命令可以用于将文件内容排序并输出，也可以用于将某些查询命令的执行结果排序后输出

```shell
# 将文件夹中的文件按大小排序，可以用以下命令
du -a|sort -rn
```

管道前面的du –a就是列出目录下所有的文件和目录的大小，后面的sort命令就是排序。

其中-r参数代表反向排序，因为sort默认是从小到大排序的，加-r是从大到小排序

-n代表按照数字排序，只认数字不认单位，本例中的数字就是文件大小，单位是默认的KB，所以这个命令不能用du -ah，这会使排序结果出现2M小于100K的情况。


参数说明：

  -b：忽略每行前面开始出的空格字符；

  -c：检查文件是否已经按照顺序排序；

  -d：排序时，处理英文字母、数字及空格字符外，忽略其他的字符；

  -f：排序时，将小写字母视为大写字母；

  -i：排序时，除了040至176之间的ASCII字符外，忽略其他的字符；

  -m：将几个排序号的文件进行合并；

  -M：将前面3个字母依照月份的缩写进行排序；

  -n：依照数值的大小排序；

  -o<输出文件>：将排序后的结果存入制定的文件；

  -r：以相反的顺序来排序；

  -t<分隔字符>：指定排序时所用的栏位分隔字符；

  +<起始栏位>-<结束栏位>：以指定的栏位来排序，范围由起始栏位到结束栏位的前一栏位。


4. 只显示前几行的命令

```bash
# 要将文件夹中的文件按大小排序，而且只看最大的几个
du -a|sort -rn|head -5

# head后面的-5表示显示前5行，不加数字则默认显示前10行
```

-n<数字>：指定显示头部内容的行数；
-c<字符数>：指定显示头部内容的字符数；
-v：总是显示文件名的头信息；
-q：不显示文件名的头信息。

5. 查看列表

list 的缩写，通过 ls 命令不仅可以查看 linux 文件夹包含的文件，而且可以查看文件权限(包括目录、文件夹、文件权限)查看目录信息等等。

```bash
ls -l

ls -a # 列出目录所有文件，包含以.开始的隐藏文件
ls -A # 列出除.及..的其它文件
ls -r # 反序排列
ls -t # 以文件修改时间排序
ls -S # 以文件大小排序
ls -h # 以易读大小显示
ls -l # 除了文件名之外，还将文件的权限、所有者、文件大小等信息详细列出来

# 按易读方式按时间反序排序，并显示文件详细信息
ls -lhrt

# 按大小反序显示文件详细信息
ls -lrS

# 列出当前目录中所有以"t"开头的目录的详细内容
ls -l t*

# 列出文件绝对路径（不包含隐藏文件）
ls | sed "s:^:`pwd`/:"

# 列出文件绝对路径（包含隐藏文件）
find $pwd -maxdepth 1 | xargs ls -ld

```

6. cd 命令

```bash
# 进入根目录
cd /

# 进入home 目录
cd ~

# 进入上一次工作路径
cd -

# 把上一个命令的参数作为cd参数使用
cd !$
```

6. pwd 命令

```bash

# 查看当前路径
pwd

# 查看软连接的实际路径
pwd -P
```

6. mkdir 

  -m: 对新建目录设置存取权限，也可以用 chmod 命令设置;
  -p: 可以是一个路径名称。此时若路径中的某些目录尚不存在,加上此选项后，系统将自动建立好那些尚不在的目录，即一次可以建立多个目录。

```bash
# 当前工作目录下创建名为 t的文件夹
mkdir t

# 在 tmp 目录下创建路径为 test/t1/t 的目录，若不存在，则创建：
mkdir -p /tmp/test/t1/t
```

6. rm

删除一个目录中的一个或多个文件或目录，如果没有使用 -r 选项，则 rm 不会删除目录。
如果使用 rm 来删除文件，通常仍可以将该文件恢复原状。

```bash
# 删除任何 .log 文件，删除前逐一询问确认：
rm -i *.log

# 删除 test 子目录及子目录中所有档案删除，并且不用一一确认：
rm -rf test

#删除以 -f 开头的文件
rm -- -f*
```

6. rmdir

从一个目录中删除一个或多个子目录项，删除某目录时也必须具有对其父目录的写权限。不能删除非空目录

```bash
# 当 parent 子目录被删除后使它也成为空目录的话，则顺便一并删除：
rmdir -p parent/child/child11
```

6. mv 命令
移动文件或修改文件名，根据第二参数类型（如目录，则移动文件；如为文件则重命令该文件）。
···

···

6. 查看系统版本

```bash
# 方式1
cat /etc/redhat-release
CentOS Linux release 7.6.1810 (Core) 

# 方式2
lsb_release -a
```

7. 查看CPU和内存

```bash
# cpu:
cat /proc/cpuinfo | grep name

# 内存(RAM)：
$ cat /proc/meminfo | grep Mem
```

8. 查看磁盘和分区

```bash
# 磁盘
fdisk -l | grep -E '.+/dev/'

# 分区:
df -TH | grep ^/dev
```

9. 查看网卡和IP地址

```shell
# 网卡:
lspci | grep -i eth

# IP
ip addr | grep -E "^[1-9]+|inet"
```


10. 查看主板制造商

```bash
  dmidecode | grep -V -A2 "System Information"
```

11. 查看防火墙服务

```bash
# 查看防火墙
systemctl list-unit-files | grep firewalld

# 开启防火墙
systemctl enable firewalld.service

# 禁用防火墙
systemctl disable firewalld.service

# 查看防火墙运行状态
firewall-cmd --state

# 打开防火墙
systemctl start firewalld.service

# 关闭防火墙
systemctl stop firewalld.service
```

12. 查看用户信息

```bash
# 查看用户信息
cat /etc/passwd | grep bash

#查看用户组
cat /etc/group | grep root

# 查看系统用户密码
cat /etc/shadow | grep root
```

13. 查看系统运行状态

```bash
# 查看当前运行的进程列表
ps aux
# 加 f 以树状显示父子进程
ps aufx

# 查看实时进程资源占用(CPU和内存)
top

# 系统运行状态监控(CPU和IO)
yum install dstat

dstat
```


### 参考文献

- [Linux 常用命令学习](https://www.runoob.com/w3cnote/linux-common-command-2.html)