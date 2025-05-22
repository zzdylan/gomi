package config

import "gomi/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("redis.host", "127.0.0.1"),
			"port":     config.Env("redis.port", "6379"),
			"password": config.Env("redis.password", ""),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"database": config.Env("redis.database", 1),

			// 缓存 cache 包使用 0 ，缓存清空理应当不影响业务
			"database_cache": config.Env("redis.database_cache", 0),
		}
	})
}
