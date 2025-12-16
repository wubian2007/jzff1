package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin 模式（开发环境使用 DebugMode，生产环境使用 ReleaseMode）
	gin.SetMode(gin.DebugMode)

	// 创建 Gin 路由引擎
	r := gin.Default()

	// 加载 HTML 模板
	r.LoadHTMLGlob("templates/*")

	// 静态文件服务
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 设置路由
	setupRoutes(r)

	// 启动服务器
	port := ":8080"
	log.Printf("🚀 服务器启动在 http://localhost%s", port)
	log.Printf("📱 访问地址: http://localhost%s", port)
	
	if err := r.Run(port); err != nil {
		log.Fatal("❌ 服务器启动失败:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// 首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "桔子分发",
			"page":  "index",
		})
	})

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})
}

