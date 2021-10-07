package seriliaze

type FoodCategoryItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
	LinkUrl string `json:"linkUrl"`
	IsInServing bool `json:"is_inServing"`
}

type FoodCategoryQuerySet struct {
	List []*FoodCategoryItem `json:"food_category"`
}
