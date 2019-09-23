## 项目介绍

这是一个 Jaeger 链路追踪的 Demo，里面包括 5 个 Service 端，如图所示：

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_1.png)

API 端为： [go-gin-api](https://github.com/xinliangnote/go-gin-api)。

5 个 Service 端 Demo 分别是：

#### 听（listen）

- 端口：9901
- 通讯：gRPC

#### 说（speak）

- 端口：9902
- 通讯：gRPC

#### 读（read）

- 端口：9903
- 通讯：gRPC

#### 写（write）

- 端口：9904
- 通讯：gRPC

#### 唱（sing）

- 端口：9905
- 通讯：HTTP

其中服务之间又相互调用：

- Speak 服务，又调用了 Listen 服务 和 Sing 服务。
- Read 服务，又调用了 Listen 服务 和 Sing 服务。
- Write 服务，又调用了 Listen 服务 和 Sing 服务。

咱们要实现就是 API 调用 5 个服务的链路，以及服务与服务之间相互调用的链路。

## 运行

#### 1、部署 jaeger 服务

下载地址：https://www.jaegertracing.io/download/

我的电脑是 macOS 选择 -> Binaries -> macOS

下载后并解压，会发现以下文件：

- example-hotrod
- jaeger-agent
- jaeger-all-in-one
- jaeger-collector
- jaeger-ingester
- jaeger-query


进入到解压后的目录执行：

```
./jaeger-all-in-one
```

目测启动后，访问地址：http://localhost:16686/

看到下图，表示启动成功。

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_4.png)

#### 2、启动 Service 服务

```
// 启动 Listen 服务
cd listen && go run main.go

// 启动 Speak 服务
cd speak && go run main.go

// 启动 Read 服务
cd read && go run main.go

// 启动 Write 服务
cd write && go run main.go

// 启动 Sing 服务
cd sing && go run main.go
```

#### 3、启动 API 服务

```
// 启用 go-gin-api 服务
cd go-gin-api && go run main.go
```

#### 4、访问路由

访问 API 项目：http://127.0.0.1:9999/jaeger_test

## 效果

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_2.png)

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_3.png)

## 学习交流

:star2: 关注微信公众号「新亮笔记」

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)
