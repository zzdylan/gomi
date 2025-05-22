// Package user 处理用户相关逻辑
package user

import (
	v1 "gomi/app/http/controllers/api/v1"
	"gomi/pkg/auth"
	"gomi/pkg/response"

	"github.com/gin-gonic/gin"
)

// UserController 登录控制器
type UserController struct {
	v1.BaseAPIController
}

// Me 获取当前用户信息
func (uc *UserController) Me(c *gin.Context) {
	user := auth.CurrentUser(c)
	response.Data(c, user)
}
