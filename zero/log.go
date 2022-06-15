package zero

import (
	"io"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
)

func New(output io.Writer, debug, trace bool) watermill.LoggerAdapter {
	return Logger{zerolog.New(output), debug, trace}
}

type Logger struct {
	L     zerolog.Logger
	debug bool
	trace bool
}

func (l Logger) Error(msg string, err error, fields watermill.LogFields) {
	ev := l.L.Error().Err(err)
	for k, v := range fields {
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}
func (l Logger) Info(msg string, fields watermill.LogFields) {
	ev := l.L.Info()
	for k, v := range fields {
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}
func (l Logger) Debug(msg string, fields watermill.LogFields) {
	if !l.debug {
		return
	}
	ev := l.L.Debug()
	for k, v := range fields {
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}
func (l Logger) Trace(msg string, fields watermill.LogFields) {
	if !l.trace {
		return
	}
	ev := l.L.Trace()
	for k, v := range fields {
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}
func (l Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	ctx := l.L.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	return &Logger{ctx.Logger(), l.debug, l.trace}
}
