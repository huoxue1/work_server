package router

import (
	"github.com/gin-gonic/gin"

	"work_server/controller"
)

func adminRouter(group *gin.RouterGroup) {
	group.POST("/create_work", controller.CreateWork())

	group.GET("/get_zip_result/:work_id", controller.GetZipResult())
}
