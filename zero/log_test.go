package zero

import (
	"testing"

	"git.sr.ht/~mgirouard/waterlog"
)

func TestKit(t *testing.T) {
	waterlog.TestLogger(t, New)
}
