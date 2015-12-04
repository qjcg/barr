package main

import (
	"fmt"
	"log"
	"syscall"
)

// LoadAvg returns the load average as a string.
func LoadAvg() string {
	si := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(si)
	if err != nil {
		log.Fatal("couldn't get sysinfo:", err)
	}

	// see https://github.com/capnm/sysinfo/blob/master/sysinfo.go
	scale := 65536.0 // magic
	load1 := float64(si.Loads[0]) / scale
	load5 := float64(si.Loads[1]) / scale
	load15 := float64(si.Loads[2]) / scale

	return fmt.Sprintf("%.2f %.2f %.2f", load1, load5, load15)
}
