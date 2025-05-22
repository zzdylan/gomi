// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	//"gomi/app/requests/validators"
)

type SigninRequest struct {
	Username  string `json:"username" valid:"username"`
	Password  string `json:"password" valid:"password"`
	Captcha   string `json:"captcha" valid:"captcha"`
	CaptchaID string `json:"captcha_id" valid:"captcha_id"`
}

func Signin(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"username":   []string{"required", "min:4", "max:30"},
		"password":   []string{"required", "min:6", "max:30"},
		"captcha":    []string{"required"},
		"captcha_id": []string{"required"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名不能为空",
			"min:用户名长度不能小于 4",
			"max:用户名长度不能大于 30",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码长度不能小于 6",
			"max:密码长度不能大于 30",
		},
		"captcha": []string{
			"required:验证码不能为空",
		},
		"captcha_id": []string{
			"required:验证码ID不能为空",
		},
	}
	return validate(data, rules, messages)
}
