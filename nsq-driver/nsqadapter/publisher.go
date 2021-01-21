package nsqadapter

import (
	"encoding/json"

	nsq "github.com/nsqio/go-nsq"
)

type nsqPublisherImpl struct {
	prod *nsq.Producer
}

func (p *nsqPublisherImpl) Publish(topic string, message interface{}) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.prod.Publish(topic, b)
	if err != nil {
		return err
	}

	return nil
}
