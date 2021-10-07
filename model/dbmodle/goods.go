package dbmodle

type Goods struct {
	Id int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	Price int `json:"price"`
	OldPrice int `json:"oldPrice"`
	ShopId int `json:"shopId"`
}
