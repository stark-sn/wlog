package types

import "time"

// A day dedicated to work.
type Day struct {
	Spans []Span
	Activities []Activity
}

// Get date index string
func Date(t time.Time) string {
	return t.Format("2006-01-02")
}

