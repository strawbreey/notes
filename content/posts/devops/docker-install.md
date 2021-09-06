---
title: "Docker Install"
date: 2020-12-10T10:56:24+08:00
draft: false
---

## 介绍 

Docker Desktop 是 Docker 在 Windows 10 和 macOS 操作系统上的官方安装方式，这个方法依然属于先在虚拟机中安装 Linux 然后再安装 Docker 的方法。

Docker Desktop 官方下载地址： https://hub.docker.com/editions/community/docker-ce-desktop-windows

### window 安装

1. 下载并安装docker

2. 安装 [window subsystem for linux](/pages/window-subsystem-for-linux)

3. 安装 Hyper-V
Hyper-V 是微软开发的虚拟机，类似于 VMWare 或 VirtualBox，仅适用于 Windows 10。这是 Docker Desktop for Windows 所使用的虚拟机。

但是，这个虚拟机一旦启用，QEMU、VirtualBox 或 VMWare Workstation 15 及以下版本将无法使用！如果你必须在电脑上使用其他虚拟机（例如开发 Android 应用必须使用的模拟器），请不要使用 Hyper-V！

4. 开启 Hyper-V

5. 设置代理  settings -> Docker Engine

```json
{
  "registry-mirrors": [
    "https://mirrors.tencent.com",
    "https://hub-mirror.c.163.com/",
    "https://<你的ID>.mirror.aliyuncs.com",
    "https://reg-mirror.qiniu.com"
  ],
  "insecure-registries": [],
  "debug": false,
  "experimental": false,
  "features": {
    "buildkit": true
  }
}
```

### linux 安装

1. 查看系统要求, 确认使用linux 版本

Docker 要求 CentOS 系统的内核版本高于 3.10, 查看CentOS的内核版本。

```bash
# linux 查看系统信息
uname -a 

# cat /etc/tlinux-release
```

2. 删除旧版本
```bash
yum remove docker  docker-common docker-selinux docker-engine
```

3. 安装需要的软件包

```bash
# yum-util 提供yum-config-manager功能，另外两个是devicemapper驱动依赖的
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

4. 设置Docker yum源

```bash
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

5. 查看所有仓库中所有docker版本

可以查看所有仓库中所有docker版本,并选择特定的版本安装。
```bash
yum list docker-ce --showduplicates | sort -r
```

6. 安装docker
```bash
# 由于repo中默认只开启stable仓库，故这里安装的是最新稳定版本。
sudo yum install docker-ce

# or
yum install docker-ce -y

# 如果要安装特定版本：
sudo yum install docker-ce-18.06.1.ce
```

7. 启动

```bash
# 设置为开机启动
systemctl enable docker

# 开启
systemctl start docker

# 查看启动状态
systemctl status docker

# 查看版本
docker version

# 重启
systemctl daemon-reload
systemctl restart docker
```

### 参考资料

- [Windows Docker 安装](https://www.runoob.com/docker/windows-docker-install.html)