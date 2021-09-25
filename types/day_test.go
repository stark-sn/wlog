package types

import (
	"testing"
	"time"
)

var testTime = time.Date(2020, time.December, 25, 8, 0, 0, 0, time.UTC)

func TestDate(t *testing.T) {
	d := time.Date(2020, time.December, 24, 6, 0, 0, 0, time.UTC)
	got := Date(d)
	want := "2020-12-24"

	if got != want {
		t.Errorf("Date failed, got %v, want %v", got, want)
	}
}

func TestIsIn(t *testing.T) {
	day := Day{}
	if day.IsIn() {
		t.Error("Should not be in")
	}
}

func TestIsOnBreak(t *testing.T) {
	day := Day{}
	if day.IsOnBreak() {
		t.Error("Should not be on break")
	}
}

func TestIsOccupied(t *testing.T) {
	day := Day{}
	if day.IsOccupied() {
		t.Error("Should not be occupied")
	}
}

func TestSpans(t *testing.T) {
	day := Day{
		Spans: []Span{
			Span{},
		},
	}

	if len(day.GetSpans(testTime)) != 1 {
		t.Error("Incorrect number of spans returned")
	}

	day.CurSpan = &Span{}

	if len(day.GetSpans(testTime)) != 2 {
		t.Error("Current span not returned from GetSpans")
	}
}

func TestBreaks(t *testing.T) {
	day := Day{
		Breaks: []Span{
			Span{},
		},
	}

	if len(day.GetBreaks(testTime)) != 1 {
		t.Error("Incorrect number of breaks returned")
	}

	day.CurBreak = &Span{}

	if len(day.GetBreaks(testTime)) != 2 {
		t.Error("Current break not returned from GetBreaks")
	}
}

func TestActivities(t *testing.T) {
	day := Day{
		Activities: []Activity{
			Activity{},
		},
	}

	if len(day.GetActivities(testTime)) != 1 {
		t.Error("Incorrect number of activities returned")
	}

	day.CurActivity = &Activity{}

	if len(day.GetActivities(testTime)) != 2 {
		t.Error("Current activity not returned from GetActivities")
	}
}
