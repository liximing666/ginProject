package controllers

import (
	"GINStudy/ecode"
	"GINStudy/model/seriliaze/request"
	"GINStudy/service"
	"GINStudy/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MemberController struct {
}

func (c MemberController) HelloMem(ctx *gin.Context )  {
	ctx.String(http.StatusOK, "ok")
}

func (c MemberController) SendSms(context *gin.Context) {
	requestParam := &request.SendSmsRequest{}
	phone, ok := context.GetQuery("phone")
	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"mes": "phone err",
		})
		return
	}
	requestParam.PhoneNum = phone

	memberService := service.NewMemberService(context)
	data := memberService.SendSms(requestParam)

	context.JSON(http.StatusOK, data)

}

func (c MemberController) Login(context *gin.Context) {
	//前端传参过来的请求参数装进结构体
	paramRequest := &request.MemberLoginRequest{}
	json := context.Request.Body
	err := tools.JsonParseTOStruct(json, paramRequest)
	if err != nil {
		context.JSON(200, gin.H{"isExists": -1, "code":ecode.JSON_ERROR.GetEcode() })
	}

	memberService := service.NewMemberService(context)
	data, errInfo := memberService.Login(paramRequest)
	//抛出自己定义的异常类
	context.JSON(200, gin.H{"isExists": data, "code":errInfo.GetEcode() })

}

func (c MemberController) UploadAvatar(context *gin.Context) {
	//接收参数
	file, err := context.FormFile("avatar")
	if err != nil {
		context.JSON(http.StatusOK, ecode.FILE_ERROR)
	}
	id, _ := strconv.Atoi(context.PostForm("id"))
	filePath := "./static/img/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	fmt.Println(filePath)

	memberService := service.NewMemberService(context)
	code := memberService.UploadAvatar(filePath, id, file)

	context.JSON(http.StatusOK, code)
}
