package sysinfo

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	wifiQuality = regexp.MustCompile(`\w:\s+\d+\s+(\d+)\.`)
)

type WifiData struct {
	ESSID   string
	Ifname  string
	Quality int
}

// String implements the fmt.Stringer interface.
func (w *WifiData) String() string {
	err := w.getConnection()
	if err != nil {
		return ""
	}

	err = w.getQuality()
	if err != nil {
		return ""
	}

	if w.ESSID == "" || w.Quality == 0 {
		return ""
	}
	return fmt.Sprintf("%s:%.0f%%", w.ESSID, float64(w.Quality)/70*100)
}

// getConnection updates the w.Ifname and w.ESSID values based on the output
// of the "iw dev" command.
// NOTE: "iw" help output says: "Do NOT screenscrape this tool, we don't
// consider its output stable." Until a better solution comes along, we'll
// screenscrape it anyway!
func (w *WifiData) getConnection() error {
	out, err := exec.Command("iw", "dev").Output()
	if e, ok := err.(*exec.ExitError); ok {
		return e
	}
	if err != nil {
		return err
	}
	r := bytes.NewReader(out)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Interface") {
			w.Ifname = strings.Fields(line)[1]
		} else if strings.Contains(line, "ssid") {
			w.ESSID = strings.Fields(line)[1]
		}
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (w *WifiData) getQuality() error {
	file, err := os.Open("/proc/net/wireless")
	if err != nil {
		return err
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	var line int
	for fscanner.Scan() {

		// Skip initial header lines.
		if line++; line < 3 {
			continue
		}
		lineBytes := fscanner.Bytes()

		// Only interested in lines containing "ifname".
		m, err := regexp.Match(w.Ifname, lineBytes)
		if !m {
			continue
		}

		result := wifiQuality.FindSubmatch(lineBytes)
		if result != nil {
			w.Quality, err = strconv.Atoi(string(result[1]))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
