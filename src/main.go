package main

import (
	"GINStudy/dao"
	"GINStudy/router"
	"GINStudy/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()
	//解析配置信息
	config := tools.ParseJson("D:\\Mycode\\Goland\\GINStudy\\config\\config.json")
	r = router.ApiBind(r)

	dao.InitDb("root", "root", "127.0.0.1", "gindb", 3306)


	r.Run(fmt.Sprintf(":%s", config.AppPort))
}
