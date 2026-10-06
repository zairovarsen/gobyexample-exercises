//go:build solution

package epoch

import "time"

func Stamps(t time.Time) (sec, milli, nano int64) {
	return t.Unix(), t.UnixMilli(), t.UnixNano()
}

func FromSeconds(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func FromNanos(nano int64) time.Time {
	return time.Unix(0, nano)
}

func FromMillis(ms int64) time.Time {
	return time.Unix(ms/1000, ms%1000*int64(time.Millisecond))
}

func Elapsed(startSec int64, now time.Time) time.Duration {
	return now.Sub(time.Unix(startSec, 0))
}

func ParseTimestamp(ts int64) time.Time {
	switch {
	case ts < 1e11:
		return time.Unix(ts, 0)
	case ts < 1e14:
		return time.UnixMilli(ts)
	case ts < 1e17:
		return time.UnixMicro(ts)
	}
	return time.Unix(0, ts)
}
