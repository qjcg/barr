package blocks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qjcg/barr/pkg/protocol"
)

// See https://www.kraken.com/en-gb/features/api#get-ticker-info
const URLTemplate = `https://api.kraken.com/0/public/Ticker?pair=%s`

type APIResponse struct {
	Result struct {
		Pair struct {
			P [2]string
		} `json:"XXBTZCAD"`
	} `json:"result"`
}

var DefaultCryptoCurrency = CryptoCurrency{
	Block: protocol.DefaultBlock,
}

type CryptoCurrency struct {
	Pair string

	protocol.Block
}

func (c *CryptoCurrency) Update() {
	url := fmt.Sprintf(URLTemplate, c.Pair)
	resp, err := http.Get(url)
	if err != nil {
		c.FullText = err.Error()
	}

	defer resp.Body.Close()

	var r APIResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		c.FullText = err.Error()
	}

	// P (volume weighted average price array): [today, last-24-hours]
	c.FullText = fmt.Sprintf("%s: %s", c.Pair, r.Result.Pair.P[0])
}
