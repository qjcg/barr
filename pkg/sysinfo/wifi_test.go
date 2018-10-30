package sysinfo

import "testing"

func TestGetConnection(t *testing.T) {
	w := WifiData{}
	err := w.getConnection()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ifname: %s\n", w.Ifname)
	t.Logf("ESSID: %s\n", w.ESSID)
}
