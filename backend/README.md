# SF6 一灯学堂后端 (Go + MySQL)

轻量级 REST API，支撑前端用户端/管理端。固定管理员账号、基础学员登录、赛事/学员/成长路线/广告点击数据。

## 快速开始

```powershell
# 1) 安装依赖
cd backend

# 2) 运行开发（读取 .env）
$env:PORT=8080; go run ./cmd/api
```

## 目录
- `cmd/api` 主入口
- `internal/handlers` HTTP 处理
- `internal/models` 数据结构
- `internal/store` 数据访问 (MySQL)
- `internal/services` 业务逻辑
- `db/migrations` SQL 初始化

## 环境变量
- `PORT` (可选) 默认 8080
- `DATABASE_URL` (必填) 例如 `user:pass@tcp(127.0.0.1:3306)/sf6?parseTime=true&charset=utf8mb4`
- `ADMIN_PASSWORD` (可选) 默认 `feigeA.5200....`

## 待实现
- 学员注册/登录/套餐权限校验（角色数/冷却时间）
- 成长路线 CRUD + 冷却控制
- 赛事/公告/分组/奖项 CRUD
- 广告点击统计
- 学员处罚/冻结
- 管理端统计聚合

保持接口简单、轻量，满足前端所需最小集合。

## 运行内存版演示

无需数据库，直接使用内存存储（演示数据）:

```powershell
cd backend
$env:PORT=8080; go run ./cmd/api
```

## 运行 MySQL 版本（待实现）

1. 创建数据库并导入迁移：
```powershell
mysql -uuser -p -h 127.0.0.1 -P 3306 -e "CREATE DATABASE IF NOT EXISTS sf6 CHARACTER SET utf8mb4;"
mysql -uuser -p sf6 < db/migrations/001_init.sql
```
2. 配置环境变量：复制 `.env.example` 为 `.env`，填写 `DATABASE_URL`。
3. 替换存储实现：后续可将 `NewMemoryStore()` 换成 MySQL 版（待补齐）。
