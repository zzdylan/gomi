package bootstrap

import (
	"gomi/app/queue"
	"gomi/pkg/config"
	"gomi/pkg/nsqhelper"
)

func SetupNsq() {
	nsqhelper.InitProducer(config.GetString("nsq.addr"))
	queue.GetMasterConsumerInstance()
}
