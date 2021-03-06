package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masato25/lambdaQuery/g"
)

var (
	db  *gorm.DB
	err error
)

func DBConn() *gorm.DB {
	return db
}

func Init() {
	conf := g.Config()
	db, err = gorm.Open("mysql", conf.GraphDB.Addr)
	if err != nil {
		log.Println(err.Error())
	}
}
