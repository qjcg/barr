// The barr command prints out a system status line.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/qjcg/barr/pkg/sysinfo"
)

var separator string

func main() {
	flag.StringVar(&separator, "s", "  ", "output field separator")
	version := flag.Bool("v", false, "print version")
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

	fmt.Println(sb.Get())
}

// StatusBar describes a statusbar.
type StatusBar struct {
	Stringers []fmt.Stringer
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
