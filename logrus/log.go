package logrus

import (
	"io"

	"git.sr.ht/~mgirouard/waterlog"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/sirupsen/logrus"
)

func New(output io.Writer, debug, trace bool) watermill.LoggerAdapter {
	fmt := &logrus.JSONFormatter{
		DisableTimestamp: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}
	return Logger{
		L: &logrus.Logger{
			Out:       output,
			Formatter: fmt,
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.TraceLevel,
		},
		Opts: waterlog.Opts{debug, trace},
	}
}

type Logger struct {
	L      *logrus.Logger
	fields logrus.Fields
	waterlog.Opts
}

func (l Logger) Error(msg string, err error, fields watermill.LogFields) {
	l.L.
		WithFields(l.fields).
		WithFields(logrus.Fields(fields)).
		WithField("error", err).
		Error(msg)
}
func (l Logger) Info(msg string, fields watermill.LogFields) {
	l.L.
		WithFields(l.fields).
		WithFields(logrus.Fields(fields)).
		Info(msg)
}
func (l Logger) Debug(msg string, fields watermill.LogFields) {
	if !l.Opts.Debug {
		return
	}
	l.L.
		WithFields(l.fields).
		WithFields(logrus.Fields(fields)).
		Debug(msg)
}
func (l Logger) Trace(msg string, fields watermill.LogFields) {
	if !l.Opts.Trace {
		return
	}
	l.L.
		WithFields(l.fields).
		WithFields(logrus.Fields(fields)).
		Trace(msg)
}
func (l Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		L:      l.L,
		fields: logrus.Fields(fields),
		Opts:   l.Opts,
	}
}
