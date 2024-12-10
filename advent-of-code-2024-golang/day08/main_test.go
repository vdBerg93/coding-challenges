package main

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed example.txt
var testData []byte

func Test_part1(t *testing.T) {
	m := parseInput(bytes.NewReader(testData))
	want := 14
	if got := solve(m, getAntiZones); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	m := parseInput(bytes.NewReader(testData))
	want := 34
	if got := solve(m, getAntiZones2); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
