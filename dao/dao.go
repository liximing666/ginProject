package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func dbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	db.SingularTable(true)
	//db.CreateTable(dbmodle.Sms{})
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}


func InitDb(MyUser, Password, Host, Db string, Port int)  {
	DB = dbConn(MyUser, Password, Host, Db, Port)
}