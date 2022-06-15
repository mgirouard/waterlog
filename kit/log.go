package kit

import (
	"io"

	"github.com/ThreeDotsLabs/watermill"
	log "github.com/go-kit/log"
)

func New(output io.Writer, debug, trace bool) watermill.LoggerAdapter {
	return Logger{log.NewJSONLogger(output), debug, trace}
}

type Logger struct {
	L     log.Logger
	debug bool
	trace bool
}

func (l Logger) Log(keyvals ...interface{}) error {
	return l.L.Log(keyvals...)
}

func (l Logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = lfs(fields)
	fields["level"] = "error"
	fields["message"] = msg
	fields["error"] = err
	l.Log(kvs(fields)...)
}

func (l Logger) Info(msg string, fields watermill.LogFields) {
	fields = lfs(fields)
	fields["level"] = "info"
	fields["message"] = msg
	l.Log(kvs(fields)...)
}

func (l Logger) Debug(msg string, fields watermill.LogFields) {
	if !l.debug {
		return
	}
	fields = lfs(fields)
	fields["level"] = "debug"
	fields["message"] = msg
	l.Log(kvs(fields)...)
}

func (l Logger) Trace(msg string, fields watermill.LogFields) {
	if !l.trace {
		return
	}
	fields = lfs(fields)
	fields["level"] = "trace"
	fields["message"] = msg
	l.Log(kvs(fields)...)
}

func (l Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{log.With(l.L, kvs(fields)...), l.debug, l.trace}
}

func kvs(m map[string]interface{}) []interface{} {
	kv := []interface{}{}
	for k, v := range m {
		kv = append(kv, k, v)
	}
	return kv
}

func lfs(fields watermill.LogFields, keyvals ...interface{}) watermill.LogFields {
	if fields == nil {
		fields = watermill.LogFields{}
	}
	return fields
}
