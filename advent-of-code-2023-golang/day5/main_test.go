package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	readData()
	os.Exit(m.Run())
}
