package cmd

import (
	"gomi/app/queue"

	"github.com/spf13/cobra"
)

var CmdTestPublishNsq = &cobra.Command{
	Use:   "test-publish-nsq",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runTestPublishNsq,
}

// 调试完成后请记得清除测试代码
func runTestPublishNsq(cmd *cobra.Command, args []string) {
	payload := queue.MasterPayload{
		Test: "test",
	}
	// queue.PublishMaster(payload)
	queue.DeferredPublishMaster(payload)
}
