package dbmodle

//商家与服务是多对多的关系
type ShopServiceRelation struct {
	ShopId int `gorm:"primaryKey" json:"shopId"`
	ServiceId int `gorm:"primaryKey" json:"serviceId"`
}
