package types

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	span := Span{
		Start: time.Date(2020, time.December, 24, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2020, time.December, 24, 0, 0, 10, 0, time.UTC),
	}

	want := int64(10_000)
	got := span.Duration().Milliseconds()

	if want != got {
		t.Errorf("Duration failed, got %v, want %v", got, want)
	}
}
