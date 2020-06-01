// swaybar is an library for implementing the swaybar-protocol(7).
// See https://i3wm.org/docs/i3bar-protocol.html
package swaybar

import (
	"os"
	"syscall"
)

// Header represents a swaybar-protocol header.
type Header struct {
	Version     int       `json:"version"`
	ClickEvents bool      `json:"click_events"`
	ContSignal  os.Signal `json:"cont_signal"`
	StopSignal  os.Signal `json:"stop_signal"`
}

var DefaultHeader = Header{
	Version:     1,
	ClickEvents: true,
	ContSignal:  syscall.SIGCONT,
	StopSignal:  syscall.SIGSTOP,
}

// Updater defines the interface.
type Updater interface {
	Update()
}

// StatusLine is a slice of Blocks representing a complete swaybar statusline.
type StatusLine struct {
	Blocks []Updater
}

// Update implements the Updater interface for StatusLine.
func (sl *StatusLine) Update() {
	for _, b := range sl.Blocks {
		b.Update()
	}
}

// Block represents a single item in a StatusLine.
type Block struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Color               string `json:"color"`
	Background          string `json:"background"`
	Border              string `json:"border"`
	BorderTop           int    `json:"border_top"`
	BorderBottom        int    `json:"border_bottom"`
	BorderLeft          int    `json:"border_left"`
	BorderRight         int    `json:"border_right"`
	MinWidth            int    `json:"min_width"`
	Align               string `json:"align"`
	Name                string `json:"name"`
	Instance            string `json:"instance"`
	Urgent              bool   `json:"urgent"`
	Separator           bool   `json:"separator"`
	SeparatorBlockWidth int    `json:"separator_block_width"`
	Markup              string `json:"markup"`
}

// DefaultBlock is a block providing default settings.
var DefaultBlock = Block{
	Background: "#0000ff",
	Color:      "#00ff00",
	MinWidth:   100,
	Align:      "right",
}

// ClickEvent represents a swaybar-protocol click event.
type ClickEvent struct {
	Name      string
	Instance  string
	X         int
	Y         int
	Button    int
	Event     int
	RelativeX int
	RelativeY int
	Width     int
	Height    int
}
