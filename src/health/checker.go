package health

import (
	"errors"
	"time"
)

// ErrTimeout indicates health check timed out.
var ErrTimeout = errors.New("Health check timed out")

// Checker represents a checker of specific protocol.
type Checker interface {
	Check(string, time.Duration) error
}
