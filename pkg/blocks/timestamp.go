package blocks

import (
	"time"

	"github.com/qjcg/barr/pkg/swaybar"
)

var DefaultTimestamp = Timestamp{
	Block: swaybar.DefaultBlock,
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
	ts.MinWidth = ts.FullText + "11"
}
