package main

import (
	"fmt"

	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/server"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry
var debug = true

func init() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	if debug {
		logrus.SetReportCaller(true)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log = logrus.WithFields(
		logrus.Fields{
			"application": "watcher-demo",
			"version":     "0.0.1",
		},
	)
}

func demo(s interface{}) {
	log.Printf("Starting demo")
	ccount := 1
	lcount := 1
	port := 8080
	log.Printf("Creating %v clients @ %v req/ea for %v events", ccount, lcount, ccount*lcount)
	for i := 0; i < ccount; i++ {
		c := client.New()
		socket := fmt.Sprintf("localhost:%d", port)
		for j := 0; j < lcount; j++ {
			c.Connect(socket)
		}
	}
}

func main() {
	s := server.New()
	go s.Start()
	demo(s)
	log.Printf("Exiting")
}
