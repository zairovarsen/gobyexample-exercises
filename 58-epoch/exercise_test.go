package epoch

import (
	"testing"
	"time"
)

// Fixed moments in UTC, so that no test depends on the clock or the time
// zone of the machine it runs on.
var (
	epochStart = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	// The moment from the gobyexample page.
	halloween = time.Date(2012, time.October, 31, 16, 13, 58, 292387000, time.UTC)
	// The same moment cut off at full seconds and full milliseconds.
	halloweenSec   = time.Date(2012, time.October, 31, 16, 13, 58, 0, time.UTC)
	halloweenMilli = time.Date(2012, time.October, 31, 16, 13, 58, 292000000, time.UTC)
)

func TestStamps(t *testing.T) {
	// The same instant in another time zone has the same Unix time.
	tokyo := halloween.In(time.FixedZone("UTC+9", 9*60*60))
	tests := []struct {
		name                    string
		t                       time.Time
		wantSec, wantMs, wantNs int64
	}{
		{"2012-10-31 16:13:58.292387 UTC", halloween, 1351700038, 1351700038292, 1351700038292387000},
		{"2012-11-01 01:13:58.292387 +09:00", tokyo, 1351700038, 1351700038292, 1351700038292387000},
		{"one day after the epoch", epochStart.Add(24 * time.Hour), 86400, 86400000, 86400000000000},
		{"1.5 seconds before the epoch", epochStart.Add(-1500 * time.Millisecond), -2, -1500, -1500000000},
	}
	for _, tt := range tests {
		sec, ms, ns := Stamps(tt.t)
		if sec != tt.wantSec || ms != tt.wantMs || ns != tt.wantNs {
			t.Errorf("Stamps(%s) = %d, %d, %d; want %d, %d, %d",
				tt.name, sec, ms, ns, tt.wantSec, tt.wantMs, tt.wantNs)
		}
	}
}

func TestFromSeconds(t *testing.T) {
	tests := []struct {
		sec  int64
		want time.Time
	}{
		{1351700038, halloweenSec},
		{86400, epochStart.Add(24 * time.Hour)},
		{1000000000, time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)},
		{-3600, epochStart.Add(-time.Hour)},
	}
	for _, tt := range tests {
		if got := FromSeconds(tt.sec); !got.Equal(tt.want) {
			t.Errorf("FromSeconds(%d) = %v, want %v", tt.sec, got.UTC(), tt.want)
		}
	}
}

func TestFromNanos(t *testing.T) {
	tests := []struct {
		nano int64
		want time.Time
	}{
		{1351700038292387000, halloween},
		{1351700038000000000, halloweenSec},
		{1, epochStart.Add(time.Nanosecond)},
		{1500000000, epochStart.Add(1500 * time.Millisecond)},
	}
	for _, tt := range tests {
		if got := FromNanos(tt.nano); !got.Equal(tt.want) {
			t.Errorf("FromNanos(%d) = %v, want %v", tt.nano, got.UTC(), tt.want)
		}
	}
}

func TestFromMillis(t *testing.T) {
	tests := []struct {
		ms   int64
		want time.Time
	}{
		{1351700038292, halloweenMilli},
		{1351700038000, halloweenSec},
		{1, epochStart.Add(time.Millisecond)},
		{86400000, epochStart.Add(24 * time.Hour)},
		{1999, epochStart.Add(1999 * time.Millisecond)},
	}
	for _, tt := range tests {
		if got := FromMillis(tt.ms); !got.Equal(tt.want) {
			t.Errorf("FromMillis(%d) = %v, want %v", tt.ms, got.UTC(), tt.want)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	sec, ms, ns := Stamps(halloween)
	if got := FromSeconds(sec); !got.Equal(halloweenSec) {
		t.Errorf("FromSeconds(seconds of %v) = %v, want %v", halloween, got.UTC(), halloweenSec)
	}
	if got := FromMillis(ms); !got.Equal(halloweenMilli) {
		t.Errorf("FromMillis(milliseconds of %v) = %v, want %v", halloween, got.UTC(), halloweenMilli)
	}
	if got := FromNanos(ns); !got.Equal(halloween) {
		t.Errorf("FromNanos(nanoseconds of %v) = %v, want %v", halloween, got.UTC(), halloween)
	}
}

func TestElapsed(t *testing.T) {
	tests := []struct {
		startSec int64
		now      time.Time
		want     time.Duration
	}{
		{1351700038, halloweenSec.Add(70 * time.Second), 70 * time.Second},
		{1351700038, halloween, 292387 * time.Microsecond},
		{1351700038, halloweenSec.Add(-time.Hour), -time.Hour},
		{0, epochStart.Add(36 * time.Hour), 36 * time.Hour},
	}
	for _, tt := range tests {
		if got := Elapsed(tt.startSec, tt.now); got != tt.want {
			t.Errorf("Elapsed(%d, %v) = %v, want %v", tt.startSec, tt.now, got, tt.want)
		}
	}
}

func TestParseTimestamp(t *testing.T) {
	y2k := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		ts   int64
		want time.Time
	}{
		{1351700038, halloweenSec},
		{1351700038292, halloweenMilli},
		{1351700038292387, halloween},
		{1351700038292387000, halloween},
		{1351700038000, halloweenSec},
		{1351700038000000, halloweenSec},
		{1351700038000000000, halloweenSec},
		{946684800, y2k},
		{946684800000, y2k},
		{946684800000000, y2k},
		{946684800000000000, y2k},
		{60, epochStart.Add(time.Minute)},
	}
	for _, tt := range tests {
		if got := ParseTimestamp(tt.ts); !got.Equal(tt.want) {
			t.Errorf("ParseTimestamp(%d) = %v, want %v", tt.ts, got.UTC(), tt.want)
		}
	}
}
