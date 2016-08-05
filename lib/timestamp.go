package barr

import (
	"time"
)

const (
	// Normal mode
	FmtNormal  = "Mon Jan 2, 3:04pm"
	FreqNormal = time.Second * 5

	// Test mode (more frequent updates)
	FmtTest  = "Mon Jan 2, 3:04:05.000pm"
	FreqTest = time.Millisecond
)

type TimeStamp struct {
	Fmt  string
	Freq time.Duration
	time
}

// Implement the BarStringer interface.
func (ts *TimeStamp) Str() string {
	return time.Format(ts.Fmt)
}

// Implement the BarStringer interface.
func (ts *TimeStamp) Update() error {
	return nil
}
