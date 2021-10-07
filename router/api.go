package router

import (
	"GINStudy/controllers"
	"github.com/gin-gonic/gin"
)

func ApiBind(r *gin.Engine) *gin.Engine {
	r.GET("/hello", controllers.HelloController{}.Hello)

	//用户模块的api
	userGroup := r.Group("/member")
	{
		userGroup.GET("/hello", controllers.MemberController{}.HelloMem)
		userGroup.GET("/send-sms", controllers.MemberController{}.SendSms)
		userGroup.POST("/login", controllers.MemberController{}.Login)
		userGroup.POST("/upload_avatar",controllers.MemberController{}.UploadAvatar)

	}

	//食品分类模块
	categoryGroup := r.Group("/category")
	{
		categoryGroup.GET("/all", controllers.CategoryController{}.GetAll)
	}

	//商户模块
	shopGroup := r.Group("/shop")
	{
		shopGroup.POST("/nearby", controllers.ShopController{}.GetNearBy)
		shopGroup.POST("/search_by_name", controllers.ShopController{}.SearchName)
		shopGroup.POST("/get_service", controllers.ShopController{}.GetAllShopService)
		shopGroup.GET("/get_menu", controllers.ShopController{}.GetShopMenu)

	}

	return r
}