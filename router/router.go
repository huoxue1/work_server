package router

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed view
var static embed.FS

// InitRouter
/**
 * @Description: 初始化路由
 * @return *gin.Engine
 */
func InitRouter() *gin.Engine {
	engine := gin.Default()
	// 配置跨域
	engine.Use(cors())

	engine.StaticFS("/static", http.FS(static))
	// 公共接口
	public := engine.Group("/public")
	publicRouter(public)

	// 管理员接口
	admin := engine.Group("/admin", func(context *gin.Context) {
		token := context.Query("token")
		if token == "qqewqeqdadadadd" {
			context.Next()
		} else {
			context.Abort()
		}
	})
	adminRouter(admin)

	return engine
}
