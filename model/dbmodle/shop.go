package dbmodle

type Shop struct {
	Id int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	Phone string `json:"phone"`
}
