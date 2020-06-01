// swaybar is an library for implementing the swaybar-protocol(7).
// See https://i3wm.org/docs/i3bar-protocol.html
package swaybar

import (
	"os"
	"syscall"
)

// Header represents a protocol header.
type Header struct {
	Version     int       `json:"version"`
	ClickEvents bool      `json:"click_events,omitempty"`
	ContSignal  os.Signal `json:"cont_signal,omitempty"`
	StopSignal  os.Signal `json:"stop_signal,omitempty"`
}

// DefaultHeader is a Header providing default settings.
var DefaultHeader = Header{
	Version:     1,
	ClickEvents: true,
	ContSignal:  syscall.SIGCONT,
	StopSignal:  syscall.SIGSTOP,
}

// Updater defines an interface for structs that can Update their values.
type Updater interface {
	Update()
}

// StatusLine is a slice of Blocks representing a complete statusline.
type StatusLine struct {
	Blocks []Updater
}

// Update updates all Updaters in a StatusLine.
func (sl *StatusLine) Update() {
	for _, b := range sl.Blocks {
		b.Update()
	}
}

// Block represents a single item in a StatusLine.
type Block struct {
	FullText            string `json:"full_text,omitempty"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"`
	Border              string `json:"border,omitempty"`
	BorderTop           int    `json:"border_top,omitempty"`
	BorderBottom        int    `json:"border_bottom,omitempty"`
	BorderLeft          int    `json:"border_left,omitempty"`
	BorderRight         int    `json:"border_right,omitempty"`
	MinWidth            string `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"`
	Name                string `json:"name,omitempty"`
	Instance            string `json:"instance,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`
}

// DefaultBlock provides a block with default settings.
var DefaultBlock = Block{
	Background:          "#000000",
	Color:               "#cccccc",
	Align:               "center",
	Separator:           false,
	SeparatorBlockWidth: 25,
}

// ClickEvent represents a protocol click event.
type ClickEvent struct {
	Name      string `json:"name,omitempty"`
	Instance  string `json:"instance,omitempty"`
	X         int    `json:"x,omitempty"`
	Y         int    `json:"y,omitempty"`
	Button    int    `json:"button,omitempty"`
	Event     int    `json:"event,omitempty"`
	RelativeX int    `json:"relative_x,omitempty"`
	RelativeY int    `json:"relative_y,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}
