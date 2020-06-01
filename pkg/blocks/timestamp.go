package blocks

import (
	"time"

	"github.com/qjcg/barr/pkg/swaybar"
)

var DefaultTimestamp = Timestamp{
	Block: swaybar.Block{
		Background: "#0000ff",
		Color:      "#00ff00",
		MinWidth:   100,
		Align:      "right",
	},
}

type Timestamp struct {
	swaybar.Block
}

const (
	fmtShort = "Mon Jan 2, 3:04pm"
	fmtLong  = "Mon Jan 2, 3:04:05pm MST"
)

// Update updates the Timestamp FullText.
func (ts *Timestamp) Update() {
	ts.FullText = time.Now().Format(fmtShort)
}
