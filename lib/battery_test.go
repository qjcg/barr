// +build integration

// NOTE: Tests in this file require that a battery be present on
// the test system to work (hence the "integration" build tag).
package barr

import (
	"testing"
)

func TestNewBattery(t *testing.T) {
	testdata := []struct {
		dir                   string
		chargeNow, chargeFull float64
	}{
		{"/sys/class/power_supply/BAT0", 0.0, 0.0},
		{"/sys/class/power_supply/BAT0", 1.0, 2.0},
	}

	for _, s := range testdata {
		d, err := NewBattery(s.dir)
		if err != nil {
			t.Error(d, err)
		}
	}
}

func TestBatteryString(t *testing.T) {
	b := &Battery{
		Dir: "/sys/class/power_supply/BAT0",
	}
	if b.String() == "" {
		t.Error("battery String() method returns empty string")
	}
}
