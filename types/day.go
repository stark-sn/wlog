package types

import "time"

// A day dedicated to work.
type Day struct {
	Spans      []Span
	Breaks     []Span
	Activities []Activity

	CurSpan     *Span
	CurBreak    *Span
	CurActivity *Activity
}

// Get date index string
func Date(t time.Time) string {
	return t.Format("2006-01-02")
}

func (d Day) IsIn() bool {
	return d.CurSpan != nil
}

func (d Day) IsOnBreak() bool {
	return d.CurBreak != nil
}

func (d Day) IsOccupied() bool {
	return d.CurActivity != nil
}

func (d Day) GetSpans(t time.Time) []Span {
	spans := d.Spans

	if d.IsIn() {
		currentSpan := *d.CurSpan
		currentSpan.End = t
		spans = append(spans, currentSpan)
	}

	return spans
}

func (d Day) GetBreaks(t time.Time) []Span {
	breaks := d.Breaks

	if d.IsOnBreak() {
		currentBreak := *d.CurBreak
		currentBreak.End = t
		breaks = append(breaks, currentBreak)
	}

	return breaks
}

func (d Day) GetActivities(t time.Time) []Activity {
	activities := d.Activities

	if d.IsOccupied() {
		currentActivity := *d.CurActivity
		currentActivity.End = t
		activities = append(activities, currentActivity)
	}

	return activities
}
