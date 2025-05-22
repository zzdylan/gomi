package config

import "gomi/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认数据库
			"connection": config.Get("database.connection", "mysql"),

			"mysql": map[string]interface{}{
				// 数据库连接信息
				"host":     config.Env("database.host", "127.0.0.1"),
				"port":     config.Env("database.port", "3306"),
				"database": config.Env("database.database", "gomi"),
				"username": config.Env("database.username", ""),
				"password": config.Env("database.password", ""),
				"charset":  config.Env("database.charset", "utf8mb4"),

				// 连接池配置
				"max_idle_connections": config.Env("database.max_idle_connections", 100),
				"max_open_connections": config.Env("database.max_open_connections", 25),
				"max_life_seconds":     config.Env("database.max_life_seconds", 300),
			},

			"sqlite": map[string]interface{}{
				"database": config.Env("database.sqlite_file", "database/database.db"),
			},
		}
	})
}
