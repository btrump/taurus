/*
Package taurus contains the entrypoint to taurus-server and taurus-api
which are state and client/server communication engines, respectively.
*/
package taurus

import (
	"github.com/btrump/taurus-server/pkg/api"
	"github.com/btrump/taurus-server/pkg/server"
)

func getConfig() map[string]string {
	return map[string]string{}
}

func main() {
	s := server.New(getConfig())
	a := api.New()
	a.Use(&s)
	a.Start()
}
