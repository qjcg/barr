// swaybar is an library for implementing the swaybar-protocol(7).
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

// StatusLine is a slice of Blocks representing a complete swaybar statusline.
type StatusLine struct {
	Blocks []Block
}

// Block represents a single item in a StatusLine.
type Block struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Color               string
	Background          string
	Border              string
	BorderTop           int `json:"border_top"`
	BorderBottom        int `json:"border_bottom"`
	BorderLeft          int `json:"border_left"`
	BorderRight         int `json:"border_right"`
	MinWidth            int `json:"min_width"`
	Align               string
	Name                string
	Instance            string
	Urgent              bool
	Separator           bool
	SeparatorBlockWidth int `json:"separator_block_width"`
	Markup              string
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
