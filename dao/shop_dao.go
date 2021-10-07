package dao

import (
	"GINStudy/model/dbmodle"
)

func GetNearByShopByLocation(longitude, latitude float64) ([]dbmodle.Shop, error) {
	var shops []dbmodle.Shop
	at := latitude + 1
	ab := latitude - 1
	ot := longitude + 1
	ob := longitude - 1
	err := DB.Model(dbmodle.Shop{}).Where("longitude BETWEEN ? AND ?", ob, ot).Where("latitude BETWEEN ? AND ?", ab, at).Find(&shops).Error

	return shops, err
}

func GetShopByName(shopName string) ([]dbmodle.Shop, error) {
	var shops []dbmodle.Shop
	err := DB.Model(dbmodle.Shop{}).Where("name LIKE ?", "%" + shopName + "%").Find(&shops).Error

	return shops, err
}

func GetShopServiceById(shopId int) ([]dbmodle.ShopService, error) {
	var shopServices []dbmodle.ShopService
	err := DB.Table("shop_service").Select("shop_service.name").
		Joins("inner join shop_service_relation on shop_service.id = shop_service_relation.serviceId").
		Where("shop_service_relation.shopId = ?", shopId).Find(&shopServices).Error

	return shopServices, err
}

func GetShopGoodsById(shopId int) ([]dbmodle.Goods, error) {
	var goods []dbmodle.Goods
	err := DB.Model(dbmodle.Goods{}).Where("shopId = ?", shopId).Find(&goods).Error

	return goods, err
}
