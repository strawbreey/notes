---
title: "Docker Install"
date: 2020-12-10T10:56:24+08:00
draft: false
---

## 介绍 

Docker Desktop 是 Docker 在 Windows 10 和 macOS 操作系统上的官方安装方式，这个方法依然属于先在虚拟机中安装 Linux 然后再安装 Docker 的方法。

Docker Desktop 官方下载地址： https://hub.docker.com/editions/community/docker-ce-desktop-windows

### 安装

1. 下载并安装docker

2. 安装 [window subsystem for linux](/pages/window-subsystem-for-linux)

3. 安装 Hyper-V
Hyper-V 是微软开发的虚拟机，类似于 VMWare 或 VirtualBox，仅适用于 Windows 10。这是 Docker Desktop for Windows 所使用的虚拟机。

但是，这个虚拟机一旦启用，QEMU、VirtualBox 或 VMWare Workstation 15 及以下版本将无法使用！如果你必须在电脑上使用其他虚拟机（例如开发 Android 应用必须使用的模拟器），请不要使用 Hyper-V！

4. 开启 Hyper-V

### 设置代理 

settings -> Docker Engine

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

### 参考资料

- [Windows Docker 安装](https://www.runoob.com/docker/windows-docker-install.html)