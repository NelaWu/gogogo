# Gin 框架示例项目

## 🔧项目功能亮点
| 功能             | 说明                       |
| -------------- | ------------------------ |
| ✅ 用户管理         | 实现注册、登录、获取用户信息等接口        |
| 🔐 JWT 认证      | 使用 JSON Web Token 进行身份验证 |
| 📄 Swagger 文档  | 自动生成 API 文档，便于测试和维护      |
| 🔁 RESTful API | 遵循 RESTful 设计规范          |
| 📁 路由分组        | 使用模块化路由进行项目结构划分          |
| 🧱 中间件支持       | 日志记录、JWT 校验、错误处理等        |
| 📚 GORM 集成     | 数据持久层采用 GORM 操作数据库       |
| 📦 企业级结构       | 遵循合理的项目组织和工程化规范          |

## 🧱 技术栈
* Gin：高性能 HTTP Web 框架
* GORM：ORM 框架，简化数据库操作
* JWT：用于用户身份验证的 Token 机制
* Swagger：API 文档生成工具
* Viper（可选）: 配置管理
* Zap（或 logrus）：结构化日志框架

## 測試方法

首先先 `go run main.go`

### 新增用戶
```curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"username":"alice","email":"alice@example.com"}'
```
### 查看所有用戶
`curl http://localhost:8080/users`