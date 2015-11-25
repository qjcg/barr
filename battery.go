package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fACOnline    string = "/sys/class/power_supply/AC/online"
	fFirstBatDir string = "/sys/class/power_supply/BAT0"
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
	// return early if batDir doesn't exist
	if _, err := os.Stat(batDir); os.IsNotExist(err) {
		return &Battery{}, err
	}

	fChargeFull := fmt.Sprintf("%s/%s", batDir, "charge_full")
	// Sometimes file prefix changes, ex sometimes is "charge_full", sometimes "energy_full".
	if _, err := os.Stat(fChargeFull); os.IsNotExist(err) {
		fChargeFull = strings.Replace(fChargeFull, "charge", "energy", -1)

		// If we still can't find the file after substituting the "energy" prefix, return
		if _, err := os.Stat(fChargeFull); os.IsNotExist(err) {
			return &Battery{}, err
		}
	}

	// Get ChargeFull
	cBytes, err := ioutil.ReadFile(fChargeFull)
	check(err)
	chargeFullStr := strings.Trim(string(cBytes), "\n")
	chargeFullFloat, err := strconv.ParseFloat(chargeFullStr, 64)
	check(err)

	return &Battery{
		Dir:         batDir,
		FChargeFull: fChargeFull,
		FChargeNow:  strings.Replace(fChargeFull, "_full", "_now", -1),
		ChargeFull:  chargeFullFloat,
		ChargeNow:   0.0,
	}, nil
}

// Str returns battery info as a string.
func (b *Battery) Str() string {
	fmtStr := "%s%%"
	if b.charging() {
		fmtStr = "AC.%s%%"
	}

	b.getChargeNow()
	return fmt.Sprintf(fmtStr, b.getChargePct())
}

// charging returns true if plugged in.
func (b *Battery) charging() bool {
	cByt, err := ioutil.ReadFile(fACOnline)
	if err != nil {
		log.Println(err)
	}
	// "0": false, "1": true
	return cByt[0] == '1'
}

// getChargeNow updates b.ChargeNow with the latest value.
func (b *Battery) getChargeNow() {
	cBytes, err := ioutil.ReadFile(b.FChargeNow)
	check(err)
	cStr := strings.Trim(string(cBytes), "\n")
	b.ChargeNow, err = strconv.ParseFloat(cStr, 64)
	check(err)
}

// getChargePct returns the charge percent as a string
func (b *Battery) getChargePct() string {
	chargePct := b.ChargeNow / b.ChargeFull * 100.0
	return strconv.FormatFloat(chargePct, 'f', 0, 64)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
