// +build integration

// NOTE: Tests in this file require that a battery be present on
// the test system to work (hence the "integration" build tag).
package barr

import (
	"testing"
)

func TestNewBattery(t *testing.T) {
	b, err := NewBattery()
	if err != nil {
		t.Error(b, err)
	}
}

func TestBatteryString(t *testing.T) {
	b, err := NewBattery()
	if err != nil {
		t.Error(b, err)
	}
	if b.String() == "" {
		t.Error("battery String() method returns empty string")
	}
}

func TestCapacity(t *testing.T) {
}
