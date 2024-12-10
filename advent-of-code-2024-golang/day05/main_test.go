package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var testData []byte

func Test_part1(t *testing.T) {
	rules, pages := readInput(testData)
	want := 143
	t.Run("example", func(t *testing.T) {
		if got := solve(rules, pages); got != want {
			t.Errorf("solve() = %v, want %v", got, want)
		}
	})
}

func Test_part2(t *testing.T) {
	rules, pages := readInput(testData)
	want := 123
	t.Run("example", func(t *testing.T) {
		if got := part2(rules, pages); got != want {
			t.Errorf("solve() = %v, want %v", got, want)
		}
	})
}
