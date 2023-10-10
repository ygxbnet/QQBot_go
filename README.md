# QQBot_go

> 正如简介所说，这是一个好玩的QQ机器人

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/YGXB-net/QQBot_go)](./go.mod) [![GitHub](https://img.shields.io/github/license/YGXB-net/QQBot_go)](./LICENSE) [![GitHub Workflow Status (with branch)](https://img.shields.io/github/actions/workflow/status/YGXB-net/QQBot_go/build_docker_image.yml?branch=dev)](https://github.com/YGXB-net/QQBot_go/actions)

## 介绍

这个项目主要是我自己在开发的QQ机器人，并且也是自己在用，纯粹是好玩的，代码水平较低，还请见谅

<!--
如果你对项目感兴趣或者想交流讨论，也非常欢迎

E-mail: [me@ygxb.net](mailto:me@ygxb.net)

QQ: 3040809965
-->

## 使用方法

### 使用 Docker（推荐）

1. 安装 Docker

2. 编写 `docker-compose.yml`

   ```yaml
   services:
     qqbot-go:
       image: ygxb/qqbot-go
       restart: always
       volumes:
         - ./qqbot-go_data:/data
   ```

3. 启动 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp)

4. 运行 `docker compose up` 启动本程序

5. 根据提示修改 `config.yml` 文件

6. 再次运行，畅(tiao)玩(xi)机器人吧

### 本地启动

1. 将代码克隆到本地

2. 运行:（请先启动 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) 然后再运行此程序）

   ```shell
   go mod download
   go run
   ```

3. 根据提示修改 `config.yml` 文件

4. 再次运行，畅(tiao)玩(xi)机器人吧

## 特别说明

**本项目还处于开发阶段，仅个人自用，代码简陋及相关教程不完善还请见谅**

如果发现项目有问题，欢迎提交 [Issues](https://github.com/ygxbnet/QQBot_go/issues)，我会尽量解决

## 相关接口

- [https://api.aa1.cn/doc/chatgpts.html（ChatGPT免费API）](https://api.aa1.cn/doc/chatgpts.html)
- [https://api.aa1.cn/doc/pyq.html（朋友圈一言免费API）](https://api.aa1.cn/doc/pyq.html)
