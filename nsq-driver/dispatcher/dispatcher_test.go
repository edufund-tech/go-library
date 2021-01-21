package dispatcher

import (
	"errors"
	"testing"

	nsqdriver "github.com/edufund-tech/go-library/nsq-driver"
	"github.com/stretchr/testify/assert"

	nsq "github.com/nsqio/go-nsq"
)

func TestDispatcher(t *testing.T) {
	result1 := ""
	step1 := func(param string) nsqdriver.Dispatch {
		return func(next nsq.HandlerFunc) nsq.HandlerFunc {
			return func(message *nsq.Message) error {
				result1 = param
				next(message)
				return nil
			}
		}
	}

	result2 := 10
	step2 := func(param int) nsqdriver.Dispatch {
		return func(next nsq.HandlerFunc) nsq.HandlerFunc {
			return func(message *nsq.Message) error {
				result2 = param
				next(message)
				return nil
			}
		}
	}

	step3 := func() nsqdriver.Dispatch {
		return func(next nsq.HandlerFunc) nsq.HandlerFunc {
			return func(message *nsq.Message) error {
				return errors.New("")
			}
		}
	}

	final := "final"
	step4 := func() nsqdriver.Dispatch {
		return func(next nsq.HandlerFunc) nsq.HandlerFunc {
			return func(message *nsq.Message) error {
				final = "should not be called"
				next(message)
				return nil
			}
		}
	}

	New().Add(step1("a param")).
		Add(step2(100)).
		Wrap().
		HandleMessage(nil)
	assert.Equal(t, "a param", result1)
	assert.Equal(t, 100, result2)

	New().Add(step1("a param")).
		Add(step2(100)).
		Add(step3()).
		Add(step4()).
		Wrap().
		HandleMessage(nil)
	assert.Equal(t, "final", final)
}
