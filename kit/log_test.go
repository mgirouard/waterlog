package kit

import (
	"bytes"
	"errors"
	"sort"
	"strings"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-kit/log"
)

func TestError(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Error("message", errors.New("error"), nil)
	assertSame(t, b.String(), "err=error msg=message")
}

func TestErrorLogFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Error("message", errors.New("error"), watermill.LogFields{"foo": "bar"})
	assertSame(t, b.String(), "err=error foo=bar msg=message")
}

func TestInfo(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Info("message", nil)
	assertSame(t, b.String(), "msg=message")
}

func TestInfoLogFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Info("message", watermill.LogFields{"foo": "bar"})
	assertSame(t, b.String(), "foo=bar msg=message")
}

func TestDebug(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b), debug: true}
	l.Debug("message", nil)
	assertSame(t, b.String(), "msg=message")
}

func TestDebugDisabled(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Debug("message", nil)
	assertSame(t, b.String(), "")
}

func TestDebugFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b), debug: true}
	l.Debug("message", watermill.LogFields{"foo": "bar"})
	assertSame(t, b.String(), "msg=message foo=bar")
}

func TestTrace(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b), trace: true}
	l.Trace("message", nil)
	assertSame(t, b.String(), "msg=message")
}

func TestTraceDisabled(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	l.Trace("message", nil)
	assertSame(t, b.String(), "")
}

func TestTraceFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b), trace: true}
	l.Trace("message", watermill.LogFields{"foo": "bar"})
	assertSame(t, b.String(), "msg=message foo=bar")
}

func TestWith(t *testing.T) {
	b := &bytes.Buffer{}
	l := Logger{L: log.NewLogfmtLogger(b)}
	w := l.With(watermill.LogFields{"foo": "bar", "baz": "bif"})
	w.Info("message", nil)
	assertSame(t, b.String(), "baz=bif foo=bar msg=message")
}

// assertSame is is a lazy and unsophisticated assertion for log lines
// it's not smart enough to handle values with spaces
func assertSame(t *testing.T, act, exp string) {
	a := strings.Split(strings.Trim(act,"\n"), " ")
	e := strings.Split(strings.Trim(exp,"\n"), " ")
	sort.Strings(a)
	sort.Strings(e)
	if strings.Join(a, " ") != strings.Join(e, " ") {
		t.Fatalf("expected %q but got %q", exp, act)
	}
}
