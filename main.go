// The barr command prints out a status line (e.g. for use with dwm(1)).
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	barr "github.com/qjcg/barr/lib"
)

var Version string

func main() {
	freq := flag.Duration("f", time.Second*5, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	testMode := flag.Bool("t", false, "test mode")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	// We will append to stringers all fmt.Stringers we want to compose
	// together to produce our final status string.
	var stringers []fmt.Stringer

	// Wifi.
	var wd barr.WifiData
	stringers = append(stringers, &wd)

	// Battery.
	// We get live battery capacities from sysfs.
	var bat barr.Battery

	matches, err := filepath.Glob("/sys/class/power_supply/BAT*/capacity")
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range matches {
		f, err := os.Open(m)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		bat.Sources = append(bat.Sources, f)
	}

	stringers = append(stringers, &bat)

	// Load average.
	var loadavg barr.LoadAvg
	stringers = append(stringers, &loadavg)

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

		// Setting X root window title sets dwm status string.
		err := exec.Command("xsetroot", "-name", output).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Get returns a status string from an OFS and list of fmt.Stringer interface
// values.
func Get(ofs string, stringers []fmt.Stringer) string {
	var data []string
	var output string

	for _, s := range stringers {
		data = append(data, s.String())
	}

	output = strings.Join(data, ofs)
	output = strings.Trim(output, " ")

	return output
}
