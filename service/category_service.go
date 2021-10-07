package service

import (
	"GINStudy/dao"
	"GINStudy/ecode"
	"GINStudy/model/dbmodle"
	"GINStudy/model/seriliaze"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	context *gin.Context
}

func NewCategoryService(c *gin.Context) *CategoryService {
	return &CategoryService{context: c}
}

func (s CategoryService) GetAll() (seriliaze.FoodCategoryQuerySet, ecode.ErrInfo) {

	queryList, err := dao.GetAllFoodCategory()
	if err != nil {
		return seriliaze.FoodCategoryQuerySet{}, ecode.MYSQL_ERROR
	}

	data := formatGetAllCategory(queryList)

	return data, ecode.OK
}

func formatGetAllCategory(queryList []dbmodle.FoodCategory) seriliaze.FoodCategoryQuerySet {
	res := seriliaze.FoodCategoryQuerySet{List: []*seriliaze.FoodCategoryItem{}}

	for _, eachItem := range queryList {
		tmp := &seriliaze.FoodCategoryItem{
			Id:          eachItem.Id,
			Title:       eachItem.Title,
			Description: eachItem.Description,
			ImageUrl:    eachItem.ImageUrl,
			LinkUrl:     eachItem.LinkUrl,
			IsInServing: eachItem.IsInServing,
		}

		res.List = append(res.List, tmp)
	}

	return res
}
