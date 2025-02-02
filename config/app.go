// Package config 站点配置信息
package config

import "gomi/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Get("app.name", "Gomi"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Get("app.env", "production"),

			// 是否进入调试模式
			"debug": config.GetBool("app.debug", false),

			// 应用服务端口
			"port": config.Get("app.port", "3000"),

			// 加密会话、JWT 加密
			"key": config.Get("app.key", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Get("app.url", "http://localhost:3000"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Get("app.timezone", "Asia/Shanghai"),

			// API 域名，未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Get("app.api_domain"),
		}
	})
}
