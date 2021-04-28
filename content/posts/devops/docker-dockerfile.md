---
title: "Docker Dockerfile"
date: 2020-12-17T16:00:13+08:00
draft: false
---

Docker can build images automatically by reading the instructions from a Dockerfile. A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. Using docker build users can create an automated build that executes several command-line instructions in succession.

FROM 

FROM [--platform=<platform>] <image> [AS <name>]
FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]
FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]


RUN

RUN <command>
RUN ["executable", "param1", "param2"]


### 参考资料

- [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)