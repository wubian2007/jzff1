# 桔子分发

基于 Go 和 Gin 框架开发的应用分发平台前端项目。

## 技术栈

- **后端**: Go 1.21+
- **框架**: Gin
- **模板**: Go html/template
- **前端**: HTML5, CSS3, JavaScript (ES6+)
- **容器化**: Docker & Docker Compose

## 快速开始

### 使用 Docker Compose（推荐）

```bash
# 构建并启动容器
docker-compose up -d --build

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### 本地开发

```bash
# 安装依赖
go mod download

# 运行项目
go run main.go
```

### 访问

打开浏览器访问: http://localhost:8080

## 项目结构

```
桔子分发/
├── main.go              # 主程序入口
├── go.mod               # Go 模块依赖
├── Dockerfile           # Docker 镜像构建文件
├── docker-compose.yml   # Docker Compose 配置
├── static/              # 静态资源
│   ├── css/            # 样式文件
│   ├── js/             # JavaScript 文件
│   └── images/         # 图片资源
└── templates/          # HTML 模板
    ├── layout.html     # 基础布局
    └── index.html      # 首页
```

## 开发说明

1. 从 Figma 导出图片资源到 `static/images/` 目录
2. 根据 Figma 设计稿调整 `static/css/style.css` 中的样式
3. 根据设计稿内容更新 `templates/index.html` 中的 HTML 结构
4. 添加交互功能到 `static/js/main.js`

## Docker 命令

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 停止服务
docker-compose stop

# 重启服务
docker-compose restart

# 查看日志
docker-compose logs -f web

# 进入容器
docker-compose exec web sh
```

## 部署

### 编译

```bash
go build -o juzi-fenfa
```

### 运行

```bash
./juzi-fenfa
```

## License

MIT

