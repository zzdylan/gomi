// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gomi/app/http/controllers/api/v1"
	"gomi/app/models/user"
	"gomi/app/requests"
	"gomi/pkg/captcha"
	"gomi/pkg/jwt"
	"gomi/pkg/response"

	"github.com/gin-gonic/gin"
)

// SigninController 登录控制器
type SigninController struct {
	v1.BaseAPIController
}

// Login使用username和password登录
func (sc *SigninController) Login(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SigninRequest{}
	if ok := requests.Validate(c, &request, requests.Signin); !ok {
		return
	}

	// 验证验证码
	if !captcha.NewCaptcha().VerifyCaptcha(request.CaptchaID, request.Captcha) {
		response.Fail(c, "验证码错误")
		return
	}

	// 尝试根据用户名获取用户
	userModel := user.GetByUsername(request.Username)

	// 用户不存在或密码错误
	if userModel.ID == 0 || !userModel.ComparePassword(request.Password) {
		response.Fail(c, "账号不存在或密码错误")
		return
	}

	// 生成 JWT Token
	token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)

	// 返回 Token 和用户数据
	response.Data(c, gin.H{
		"token": token,
	})
}
