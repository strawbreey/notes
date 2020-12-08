---
title: "Git Pull Slow"
date: 2020-11-22T20:53:33+08:00
draft: false
---

### 解决git节点清点耗时较长的方法

出现节点清点耗时较长的原因

由于remote与本地仓库长时间运行，导致两遍的仓库节点差异较大。git 每次fetch都会比对差异节点，将remote存在而本地不存在的节点fetch到本地，因此每次清点需要耗较长时间。

- 删除本地仓库重新clone

将本地仓库所有分支和变更都commit并push后。重新clone仓库。新拉下来的仓库与remote节点完全一致。这样清点节点耗时就非常短了，能大大加快清点耗时。

- 本地仓库执行节点清理命令
节点清理压缩（需较多内存）：在git所在目录执行命令：`git gc --aggressive --prune=now`。将本地节点清理压缩，减小差异节点清点耗时。


在网站https://www.ipaddress.com/分别搜索`github.global.ssl.fastly.net`和`github.com`，查到的IP与我之前在hosts存储的对应IP不同，使用新IP替换旧IP
使用`sudo killall -HUP mDNSResponder`命令刷新mac的DNS
由于看到github的两个IP属地为美国，因此将代理服务器切换至美国，并开启全局模式（实际验证PAC模式问题依旧），重新打开一个终端，git pull速度大幅提升


git clone特别慢是因为github.global.ssl.fastly.net域名被限制了。

只要找到这个域名对应的ip地址，然后在hosts文件中加上ip–>域名的映射，刷新DNS缓存便可。


在网站 https://www.ipaddress.com/ 分别搜索：
```
github.global.ssl.fastly.net
github.com
```

打开hosts文件

- Windows上的hosts文件路径在C:\Windows\System32\drivers\etc\hosts
- Linux的hosts文件路径在：sudo vim /etc/hosts

在hosts文件末尾添加两行(对应上面查到的ip)


```
199.232.69.194  github.global-ssl.fastly.net
140.82.113.3 github.com
```

保存更新DNS

- Winodws: CMD，输入ipconfig /flushdns
- Linux: 终端 输入sudo /etc/init.d/networking restart


