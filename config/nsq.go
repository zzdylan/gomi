package config

import "gomi/pkg/config"

func init() {
	config.Add("nsq", func() map[string]interface{} {
		return map[string]interface{}{
			"addr":    config.Env("nsq.addr", "127.0.0.1:4150"),
			"lookupd": config.Env("nsq.lookupd", "127.0.0.1:4161"),
			"consumer": map[string]interface{}{
				"master_concurrency": config.Env("nsq.consumer.master_concurrency", 1),
			},
		}
	})
}
