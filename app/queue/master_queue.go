package queue

import (
	"encoding/json"
	"fmt"
	"gomi/pkg/config"
	"gomi/pkg/logger"
	"gomi/pkg/nsqhelper"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

const MasterTask = "master_task"

// DetectionAccountQueue 用于管理私信任务队列
type MasterQueue struct {
	consumer *nsq.Consumer
}

var (
	masterConsumer     *MasterQueue
	masterConsumerOnce sync.Once
)

type MasterHandler struct{}

func GetMasterConsumerInstance() *MasterQueue {
	masterConsumerOnce.Do(func() {
		consumer, err := nsqhelper.AddConsumer(
			config.GetString("nsq.addr"),
			MasterTask, MasterTask,
			config.GetInt("nsq.consumer.master_concurrency"),
			&MasterHandler{},
		)
		if err != nil {
			logger.FatalString("MasterQueue", "GetMasterConsumerInstance", err.Error())
		}
		masterConsumer = &MasterQueue{consumer: consumer}
	})
	return masterConsumer
}

type MasterPayload struct {
	Test string `json:"test"`
}

func (mq *MasterHandler) HandleMessage(message *nsq.Message) error {
	var payload MasterPayload
	if err := json.Unmarshal(message.Body, &payload); err != nil {
		return err
	}
	fmt.Println(payload)

	return nil
}

func PublishMaster(payload MasterPayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return nsqhelper.Producer.Publish(MasterTask, data)
}

func DeferredPublishMaster(payload MasterPayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return nsqhelper.Producer.DeferredPublish(MasterTask, 1*time.Minute, data)
}
