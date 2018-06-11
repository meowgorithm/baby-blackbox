package main_test

import (
	"blackbox"
	"os"
	"testing"

	example "blackbox/example"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCoolness(t *testing.T) {
	req := example.Response{}

	api := blackbox.New(t, example.Init())
	api.Request("GET", "/", nil).
		OK().
		JSON(&req)

	if !req.Cool {
		t.Error("expected things to be cool, but they were not")
	}
}
