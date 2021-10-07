package dao

import (
	"GINStudy/model/dbmodle"
)

func GetSmsCodeByPhone(phone string) (dbmodle.Sms, error) {
	var obj dbmodle.Sms
	err := DB.Model(dbmodle.Sms{}).Where("phone = ?", phone).Take(&obj).Error

	return obj, err
}

func GetIsLoginByPassword(userName, password string) (int, error) {
	var count int64
	db := DB.Model(dbmodle.Member{}).Where("user_name = ?", userName).Where("password = ?", password)
 	err := db.Count(&count).Error

 	return int(count), err
}

func IsExistsById(id int) (int, error) {
	var count int64

	err := DB.Model(dbmodle.Member{}).Where("id = ?", id).Count(&count).Error

	return int(count), err
}

func UpdateAvatarByPath(id int, path string) error {
 	err := DB.Model(dbmodle.Member{}).Where("id = ?", id).Update("avatar", path[1:]).Error

	return err
}



