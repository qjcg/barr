// swaybar is an library for implementing the swaybar-protocol(7).
package swaybar

import (
	"os"
)

// Header represents a swaybar-protocol header.
type Header struct {
	Version     int
	ClickEvents bool
	ContSignal  os.Signal
	StopSignal  os.Signal
}

// Body represents a swaybar-protocol body.
type Body struct {
	StatusLines []StatusLine
}

// StatusLine is a slice of Blocks representing a complete swaybar statusline.
type StatusLine struct {
	Blocks []Block
}

// Block represents a single item in a StatusLine.
type Block struct {
	FullText            string
	ShortText           string
	Color               string
	Background          string
	Border              string
	BorderTop           int
	BorderBottom        int
	BorderLeft          int
	BorderRight         int
	MinWidth            int
	Align               string
	Name                string
	Instance            string
	Urgent              bool
	Separator           bool
	SeparatorBlockWidth int
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
