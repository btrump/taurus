package main

import (
	"github.com/btrump/taurus/pkg/server"
)

func getConfig() map[string]string {
	return map[string]string{
		"key": "value",
	}
}

func main() {
	server.Start(getConfig())
}
