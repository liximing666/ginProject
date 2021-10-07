package dbmodle

type FoodCategory struct {
	Id int `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
	LinkUrl string `json:"linkUrl"`
	IsInServing bool `json:"is_inServing"`
}
