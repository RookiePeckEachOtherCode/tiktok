# 字节挑动青训营-Go语言实战项目 极简抖音

[本地存储版本](https://github.com/RookiePeckEachOtherCode/tiktok/main)

## 技术选型

- GORM
- GIN
- MySQL
- Redis
- Docker

## 项目结构

```
.
├── configs
│   ├── config.go
│   ├── dict.txt
│   └── tiktok.sql
├── controller
│   ├── commentController.go
│   ├── favoriteController.go
│   ├── followController.go
│   ├── messageController.go
│   ├── userController.go
│   └── videoController.go
├── dao
│   ├── commentDao.go
│   ├── InitDB.go
│   ├── messageDao.go
│   ├── response.go
│   ├── userInfoDao.go
│   ├── userLoginDao.go
│   └── videoDao.go
├── doc
├── docker-compose.yml
├── Dockerfile
├── gin.log
├── go.mod
├── go.sum
├── go_test.sh
├── LICENSE
├── main.go
├── middleware
│   ├── hash
│   │   └── sha1.go
│   ├── jwt
│   │   └── jwt.go
│   └── redis
│       └── redis.go
├── README.md
├── router
│   └── router.go
├── service
│   ├── commentService.go
│   ├── favoriteService.go
│   ├── followService.go
│   ├── messageService.go
│   ├── userService.go
│   └── videoService.go
├── test
├── tiktok.log
└── util
    ├── log
    │   └── log.go
    ├── oos
    │   └── oss.go
    └── util.go
```

## 开发设计

### 架构设计

![oss.png](https://img1.imgtp.com/2023/08/23/DfOOjLiv.png)

### 数据库设计

ER图
![2646d572-5456-4107-957f-7c3c4bfbffbe.png](https://img1.imgtp.com/2023/08/23/vxQFOzRQ.png)

### 接口文档

- [Apifox](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/)

  > Apifox上的接口文档有遗漏

- [Protobuf](./doc/接口文档.md)

## 项目部署

默认运行在8080端口

### 本地部署

1. 导入`configs/tiktok.sql`到MySQL
2. 修改`configs/config.go`中的配置
3. 运行`go build -o tiktok && ./tiktok`

### Docker部署

1. 修改`docker-compose.yml`中的配置
2. 运行`docker-compose up -d`

## 后续期望

- 使用微服务
- 使用消息队列
- 使用k8s
- .........
