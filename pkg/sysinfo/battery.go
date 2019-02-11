package sysinfo

import (
	"fmt"
	"strconv"

	"github.com/distatus/battery"
)

// Battery represents battery information.
type Battery struct{}

// String returns battery info in a textual format.
func (b *Battery) String() string {
	batteries, err := battery.GetAll()
	if err != nil {
		return ""
	}

	if len(batteries) == 0 {
		return ""
	}

	var charging bool
	var capsCur, capsFull float64
	for _, bat := range batteries {
		capsCur += bat.Current
		capsFull += bat.Full

		if bat.State == battery.Charging || bat.State == battery.Full {
			charging = true
		}
	}
	pctBatRemaining := capsCur / capsFull * 100

	if charging {
		return fmt.Sprintf("AC %s%%", strconv.FormatFloat(pctBatRemaining, 'f', 0, 64))
	} else {
		return fmt.Sprintf("%s%%", strconv.FormatFloat(pctBatRemaining, 'f', 0, 64))
	}
}
