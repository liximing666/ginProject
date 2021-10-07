package seriliaze

type ShopItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type NearByShopQuerySet struct {
	List []*ShopItem `json:"list"`
}

type ShopServiceItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	IconName string `json:"iconName"`
	IconColor string `json:"iconColor"`
}

type ShopServiceQuerySet struct {
	List []*ShopServiceItem `json:"list"`
}

type GoodsItem struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

type GoodsItemQuerySet struct {
	List []*GoodsItem `json:"list"`
}