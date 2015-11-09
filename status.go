package main

import "strings"

type Status struct {
	OFS      string
	Elements []string
}

// Str returns a Status line string.
func (s *Status) Str() string {
	return strings.Join(s.Elements, s.OFS)
}
