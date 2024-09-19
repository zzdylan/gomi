package test

import (
	"github.com/gin-gonic/gin"
	v1 "gomi/app/http/controllers/api/v1"
	"gomi/pkg/cache"
	"net/http"
	"time"
)

// TestController 注册控制器
type TestController struct {
	v1.BaseAPIController
}

func (sc *TestController) Index(c *gin.Context) {
	//panic("这是 panic 测试")
	// 获取请求参数，并做表单验证
	//request := requests.SignupPhoneExistRequest{}
	//if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
	//	return
	//}
	////  检查数据库并返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"success": true,
	//})
	type MyTestObj struct {
		Test1 string `json:"test1"`
		Test2 string `json:"test2"`
	}
	myTestObj := MyTestObj{Test1: "test1", Test2: "test2"}
	// 设置缓存
	cache.Set("mytest", myTestObj, 2*time.Minute)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
