package logrusbadger

import (
	"fmt"

	honeybadger "github.com/honeybadger-io/honeybadger-go"
	"github.com/sirupsen/logrus"
)

type notifier interface {
	Notify(err interface{}, extra ...interface{}) (string, error)
}

// DefaultLevels control the levels that will notify Honeybadger.
var DefaultLevels = []logrus.Level{
	logrus.ErrorLevel,
	logrus.FatalLevel,
	logrus.PanicLevel,
}

// Hook is a logrus.Hook implementor that notifies Honeybadger.
type Hook struct {
	notifier notifier
}

// NewHook instantiates a Hook configured to use honeybadger.DefaultClient.
func NewHook() *Hook {
	return &Hook{
		notifier: honeybadger.DefaultClient,
	}
}

// Fire notifies Honeybadger with the passed-in entry. This will be called
// automatically by logrus for messages of the appropriate level.
func (h *Hook) Fire(entry *logrus.Entry) error {
	ctx := honeybadger.Context{}
	for k, v := range entry.Data {
		ctx[k] = fmt.Sprintf("%v", v)
	}

	msg, extra := msgAndClass(entry)
	extra = append(extra, ctx)

	_, err := h.notifier.Notify(msg, extra...)
	return err
}

// Levels tells logrus what message level we care about, and can be tweaked via
// the DefaultLevels pacakge-level var.
func (h *Hook) Levels() []logrus.Level {
	return DefaultLevels
}

func msgAndClass(entry *logrus.Entry) (string, []interface{}) {
	if err := entry.Data["error"]; err != nil {
		msg := fmt.Sprintf("%v", err)
		class := honeybadger.ErrorClass{Name: entry.Message}
		return msg, []interface{}{class}
	}

	return entry.Message, nil
}
