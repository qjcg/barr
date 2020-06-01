// +build integration

package blocks

import (
	"testing"
)

func TestCryptoCurrency(t *testing.T) {
	c := DefaultCryptoCurrency
	c.Pair = "xbtcad"
	c.Update()

	t.Log(c.FullText)
}
