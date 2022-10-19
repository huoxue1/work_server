package controller

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

func uploadFile(context *gin.Context) {
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

func uploadDir(context *gin.Context) {
	p := new(pojo.Files)
	form, _ := context.MultipartForm()
	files := form.File["file"]
	fileNames := strings.Split(context.PostForm("fileNames"), ",")
	fileName := strings.Split(fileNames[0], "/")[0]
	var fileSize int64 = 0
	for i, file := range files {
		fileSize += file.Size
		file.Filename = fileNames[i]
	}
	tempID, _ := context.GetPostForm("work_id")
	workId, _ := strconv.Atoi(tempID)
	token, _ := context.GetPostForm("token")
	p.WorkID = workId
	p.Size = fileSize
	p.FileName = fileName
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
	for _, file := range files {
		s := "./work/" + w.Name + "/" + file.Filename
		_ = os.MkdirAll(filepath.Dir(s), os.ModePerm)
		err := context.SaveUploadedFile(file, s)
		if err != nil {
			context.JSON(403, ok(403, "文件保存失败"))
			session.Rollback()
			return
		}
	}

	err = session.Commit()
	if err != nil {
		session.Rollback()
		context.JSON(403, ok(403, "事务提交失败"))
		return
	}
	context.JSON(200, ok(200, nil))
}

// Upload
/**
 * @Description: 上传文件
 * @return gin.HandlerFunc
 */
func Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		fileType, _ := context.GetPostForm("type")
		if fileType == "file" {
			uploadFile(context)
		} else {
			uploadDir(context)
		}

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

func DownloadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileId, _ := ctx.Params.Get("file_id")
		workId, _ := ctx.Params.Get("work_id")
		fileID, _ := strconv.Atoi(fileId)
		workID, _ := strconv.Atoi(workId)
		token := ctx.Query("token")
		file := new(pojo.Files)
		_, err := dao.GetDB().ID(fileID).Get(file)
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			return
		}
		if token != TOKEN && token != file.Token {
			ctx.JSON(401, gin.H{"code": 401})
			return
		}
		work := new(pojo.Work)
		_, err = dao.GetDB().ID(workID).Get(work)
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			return
		}
		ctx.Header("content-disposition", "attachment;filename="+file.FileName)
		path := fmt.Sprintf("./work/%s/%s", work.Name, file.FileName)
		info, _ := os.Stat(path)
		if info.IsDir() {
			_ = util.Compress(path, "./temp/"+file.FileName+".zip")
			ctx.File("./temp/" + file.FileName + ".zip")
			defer os.Remove("./temp/" + file.FileName + ".zip")
		} else {
			ctx.File(path)
		}

	}
}

func RenameFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileId, _ := ctx.Params.Get("file_id")
		workId, _ := ctx.Params.Get("work_id")
		fileID, _ := strconv.Atoi(fileId)
		workID, _ := strconv.Atoi(workId)
		token := ctx.Query("token")
		newName := ctx.PostForm("new_name")
		file := new(pojo.Files)
		_, err := dao.GetDB().ID(fileID).Get(file)
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			return
		}
		if token != TOKEN && token != file.Token {
			ctx.JSON(401, gin.H{"code": 401})
			return
		}
		work := new(pojo.Work)
		_, err = dao.GetDB().ID(workID).Get(work)
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			return
		}
		path := fmt.Sprintf("./work/%v/%v", work.Name, file.FileName)
		err = os.Rename(path, fmt.Sprintf("./work/%v/%v", work.Name, newName))
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			log.Errorln(err.Error())
			return
		}
		file.FileName = newName
		_, err = dao.GetDB().ID(fileID).Update(file)
		if err != nil {
			ctx.JSON(503, gin.H{"code": 503, "err": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"code": 200})
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
