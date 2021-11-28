package router

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	TOKEN = "qqewqeqdadadadd"
)

func SetToken(token string) {
	TOKEN = token
}

//go:embed view/dist
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
	engine.GET("/", func(context *gin.Context) {
		context.Redirect(301, "/static/view/dist/home.html")
	})
	// 公共接口
	public := engine.Group("/public")
	publicRouter(public)

	// 管理员接口
	admin := engine.Group("/admin", func(context *gin.Context) {
		token := context.Query("token")
		if token == TOKEN {
			context.Next()
		} else {
			context.Abort()
		}
	})
	adminRouter(admin)

	return engine
}
