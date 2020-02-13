package server_test

import (
	"testing"

	"github.com/btrump/taurus-server/pkg/server"
)

func TestNew(t *testing.T) {
	s := server.New()
	if s.Engine == nil {
		t.Errorf("Got nil; want value")
	}
	if s.Name != "taurus-server-default" {
		t.Errorf("Got '%s'; want 'taurus-server-default'", s.Name)
	}
	if s.Version != "development" {
		t.Errorf("Got '%s'; want 'development'", s.Version)
	}
}

func TestConfigure(t *testing.T) {
	config := server.Config{
		Name:    "test-name",
		Version: "test-version",
	}
	s := server.New()
	s.Configure(config)
	if s.Engine == nil {
		t.Errorf("Got nil; want value")
	}
	if s.Name != config.Name {
		t.Errorf("Got '%s'; want '%s'", s.Name, config.Name)
	}
	if s.Version != config.Version {
		t.Errorf("Got '%s'; want '%s'", s.Version, config.Version)
	}
}

func TestProcessRequest(t *testing.T) {
	// t.Errorf("error")
}
