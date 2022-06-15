package logrus

import (
	"testing"

	"git.sr.ht/~mgirouard/waterlog"
)

func TestLogrus(t *testing.T) {
	waterlog.TestLogger(t, New)
}
