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
	"github.com/qjcg/barr/pkg/protocol"
)

func main() {
	flagConfig := flag.String("c", "", "config file")
	flagPrettyPrint := flag.Bool("p", false, "pretty print JSON output")
	flagVersion := flag.Bool("v", false, "print version")
	flag.Parse()

	if *flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	var config Config
	if *flagConfig != "" {
		c, err := ioutil.ReadFile(*flagConfig)
		err = toml.Unmarshal(c, &config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		disk := blocks.DefaultDisk
		disk.Dir = "/"
		config.Blocks = []protocol.Updater{
			&disk,
			&blocks.DefaultLoadAvg,
			&blocks.DefaultTimestamp,
		}
	}

	go func() {
		dec := json.NewDecoder(os.Stdin)
		for dec.More() {
			var event protocol.ClickEvent
			err := dec.Decode(&event)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v\n", event)
		}
	}()

	// Create encoder and write the header.
	enc := json.NewEncoder(os.Stdout)
	if *flagPrettyPrint {
		enc.SetIndent("", "  ")
	}
	enc.Encode(protocol.DefaultHeader)

	// Create a statusline.
	sl := protocol.StatusLine{}
	sl.Blocks = append(sl.Blocks, config.Blocks...)

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
