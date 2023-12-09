package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	data, err = os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
