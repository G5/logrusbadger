package logrusbadger

import (
	"errors"
	"testing"

	honeybadger "github.com/honeybadger-io/honeybadger-go"
	"github.com/sirupsen/logrus"
)

type Notification struct {
	Message string
	Extra   []interface{}
}

type FakeNotifier struct {
	T             *testing.T
	Notifications []Notification
	NotifyErr     error
}

func NewFakeNotifer(t *testing.T) *FakeNotifier {
	return &FakeNotifier{T: t}
}

func (n *FakeNotifier) Notify(msg interface{}, extra ...interface{}) (string, error) {
	s, ok := msg.(string)
	if !ok {
		n.T.Errorf("message must be a string")
	}
	notification := Notification{Message: s, Extra: extra}
	n.Notifications = append(n.Notifications, notification)

	return "", n.NotifyErr
}

func TestLevels(t *testing.T) {
	DefaultLevels = []logrus.Level{logrus.PanicLevel}
	h := NewHook()
	actual := h.Levels()

	if i := len(actual); i != 1 {
		t.Errorf("unexpected length %d of Levels", i)
	}
	if actual[0] != logrus.PanicLevel {
		t.Errorf("expected PanicLevel, got %v", actual[0])
	}
}

func TestFire_WithErrorField(t *testing.T) {
	h := NewHook()
	fake := NewFakeNotifer(t)
	h.notifier = fake
	e := logrus.WithField("test", "value").WithError(errors.New("test error"))
	e.Message = "test message"

	err := h.Fire(e)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if i := len(fake.Notifications); i != 1 {
		t.Errorf("expected 1 notification, got %d", i)
	}
	n := fake.Notifications[0]
	if expected := "test error"; n.Message != expected {
		t.Errorf("expected '%s', got '%s'", expected, n.Message)
	}
	if i := len(n.Extra); i != 2 {
		t.Errorf("expected 2 extras, got %d", i)
	}
	class, ok := n.Extra[0].(honeybadger.ErrorClass)
	if !ok {
		t.Error("expected first extra to be honeybadger.ErrorClass")
	}
	if expected := "test message"; class.Name != expected {
		t.Errorf("expected error class '%s', got '%s'", expected, class.Name)
	}
	ctx, ok := n.Extra[1].(honeybadger.Context)
	if !ok {
		t.Error("expected second extra to be honeybadger.Context")
	}
	if i := len(ctx); i != 2 {
		t.Errorf("expected 2 entries in context, got %d: %v", i, ctx)
	}
	if v := ctx["test"]; v != "value" {
		t.Errorf("expected test='value' in context, got %v", v)
	}
	if v := ctx["error"]; v != "test error" {
		t.Errorf("expected error='test error' in context, got %v", v)
	}
}

func TestFire_NoErrorField(t *testing.T) {
	h := NewHook()
	fake := NewFakeNotifer(t)
	h.notifier = fake
	e := logrus.WithField("test", "value")
	e.Message = "test message"

	err := h.Fire(e)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if i := len(fake.Notifications); i != 1 {
		t.Errorf("expected 1 notification, got %d", i)
	}
	n := fake.Notifications[0]
	if n.Message == "" {
		t.Errorf("expected empty message, got '%s'", n.Message)
	}
	if i := len(n.Extra); i != 1 {
		t.Errorf("expected 1 extras, got %d: %+v", i, n.Extra)
	}
	ctx, ok := n.Extra[0].(honeybadger.Context)
	if !ok {
		t.Error("expected extra to be honeybadger.Context")
	}
	if i := len(ctx); i != 1 {
		t.Errorf("expected 1 entry in context, got %d: %v", i, ctx)
	}
	if v := ctx["test"]; v != "value" {
		t.Errorf("expected test='value' in context, got %v", v)
	}
}

func TestFire_NotifyError(t *testing.T) {
	h := NewHook()
	fake := NewFakeNotifer(t)
	fake.NotifyErr = errors.New("test error")
	h.notifier = fake
	e := logrus.WithField("test", "value")
	e.Message = "test message"

	err := h.Fire(e)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "test error" {
		t.Errorf("expected 'test error', got '%s'", err.Error())
	}
}
