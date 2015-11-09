/* The barr command prints out a status line (e.g. for use with dwm(1)).

Functions return strings usable in a statusbar context.
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

// The loopStatus function sends status strings to a channel at a specified interval.
func loopStatus(s *StatusStringer, c chan<- string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c <- s.Str()
	}
}

func main() {
	batdir := flag.String("b", "/sys/class/power_supply/BAT0/", "base directory for battery info")
	freq := flag.Duration("f", time.Second*5, "update frequency")
	ofs := flag.String("s", "  ", "output field separator")
	testMode := flag.Bool("t", false, "test mode")
	flag.Parse()

	//tsfmt := "Mon Jan 2, 3:04pm"
	if *testMode {
		*freq = time.Millisecond
		//tsfmt = "Mon Jan 2, 3:04:05.000pm"
	}

	// BEGIN WIP
	batchan := make(chan string)
	loadchan := make(chan string)

	go loopStatus(Battery, batchan, time.Second*10)
	go loopStatus(LoadAvg, loadchan, time.Second*5)

	var bat, load string
	// When a message is received on a channel, update string
	for {
		select {
		case bat = <-batchan:
			fmt.Printf("%s\n", bat) //DEBUG
			<-update
			fallthrough
		case load = <-loadchan:
			fmt.Printf("%s\n", load) //DEBUG
			<-update
			fallthrough
		case <-update:
			s := &Status{
				OFS:      *ofs,
				Elements: []string{bat, load},
			}
			fmt.Printf("\r%s", s.Str())
		}
	}
	// END WIP

	//go func() {
	//	ticker := time.NewTicker(*freq)
	//	for t := range ticker.C {
	//		data := []string{
	//			Battery(),
	//			LoadAvg(),
	//			t.Format(tsfmt),
	//		}
	//		output := strings.Join(data, *ofs)

	//		if *testMode {
	//			fmt.Printf("\r%s ", output)
	//		} else {
	//			_ = exec.Command("xsetroot", "-name", output).Run()
	//		}
	//	}
	//}()

	fmt.Scanln()
}
