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

// Get current work span of day.
// Zero value will be returned when day does not has any spans yet.
// In that case the returned bool will be false
func (d *Day) CurrentSpan() (Span, bool) {
	if len(d.Spans) == 0 {
		return Span{}, false
	}

	return d.Spans[len(d.Spans) - 1], true
}

