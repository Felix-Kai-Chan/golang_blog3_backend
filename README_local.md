# Golang_Blog3 API

一个基于 Go 语言、Gin 框架和 GORM 库开发的轻量级个人博客系统后端 API。该项目实现了用户认证、文章管理、评论系统以及数据库迁移等核心功能。

## 功能特性

- **用户系统**：支持用户注册、登录，密码使用 `bcrypt` 加密存储。
- **认证授权**：基于 JWT 的身份验证机制，保护私有路由。
- **文章管理**：支持文章的发布、获取详情、更新和删除（仅作者可操作）。
- **评论系统**：支持对文章进行评论。
- **基础设施**：
    - 集成 `Logrus` 进行日志记录。
    - 使用 GORM 进行数据库操作及自动迁移。
    - 统一 JSON 响应格式。

## 技术栈

- **语言**: Go (Golang)
- **Web 框架**: Gin
- **ORM 库**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **日志**: Logrus
- **工具**: Viper (配置管理), bcrypt (密码加密)

## 项目结构

```text
golang_blog3/
├── cmd/
│   └── server/
│       └── main.go          # 程序入口，初始化路由和服务
├── config/
│   ├── db.go                # 数据库连接配置
│   └── logger.go            # 日志配置
├── controllers/             # 业务逻辑处理
│   ├── auth.go              # 认证逻辑
│   ├── comment.go           # 评论逻辑
│   ├── post.go              # 文章逻辑
│   └── user.go              # 用户逻辑
├── middleware/              # 中间件
│   ├── auth.go              # JWT 验证中间件
│   └── jwt.go               # JWT 生成与解析工具
├── migrate/                 # 数据库迁移脚本
├── models/                  # 数据库模型定义
│   ├── comment.go
│   ├── hooks.go             # GORM 钩子（如密码加密）
│   ├── post.go
│   └── user.go
├── routes/                  # 路由定义
│   ├── comment_router.go
│   ├── post_routes.go
│   ├── router.go            # 总路由入口
│   └── user_routes.go
├── utils/                   # 工具类
│   └── jwt.go
├── .env                     # 环境变量配置
├── .env.example             # 环境变量示例
├── go.mod
└── go.sum

