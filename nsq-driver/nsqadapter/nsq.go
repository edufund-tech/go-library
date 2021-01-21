package nsqadapter

import (
	nsqdriver "github.com/edufund-tech/go-library/nsq-driver"
	nsq "github.com/nsqio/go-nsq"
)

type nopLogger struct{}

func (*nopLogger) Output(int, string) error {
	return nil
}

type nsqAdapterImpl struct {
	nopl       *nopLogger
	nsqconfig  *nsq.Config
	nsqd       string
	nsqlookupd string
}

//NewNSQAdapter new instance of MQ adapter using NSQ
func NewNSQAdapter(nsqd, nsqlookupd string) nsqdriver.Adapter {
	return &nsqAdapterImpl{
		nopl:       &nopLogger{},
		nsqconfig:  nsq.NewConfig(),
		nsqd:       nsqd,
		nsqlookupd: nsqlookupd,
	}
}

//Listen to a message from a topic in NSQ
func (a *nsqAdapterImpl) AddListener(topic, channel string, handler nsq.HandlerFunc) (*nsq.Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, a.nsqconfig)
	if err != nil {
		return nil, err
	}

	consumer.SetLogger(a.nopl, 0)
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQLookupd(a.nsqlookupd)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (a *nsqAdapterImpl) AddPublisher() (nsqdriver.Publisher, error) {
	prod, err := nsq.NewProducer(a.nsqd, a.nsqconfig)
	if err != nil {
		return nil, err
	}

	return &nsqPublisherImpl{prod: prod}, nil
}
