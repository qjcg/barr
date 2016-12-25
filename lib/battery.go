package barr

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FACOnline = "/sys/class/power_supply/AC/online"
	BatDir    = "/sys/class/power_supply/BAT0"
)

// Battery represents system battery information.
type Battery struct {
	// files
	Dir         string
	FChargeFull string
	FChargeNow  string

	// values
	ChargeFull float64
	ChargeNow  float64
}

// NewBattery returns a *Battery based on input batDir.
func NewBattery(batDir string) (*Battery, error) {

	// Return early if batDir doesn't exist.
	if _, err := os.Lstat(batDir); os.IsNotExist(err) {
		return nil, err
	}

	fChargeFull := fmt.Sprintf("%s/%s", batDir, "charge_full")

	// Sometimes file prefix changes, ex sometimes is "charge_full",
	// sometimes "energy_full".
	if _, err := os.Stat(fChargeFull); os.IsNotExist(err) {
		fChargeFull = strings.Replace(fChargeFull, "charge", "energy", -1)

		// Return if we still can't find the file after substituting the
		// "energy" prefix.
		if _, err := os.Stat(fChargeFull); os.IsNotExist(err) {
			return nil, err
		}
	}

	// Get ChargeFull
	cBytes, err := ioutil.ReadFile(fChargeFull)
	if err != nil {
		return nil, err
	}
	chargeFullStr := strings.TrimSpace(string(cBytes))
	chargeFullFloat, err := strconv.ParseFloat(chargeFullStr, 64)
	if err != nil {
		return nil, err
	}

	return &Battery{
		Dir:         batDir,
		FChargeFull: fChargeFull,
		FChargeNow:  strings.Replace(fChargeFull, "_full", "_now", -1),
		ChargeFull:  chargeFullFloat,
		ChargeNow:   0.0,
	}, nil
}

// Str returns battery info as a string.
func (b *Battery) String() string {
	symbol := "🔋"
	if b.charging() {
		symbol = "🔌"
	}

	b.getChargeNow()
	return fmt.Sprintf("%s %s%%", symbol, b.getChargePct())
}

// Spark returns battery info as a sparkline.
func (b *Battery) Spark() string {
	fmtStr := "🔋  %s%%"
	if b.charging() {
		fmtStr = "🔌 %s%%"
	}

	err := b.getChargeNow()
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf(fmtStr, b.getChargePct())
}

// charging returns true if plugged in.
func (b *Battery) charging() bool {
	if _, err := os.Stat(FACOnline); os.IsNotExist(err) {
		return false
	}

	cByt, err := ioutil.ReadFile(FACOnline)
	if err != nil {
		log.Println(err)
		return false
	}
	// "0": false, "1": true
	return cByt[0] == '1'
}

// getChargeNow updates b.ChargeNow with the latest value.
func (b *Battery) getChargeNow() error {
	cBytes, err := ioutil.ReadFile(b.FChargeNow)
	if err != nil {
		return err
	}
	cStr := strings.Trim(string(cBytes), "\n")
	b.ChargeNow, err = strconv.ParseFloat(cStr, 64)
	if err != nil {
		return err
	}
	return nil
}

// getChargePct returns the charge percent as a string
func (b *Battery) getChargePct() string {
	chargePct := b.ChargeNow / b.ChargeFull * 100.0
	return strconv.FormatFloat(chargePct, 'f', 0, 64)
}
