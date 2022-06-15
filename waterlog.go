package waterlog

import (
	"io"

	"github.com/ThreeDotsLabs/watermill"
)

// NewFunc matches constructors for every logger implementation
type NewFunc func(output io.Writer, debug, trace bool) watermill.LoggerAdapter
