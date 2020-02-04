package main

import (
	"github.com/btrump/taurus-server/pkg/api"
	"github.com/btrump/taurus-server/pkg/server"
)

func getConfig() map[string]string {
	return map[string]string{
		"key": "value",
	}
}

func main() {
	s := server.New(getConfig())
	api.Use(&s)
	api.Start()
}
