// The barr command prints out a status line (e.g. for use with dwm(1)).
package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"

	barr "github.com/qjcg/barr/lib"
)

// Get returns a status string from an OFS and list of BarStringers.
func Get(ofs string, bs []*BarStringer) string {
	var data []string
	var output string

	for _, b := range bs {
		b.Update()
		data = append(data, b.Str())
	}

	output = strings.Join(data, ofs)
	output = strings.Trim(output, " ")

	return output
}

func main() {
	batdir := flag.String("b", barr.fFirstBatDir, "base directory for battery info")
	freq := flag.Duration("f", barr.freqNormal, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	wifiIface := flag.String("w", "", "wifi card interface name")
	testMode := flag.Bool("t", false, "test mode")
	flag.Parse()

	// append to BarStringers all we want to compose together as output
	var bstrs []*barr.BarStringer

	bstrs = append(bstrs, &barr.Battery{})

	bstrs = append(bstrs, &barr.LoadAvg{})

	if *wifiIface != "" {
		bstrs = append(bstrs, &barr.WifiData{Iface: *wifiIface})
	}

	if *testMode {
		*freq = barr.freqTest
		tsfmt = "Mon Jan 2, 3:04:05.000pm"
	}

	var output string
	ticker := time.NewTicker(*freq)
	for t := range ticker.C {
		output = Get(*ofs, bs)

		if *testMode {
			fmt.Printf("\r%s ", output)
		} else {
			_ = exec.Command("xsetroot", "-name", output).Run()
		}
	}
}
