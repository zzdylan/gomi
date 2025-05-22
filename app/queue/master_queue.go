package queue

import (
	"encoding/json"
	"gomi/pkg/config"
	"gomi/pkg/logger"
	"gomi/pkg/nsqhelper"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
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
	logger.Info("MasterQueue", zap.Any("payload", payload))
	// 每30秒执行一次Touch(防止nsq超时重新进入队列，但是要注意的是就算touch也无法超过nsq启动参数中的--max-msg-timeout的时间)
	done := make(chan struct{})
	defer close(done)
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// message.Touch()
				logger.Info("MasterQueue", zap.String("status", "processing"))
			case <-done:
				return
			}
		}
	}()

	time.Sleep(1 * time.Hour)
	//结束定时touch
	done <- struct{}{}
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
