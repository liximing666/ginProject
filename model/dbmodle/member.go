package dbmodle

type Member struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	UserName string `json:"user_name"`
	Mobile   string	`json:"mobile"`
	PassWord string	`json:"password"`
	Avatar string	`json:"avatar"`
	Balance float64	`json:"blance"`
	IsActive int8	`json:"is_active"`
	City string 	`json:"city"`
	UpdateTime   int64 `gorm:"autoUpdateTime:milli" json:"update_time"` // 使用时间戳毫秒数填充更新时间
	CreateTime   int64 `gorm:"autoCreateTime" json:"create_time"`      // 使用时间戳秒数填充创建时间
}
