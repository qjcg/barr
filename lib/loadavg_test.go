package barr

import (
	"regexp"
	"testing"
)

func TestLoadAvg(t *testing.T) {
	var la LoadAvg
	la.Update()
	if m, _ := regexp.MatchString("([0-9]+.[0-9]{2} ?){3}", la.Str()); !m {
		t.Fatalf("actual output does not match regexp: %s", la.Str())
	}
}
