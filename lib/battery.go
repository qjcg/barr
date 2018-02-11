package barr

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gonum/stat"
)

// FIXME: Generalize to use *io.Reader instead of *os.File to improve
// testability, either here or in Capacity().
type Battery struct {

	// Each source will contain the current battery capacity as a
	// percentage (string), as in /sys/class/power_supply/BAT0/capacity .
	Sources []*os.File
}

// Str returns battery info as a string.
func (b *Battery) String() string {
	symbol := "🔋"
	if Charging() {
		symbol = "🔌"
	}

	return fmt.Sprintf("%s %s%%", symbol, strconv.FormatFloat(b.Capacity(), 'f', 0, 64))
}

// Spark returns battery info as a sparkline.
func (b *Battery) Spark() string {
	fmtStr := "🔋  %s%%"
	if Charging() {
		fmtStr = "🔌 %s%%"
	}

	return fmt.Sprintf(fmtStr, b.Capacity())
}

// Charging returns true if AC power plugged in.
func Charging() bool {
	f := "/sys/class/power_supply/AC/online"

	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}

	cByt, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return false
	}

	// "0": false, "1": true
	return cByt[0] == '1'
}

// Capacity returns the percentage battery capacity remaining, *averaged*
// accross all sources.
func (b *Battery) Capacity() float64 {

	var capacities []float64
	for _, src := range b.Sources {

		// Create a reader from src *os.File.
		r := bufio.NewReader(src)

		// Read until newline, and then trim the newline.
		s, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s = strings.TrimSpace(s)

		// Parse the string as an fp number and append to capacities slice.
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatal(err)
		}
		capacities = append(capacities, num)

		// Move file handle cursor back to beginning. This avoids the
		// next call to Capacity() failing upon finding no data at end
		// of file (where the cursor is before this call to Seek()).
		_, err = src.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	return stat.Mean(capacities, nil)
}
