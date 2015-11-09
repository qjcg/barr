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
	fACOnline string = "/sys/class/power_supply/AC/online"
)

// Battery represents system battery information.
type Battery struct {
	ChargeFull float64
	ChargeNow  float64
}

type BatteryFiles struct {
	Dir        string
	ChargeFull string
	ChargeNow  string
}

// Return *BatteryFiles based on input batDir.
func NewBatteryFiles(batDir string) *BatteryFiles {
	prefix := "charge"
	chargePath := fmt.Sprintf("%s/%s%s", batDir, prefix, "_full")
	// Sometimes file prefix changes, ex sometimes is "charge_full", sometimes "energy_full".
	// Why? Who knows!
	if _, err := os.Stat(chargePath); err != nil {
		prefix = "energy"
	}
	return &BatteryFiles{
		Dir:        batDir,
		ChargeFull: fmt.Sprintf("%s/%s%s", batDir, prefix, "_full"),
		ChargeNow:  fmt.Sprintf("%s/%s%s", batDir, prefix, "_now"),
	}
}

func (bf *BatteryFiles) NewBattery() *Battery {
	// get full charge here (only need to read it once)
	cBytes, err := ioutil.ReadFile(bf.ChargeFull)
	check(err)
	b.ChargeFullStr = strings.Trim(string(cBytes), "\n")
	chargeFullFloat, err := strconv.ParseFloat(chargeFullStr, 64)
	check(err)

	return &Battery{Dir: bf.Dir, bf.ChargeFull, bf.ChargeNow}
}

// Str returns battery info as a string.
// Implements the StatusStringer interface.
func (b *Battery) Str() string {
	b.chargeFull()
	b.chargeNow()
	if b.charging() {
		return fmt.Sprintf("AC.%s%%", b.ChargePctStr)
	}
	return fmt.Sprintf("%s%%", b.ChargePctStr)
}

// charging returns true if plugged in.
func (b *Battery) charging() bool {
	cByt, err := ioutil.ReadFile(fACOnline)
	if err != nil {
		log.Println(err)
	}
	// "0": false, "1": true
	return string(cByt[0]) == "1"
}

func (b *Battery) getChargeNow() {
	cBytes, err := ioutil.ReadFile(b.Dir + "/" + "charge_now")
	check(err)
	cStr := strings.Trim(string(cBytes), "\n")
	b.ChargeNow, err = strconv.ParseFloat(cStr, 64)
	check(err)
}

func (b *Battery) getChargePct() {
	chargePct := chargeNowFloat / chargeFullFloat * 100.0
	chargePctStr := strconv.FormatFloat(chargePct, 'f', 0, 64)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
