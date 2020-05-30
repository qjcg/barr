// The barr command prints out a system status line.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	toml "github.com/pelletier/go-toml"

	"github.com/qjcg/barr/pkg/blocks"
	"github.com/qjcg/barr/pkg/swaybar"
)

var defaultBlock = swaybar.Block{
	Color:      "#00ff00",
	Background: "#0000ff",
	Border:     "#ff0000",
	Markup:     "",
}

func main() {
	flagConfig := flag.String("c", "", "config file")
	flagVersion := flag.Bool("v", false, "print version")
	flag.Parse()

	if *flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	var config Config
	c, err := ioutil.ReadFile(*flagConfig)
	err = toml.Unmarshal(c, &config)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		dec := json.NewDecoder(os.Stdin)
		for dec.More() {
			var event swaybar.ClickEvent
			err := dec.Decode(&event)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v\n", event)
		}
	}()

	// Create encoder and write the header.
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(swaybar.DefaultHeader)

	// Create a statusline.
	sl := swaybar.StatusLine{
		Blocks: []swaybar.Updater{c.Blocks...},
	}

	fmt.Fprintln(os.Stdout, "[")
	sl.Update()
	enc.Encode(sl.Blocks)
	fmt.Fprintln(os.Stdout, ",")
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		sl.Update()
		enc.Encode(sl.Blocks)
		fmt.Fprintln(os.Stdout, ",")
	}
	defer fmt.Fprintln(os.Stdout, "]")
}
