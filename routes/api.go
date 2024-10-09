// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gomi/app/http/controllers/api/v1/auth"
	"gomi/app/http/controllers/api/v1/test"
	"gomi/app/http/middlewares"
	"gomi/pkg/config"
	"gomi/pkg/response"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	var v1 *gin.RouterGroup

	r.GET("/ping", func(c *gin.Context) {
		response.Success(c)
	})

	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	// 注册一个路由
	//v1.GET("/", func(c *gin.Context) {
	//	// 以 JSON 格式响应
	//	c.JSON(http.StatusOK, gin.H{
	//		"Hello": "World!",
	//	})
	//})
	v1.Use(middlewares.LimitIP("5-M"))
	{
		testGroup := v1.Group("/test")
		{
			t := new(test.TestController)
			// 判断手机是否已注册
			testGroup.POST("/index", t.Index)
		}
	}
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}

}
