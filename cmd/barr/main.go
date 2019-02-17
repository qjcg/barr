// The barr command prints out a status line for use with minimalistic window managers.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/qjcg/barr/pkg/sysinfo"
)

// StatusBar describes a statusbar.
type StatusBar struct {
	Stringers []fmt.Stringer
}

var separator string

func main() {
	flag.StringVar(&separator, "s", "  ", "output field separator")

	version := flag.Bool("v", false, "print version")
	xSetRootMode := flag.Bool("x", false, "xsetroot mode (loop and update)")
	xSetRootModeFreq := flag.Duration("xf", time.Second*5, "xsetroot mode update frequency")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	// Create a new StatusBar.
	sb := StatusBar{
		Stringers: []fmt.Stringer{
			&sysinfo.WifiData{},
			&sysinfo.Battery{},
			&sysinfo.Disk{Dir: "/"},
			// FIXME: Not working! Enable when fixed.
			//&sysinfo.CryptoCurrency{Pair: "xbtcad"},
			&sysinfo.LoadAvg{},
			&sysinfo.DefaultTimeStamp,
		},
	}

	if *xSetRootMode {

		// Clear screen and exit when interrupt signal received.
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		go func(c chan os.Signal) {
			for range c {
				exec.Command("/usr/bin/xsetroot", "-name", "barr stopped").Run()
				os.Exit(0)
			}
		}(c)

		// Print once right away.
		err := sb.UpdateXRootWindowTitle()
		if err != nil {
			log.Fatal(err)
		}

		// After printing once above, loop and update.
		// See https://github.com/golang/go/issues/17601
		ticker := time.NewTicker(*xSetRootModeFreq)
		for range ticker.C {
			err := sb.UpdateXRootWindowTitle()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println(sb.Get())
	}
}

// Get returns a status string.
func (sb *StatusBar) Get() string {
	var fields []string
	var output string

	for _, s := range sb.Stringers {
		str := s.String()
		if str == "" {
			continue
		}
		fields = append(fields, str)
	}

	output = strings.Join(fields, separator)
	output = strings.Trim(output, " ")

	return output
}

// UpdateXRootWindowTitle sets the X root window title.
// In dwm, this is also the way to set the WM status string.
func (sb *StatusBar) UpdateXRootWindowTitle() error {
	err := exec.Command("/usr/bin/xsetroot", "-name", sb.Get()).Run()
	if err != nil {
		return err
	}
	return nil
}
