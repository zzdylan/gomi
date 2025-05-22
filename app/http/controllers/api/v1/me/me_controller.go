// Package me 处理用户相关逻辑
package me

import (
	v1 "gomi/app/http/controllers/api/v1"
	"gomi/pkg/auth"
	"gomi/pkg/response"

	"github.com/gin-gonic/gin"
)

// MeController 登录控制器
type MeController struct {
	v1.BaseAPIController
}

// Login使用username和password登录
func (mc *MeController) Info(c *gin.Context) {
	user := auth.CurrentUser(c)
	response.Data(c, user)
}
