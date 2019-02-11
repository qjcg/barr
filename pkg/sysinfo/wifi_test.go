package sysinfo

import "testing"

func TestGetESSID(t *testing.T) {
	var w WifiData
	err := w.getESSID()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ESSID: %s\n", w.ESSID)
}
