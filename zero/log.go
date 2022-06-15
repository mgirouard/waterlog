package zero

import (
	"io"

	"git.sr.ht/~mgirouard/waterlog"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
)

func New(output io.Writer, debug, trace bool) watermill.LoggerAdapter {
	return Logger{
		L:    zerolog.New(output),
		Opts: waterlog.Opts{debug, trace},
	}
}

type Logger struct {
	L zerolog.Logger
	waterlog.Opts
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
	if !l.Opts.Debug {
		return
	}
	ev := l.L.Debug()
	for k, v := range fields {
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}
func (l Logger) Trace(msg string, fields watermill.LogFields) {
	if !l.Opts.Trace {
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
	return &Logger{
		L:    ctx.Logger(),
		Opts: l.Opts,
	}
}
