package main

import "github.com/qjcg/barr/pkg/protocol"

// Config holds our application's configuration settings.
type Config struct {
	Blocks []protocol.Updater
}
