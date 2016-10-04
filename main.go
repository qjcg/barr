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

func main() {
	//batdir := flag.String("b", barr.BatDir, "base directory for battery info")
	freq := flag.Duration("f", barr.FreqNormal, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	wifiIface := flag.String("w", "", "wifi card interface name")
	testMode := flag.Bool("t", false, "test mode")
	flag.Parse()

	// append to BarStringers all we want to compose together as output
	var bstrs []BarStringer
	//bstrs = append(bstrs, &barr.Battery{})
	bstrs = append(bstrs, &barr.LoadAvg{})

	if *wifiIface != "" {
		bstrs = append(bstrs, &barr.WifiData{Ifname: *wifiIface})
	}

	if *testMode {
		*freq = barr.FreqTest
		//tsfmt = "Mon Jan 2, 3:04:05.000pm"
	}

	var output string
	ticker := time.NewTicker(*freq)
	for range ticker.C {
		output = Get(*ofs, bstrs)

		if *testMode {
			fmt.Printf("\r%s ", output)
		} else {
			_ = exec.Command("xsetroot", "-name", output).Run()
		}
	}
}

// Get returns a status string from an OFS and list of BarStringers.
func Get(ofs string, bs []BarStringer) string {
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
