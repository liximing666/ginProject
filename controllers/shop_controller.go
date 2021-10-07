package controllers

import (
	"GINStudy/ecode"
	"GINStudy/model/seriliaze/request"
	"GINStudy/service"
	"GINStudy/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShopController struct {

}

func (c ShopController) GetNearBy(context *gin.Context) {
	json := context.Request.Body
	parmaRequest := request.NearByRequest{}
	err := tools.JsonParseTOStruct(json, &parmaRequest)
	if err != nil {
		context.JSON(http.StatusOK, ecode.JSON_ERROR)
	}

	shopService := service.NewShopService(context)
	data, code := shopService.GetNearBy(parmaRequest)

	context.JSON(http.StatusOK, gin.H{
		"code": code.GetEcode(),
		"data": data,
	})

}

func (c ShopController) SearchName(context *gin.Context) {
	json := context.Request.Body
	parmaRequest := request.SearchNameRequest{}
	err := tools.JsonParseTOStruct(json, &parmaRequest)
	if err != nil {
		context.JSON(http.StatusOK, ecode.JSON_ERROR)
	}

	shopService := service.NewShopService(context)
	data, code := shopService.SearchByName(parmaRequest)

	context.JSON(http.StatusOK, gin.H{
		"code": code.GetEcode(),
		"data": data,
	})
}

func (c ShopController) GetAllShopService(context *gin.Context) {
	json := context.Request.Body
	parmaRequest := request.GetShopServiceRequest{}
	err := tools.JsonParseTOStruct(json, &parmaRequest)
	if err != nil {
		context.JSON(http.StatusOK, ecode.JSON_ERROR)
	}

	shopService := service.NewShopService(context)
	data, code := shopService.GetAllShopService(parmaRequest)

	context.JSON(http.StatusOK, gin.H{
		"code": code.GetEcode(),
		"data": data,
	})
}

func (c ShopController) GetShopMenu(context *gin.Context) {
	parmaRequest := request.GetShopServiceRequest{}
	err := context.ShouldBindJSON(&parmaRequest)
	if err != nil {
		context.JSON(http.StatusOK, ecode.JSON_ERROR)
	}

	shopService := service.NewShopService(context)
	data, code := shopService.GetShopGoods(parmaRequest)

	context.JSON(http.StatusOK, gin.H{
		"code": code.GetEcode(),
		"data": data,
	})
}


