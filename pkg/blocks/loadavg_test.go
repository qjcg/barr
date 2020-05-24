package blocks

import (
	"regexp"
	"testing"
)

func TestLoadAvg(t *testing.T) {
	var la LoadAvg
	got := la.String()
	reLoadAvg := `([0-9]+.[0-9]{2} ?){3}`
	if m, _ := regexp.MatchString(reLoadAvg, got); !m {
		t.Fatalf("Wanted regex match for %s, got %s", reLoadAvg, got)
	}
	t.Logf("Current load average: %s\n", got)
}
