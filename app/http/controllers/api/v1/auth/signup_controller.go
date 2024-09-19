// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gomi/app/http/controllers/api/v1"
	"gomi/app/models/user"
	"gomi/app/requests"
	"gomi/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// SignupUsingPhone 使用手机和验证码进行注册
//func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
//
//	// 1. 验证表单
//	request := requests.SignupUsingPhoneRequest{}
//	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
//		return
//	}
//
//	// 2. 验证成功，创建数据
//	userModel := user.User{
//		Name:     request.Name,
//		Phone:    request.Phone,
//		Password: request.Password,
//	}
//	userModel.Create()
//
//	if userModel.ID > 0 {
//		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
//		response.CreatedJSON(c, gin.H{
//			"token": token,
//			"data":  userModel,
//		})
//	} else {
//		response.Abort500(c, "创建用户失败，请稍后尝试~")
//	}
//}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
