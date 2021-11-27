package router

import (
	"github.com/gin-gonic/gin"

	"work_server/controller"
)

func publicRouter(group *gin.RouterGroup) {
	group.POST("/upload", controller.Upload())
	group.POST("/get_works", controller.GetWorks())
	group.POST("/get_files/:work_id", controller.GwtFiles())
	group.POST("/get_work/:work_id", controller.GetWork())
	group.POST("/get_file/:file_id", controller.GetFile())
}
