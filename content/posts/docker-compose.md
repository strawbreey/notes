---
title: "Docker Compose"
date: 2020-12-17T16:00:03+08:00
draft: false
---

### Compose 简介

Compose 是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，您可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

### Docker Compose 常用命令和配置

```bash
# 1. ps：列出所有运行容器
docker-compose ps

# 2. logs：查看服务日志输出
docker-compose logs

# 3. port：打印绑定的公共端口，下面命令可以输出 eureka 服务 8761 端口所绑定的公共端口
docker-compose port eureka 8761

# 4. build：构建或者重新构建服务
docker-compose build

# 5. start：启动指定服务已存在的容器
docker-compose start eureka

# 6. stop：停止已运行的服务的容器
docker-compose stop eureka

# 7. rm：删除指定服务的容器
docker-compose rm eureka

# 8. up：构建、启动容器
docker-compose up
 
# 9. kill：通过发送 SIGKILL 信号来停止指定服务的容器
docker-compose kill eureka

# 10. pull：下载服务镜像

# 11 scale：设置指定服务运气容器的个数，以 service=num 形式指定
docker-compose scale user=3 movie=3

#  12. run：在一个服务上执行一个命令
docker-compose run web bash

```



### Compose 使用的三个步骤：

1. 准备项目

```bash
mkdir composetest
cd composetest
```

```js
import time

import redis
from flask import Flask

app = Flask(__name__)
cache = redis.Redis(host='redis', port=6379)


def get_hit_count():
    retries = 5
    while True:
        try:
            return cache.incr('hits')
        except redis.exceptions.ConnectionError as exc:
            if retries == 0:
                raise exc
            retries -= 1
            time.sleep(0.5)


@app.route('/')
def hello():
    count = get_hit_count()
    return 'Hello World! I have been seen {} times.\n'.format(count)
```

2. 使用 Dockerfile 定义应用程序的环境。

```dockerfile
# FROM python:3.7-alpine: 从 Python 3.7 映像开始构建镜像。
FROM python:3.7-alpine
# 将工作目录设置为 /code。
WORKDIR /code
# RUN apk add --no-cache gcc musl-dev linux-headers: 安装 gcc，以便诸如 MarkupSafe 和 SQLAlchemy 之类的 Python 包可以编译加
ENV FLASK_APP app.py
ENV FLASK_RUN_HOST 0.0.0.0
RUN apk add --no-cache gcc musl-dev linux-headers
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
# 将 . 项目中的当前目录复制到 . 镜像中的工作目录。
COPY . .
# CMD ["flask", "run"]: 容器提供默认的执行命令为：flask run。
CMD ["flask", "run"]
```
2.  使用 docker-compose.yml 定义构成应用程序的服务，这样它们可以在隔离环境中一起运行。

```yml
version: '3'
services:
  web:
    build: .
    ports:
     - "5000:5000"
  redis:
    image: "redis:alpine"
```

该 Compose 文件定义了两个服务：web 和 redis。

- web：该 web 服务使用从 Dockerfile 当前目录中构建的镜像。然后，它将容器和主机绑定到暴露的端口 5000。此示例服务使用 Flask Web 服务器的默认端口 5000 。
- redis：该 redis 服务使用 Docker Hub 的公共 Redis 映像。

4. 最后，执行 docker-compose up 命令来启动并运行整个应用程序。

```bash
# 在测试目录中，执行以下命令来启动应用程序：
docker-compose up

# 如果你想在后台执行该服务可以加上 -d 参数：
docker-compose up -d
```

### yml 配置指令参考

```yaml
# version 指定本 yml 依从的 compose 哪个版本制定的。
version: "3.7"
# 多个容器集合
services:
  webapp:
    # build: 指定为构建镜像上下文路径：例如 webapp 服务，指定为从上下文路径 ./dir/Dockerfile 所构建的镜像：
    build:
      # context: 上下文路径。
      context: ./dir
      # dockerfile: 指定构建镜像的 Dockerfile 文件名。
      dockerfile: Dockerfile-alternate
      # args: 添加构建参数，这是只能在构建过程中访问的环境变量。
      args:
        buildno: 1
    # 设置构建镜像的标签。

    # command：覆盖容器启动后默认执行的命令
    command: bundle exec thin -p 3000
    labels:
      - "com.example.description=Accounting webapp"
      - "com.example.department=Finance"
      - "com.example.label-with-empty-value"
    # 多层构建，可以指定构建哪一层
    target: prod
    # 设置依赖关系。
      # docker-compose up ：以依赖性顺序启动服务。在以下示例中，先启动 db 和 redis ，才会启动 web。
      # docker-compose up SERVICE ：自动包含 SERVICE 的依赖项。在以下示例中，docker-compose up web 还将创建并启动 db 和 redis。
      # docker-compose stop ：按依赖关系顺序停止服务。在以下示例中，web 在 db 和 redis 之前停止。
    depends_on:
      - db
      - redis
  redis:
    # image: 指定容器运行的镜像。以下格式都可以：
    image: redis:alpine
    # devices: 指定设备映射列表。
    devices:  
      - "/dev/ttyUSB0:/dev/ttyUSB0"
    # dns: 自定义 DNS 服务器，可以是单个值或列表的多个值。
    dns: 8.8.8.8
    # dns_search:自定义 DNS 搜索域。可以是单个值或列表。
    dns_search:  example.com
    # entrypoint: 覆盖容器默认的 entrypoint。
    entrypoint: /code/entrypoint.sh
    # env_file: 从文件添加环境变量。可以是单个值或列表的多个值。
    env_file:
      - ./common.env
      - ./apps/web.env
      - /opt/secrets.env
    # environment: 添加环境变量。您可以使用数组或字典、任何布尔值，布尔值需要用引号引起来，以确保 YML 解析器不会将其转换为 True 或 False。
    environment: 
    # expose: 暴露端口，但不映射到宿主机，只被连接的服务访问。
    expose:
      - "3000"
      - "8000" 
    # extra_hosts： 添加主机名映射。类似 docker client --add-host。
    extra_hosts: 
      - "somehost:162.242.195.82"
      - "otherhost:50.31.209.229"
    # healthcheck: 用于检测 docker 服务是否健康运行。
    healthcheck: 
      test: ["CMD", "curl", "-f", "http://localhost"] # 设置检测程序
      interval: 1m30s # 设置检测间隔
      timeout: 10s # 设置检测超时时间
      retries: 3 # 设置重试次数
      start_period: 40s # 启动后，多少秒开始启动检测程序
    # logging 服务的日志记录配置。
    logging:
      driver: json-file
      options:
        max-size: "200k" # 单个文件大小为200k
        max-file: "10" # 最多10个文件
    # 设置网络模式。
    # network_mode: "bridge"
    # network_mode: "host"
    # network_mode: "none"
    # network_mode: "service:[service name]"
    # network_mode: "container:[container name/id]"
    network_mode: "bridge"

    # restart: o：是默认的重启策略，在任何情况下都不会重启容器。always：容器总是重新启动。on-failure：在容器非正常退出时（退出状态非0），才会重启容器。unless-stopped：在容器退出时总是重启容器，但是不考虑在Docker守护进程启动时就已经停止了的容器
    restart: unless-stopped

    # 存储敏感数据，例如密码：
    secrets:
      - my_secret

    # 修改容器默认的 schema 标签。
    security-opt:
      - label:user:USER   # 设置容器的用户标签
      - label:role:ROLE   # 设置容器的角色标签
      - label:type:TYPE   # 设置容器的安全策略标签
      - label:level:LEVEL  # 设置容器的安全等级标签

    # 指定在容器无法处理 SIGTERM (或者任何 stop_signal 的信号)，等待多久后发送 SIGKILL 信号关闭容器。
    stop_grace_period: 1s

    # stop_signal
    stop_signal: SIGUSR1

    # 设置容器中的内核参数，可以使用数组或字典格式。
    sysctls:
      - net.core.somaxconn=1024
      - net.ipv4.tcp_syncookies=0

    # 在容器内安装一个临时文件系统。可以是单个值或列表的多个值。
    tmpfs:
      - /run
      - /tmp

    # 覆盖容器默认的 ulimit。
    ulimits:
      nproc: 65535
      nofile:
        soft: 20000
        hard: 40000
    # 将主机的数据卷或着文件挂载到容器里。
    volumes: 
      - "/localhost/postgres.sock:/var/run/postgres/postgres.sock"
      - "/localhost/data:/var/lib/postgresql/data"
    # 指定与服务的部署和运行有关的配置。只在 swarm 模式下才会有用。
    deploy:
      # mode: 指定服务提供的模式。 1. replicated：复制服务，复制指定服务到集群的机器上。 2. global：全局服务，服务将部署至集群的每个节点。
      mode: replicated
      replicas: 6
      # endpoint_mode: 访问集群服务的方式 1. vip: Docker 集群服务一个对外的虚拟 ip。所有的请求都会通过这个虚拟 ip 到达集群服务内部的机器。2. dnsrr DNS 轮询（DNSRR）。所有的请求会自动轮询获取到集群 ip 列表中的一个 ip 地址。
      endpoint_mode: dnsrr

      #labels 在服务上设置标签。可以用容器上的 labels（跟 deploy 同级的配置） 覆盖 deploy 下的 labels。
      labels: 
        description: "This redis service label"
      # 配置服务器资源使用的限制，例如上例子，配置 redis 集群运行需要的 cpu 的百分比 和 内存的占用。避免占用资源过高出现异常
      resources:
        limits:
          cpus: '0.50'
          memory: 50M
        reservations:
          cpus: '0.25'
          memory: 20M
      # 配置如何在退出容器时重新启动容器
      restart_policy:
        # condition：可选 none，on-failure 或者 any（默认值：any）。
        condition: on-failure
        # delay：设置多久之后重启（默认值：0）。
        delay: 5s
        # max_attempts：尝试重新启动容器的次数，超出次数，则不再尝试（默认值：一直重试）。
        max_attempts: 3
        # window：设置容器重启超时时间（默认值：0）。
        window: 120s
      # 配置在更新失败的情况下应如何回滚服务。
      rollback_config: 
        # parallelism: 一次要回滚的容器数。如果设置为0，则所有容器将同时回滚。
        parallelism: 0
        # delay: 每个容器组回滚之间等待的时间（默认为0s）。
        delay: 0
        # failure_action: 如果回滚失败，该怎么办。其中一个 continue 或者 pause（默认pause）。
        failure_action: continue
        # monitor: 每个容器更新后，持续观察是否失败了的时间 (ns|us|ms|s|m|h)（默认为0s）。
        monitor: 0
        # max_failure_ratio: 在回滚期间可以容忍的故障率（默认为0）。
        max_failure_ratio: 0
        # order: 回滚期间的操作顺序。其中一个 stop-first（串行回滚），或者 start-first（并行回滚）（默认 stop-first ）。
        order: start-first
      # 配置应如何更新服务，对于配置滚动更新很有用。
      update_config:
        # parallelism： 一次更新的容器数。
        parallelism: 5
        # delay：在更新一组容器之间等待的时间。
        delay: 0
        # failure_action：如果更新失败，该怎么办。其中一个 continue，rollback 或者pause （默认：pause）。
        failure_action: pause
        # monitor：每个容器更新后，持续观察是否失败了的时间 (ns|us|ms|s|m|h)（默认为0s）。
        monitor: 0
        # max_failure_ratio：在更新过程中可以容忍的故障率。
        max_failure_ratio: 0
        # order回滚期间的操作顺序。其中一个 stop-first（串行回滚），或者 start-first（并行回滚）（默认stop-first）。
        order: stop-first
```

### 参考资料

- [docker-compose](https://www.runoob.com/docker/docker-compose.html)