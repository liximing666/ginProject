package controllers

import (
	"GINStudy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {

}

func (c CategoryController) GetAll(context *gin.Context)  {

	categoryService := service.NewCategoryService(context)
	data, code := categoryService.GetAll()

	context.JSON(http.StatusOK, gin.H{
		"code": code.GetEcode(),
		"data": data,
	})
}

