package main

import (
	"regexp"
	"testing"
)

func TestLoadAvg(t *testing.T) {
	actual := LoadAvg()
	m, _ := regexp.MatchString("([0-9]+.[0-9]{2} ?){3}", actual)
	if !m {
		t.Fatalf("actual output does not match regexp: %s", actual)
	}
}
