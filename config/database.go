package config

import "gomi/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认数据库
			"connection": config.Get("database.connection", "mysql"),

			"mysql": map[string]interface{}{
				// 数据库连接信息
				"host":     config.Get("database.host", "127.0.0.1"),
				"port":     config.Get("database.port", "3306"),
				"database": config.Get("database.database", "gomi"),
				"username": config.Get("database.username", ""),
				"password": config.Get("database.password", ""),
				"charset":  config.Get("database.charset", "utf8mb4"),

				// 连接池配置
				"max_idle_connections": config.GetInt("database.max_idle_connections", 100),
				"max_open_connections": config.GetInt("database.max_open_connections", 25),
				"max_life_seconds":     config.GetInt("database.max_life_seconds", 300),
			},

			"sqlite": map[string]interface{}{
				"database": config.Get("database.sqlite_file", "database/database.db"),
			},
		}
	})
}
