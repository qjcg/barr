package main

// Barstringers should return a string describing the current state of the
// data, or otherwise if no data is available, the empty string.
type BarStringer interface {
	Str() string
	Update() error
}
