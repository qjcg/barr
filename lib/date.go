package barr

import (
	"time"
)

const (
	// Normal mode
	fmtNormal  = "Mon Jan 2, 3:04pm"
	freqNormal = time.Second * 5

	// Test mode (more frequent updates)
	fmtTest  = "Mon Jan 2, 3:04:05.000pm"
	freqTest = time.Millisecond
)

type Date struct {
	Fmt  string
	Freq time.Duration
}

// Implement the BarStringer interface.
func (d *Date) Str() string {
	return time.Format(d.Fmt)
}

// Implement the BarStringer interface.
func (d *Date) Update() error {
	return nil
}
