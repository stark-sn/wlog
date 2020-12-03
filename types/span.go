package types

import "time"

// A Span represents a period of time between two instants in time.
type Span struct {
	// The beginning of the period.
	Start time.Time
	// The end of the period.
	// A zero value indicates that the period is not considered
	// to be completed yet.
	End time.Time
}
