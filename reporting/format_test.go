package reporting

import (
	"testing"
	"time"
)

func TestFmtDuration(t *testing.T) {
	testDuration, _ := time.ParseDuration("1h15m10s")
	want := "01h 15m 10s"
	got := fmtDuration(testDuration)

	if want != got {
		t.Errorf("fmtDuration failed, got %v, want %v", got, want)
	}
}

func TestFmtTime(t *testing.T) {
	// 12:00 UTC equals 14:00 Europe/Berlin (+02:00) during daylight saving time
	want := "14:30:10"
	got := fmtTime(time.Date(2021, time.September, 1, 12, 30, 10, 0, time.UTC))

	if want != got {
		t.Errorf("fmtTime failed, got %v, want %v", got, want)
	}
}
