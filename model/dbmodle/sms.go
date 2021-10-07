package dbmodle

//手机验证码表
type Sms struct {
	ID int `gorm:"primaryKey" json:"id"`
	Phone string `json:"phone"`
	Code string `json:"code"`
	UpdateTime   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	CreateTime   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}
