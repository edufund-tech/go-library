package dispatcher

import (
	nsqdriver "github.com/edufund-tech/go-library/nsq-driver"
	nsq "github.com/nsqio/go-nsq"
)

type dispatcher struct {
	targets []nsqdriver.Dispatch
}

//New instance of Dispatcher
func New() nsqdriver.Dispatcher {
	return &dispatcher{targets: []nsqdriver.Dispatch{}}
}

func (d *dispatcher) Add(target nsqdriver.Dispatch) nsqdriver.Dispatcher {
	d.targets = append([]nsqdriver.Dispatch{target}, d.targets...)
	return d
}

func (d *dispatcher) Wrap() nsq.HandlerFunc {
	next := func(m *nsq.Message) error { return nil }
	for _, dispatch := range d.targets {
		next = dispatch(next)
	}

	return next
}
