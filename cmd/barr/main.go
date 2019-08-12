// The barr command prints out a system status line.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/qjcg/barr/pkg/sysinfo"
)

func main() {
	flagSeparator := flag.String("s", "  ", "output field separator")
	flagVersion := flag.Bool("v", false, "print version")
	flag.Parse()

	if *flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	// Create a new StatusBar.
	// TODO: Accept config from environment and/or config file as well as
	// defining a default.
	sb := StatusBar{
		Separator: *flagSeparator,
		Stringers: []fmt.Stringer{
			&sysinfo.WifiData{},
			&sysinfo.Battery{},
			&sysinfo.Disk{Dir: "/"},
			// FIXME: Not working! Enable when fixed.
			//&sysinfo.CryptoCurrency{Pair: "xbtcad"},
			&sysinfo.LoadAvg{},
			&sysinfo.Volume{},
			&sysinfo.DefaultTimeStamp,
		},
	}

	fmt.Println(sb.Get())
}

// StatusBar describes a statusbar.
type StatusBar struct {
	Separator string
	Stringers []fmt.Stringer
}

// Get returns a status string.
func (sb *StatusBar) Get() string {
	var fields []string
	for _, s := range sb.Stringers {
		str := s.String()
		// If a Stringer returns the empty string, it's skipped.
		if str == "" {
			continue
		}
		fields = append(fields, strings.TrimSpace(str))
	}

	return strings.Join(fields, sb.Separator)
}
