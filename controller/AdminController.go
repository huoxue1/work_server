package controller

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"work_server/dao"
	"work_server/pojo"
	"work_server/util"
)

func DeleteWork() gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("work_id")
		workId, _ := strconv.Atoi(param)
		db := dao.GetDB()
		session := db.NewSession()
		defer session.Close()
		work := new(pojo.Work)
		_, err := session.ID(workId).Get(work)
		if err != nil {
			session.Rollback()
			context.JSON(403, response{
				Code:  400,
				Msg:   "",
				Data:  nil,
				Error: err.Error(),
			})
			return
		}
		_, err = session.Delete(&pojo.Files{WorkID: workId})
		if err != nil {
			session.Rollback()
			context.JSON(403, response{
				Code:  400,
				Msg:   "",
				Data:  nil,
				Error: err.Error(),
			})
			return
		}
		_, err = session.Delete(&pojo.Work{Id: int64(workId)})
		if err != nil {
			session.Rollback()
			context.JSON(403, response{
				Code:  400,
				Msg:   "",
				Data:  nil,
				Error: err.Error(),
			})
			return
		}
		dir, err := os.ReadDir("./work/" + work.Name + "/")
		for _, d := range dir {
			_ = os.RemoveAll(path.Join([]string{"tmp", work.Name, d.Name()}...))
		}
		err = session.Commit()
		if err != nil {
			session.Rollback()
			context.JSON(403, response{
				Code:  400,
				Msg:   "",
				Data:  nil,
				Error: err.Error(),
			})
			return
		}
		context.JSON(200, ok(200, nil))
	}
}

// RemoveFile
/**
 * @Description: 删除某个文件
 * @return gin.HandlerFunc
 */
func RemoveFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("file_id")
		token := context.Query("token")
		fileId, _ := strconv.Atoi(param)
		db := dao.GetDB()
		session := db.NewSession()
		defer session.Close()
		file := new(pojo.Files)
		_, err := session.ID(fileId).Get(file)
		if err != nil {
			session.Rollback()
			return
		}
		if token != file.Token && token != TOKEN {
			context.JSON(404, response{
				Code:  404,
				Msg:   "token check failed",
				Data:  nil,
				Error: "token check failed",
			})
			return
		}
		_, err = session.ID(fileId).Delete(file)
		if err != nil {
			session.Rollback()
			return
		}
		work := new(pojo.Work)
		_, err = db.ID(file.WorkID).Get(work)
		if err != nil {
			session.Rollback()
			return
		}

		err = os.RemoveAll("./work/" + work.Name + "/" + file.FileName)
		if err != nil {
			session.Rollback()
			return
		}
		err = session.Commit()
		if err != nil {
			session.Rollback()
			return
		}
	}
}

// GetZipResult
/**
 * @Description: 获取一个项目zip文件
 * @return gin.HandlerFunc
 */
func GetZipResult() gin.HandlerFunc {
	return func(context *gin.Context) {
		util.CheckDir("./temp/")
		workId, b := context.Params.Get("work_id")
		if !b {
			return
		}
		id, _ := strconv.Atoi(workId)
		db := dao.GetDB()
		work := new(pojo.Work)
		_, err := db.ID(id).Get(work)
		if err != nil {
			return
		}
		err = util.Compress("./work/"+work.Name+"/", "./temp/"+work.Name+".zip")
		if err != nil {
			return
		}
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", work.Name+".zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.File("./temp/" + work.Name + ".zip")
		defer os.Remove("./temp/" + work.Name + ".zip")
	}
}

// CreateWork
/**
 * @Description: 创建一个项目
 * @return gin.HandlerFunc
 */
func CreateWork() gin.HandlerFunc {
	return func(context *gin.Context) {
		work := new(pojo.Work)
		err := context.BindJSON(work)
		if err != nil {
			log.Errorln("绑定数据失败")
			log.Errorln(err.Error())
			return
		}
		db := dao.GetDB()
		_, err = db.Insert(work)
		if err != nil {
			return
		}
		util.CheckDir("./work/" + work.Name + "/")
		context.JSON(200, ok(200, nil))
	}
}
