package cmd

import (
	"gomi/pkg/logger"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {

	// // 存进去 redis 中
	// redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// // 从 redis 里取出
	// console.Success(redis.Redis.Get("hello"))
	logger.Info("test", zap.String("key", "value"))
	logger.Infof("test %d", 123)
}
