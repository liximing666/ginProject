package request

type SendSmsRequest struct {
	PhoneNum string `json:"phone_num"`
}

type MemberLoginRequest struct {
	UserName string `json:"userName"`
	Password string	`json:"password"`
}

type AvatarRequest struct {
	Id int `json:"id"`
	FilePath string `json:"filePath"`
}