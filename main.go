/*
Package taurus contains the entrypoint to taurus-server and taurus-api
which are state and client/server communication engines, respectively.
*/
package main

import (
	"github.com/btrump/taurus-server/pkg/api"
	"github.com/btrump/taurus-server/pkg/server"
)

func main() {
	s := server.New()
	a := api.New()
	a.Use(s)
	a.Start()
}
