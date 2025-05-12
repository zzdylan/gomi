package nsqhelper

import (
	"github.com/nsqio/go-nsq"
)

type Option func(o *Options)

type Options struct {
	config *nsq.Config
}

func WithConfig(c *nsq.Config) Option {
	return func(o *Options) {
		o.config = c
	}
}
