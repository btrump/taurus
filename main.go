package main

import (
	"fmt"

	"github.com/btrump/taurus/pkg/client"
	"github.com/btrump/taurus/pkg/server"
)

func main() {
	// fmt.Println(taurus.Config())
	fmt.Println(client.Hello())
	fmt.Println(server.Hello())
	serverConfig := map[string]string{
		"hello": "hello",
	}
	server.Start(serverConfig)
}
