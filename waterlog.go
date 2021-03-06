package waterlog

import (
	"io"

	"github.com/ThreeDotsLabs/watermill"
)

// NewFunc matches constructors for every logger implementation
type NewFunc func(output io.Writer, debug, trace bool) watermill.LoggerAdapter

// Opts controls global features across all loggers
type Opts struct {
	Debug bool
	Trace bool
}
