package nsqdriver

import nsq "github.com/nsqio/go-nsq"

//Dispatch is a middleware to distribute MQ messages to another handler
type Dispatch func(h nsq.HandlerFunc) nsq.HandlerFunc

//Dispatcher is the pipeline that dispatch the message sequentially
type Dispatcher interface {
	Add(d Dispatch) Dispatcher
	Wrap() nsq.HandlerFunc
}
