package blocks

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/qjcg/barr/pkg/protocol"
)

var DefaultDisk = Disk{
	Block: protocol.DefaultBlock,
}

// DiskInfo represents hard disk info.
type Disk struct {
	Dir string

	protocol.Block
}

// Update disk information.
func (d *Disk) Update() {
	output, err := exec.Command("df", "-h", d.Dir).Output()
	if err != nil {
		d.FullText = err.Error()
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var i int
	for scanner.Scan() {
		// Skip header line.
		if i == 0 {
			i++
			continue
		}

		line := scanner.Text()
		avail := strings.Fields(line)[3]

		d.FullText = fmt.Sprintf("%s %s", d.Dir, avail)
	}

	// Handle any scanning errors.
	if err := scanner.Err(); err != nil {
		d.FullText = err.Error()
	}
}
