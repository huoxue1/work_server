package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	_ "work_server/dao"
	"work_server/router"
	"work_server/util"
)

func init() {
	util.CheckDir("./work/")
	util.CheckDir("./db/")
	_, err := os.Stat("./db/data.db")
	if err != nil {
		os.Create("./db/data.db")
		return
	}
}

func main() {
	if err := router.InitRouter().Run("127.0.0.1:8085"); err != nil {
		log.Panicln(err.Error())
	}
}
