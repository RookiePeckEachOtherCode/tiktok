# Tiktok

## 部署运行

- 直接部署运行:

  - 导入`configs/tiktok.sql`
  - 修改`configs/config.go`
  - `go build -o tiktok ./main.go && ./tiktok`

- 使用docker部署

```sh
docker-compose up -d
```
