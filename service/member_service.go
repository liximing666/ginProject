package service

import (
	"GINStudy/dao"
	"GINStudy/ecode"
	"GINStudy/model/dbmodle"
	"GINStudy/model/seriliaze"
	"GINStudy/model/seriliaze/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type MemberService struct {
	//组合gin.Context 方便方法使用
	context *gin.Context
}

func NewMemberService(c *gin.Context) *MemberService {
	return &MemberService{context: c}
}

func FormatSmsItem(res dbmodle.Sms) *seriliaze.SmsItem {
	data := &seriliaze.SmsItem{
		ID:    res.ID,
		Phone: res.Phone,
		Code:  res.Code,
	}

	return data
}

func (s MemberService) SendSms(requestParam *request.SendSmsRequest) seriliaze.SmsItem {
	phoneNum := requestParam.PhoneNum
	//产生验证码
	//configInfo := tools.ParseJson("D:\\Mycode\\Goland\\GINStudy\\config\\config.json")
	//code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	res , err := dao.GetSmsCodeByPhone(phoneNum)
	if err != nil {
		return seriliaze.SmsItem{}
	}

	return *FormatSmsItem(res)
	////调用阿里云sdk
	//client, err := dysmsapi.NewClientWithAccessKey("cn-qingdao", "LTAI5t7iBPNYBnTCfWjsFFJK", "Y3JhnESe3rwXaaKi2Von5xQuHCViCR")
	///* use STS Token
	//client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	//*/
	//if err != nil {
	//	return false
	//}
	//
	//request1 := dysmsapi.CreateSendSmsRequest()
	//request1.Scheme = "https"
	//request1.SignName = "tname"
	//request1.TemplateCode = "1234"
	//request1.PhoneNumbers = phoneNum
	//par, err := json.Marshal(map[string]interface{}{
	//	"code": code,
	//})
	//request1.TemplateParam = string(par)
	//res, err := client.SendSms(request1)
	//if err != nil {
	//	return false
	//}
	//
	//if res.Code == "OK" {
	//	return true
	//}else {
	//	return false
	//}

}

func (s MemberService) Login(paramRequest *request.MemberLoginRequest) (int, ecode.ErrInfo) {
	userName := paramRequest.UserName
	password := paramRequest.Password

	isExists, err := dao.GetIsLoginByPassword(userName, password)

	if err != nil {
		return -1, ecode.MYSQL_ERROR
	}

	return isExists, ecode.OK

}

func (s MemberService) UploadAvatar(filePath string, id int, file *multipart.FileHeader) ecode.ErrInfo {
	count, err := dao.IsExistsById(id)
	if err != nil || count == 0 {
		return ecode.MYSQL_ERROR
	}

	err = dao.UpdateAvatarByPath(id, filePath)
	if err != nil {
		return ecode.MYSQL_ERROR
	}

	err = s.context.SaveUploadedFile(file, filePath)
	fmt.Println(err)

	if err != nil {
		return ecode.FILE_ERROR
	}

	return ecode.OK
}






