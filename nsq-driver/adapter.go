package nsqdriver

import nsq "github.com/nsqio/go-nsq"

//Adapter interface to MQ service
type Adapter interface {
	AddListener(topic, channel string, handler nsq.HandlerFunc) (*nsq.Consumer, error)
	AddPublisher() (Publisher, error)
}

//Publisher interface to MQ service
type Publisher interface {
	Publish(topic string, message interface{}) error
}
