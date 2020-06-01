package blocks

import (
	"time"

	"github.com/qjcg/barr/pkg/protocol"
)

var DefaultTimestamp = Timestamp{
	Block: protocol.DefaultBlock,
}

type Timestamp struct {
	protocol.Block
}

const (
	fmtShort = "Mon Jan 2, 3:04pm"
	fmtLong  = "Mon Jan 2, 3:04:05pm MST"
)

// Update updates the Timestamp FullText.
func (ts *Timestamp) Update() {
	ts.FullText = time.Now().Format(fmtShort)
}
