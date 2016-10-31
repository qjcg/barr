// The barr command prints out a status line (e.g. for use with dwm(1)).
package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	barr "github.com/qjcg/barr/lib"
)

func main() {
	freq := flag.Duration("f", time.Second*5, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	testMode := flag.Bool("t", false, "test mode")
	flag.Parse()

	// We will append to stringers all fmt.Stringers we want to compose
	// together to produce our final status string.
	var stringers []fmt.Stringer

	// Wifi.
	var wd barr.WifiData
	stringers = append(stringers, &wd)

	// Battery.
	bat, err := barr.NewBattery(barr.BatDir)
	if err == nil {
		stringers = append(stringers, bat)
	}

	// Load average.
	stringers = append(stringers, &barr.LoadAvg{})

	// Timestamp.
	ts := barr.DefaultTimeStamp
	if *testMode {
		ts = barr.TestTimeStamp
	}
	stringers = append(stringers, ts)

	// Set ticker frequency.
	ticker := time.NewTicker(*freq)
	if *testMode {
		ticker = time.NewTicker(time.Second)
	}

	// Loop and update.
	var output string
	for range ticker.C {
		output = Get(*ofs, stringers)

		if *testMode {
			fmt.Printf("\r%s ", output)
			continue
		}

		err := exec.Command("xsetroot", "-name", output).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Get returns a status string from an OFS and list of fmt.Stringer interface
// values.
func Get(ofs string, bs []fmt.Stringer) string {
	var data []string
	var output string

	for _, b := range bs {
		data = append(data, b.String())
	}

	output = strings.Join(data, ofs)
	output = strings.Trim(output, " ")

	return output
}
