package types

import "time"

// A day dedicated to work.
type Day struct {
	Spans []Span
	Breaks []Span
	Activities []Activity

	CurSpan *Span
	CurBreak *Span
	CurActivity *Activity
}

// Get date index string
func Date(t time.Time) string {
	return t.Format("2006-01-02")
}

func (d *Day) IsIn() bool {
	return d.CurSpan != nil
}

func (d *Day) IsOnBreak() bool {
	return d.CurBreak != nil
}

func (d *Day) IsOccupied() bool {
	return d.CurActivity != nil
}

