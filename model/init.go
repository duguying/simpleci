package model

import (
	"github.com/duguying/simpleci/global"
	"github.com/go-xorm/xorm"
	"github.com/gogather/com/log"
	_ "github.com/mattn/go-sqlite3"
)

func InitModel() {
	var err error
	dbname := "./test.db"
	global.En, err = xorm.NewEngine("sqlite3", dbname)
	if err != nil {
		log.Redln("init database error: " + err.Error())
	} else {
		log.Greenln("using database: " + dbname)
	}
}
