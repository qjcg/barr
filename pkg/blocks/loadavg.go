package blocks

import (
	"fmt"
	"log"
	"syscall"

	"github.com/qjcg/barr/pkg/swaybar"
)

// See https://github.com/capnm/sysinfo/blob/master/sysinfo.go
const scale = 65536.0

var DefaultLoadAvg = LoadAvg{
	Block: swaybar.DefaultBlock,
}

// LoadAvg represents the system load average (1, 5, and 15 minutes).
type LoadAvg struct {
	load1  float64
	load5  float64
	load15 float64

	swaybar.Block
}

// Update updates the LoadAvg FullText.
func (la *LoadAvg) Update() {
	si := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(si)
	if err != nil {
		log.Println("couldn't get sysinfo:", err)
		la.FullText = err.Error()
		la.MinWidth = la.FullText + "11"
	}

	la.load1 = float64(si.Loads[0]) / scale
	la.load5 = float64(si.Loads[1]) / scale
	la.load15 = float64(si.Loads[2]) / scale

	la.FullText = fmt.Sprintf("%.2f %.2f %.2f", la.load1, la.load5, la.load15)
	la.MinWidth = la.FullText + "11"
}
