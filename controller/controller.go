package controller

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"work_server/dao"
	"work_server/pojo"
	"work_server/util"
)

var (
	TOKEN = "qqewqeqdadadadd"
)

func SetToken(token string) {
	TOKEN = token
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

// Upload
/**
 * @Description: 上传文件
 * @return gin.HandlerFunc
 */
func Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		p := new(pojo.Files)
		err := context.BindJSON(p)
		if err != nil {
			return
		}
		content := p.Data
		p.Data = ""
		db := dao.GetDB()
		p2 := new(pojo.Files)
		session := db.NewSession()
		defer session.Close()
		defer func() {
			err := recover()
			if err != nil {
				_ = session.Rollback()
			}
		}()
		exist, err := session.Where("work_id=? and file_name=?", p.WorkID, p.FileName).Exist(p2)
		if exist {
			_, err := session.Where("work_id=? and file_name=?", p.WorkID, p.FileName).Get(p2)
			if err != nil {
				session.Rollback()
				return
			}
			_, err = session.ID(p2.Id).Update(p)
			if err != nil {
				session.Rollback()
				log.Errorln(err.Error())
				return
			}
		} else {
			_, err = session.Insert(p)
			if err != nil {
				session.Rollback()
				return
			}
		}

		w := new(pojo.Work)
		_, err = session.ID(p.WorkID).Get(w)
		if err != nil {
			session.Rollback()
			return
		}
		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			session.Rollback()
			return
		}
		util.CheckDir("./work/" + w.Name + "/")
		err = os.WriteFile("./work/"+w.Name+"/"+p.FileName, data, 0666)
		if err != nil {
			session.Rollback()
			return
		}
		err = session.Commit()
		if err != nil {
			session.Rollback()
			return
		}
		context.JSON(200, ok(200, nil))
	}
}

// GetWorks
/**
 * @Description: 获取作业项目
 * @return gin.HandlerFunc
 */
func GetWorks() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := dao.GetDB()
		works := make([]pojo.Work, 0)
		err := db.Find(&works)
		if err != nil {
			return
		}
		context.JSON(200, response{
			Code:  200,
			Msg:   "",
			Data:  works,
			Error: "",
		})
	}
}

func GetWork() gin.HandlerFunc {
	return func(context *gin.Context) {
		workId, _ := context.Params.Get("work_id")
		id, _ := strconv.Atoi(workId)
		work := new(pojo.Work)
		_, err := dao.GetDB().ID(id).Get(work)
		if err != nil {
			return
		}
		context.JSON(200, response{
			Code:  200,
			Msg:   "",
			Data:  work,
			Error: "",
		})
	}
}

func GetFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		fileId, _ := context.Params.Get("file_id")
		id, _ := strconv.Atoi(fileId)
		file := new(pojo.Files)
		_, err := dao.GetDB().ID(id).Get(file)
		if err != nil {
			return
		}
		file.Data = ""
		context.JSON(200, response{
			Code:  200,
			Msg:   "",
			Data:  file,
			Error: "",
		})
	}
}

// GwtFiles
/**
 * @Description: 获取某个项目的所有文件
 * @return gin.HandlerFunc
 */
func GwtFiles() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := dao.GetDB()
		workId, b := context.Params.Get("work_id")
		if !b {
			return
		}
		id, _ := strconv.Atoi(workId)
		files := make([]pojo.Files, 0)
		err := db.Where("work_id=?", id).Find(&files)
		if err != nil {
			return
		}
		for _, file := range files {
			file.Data = ""
		}
		context.JSON(200, response{
			Code:  200,
			Msg:   "",
			Data:  files,
			Error: "",
		})
	}
}

func GetZipResult() gin.HandlerFunc {
	return func(context *gin.Context) {
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
		err = util.Zip("./work/"+work.Name+"/", "./temp/"+work.Name+".zip")
		if err != nil {
			return
		}
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", work.Name+".zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.File("./temp/" + work.Name + ".zip")
		defer os.Remove("./temp/" + work.Name + ".zip")
	}
}

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
		err = os.Remove("./work/" + work.Name + "/" + file.FileName)
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
