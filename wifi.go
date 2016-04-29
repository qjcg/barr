package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type WifiData struct {
	ESSID   string
	Quality float64
}

func getESSID() (string, error) {
	cmd := exec.Command("iwgetid", "--raw")
	cmdOut, err := cmd.Output()
	if err != nil {
		return "", err
	}

	essid := strings.Trim(string(cmdOut), "\n")
	return essid, nil
}

// TODO: use /proc/net/wireless
func getQuality(ifname string) (float64, error) {
	return 0.0, nil
}

func NewWifiData(ifname string) (*WifiData, error) {
	return nil, nil
}

func (w *WifiData) Str() string {
	return fmt.Sprintf("%s:%d%%", w.ESSID, w.Quality*100)
}
