package waterlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
)

// AssertJSON asserts that two input strings are the same JSON value
func AssertJSON(t *testing.T, act, exp string) {
	var actjs, expjs map[string]interface{}
	if act == exp {
		return
	}
	if err := json.Unmarshal([]byte(act), &actjs); err != nil {
		t.Fatalf("unexpected err: %s; act=%q exp=%q", err, act, exp)
	}
	if err := json.Unmarshal([]byte(exp), &expjs); err != nil {
		t.Fatalf("unexpected err: %s; act=%q exp=%q", err, act, exp)
	}
	if !reflect.DeepEqual(expjs, actjs) {
		t.Fatalf("expected %s but got %s", exp, act)
	}
}

func TestLogger(t *testing.T, newLogger NewFunc) {
	tests := map[string]func(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T){
		"TestError":       TestError,
		"TestErrorFields": TestErrorFields,
		"TestInfo":        TestInfo,
		"TestInfoFields":  TestInfoFields,
		"TestDebug":       TestDebug,
		"TestDebugFields": TestDebugFields,
		"TestTrace":       TestTrace,
		"TestTraceFields": TestTraceFields,
		"TestWith":        TestWith,
	}
	for name, test := range tests {
		output := &bytes.Buffer{}
		debug := name == "TestDebug" || name == "TestDebugFields"
		trace := name == "TestTrace" || name == "TestTraceFields"
		logger := newLogger(output, debug, trace)
		t.Run(name, test(output, logger))
	}
}

func TestError(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Error("message", errors.New("error"), nil)
		AssertJSON(t, b.String(), `{"error":"error","message":"message","level":"error"}`)
	}
}

func TestErrorFields(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Error("message", errors.New("error"), watermill.LogFields{"foo": "bar"})
		AssertJSON(t, b.String(), `{"error":"error","message":"message","level":"error","foo":"bar"}`)
	}
}

func TestInfo(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Info("message", nil)
		AssertJSON(t, b.String(), `{"message":"message","level":"info"}`)
	}
}

func TestInfoFields(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Info("message", watermill.LogFields{"foo": "bar"})
		AssertJSON(t, b.String(), `{"message":"message","level":"info","foo":"bar"}`)
	}
}

func TestDebug(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Debug("message", nil)
		AssertJSON(t, b.String(), `{"message":"message","level":"debug"}`)
	}
}

func TestDebugDisabled(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Debug("message", nil)
		AssertJSON(t, b.String(), "")
	}
}

func TestDebugFields(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Debug("message", watermill.LogFields{"foo": "bar"})
		AssertJSON(t, b.String(), `{"message":"message","level":"debug","foo":"bar"}`)
	}
}

func TestTrace(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Trace("message", nil)
		AssertJSON(t, b.String(), `{"message":"message","level":"trace"}`)
	}
}

func TestTraceDisabled(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Trace("message", nil)
		AssertJSON(t, b.String(), "")
	}
}

func TestTraceFields(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		l.Trace("message", watermill.LogFields{"foo": "bar"})
		AssertJSON(t, b.String(), `{"message":"message","level":"trace","foo":"bar"}`)
	}
}

func TestWith(b *bytes.Buffer, l watermill.LoggerAdapter) func(t *testing.T) {
	return func(t *testing.T) {
		w := l.With(watermill.LogFields{"foo": "bar", "baz": "bif"})
		w.Info("message", nil)
		AssertJSON(t, b.String(), `{"message":"message","level":"info","foo":"bar","baz":"bif"}`)
	}
}
