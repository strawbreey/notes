---
title: "Docker Mirror"
date: 2020-12-17T14:49:12+08:00
draft: false
---

## 问题背景

Docker hub自2020年11月开始，对匿名和免费账户增加了pull镜像频率限制

### Docker Hub限频计划

[https://www.docker.com/increase-rate-limits](https://www.docker.com/increase-rate-limits)

Docker Hub将限制匿名和免费账户的docker pull频率，付费账户的pull请求则不受影响

具体pull频率限制

匿名用户：100次/6小时
免费用户：200次/6小时

### 判断逻辑

- 如果没有docker login，则为匿名，按客户端IP判别
- 如果docker login了（使用用户在Docker Hub上注册的账户），则按账户判别
    - 如同一IP既有匿名pull请求，也有已登录的pull请求，是分别判断pull频率的
- 按docker pull次数计数，一次pull中即使有多个layer的下载，也仅计为一次

### 解决方案
可通过软件源提供的docker代理 `https://mirror.baidubce.com` 拉取镜像

配置方法

配置registry-mirrors进行加速并绕过限制
```json
{
  "registry-mirrors": [
    "https://mirror.baidubce.com"
  ]
}
```
各操作系统配置方法详见 [https://yeasy.gitbook.io/docker_practice/install/mirror](https://yeasy.gitbook.io/docker_practice/install/mirror)

devcloud机器需要将dockerhub.woa.com域名加入系统和docker服务no_proxy列表中，可能需要修改如下文件的代理配置：

```
/etc/systemd/system/docker.service.d/http-proxy.conf
/etc/profile
```
### 验证

如果能成功拉取hello-world镜像，说明配置生效。

```bash
# 拉取镜像
docker pull haifangwang/hello-world

# Using default tag: latest
# latest: Pulling from haifangwang/hello-world
# Digest: sha256:71c1e285dd1ad9a509fe005fdcddee80b5ba59bfae03cbf5f9f77af1723a4edd
# Status: Image is up to date for haifangwang/hello-world:latest
```