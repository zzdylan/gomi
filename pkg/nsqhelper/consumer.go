package nsqhelper

import (
	"time"

	"github.com/nsqio/go-nsq"
)

func AddConsumer(addr, topic, channel string, concurrency int, handler nsq.Handler, opts ...Option) (*nsq.Consumer, error) {
	myconfig := nsq.NewConfig()
	// myconfig.MaxInFlight = 100
	myconfig.MaxInFlight = concurrency
	myconfig.MaxAttempts = 0 // 消息最大重试次数
	myconfig.MsgTimeout = 5 * time.Minute
	// myconfig.MsgTimeout = 10 * time.Hour
	o := Options{
		config: myconfig,
	}

	for _, opt := range opts {
		opt(&o)
	}
	consumer, err := nsq.NewConsumer(topic, channel, o.config)
	if err != nil {
		return nil, err
	}
	consumer.AddConcurrentHandlers(handler, concurrency)
	// if err := consumer.ConnectToNSQLookupd(lookupd); err != nil {
	if err := consumer.ConnectToNSQD(addr); err != nil {
		return nil, err
	}
	return consumer, nil
}
