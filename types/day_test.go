package types

import (
	"testing"
	"time"
)

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
		t.Errorf("Should not be in")
	}
}

func TestIsOnBreak(t *testing.T) {
	day := Day{}
	if day.IsOnBreak() {
		t.Errorf("Should not be on break")
	}
}

func TestIsOccupied(t *testing.T) {
	day := Day{}
	if day.IsOccupied() {
		t.Errorf("Should not be occupied")
	}
}
