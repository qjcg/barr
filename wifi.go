package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	wifiDataFile = "/proc/net/wireless"
)

var (
	wifiQualityVal = regexp.MustCompile(`\w:\s+\d+\s+(\d+)\.`)
)

type WifiData struct {
	ESSID   string
	Ifname  string
	Quality int
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

func getQuality(ifname string) (int, error) {
	file, _ := os.Open(wifiDataFile)
	fscanner := bufio.NewScanner(file)

	var line, quality int

	for fscanner.Scan() {
		// skip initial header lines
		if line++; line < 3 {
			continue
		}
		lineBytes := fscanner.Bytes()

		// only interested in lines containing "ifname"
		m, err := regexp.Match(ifname, lineBytes)
		if m {
			result := wifiQualityVal.FindSubmatch(lineBytes)
			if result != nil {
				quality, err = strconv.Atoi(string(result[1]))
				if err != nil {
					return quality, err
				}
			}
		}
	}

	return quality, nil
}

func NewWifiData(ifname string) (*WifiData, error) {
	essid, err := getESSID()
	if err != nil {
		return nil, err
	}

	quality, err := getQuality(ifname)
	if err != nil {
		return nil, err
	}

	return &WifiData{ESSID: essid, Ifname: ifname, Quality: quality}, nil
}

func (w *WifiData) Update() error {
	var err error
	w.ESSID, err = getESSID()
	if err != nil {
		return err
	}

	w.Quality, err = getQuality(w.Ifname)
	if err != nil {
		return err
	}

	return nil
}

func (w *WifiData) Str() string {
	err := w.Update()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s:%.0f%%", w.ESSID, float64(w.Quality)/70*100)
}
