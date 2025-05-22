package config

import "gomi/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// 过期时间，单位是分钟
			"expire_time": config.Env("jwt.expire_time", 120),

			// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
			"max_refresh_time": config.Env("jwt.max_refresh_time", 86400),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": config.Env("jwt.debug_expire_time", 86400),

			// JWT 密钥
			"secret": config.Env("jwt.secret", "your-secret-key"),
		}
	})
}
