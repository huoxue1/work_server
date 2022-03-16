package controller

import (
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

// Upload
/**
 * @Description: 上传文件
 * @return gin.HandlerFunc
 */
func Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		p := new(pojo.Files)
		file, _ := context.FormFile("file")
		tempID, _ := context.GetPostForm("work_id")
		workId, _ := strconv.Atoi(tempID)
		token, _ := context.GetPostForm("token")
		p.WorkID = workId
		p.Size = file.Size
		p.FileName = file.Filename
		p.Token = token
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

		util.CheckDir("./work/" + w.Name + "/")
		err = context.SaveUploadedFile(file, "./work/"+w.Name+"/"+p.FileName)
		if err != nil {
			context.JSON(403, ok(403, "文件保存失败"))
			session.Rollback()
			return
		}

		err = session.Commit()
		if err != nil {
			session.Rollback()
			context.JSON(403, ok(403, "事务提交失败"))
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
