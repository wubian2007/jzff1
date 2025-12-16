# 桔子分发 - 纯静态版本

纯静态 HTML 网站，无需 Go 环境即可部署。

## 快速开始

### 本地预览

直接用浏览器打开 `index.html` 文件即可。

### 部署

将整个项目文件夹上传到任何静态网站托管服务：
- GitHub Pages
- Netlify
- Vercel
- 阿里云 OSS
- 腾讯云 COS
- Nginx/Apache 服务器

## 修改跳转地址

**重要：所有按钮和链接的跳转地址都在 `index.html` 文件顶部配置**

打开 `index.html` 文件，找到以下代码（大约在第 12-16 行）：

```html
<script>
    // 目标站点地址 - 修改这里即可更新所有链接
    const TARGET_SITE = 'https://www.fenfa999.com';
</script>
```

**只需要修改 `TARGET_SITE` 变量的值，所有按钮和链接都会自动跳转到新地址。**

## 项目结构

```
桔子分发/
├── index.html              # 主页面（纯静态 HTML）
├── static/                 # 静态资源
│   ├── css/
│   │   └── style.css      # 样式文件
│   ├── js/
│   │   └── main.js        # JavaScript 文件
│   └── images/
│       └── hero-icon.png  # 图片资源
└── README.md              # 说明文档
```

## 功能特点

- ✅ 纯静态 HTML，无需服务器环境
- ✅ 可配置跳转地址（只需修改一个变量）
- ✅ 响应式设计，支持移动端
- ✅ 所有链接统一跳转到配置的目标站点

## 部署到 GitHub Pages

1. 将代码推送到 GitHub 仓库
2. 在仓库设置中启用 GitHub Pages
3. 选择 `main` 分支和 `/` 根目录
4. 访问 `https://你的用户名.github.io/jzff1/`

## License

MIT
