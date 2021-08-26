---
title: "Linux Install Svn"
date: 2020-12-29T17:03:17+08:00
draft: false
---

SVN（Subversion）是一个自由、开源的项目源代码版本控制工具。目前，绝大多数开源软件和企业代码管理，都使用SVN作为代码版本管理软件。

Subversion将文件存放在中心版本库里，这个版本库很像一个普通的文件服务器。不同的是，它可以记录每一次文件和目录的修改情况，这样就可以在需要的回滚时，将数据恢复到以前的版本，并可以查看数据的更改细节。

SVN 官方网址：Apache Subversion

SVN 服务器配置：安装SVN服务器

### svn 安装

查看svn版本

```bash
svnserve --version
```

ubuntu 安装svn 1.7、1.8

当前 ubuntu 12.04 中的 svn 版本为 1.6，这个版本会在每个子目录新建一个.svn 的目录保存版本文件，很不爽。找到一个第三方编译的 ubuntu 源：

```bash
# 打开source.list源：
sudo vi /etc/apt/sources.list   

# 添加源
deb http://ppa.launchpad.net/dominik-stadler/subversion-1.7/ubuntu precise main
deb-src http://ppa.launchpad.net/dominik-stadler/subversion-1.7/ubuntu precise main

# 执行安装命令
sudo apt-get install subversion
```

### SVN 常用命令

```bash

# 1. 将文件checkout到本地目录
svn checkout path
svn co path

# 2. 往版本库中添加新的文件
svn add file

# 3. 将改动的文件提交到版本库
svn commit -m “LogMessage” [-N] [--no-unlock] PATH(如果选择了保持锁，就使用–no-unlock开关)
# 简写：svn ci

# 4. 加锁/解锁

svn lock -m “LockMessage” [--force] PATH
svn unlock PATH

# 5. 更新到某个版本
svn update -r m path # 简写：svn up



# 6. 查看文件或者目录状态
1）svn status path（目录下的文件和子目录的状态，正常状态不显示）
2）svn status -v path(显示文件和子目录状态)

# 简写：svn st

# 7. 删除文件
svn delete path -m “delete test fle”
# 简写：svn (del, remove, rm)

# 8. 查看日志
svn log path

# 9. 查看文件详细信息
svn info path

# 10. 比较差异

svn diff path (将修改的文件与基础版本比较)
svn diff -r m:n path (对版本m和版本n比较差异)

# 简写：svn di

# 11. 将两个版本之间的差异合并到当前文件
svn merge -r m:n path

# 12. SVN 帮助
svn help
svn help ci

# 13. 版本库下的文件和目录列表

svn list path    

显示path目录下的所有属于版本库的文件和目录简写：svn ls

# 14. 创建纳入版本控制下的新目录

svn mkdir: 创建纳入版本控制下的新目录。

用法: 
1、mkdir PATH... 每一个以工作副本 PATH 指定的目录，都会创建在本地端，并且加入新增调度，以待下一次的提交。
2、mkdir URL... 创建版本控制的目录。 每个以URL指定的目录，都会透过立即提交于仓库中创建。在这两个情况下，所有的中间目录都必须事先存在。

# 15. 恢复本地修改
svn revert: 恢复原始未改变的工作副本文件 (恢复大部份的本地修改)。
用法: revert PATH... 注意: 本子命令不会存取网络，并且会解除冲突的状况。但是它不会恢复被删除的目录

# 16. 代码库URL变更
svn switch (sw): 更新工作副本至不同的URL。
用法: 
1、switch URL [PATH]        
更新你的工作副本，映射到一个新的URL，其行为跟“svn update”很像，也会将      服务器上文件与本地文件合并。这是将工作副本对应到同一仓库中某个分支或者标记的方法。 
2、switch --relocate FROM TO [PATH...]   
改写工作副本的URL元数据，以反映单纯的URL上的改变。当仓库的根URL变动     (比如方案名或是主机名称变动)，但是工作副本仍旧对映到同一仓库的同一目录时使用     这个命令更新工作副本与仓库的对应关系。

# 17. 解决冲突
svn resolved: 移除工作副本的目录或文件的“冲突”状态。
用法: resolved PATH... 注意: 本子命令不会依语法来解决冲突或是移除冲突标记；它只是移除冲突的相关文件，然后让 PATH 可以再次提交。

# 18. 输出指定文件或URL的内容。
svn cat 目标[@版本]...如果指定了版本，将从指定的版本开始查找。 svn cat -r PREV filename > filename (PREV 是上一版本,也可以写具体版本号,这样输出结果是可以提交的）
 
# 19. svn cleanup
当Subversion修改你的工作副本时（或者任何在.svn中的信息），它尝试尽可能做到安全。在改变一个工作副本前，Subversion把它的意 图写到一个日志文件中。接下来它执行日志文件中的命令来应用要求的修改。最后，Subversion删除日志文件。从架构上来说，这与一个日志文件系统 （journaled filesystem）类似。如果一个 Subversion操作被打断（例如，进程被杀掉了，或机器当掉了）了，日志文件仍在硬盘上。重新执行日志文件，Subversion可以完成先前开始 的操作，这样你的工作副本能回到一个可靠的状态。 

# 20. svn import
使用svn import是把未版本化的文件树复制到资料库的快速办法，它需要创建一个临时目录。 

```

### SVN 常用命令一览表

| 命令     | 功能	    | 使用格式   |
| ---      | ---     | ---        |
| checkout | 检出	    |svn co URL |
| up       | 更新到当前URL的末端 |svn  up |
| switch	 | 更新到某一tag/branch	svn  switch  (tag/分支)URL
| add	      |增加	svn  add  文件名
| rm	    |删除文件	svn  rm 文件名
| 删除目录	svn  rm 目录名
| diff	与base版本（最后检出或者更新到的版本）对比	svn  diff
| 与版本库中最新版本对比	svn  diff  -r  head
| 当前工作副本，两个版本之间对比	svn  diff  -r  reversion1:reversion2
| 版本库中任意两个tag做对比	svn   diff    (tag1)URL    (tag2)URL
| ci	提交	svn ci -m "commit log"
| log	查看当前工作副本log	svn  log
| 只查看指定版本的log	svn  log  -r
| 打印log所有附加信息	svn  log  -v
| 查看当前tag/branch版本详情	svn  log --stop-on-copy -v
| info	查看当前工作副本所在URL	svn  info
| status	查看工作副本的状态	svn st
| 查看文件的taglist	svn命令不支持，可执行cs taglist
| tag	新增tag	svn cp . （tag）URL
| 删除tag	svn rm （tag）URL -m "commit log"
| 覆盖已经存在的tag	不支持
| 分支开发	创建branch	svn  cp  （基线版本）URL （分支）URL  -m "commit log"
| 删除branch	svn rm （分支）URL   -m "commit log"
| 同步	svn co （主干）URL
| cd ~/wc
| svn merge （主干）URL （待同步tag）URL
| svn ci -m "commit log"
| svn cp （主干）URL （以_PD_BL_MAIN结尾的tag）URL -m"commit log"
| 合并	svn co （合并目标）URL
| cd ~/wc
svn merge （基线版本tag）URL  （上线tag）URL
svn ci -m "commit log"
svn cp （合并目标）URL （上线tag_MERGE_的tag对应）URL -m"commit log"



SVN实例

删除目录下所有的 .svn 隐藏子目录

find . -name ".svn" -print0 | xargs -0 rm -rf



tags打分支

svn  cp  trunk/    tags/platform_2011.11.11    （或 svn  cp  http://192.168.1.100/platform/trunk/    http://192.168.1.100/platform/tags/platform_2011.11.11）

svn  ci  -m  "svn  cp  trunk/    tags/platform_2011.11.11"         // 提交，并给出提交记录（-m  "svn  cp  trunk/    tags/platform_2011.11.11"）



svn 改名

svn  mv  platform_2011.11.11   platform_20111111

svn  ci  -m  "svn  mv  platform_2011.11.11   platform_20111111"        // 提交



svn directory is missing

1） svn up missingDirName
2） svn del missingDirName
3） svn ci



svn chech version

svn  co  http://192.168.1.100/platform/branch  -r  12  platform_branch_v12



svn log

svn  log  http://192.168.1.100/platform/branch  -l10              // svn  文字注释log

svn  log  http://192.168.1.100/platform/branch  -l10  -v          // svn 文字注释log + 文件更新log（增，删，改）



svn diff -r  v_1 : v_2 svn_path

svn diff -r 200:201 test.php









linux安装svn服务器（yum方式）
分类 编程技术
1.查看yum是否安装
在终端中输入yum即可如果已经安装，会显示yum的参数,如果没有安装，会提示yum未安装或无效命令……

2.安装svn
yum -y install subversion
出现下面截图后，表示安装完成。


3.创建目录并配置
建立版本库目录

mkdir -p /data/svn/repo
配置 /p>

svnadmin create  /data/svn/repo
执行上面的命令后，自动建立repo测试库，查看/data/svn/repo 文件夹发现包含了conf, db,format,hooks, locks, README.txt等文件，说明一个SVN库已经建立。

4.用户密码passwd配置
[root@localhost password]# cd /data/svn/repos/conf
[root@admin conf]# vim passwd
修改passwd为以下内容：

[users]
# harry = harryssecret
# sally = sallyssecret
hello=123
用户名=密码

这样我们就建立了hello用户， 123密码

以上语句都必须顶格写, 左侧不能留空格, 否则会出错.

5.权限控制authz配置
[root@admin conf]# vi  authz
目的是设置哪些用户可以访问哪些目录，向authz文件追加以下内容：

#设置[/]代表根目录下所有的资源   或者写成[repl:/]
[/]
hello = rw
意思是hello用户对repo测试库下所有的目录有读写权限，当然也可以限定。 如果是自己用，就直接是读写吧。

以上语句都必须顶格写, 左侧不能留空格, 否则会出错.

6.服务svnserve.conf配置
[root@admin conf]# vim svnserve.conf
追加以下内容：

[general]
#匿名访问的权限，可以是read,write,none,默认为read
anon-access=none
#使授权用户有写权限 
auth-access=write
#密码数据库的路径 
password-db=passwd
#访问控制文件 
authz-db=authz
#认证命名空间，subversion会在认证提示里显示，并且作为凭证缓存的关键字 
realm=/data/svn/repositories
以上语句都必须顶格写, 左侧不能留空格, 否则会出错.

7.启动svn服务
查看svn的服务是否启动

ps -ef|grep svn
如果没有启动执行 下面命令即可：

svnserve -d -r /data/svn/repo  --listen-port=3690 
8.开启端口
 /sbin/iptables -I INPUT -p tcp --dport 3690 -j ACCEPT
 /etc/rc.d/init.d/iptables save
 /etc/init.d/iptables restart
 /etc/init.d/iptables status
9.访问
svn://ip地址

端口默认为3690，输入配置好的用户名和密码即可。

注意：此处千万不能 svn://ip地址/repo，否则无法访问。

### 参考资料

- [linux安装svn服务器（yum方式）](https://www.runoob.com/w3cnote/linux-subversion-yum.html)