---
title: "Linux Command Tar"
date: 2021-01-13T11:21:48+08:00
draft: false
---

Windows下最常见的压缩文件就只有两种，一是,zip，另一个是.rar。但Linux就不同，它有.gz、.tar.gz、tgz、bz2、.Z、.tar等众多的压缩文件名，此外windows下的.zip和.rar也可以在Linux下使 用，不过在Linux使用.zip和.rar的人就太少了。（因为 7z 和 zip 压缩格式都不能保留 unix 风格的文件权限，比如解压出个可执行文件要重新 chmod chown 才能恢复正常。而 tar 格式可以。而 tar 本身不提供压缩，无非就是把包括所有文件的內容和权限拼成一个文件而己，所以用另外如 gzip 格式压缩。为什么是 gzip，因为几乎所有 linux 都支持而已。)

两个概念：打包和压缩。

- 打包是指将一大堆文件或目录什么的变成一个总的文件
- 压缩则是将一个大的文件通过一些压缩算法变成一个小文件。

为什么要区分这 两个概念呢？其实这源于Linux中的很多压缩程序只能针对一个文件进行压缩，这样当你想要压缩一大堆文件时，你就得先借助另外的工具将这一大堆文件先打 成一个包，然后再就原来的压缩程序进行压缩。

Linux下最常用的打包程序就是tar了，使用tar程序打出来的包我们常称为tar包，tar包文件的命令通常都是以.tar结尾的。生成tar包后，就可以用其它的程序来进行压缩了，所以首先就来讲讲tar命令的

### 基本用法：

tar命令的选项有很多(用man tar可以查看到)，但常用的就那么几个选项，下面来举例说明一下：

```bash
tar -cf all.tar *.jpg
# 这条命令是将所有.jpg的文件打成一个名为all.tar的包。-c是表示产生新的包，-f指定包的文件名。

tar -rf all.tar *.gif
# 这条命令是将所有.gif的文件增加到all.tar的包里面去。-r是表示增加文件的意思。

tar -uf all.tar logo.gif
# 这条命令是更新原来tar包all.tar中logo.gif文件，-u是表示更新文件的意思。

tar -tf all.tar
# 这条命令是列出all.tar包中所有文件，-t是列出文件的意思
tar -xf all.tar
# 这条命令是解出all.tar包中所有文件，-x是解包的意思

```

以上就是tar的最基本的用法。为了方便用户在打包解包的同时可以压缩或解压文件，tar提供了一种特殊的功能。这就是tar可以在打包或解包的同时调用其它的压缩程序，比如调用gzip、bzip2等。

1. tar调用gzip

gzip是GNU组织开发的一个压缩程序，.gz结尾的文件就是gzip压缩的结果。与gzip相对的解压程序是gunzip。tar中使用-z这个参数来调用gzip。下面来举例说明一下：

```bash　　
tar -czf all.tar.gz *.jpg
# 这条命令是将所有.jpg的文件打成一个tar包，并且将其用gzip压缩，生成一个gzip压缩过的包，包名为all.tar.gz

tar -xzf all.tar.gz
# 这条命令是将上面产生的包解开。
```

2. tar调用bzip2

bzip2是一个压缩能力更强的压缩程序，.bz2结尾的文件就是bzip2压缩的结果。与bzip2相对的解压程序是bunzip2。tar中使用-j这个参数来调用bzip2。下面来举例说明一下：

```bash
tar -cjf all.tar.bz2 *.jpg
# 这条命令是将所有.jpg的文件打成一个tar包，并且调用bzip2压缩，生成一个bzip2压缩过的包，包名为all.tar.bz2
tar -xjf all.tar.bz2
# 这条命令是将上面产生的包解开。
```

3. tar调用compress

compress也是一个压缩程序，但是好象使用compress的人不如gzip和bzip2的人多。.Z结尾的文件就是bzip2压缩的结果。与compress相对的解压程序是uncompress。tar中使用-Z这个参数来调用gzip。下面来举例说明一下：

```bash
tar -cZf all.tar.Z *.jpg
# 这条命令是将所有.jpg的文件打成一个tar包，并且调用compress压缩，生成一个uncompress压缩过的包，包名为all.tar.Z

tar -xZf all.tar.Z
# 这条命令是将上面产生的包解开
```

### Linux下最常用的压缩工具

1、tar

```bash
# 压缩
tar -zcvf too.tar.gz too

# 解压
.tar.gz tar -zxvf  too.tar.gz
.tar.gz2 tar -jxvf   too.tar.gz2

# 把too目录打包成too.tar.gz，除logs目录；注这里的too/logs后面不能加/，如果加的话还是会打包进去。
tar -czvf too.tar.gz --exclude=too/logs too

# 加-C参数， 这样的话可以只打包api3.0 ，而不会从/usr开始一个一个目录都打包进去
tar -zcf api3.0_`date +%Y%m%d%H%M%S`.tar.gz -C /usr/local/tomcat/webapps api3.0
```

2、gzip

gzip工具是Linux中最流行、最快的文件压缩工具，Gzip工具保留原始文件名称压缩文件的扩展名.gz和时间戳。

```bash
# 打包
gzip filename

# 解压
gzip -d filename # 打包的文件会被删除
```

3、bzip2

Bzip2实用程序执行更快的gzip，它压缩文件和文件夹更紧凑。压缩文件时需要更多的内存，为了减少内存消耗，在选项中通过-s标志。

```bash
# 压缩
bzip2 examplefile or bzip2 -s examplefile

# 解压
bzip2 -d examplefile.bz2 or bunzip2 examplefile.bz2

# 详细说明
bzip2 -v examplefile
```
4、lzma

Lzma是一种压缩工具，与zip或tar类似，但与bzip相比，它的执行速度更快，虽然lzma是一个强大的工具，但它在Linux用户中并不流行。

```bash
# 压缩
lzma -c --stdout examplefile> examplefile.lzma

# 解压
lzma -d --stdout examplefile.lzma >examplefile
```

5、xz

XZ是lzma实用程序的继承者，它只能压缩单个文件，但不能在一个命令中压缩多个文件，它将自动为压缩文件添加.xz扩展名。
```bash
# 压缩
xz examplefile

# 解压
xz -d examplefile

```
6、pax

Pax它的执行速度很快，而且它不仅仅是一个压缩器，它可以真正的归档它可以远程复制文件，在Ubuntu/Mint Linux中，默认情况下Pax没有安装。

```bash
# 压缩
pax -wf examplefile.tar examplefile
pax -wf examplefile.tar.gz examplefile 

# 解压
pax -r <examplefile.tar

# 查看压缩包文件清单
pax -f examplefile.tar
```

7、7zip

7Zip文件压缩器是一个开源工具，它最初是为微软Windows开发的，它支持多种文件压缩格式和高文件压缩，它可以用一个命令压缩多个文件。

```bash
# 安装7zip
wget https://www.mirrorservice.org/sites/dl.fedoraproject.org/pub/epel/7/x86\_64/Packages/p/p7zip-16.02-10.el7.x86\_64.rpm
wget https://www.mirrorservice.org/sites/dl.fedoraproject.org/pub/epel/7/x86\_64/Packages/p/p7zip-plugins-16.02-10.el7.x86\_64.rpm
 
sudo rpm -U --quiet p7zip-16.02-10.el7.x86_64.rpm 
sudo rpm -U --quiet p7zip-plugins-16.02-10.el7.x86_64.rpm 

# 压缩
7z  a examplefile.7z examplefile

# 解压
7z  a examplefile.7z examplefile
```

8、shar

Shar是一个命令行工具，可以用来压缩测试文件，Shar可以定义为“shell archive”。一个简单而快速的文件存档实用程序对于获取shell脚本的存档非常有用。

```bash
# 安装shar工具
yum -y install sharutils

# 压缩
shar examplefile > examplefile.shar

# 解压
unshar examplefile.shar
```

9、cpio

可以定义为复制输入和输出，它在输入中逐行读取文件名列表，在输出中读取归档文件。这是一个内置的经典命令。

```bash
# 压缩
ls | cpio -ov >/home/username/backup.cpio

# 解压
cpio -idv <backup.cpio
```

10、ar

rar的前身，仍然在Debian及其衍生物中使用，它是一个简单的归档工具，但并不是很流行。

```bash
# 压缩
ar cvsr examplefile.a examplefile

# 解压
ar -xv examplefile.a
```

11、iso

```bash
# ISO制作iso镜像
dd if=/media/dvd of=/home/username/filename.iso
```


