package main

import (
	_ "embed"
	"fmt"
	"log"
)

func main() {
	// Test data
	got := Part1(testData)
	want1 := 4361
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Solution part 1: %d\n", Part1(data))

	got = Part2(testData)
	want2 := 467835
	if got != want2 {
		log.Fatalf("expected %d, got %d", want2, got)
	}
	fmt.Printf("Solution part 2: %d\n", Part2(data))
}
