package barr

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const wifiDataFile = "/proc/net/wireless"

var (
	wifiQuality = regexp.MustCompile(`\w:\s+\d+\s+(\d+)\.`)
)

type WifiData struct {
	ESSID   string
	Ifname  string
	Quality int
}

// Implement the fmt.Stringer interface.
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

// getConnection updates the w.Interface and w.ESSID values based on the output
// of the "iwgetid" command.
func (w *WifiData) getConnection() error {
	out, err := exec.Command("iwgetid").Output()
	if e, ok := err.(*exec.ExitError); ok {
		return e
	}

	if err != nil {
		return err
	}

	data := strings.Split(string(out), `ESSID:"`)

	w.Ifname = strings.TrimSpace(data[0])
	data[1] = strings.TrimSpace(data[1])
	w.ESSID = strings.Trim(data[1], `"`)

	return nil
}

func (w *WifiData) getQuality() error {
	file, err := os.Open(wifiDataFile)
	if err != nil {
		return err
	}

	fscanner := bufio.NewScanner(file)

	var line int
	for fscanner.Scan() {
		// skip initial header lines
		if line++; line < 3 {
			continue
		}
		lineBytes := fscanner.Bytes()

		// only interested in lines containing "ifname"
		m, err := regexp.Match(w.Ifname, lineBytes)
		if !m {
			return nil
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
