package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	batBaseDir        = "/status"
	batStatusFile     = *batDir + "/status"
	batChargeNowFile  = *batDir + "/charge_now"
	batChargeFullFile = *batDir + "/charge_full"
)

// The Battery function returns the percentage of laptop battery charge remaining as a string.
func Battery(statusFile string) string {
	status, err := ioutil.ReadFile(batStatusFile)
	check(err)
	statusStr := string(status)

	var prefix string
	if statusStr == "Charging" {
		prefix = "AC."
	}

	chargeNow, err := ioutil.ReadFile(batChargeNowFile)
	check(err)
	chargeNowStr := strings.Trim(string(chargeNow), "\n")
	chargeNowFloat, err := strconv.ParseFloat(chargeNowStr, 64)
	check(err)

	chargeFull, err := ioutil.ReadFile(batChargeFullFile)
	check(err)
	chargeFullStr := strings.Trim(string(chargeFull), "\n")
	chargeFullFloat, err := strconv.ParseFloat(chargeFullStr, 64)
	check(err)

	chargePct := chargeNowFloat / chargeFullFloat * 100.0
	chargePctStr := strconv.FormatFloat(chargePct, 'f', 0, 64)

	return fmt.Sprintf("%s%s%%", prefix, chargePctStr)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
