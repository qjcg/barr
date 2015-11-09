/* The barr command prints out a status line (e.g. for use with dwm(1)).

Functions return strings usable in a statusbar context.
*/
package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

// The LoopFunc function allows functions to continuously return values to a channel at a specified delay.
func LoopFunc(f func() string, c chan<- string, delay time.Duration) {
	ticker := time.NewTicker(delay)
	for range ticker.C {
		c <- f()
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

	go LoopFunc(Battery, batchan, time.Second*10)
	go LoopFunc(LoadAvg, loadchan, time.Second*5)

	for {
		select {
		case bat := <-batchan:
			fmt.Printf("%s", bat) //DEBUG
			<-update
			fallthrough
		case load := <-loadchan:
			fmt.Printf("%s", load) //DEBUG
			<-update
			fallthrough
		case <-update:
			data := []string{bat, load}
			dataStr := strings.Join(data, *ofs)
			fmt.Printf("\r%s ", dataStr)
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
