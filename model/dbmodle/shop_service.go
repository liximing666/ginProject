package dbmodle

type ShopService struct {
	Id int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	IconName string `json:"iconName"`
	IconColor string `json:"iconColor"`
}
