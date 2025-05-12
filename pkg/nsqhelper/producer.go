package nsqhelper

import (
	"github.com/nsqio/go-nsq"
)

var Producer *nsq.Producer

// InitProducer 初始化生产者
func InitProducer(addr string) error {
	var err error
	Producer, err = nsq.NewProducer(addr, nsq.NewConfig())
	return err
}
