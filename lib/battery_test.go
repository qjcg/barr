package main

import (
	"testing"
)

func TestNewBattery(t *testing.T) {
	testdata := []struct {
		dir                   string
		chargeNow, chargeFull float64
	}{
		{"/sys/class/power_supply/BAT0"},
		{"/sys/class/power_supply/BAT0", 0.0, 0.0},
		{"/sys/class/power_supply/BAT0", 1.0, 2.0},
	}

	for _, s := range testdata {
		d := NewBattery(s.dir)
		_, ok := d.(*Battery)
		if !ok {
			t.Error("expected: *Battery type")
		}
	}
}

func TestStr(t *testing.T) {
	b := &Battery{
		Dir: "/sys/class/power_supply/BAT0",
	}
	if b.Str() == "" {
		t.Error("battery Str() method returns empty string")
	}
}
