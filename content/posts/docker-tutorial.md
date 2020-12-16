---
title: "Docker Tutorial"
date: 2020-12-10T10:51:20+08:00
draft: false
---

Docker 是一个开源的应用容器引擎，基于 Go 语言 并遵从 Apache2.0 协议开源。

Docker 可以让开发者打包他们的应用以及依赖包到一个轻量级、可移植的容器中，然后发布到任何流行的 Linux 机器上，也可以实现虚拟化。

容器是完全使用沙箱机制，相互之间不会有任何接口（类似 iPhone 的 app）,更重要的是容器性能开销极低。

## Docker

### Docker的应用场景

- Web 应用的自动化打包和发布。

- 自动化测试和持续集成、发布。

- 在服务型环境中部署和调整数据库或其他的后台应用。

- 从头编译或者扩展现有的 OpenShift 或 Cloud Foundry 平台来搭建自己的 PaaS 环境。

### Docker 能做什么

- 创建应用依赖
- 创建应用镜像并进行复制
- 创建容易分发的应用
- 允许实例简单，快速的扩展

### Docker 镜像加速

国内从 DockerHub 拉取镜像有时会遇到困难，此时可以配置镜像加速器。Docker 官方和国内很多云服务商都提供了国内加速器服务，例如：

- 网易：https://hub-mirror.c.163.com/
- 阿里云：https://<你的ID>.mirror.aliyuncs.com
- 七牛云加速器：https://reg-mirror.qiniu.com

当配置某一个加速器地址之后，若发现拉取不到镜像，请切换到另一个加速器地址。国内各大云服务商均提供了 Docker 镜像加速服务，建议根据运行 Docker 的云平台选择对应的镜像加速服务。


### docker 常用命令

```bash
# 安装docker

# 查看docker 信息
docker version
# 或者
docker info

# 查看docker 版本
docker --version

# 开启docker 服务
# service 命令的用法
sudo service docker start

# systemctl 命令的用法
sudo systemctl start docker

# 拉取镜像
docker pull ubuntu:latest
# 等于
docker pull ubuntu

# 查看镜像
docker images

# 查看所有镜像
docker images ls

# 查看所有镜像, 包括终止运行的容器
docker container ls --all

# 删除镜像
docker image rm [imageName]

# 从镜像上创建容器
docker run --rm -ti ubuntu /bin/bash
docker run -d ubuntu ping 8.8.8.8

# 运行hello-world容器
docker container run hello-world

# 检查容器是否运行
docker ps
docker exec -ti 
```

简要说明:

1. --rm: 告诉docker，一旦运行进程推出就删除容器
2. -ti 告诉docker分配一个伪终端并进入交互模式


### 制作自己的Docker容器

```bash

# 创建dockerfile 文件

# 更具dockerfile生成镜像
docker image build -t koa-demo
# 或者
docker image build -t koa-demo:0.0.1 .

# 查看文件
docker image ls

# 生成容器
docker container run -p 8000:3000 -it koa-demo /bin/bash
# 或者
docker container run -p 8000:3000 -it koa-demo:0.0.1 /bin/bash
```
注意:

- -t参数用来指定 image 文件的名字，后面还可以用冒号指定标签。如果不指定，默认的标签就是latest。最后的那个点表示 Dockerfile 文件所在的路径，上例是当前路径，所以是一个点。

- -p参数：容器的 3000 端口映射到本机的 8000 端口。
- -it参数：容器的 Shell 映射到当前的 Shell，然后你在本机窗口输入的命令，就会传入容器。
- koa-demo:0.0.1：image 文件的名字（如果有标签，还需要提供标签，默认是 latest 标签）。
- /bin/bash：容器启动以后，内部第一个执行的命令。这里是启动 Bash，保证用户可以使用 Shell。


### 参考资料

- [Docker 入门教程](https://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html)
- [Docker 微服务教程](http://www.ruanyifeng.com/blog/2018/02/docker-wordpress-tutorial.html)