// +build integration

// NOTE: Tests in this file require that the disk be present on
// the test system to work (hence the "integration" build tag).
package sysinfo

import (
	"testing"
)

func TestDiskString(t *testing.T) {
	d := Disk{Dir: "/"}
	t.Log(d.String())
}
