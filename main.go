package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"work_server/controller"
	_ "work_server/dao"
	"work_server/router"
	"work_server/util"
)

func init() {
	util.CheckDir("./work/")
	util.CheckDir("./db/")
	_, err := os.Stat("./db/data.db")
	if err != nil {
		_, err := os.Create("./db/data.db")
		if err != nil {
			return
		}
		return
	}
}

var (
	token string
	port  int
	ssl   bool
)

func init() {
	flag.StringVar(&token, "t", "qqewqeqdadadadd", "set a token")
	flag.IntVar(&port, "p", 8085, "set port")
	flag.BoolVar(&ssl, "ssl", false, "use a ssl")
	flag.Parse()

	controller.SetToken(token)
	router.SetToken(token)
}

func main() {
	if ssl {
		if err := router.InitRouter().RunTLS(fmt.Sprintf("0.0.0.0:%d", port), "key.pem", "key.key"); err != nil {
			log.Panicln(err.Error())
		}
	} else {
		if err := router.InitRouter().Run(fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
			log.Panicln(err.Error())
		}
	}
}
