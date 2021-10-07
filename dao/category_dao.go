package dao

import (
	"GINStudy/model/dbmodle"
)

func GetAllFoodCategory() ([]dbmodle.FoodCategory, error) {
	var allCategory []dbmodle.FoodCategory
	err := DB.Model(dbmodle.FoodCategory{}).Find(&allCategory).Error
	
	return allCategory, err
}
