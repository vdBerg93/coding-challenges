package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readData()

	// Test data
	got := Part1(testData)
	want1 := 288
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Solution part 1: %d\n", Part1(data))

	got = Part2(testData)
	want2 := 71503
	if got != want2 {
		log.Fatalf("expected %d, got %d", want2, got)
	}
	got = Part2(data)
	fmt.Printf("Solution part 2: %d\n", got)
}

var testData []byte

var data []byte

func readData() {
	var err error
	data, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample")
	if err != nil {
		panic(err)
	}
}

func parseInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, word := range strs {
		val, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		ints = append(ints, val)
	}
	return ints
}

func Part1(data []byte) int {
	rows := strings.Split(string(data), "\n")
	times := parseInts(strings.Fields(strings.Split(rows[0], ":")[1]))
	distances := parseInts(strings.Fields(strings.Split(rows[1], ":")[1]))

	var optionsPerRace []int
	for i := range times {
		options := 0
		tRace := times[i]
		for tHold := 0; tHold <= tRace; tHold++ {
			d := getDistance(tHold, tRace)
			if d > distances[i] {
				options++
			}
		}
		optionsPerRace = append(optionsPerRace, options)
	}
	product := 1
	for _, r := range optionsPerRace {
		product *= r
	}

	return product
}

func getDistance(tHold, tRace int) int {
	return (tRace - tHold) * tHold
}

func Part2(data []byte) int {
	rows := strings.Split(string(data), "\r\n")

	timeText := strings.Replace(strings.Split(rows[0], ":")[1], " ", "", -1)
	tRace, err := strconv.Atoi(timeText)
	if err != nil {
		panic(err)
	}

	distanceText := strings.Replace(strings.Split(rows[1], ":")[1], " ", "", -1)
	distance, err := strconv.Atoi(distanceText)
	if err != nil {
		panic(err)
	}

	options := 0
	for tHold := 0; tHold <= tRace; tHold++ {
		d := getDistance(tHold, tRace)
		if d > distance {
			options++
		}
	}

	return options
}
