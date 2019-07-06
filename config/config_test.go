package config

import (
	"testing"
)

func TestCfg(t *testing.T) {
	want := config{
		server{"debug", "0.0.0.0:8080", 60, 60},
		database{"mysql", "root", "root", "localhost:3306", "forum", "tcp", "utf8", "True", "Local", 10, 100},
	}
	if got := Cfg; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
