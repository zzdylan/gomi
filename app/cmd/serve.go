package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gomi/bootstrap"
	"gomi/pkg/config"
	"gomi/pkg/logger"
	"net/http"
	"os/signal"
	"syscall"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":" + config.Get("app.port"),
		Handler: router,
	}

	// 在启动服务器前打印提示信息
	logger.Info("Starting web server on port " + config.Get("app.port"))
	defer logger.Info("exit web server")

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatal(err.Error())
		}
		logger.Info("Shutdown web server")
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}
}
