package zero

import (
	"testing"

	"git.sr.ht/~mgirouard/waterlog"
)

func TestZero(t *testing.T) {
	waterlog.TestLogger(t, New)
}
