# Gin Hello World Project

基于字节跳动开发规范的简单 Gin 项目

## 项目结构

```
.
├── cmd/
│   └── server/          # 主应用入口
├── internal/
│   ├── handler/         # HTTP 处理器
│   └── service/         # 业务逻辑层
├── pkg/                 # 公共库代码
├── api/                 # API 协议定义
├── configs/             # 配置文件
└── test/               # 测试文件
```

## 运行项目

```bash
go run cmd/server/main.go
```

## API 接口

- `GET /api/v1/hello` - 返回 Hello World 消息

## 响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": "Hello World"
}
```