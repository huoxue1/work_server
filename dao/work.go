package dao

import ( //nolint:gci
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"work_server/pojo"
)

var (
	db *xorm.Engine
)

func init() {
	daoInit()
}

// GetDB
/**
 * @Description: 获取数据库引擎
 * @return *xorm.Engine
 */
func GetDB() *xorm.Engine {
	if db.Ping() != nil {
		daoInit()
	}
	return db
}

func daoInit() {
	var err error
	db, err = xorm.NewEngine("sqlite3", "./db/data.db")
	if err != nil {
		log.Panicln("连接数据库失败" + err.Error())
		return
	}
	err = db.Sync2(new(pojo.Work), new(pojo.Files))
	if err != nil {
		log.Errorln("数据库同步失败" + err.Error())
		return
	}
}
