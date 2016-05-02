/*The barr command prints out a status line (e.g. for use with dwm(1)).

Functions return strings usable in a statusbar context.
*/
package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// StrFuncs returns a joined string from an OFS and list of functions returning strings.
func StrFuncs(ofs string, fns ...func() string) string {
	var data []string
	for _, f := range fns {
		data = append(data, f())
	}
	return strings.Join(data, ofs)
}

func emptyString() string {
	return ""
}

func main() {
	batdir := flag.String("b", fFirstBatDir, "base directory for battery info")
	freq := flag.Duration("f", time.Second*5, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	wifiIface := flag.String("w", "", "wifi card interface name")
	testMode := flag.Bool("t", false, "test mode")
	flag.Parse()

	tsfmt := "Mon Jan 2, 3:04pm"
	if *testMode {
		*freq = time.Millisecond
		tsfmt = "Mon Jan 2, 3:04:05.000pm"
	}

	// FIXME: The below two stanzas are hideous. Refactor so that all bar
	// items use an interface with predictable behaviour.

	batFn := emptyString
	b, err := NewBattery(*batdir)
	if err == nil {
		batFn = b.Str
	}

	wifiFn := emptyString
	if *wifiIface != "" {
		wd, err := NewWifiData(*wifiIface)
		if err == nil {
			wifiFn = wd.Str
		}
	}

	var data []string
	var output string
	ticker := time.NewTicker(*freq)
	for t := range ticker.C {
		data = []string{
			wifiFn(),
			batFn(),
			LoadAvg(),
			t.Format(tsfmt),
		}
		output = strings.Join(data, *ofs)
		output = strings.Trim(output, " ")

		if *testMode {
			fmt.Printf("\r%s ", output)
		} else {
			_ = exec.Command("xsetroot", "-name", output).Run()
		}
	}
}
