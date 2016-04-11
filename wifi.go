package main

import (
	"fmt"
	"os/exec"
)

type WifiData struct {
	ESSID   string
	Quality float64
}

func getESSID() (string, error) {
	cmd := exec.Command("iwgetid", "--raw")
	cmdOut, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return string(cmdOut), nil
}

// TODO: use /proc/net/wireless
func getQuality(ifname string) (float64, error) {
}

func NewWifiData(ifname string) (*WifiData, error) {
}

func (w *WifiData) Str() string {
	return fmt.Sprintf("%s:%d%%", w.ESSID, w.Quality*100)
}
