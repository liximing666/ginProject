package service

import (
	"GINStudy/dao"
	"GINStudy/ecode"
	"GINStudy/model/dbmodle"
	"GINStudy/model/seriliaze"
	"GINStudy/model/seriliaze/request"
	"github.com/gin-gonic/gin"
)

type ShopService struct {
	c *gin.Context
}

func NewShopService(c *gin.Context) *ShopService {
	return &ShopService{c: c}
}

func (s ShopService) GetNearBy(parmaRequest request.NearByRequest) (seriliaze.NearByShopQuerySet, ecode.ErrInfo) {
	longitude := parmaRequest.Longitude
	latitude := parmaRequest.Latitude

	shops, err := dao.GetNearByShopByLocation(longitude, latitude)
	if err != nil || len(shops) == 0 {
		return seriliaze.NearByShopQuerySet{List: []*seriliaze.ShopItem{}}, ecode.MYSQL_ERROR
	}

	data := formatShops(shops)

	return data, ecode.OK
}

func (s ShopService) SearchByName(parmaRequest request.SearchNameRequest) (seriliaze.NearByShopQuerySet, ecode.ErrInfo) {

	shopName := parmaRequest.Name

	shops, err := dao.GetShopByName(shopName)
	if err != nil {
		return seriliaze.NearByShopQuerySet{List: []*seriliaze.ShopItem{}}, ecode.MYSQL_ERROR
	}

	data := formatShops(shops)

	return data, ecode.OK
}

func (s ShopService) GetAllShopService(parmaRequest request.GetShopServiceRequest) (seriliaze.ShopServiceQuerySet, ecode.ErrInfo) {
	id := parmaRequest.Id

	querySet, err := dao.GetShopServiceById(id)
	if err != nil {
		return seriliaze.ShopServiceQuerySet{}, ecode.MYSQL_ERROR
	}

	data := formatShopServiceQuerySet(querySet)

	return data, ecode.OK
}

func (s ShopService) GetShopGoods(parmaRequest request.GetShopServiceRequest) (seriliaze.GoodsItemQuerySet, ecode.ErrInfo) {
	shopId := parmaRequest.Id

	querySet, err := dao.GetShopGoodsById(shopId)
	if err != nil {
		return seriliaze.GoodsItemQuerySet{}, ecode.MYSQL_ERROR
	}

	data := formatGoodsQuerySet(querySet)

	return data, ecode.OK
}

func formatGoodsQuerySet(querySet []dbmodle.Goods) seriliaze.GoodsItemQuerySet {
	res := seriliaze.GoodsItemQuerySet{List: []*seriliaze.GoodsItem{}}
	for _, eachItem := range querySet {
		tmp := &seriliaze.GoodsItem{
			Name:  eachItem.Name,
			Price: eachItem.Price,
		}

		res.List = append(res.List, tmp)
	}

	return res
}

func formatShopServiceQuerySet(querySet []dbmodle.ShopService) seriliaze.ShopServiceQuerySet {
	var res = seriliaze.ShopServiceQuerySet{List: []*seriliaze.ShopServiceItem{}}
	for _, eachItem := range querySet {
		tmp := &seriliaze.ShopServiceItem{
			Id:        eachItem.Id,
			Name:      eachItem.Name,
			Desc:      eachItem.Desc,
			IconName:  eachItem.IconName,
			IconColor: eachItem.IconColor,
		}

	 	res.List = append(res.List, tmp)
	}

	return res
}

func formatShops(shops []dbmodle.Shop) seriliaze.NearByShopQuerySet {
	res := seriliaze.NearByShopQuerySet{List: []*seriliaze.ShopItem{}}
	for _, eachItem := range shops {
		tmp := &seriliaze.ShopItem{
			Id:   eachItem.Id,
			Name: eachItem.Name,
			Phone: eachItem.Phone,
		}

		res.List = append(res.List, tmp)
	}

	return res
}
