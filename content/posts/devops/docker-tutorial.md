---
title: "Docker Tutorial"
date: 2020-12-10T10:51:20+08:00
draft: false
---

Docker 是一个开源的应用容器引擎，基于 Go 语言 并遵从 Apache2.0 协议开源。

Docker 可以让开发者打包他们的应用以及依赖包到一个轻量级、可移植的容器中，然后发布到任何流行的 Linux 机器上，也可以实现虚拟化。

容器是完全使用沙箱机制，相互之间不会有任何接口（类似 iPhone 的 app）,更重要的是容器性能开销极低。


### Docker的应用场景

- Web 应用的自动化打包和发布。

- 自动化测试和持续集成、发布。

- 在服务型环境中部署和调整数据库或其他的后台应用。

- 从头编译或者扩展现有的 OpenShift 或 Cloud Foundry 平台来搭建自己的 PaaS 环境。

### Docker架构

- Docker daemon（ Docker守护进程）

Docker daemon是一个运行在宿主机（ DOCKER-HOST）的后台进程。可通过 Docker客户端与之通信。

- Client（ Docker客户端）

Docker客户端是 Docker的用户界面，它可以接受用户命令和配置标识，并与 Docker daemon通信。图中， docker build等都是 Docker的相关命令。

- Images（ Docker镜像）

Docker镜像是一个只读模板，它包含创建 Docker容器的说明。它和系统安装光盘有点像，使用系统安装光盘可以安装系统，同理，使用Docker镜像可以运行 Docker镜像中的程序。

- Container（容器）

容器是镜像的可运行实例。镜像和容器的关系有点类似于面向对象中，类和对象的关系。可通过 Docker API或者 CLI命令来启停、移动、删除容器。

- Registry

Docker Registry是一个集中存储与分发镜像的服务。构建完 Docker镜像后，就可在当前宿主机上运行。但如果想要在其他机器上运行这个镜像，就需要手动复制。此时可借助 Docker Registry来避免镜像的手动复制。

### Docker 能做什么

- 创建应用依赖
- 创建应用镜像并进行复制
- 创建容易分发的应用
- 允许实例简单，快速的扩展

### Docker 镜像加速

国内从 DockerHub 拉取镜像有时会遇到困难，此时可以配置镜像加速器。Docker 官方和国内很多云服务商都提供了国内加速器服务，例如：

- 网易：https://hub-mirror.c.163.com/
- 阿里云：https://*.mirror.aliyuncs.com
- 七牛云加速器：https://reg-mirror.qiniu.com

当配置某一个加速器地址之后，若发现拉取不到镜像，请切换到另一个加速器地址。国内各大云服务商均提供了 Docker 镜像加速服务，建议根据运行 Docker 的云平台选择对应的镜像加速服务。


### docker 常用命令

#### 1. 安装docker

```bash

# 查看内核版本
uname -r

# 更新yum包
yum -y update

# 卸载旧版本
yum remove docker docker-common docker-selinux docker-engine

# 安装需要的软件包
yum install -y yum-utils device-mapper-persistent-data lvm2

# 设置 yum 源
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# 可以查看仓库中所有docker版本，并选择特定版本安装
yum list docker-ce --showduplicates | sort -r 

# 安装docker
yum install -y docker-ce

# 查看docker 信息
docker version

# 或者
docker info

# 查看docker 版本
docker --version

# 开启docker 服务

# 启动并加入开机启动 ， systemctl 命令的用法
systemctl start docker
systemctl enable docker

```

#### 2. 镜像相关
```bash
# 1. 搜索镜像
docker search java

# NAME:镜像仓库名称。
# DESCRIPTION:镜像仓库描述。
# STARS：镜像仓库收藏数，表示该镜像仓库的受欢迎程度，类似于 GitHub的 stars0
# OFFICAL:表示是否为官方仓库，该列标记为[0K]的镜像均由各软件的官方项目组创建和维护。
# AUTOMATED：表示是否是自动构建的镜像仓库。


# 2. 下载镜像
docker pull java

docker pull ubuntu:latest # 等于
docker pull ubuntu 


# 3. 查看镜像
docker image
docker images ls # 查看所有镜像

# REPOSITORY：镜像所属仓库名称。
# TAG:镜像标签。默认是 latest,表示最新。
# IMAGE ID：镜像 ID，表示镜像唯一标识。
# CREATED：镜像创建时间。
# SIZE: 镜像大小。

# 4. 删除镜像
docker rmi [imageName] # 旧
docker image rm [imageName] # 新
```


#### 3. 容器相关
```bash
# 1. 新建并启动容器
docker run -d -p 91:80 nginx

docker run --rm -ti ubuntu /bin/bash
docker run -d ubuntu ping 8.8.8.8

docker container run hello-world

# -d 后台运行
# -p 宿主机端口:容器端口 => 开放容器端口到宿主机端口
# 使用 docker run命令创建容器时，会先检查本地是否存在指定镜像。如果本地不存在该名称的镜像， Docker就会自动从 Docker Hub下载镜像并启动一个 Docker容器。

# 2. 查看容器
docker ps # 旧
docker container ls --all # 新

# CONTAINER_ID：表示容器 ID。
# IMAGE:表示镜像名称。
# COMMAND：表示启动容器时运行的命令。
# CREATED：表示容器的创建时间。
# STATUS：表示容器运行的状态。UP表示运行中， Exited表示已停止。
# PORTS:表示容器对外的端口号。
# NAMES:表示容器名称。该名称默认由 Docker自动生成，也可使用 docker run命令的--name选项自行指定。

# 3. 停止容器
docker stop f0b1c8ab3633
docker kill f0b1c8ab3633 #强制停止容器

# 4. 重启容器
docker start f0b1c8ab3633

# 5. 查看容器所有信息
docker inspect f0b1c8ab3633

# 7. 查看容器日志
docker container logs f0b1c8ab3633

# 8. 查看容器进程
docker top f0b1c8ab3633

# 9. 进入容器
docker container exec -it f0b1c8ab3633 /bin/bash

# 10. 删除容器
docker rm f0b1c8ab3633
```
简要说明:

1. --rm: 告诉docker，一旦运行进程推出就删除容器
2. -ti 告诉docker分配一个伪终端并进入交互模式

#### 4.  制作自己的Docker容器

```bash

# 创建dockerfile 文件

# 更具dockerfile生成镜像
docker image build -t koa-demo
# 或者
docker image build -t koa-demo:0.0.1 .

# 查看文件
docker image ls

# 生成并执行容器
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