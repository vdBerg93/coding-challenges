package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day2-input-sample.txt
var testData []byte

//go:embed day2-input.txt
var data []byte

func main() {
	// Test data
	got := Part1(testData)
	if got != 8 {
		log.Fatalf("expected 8, got %d", got)
	}
	fmt.Printf("Solution exercise 1: %d\n", Part1(data))

	got = Part2(testData)
	if got != 2286 {
		log.Fatalf("expected 2286, got %d", got)
	}
	fmt.Printf("Solution part 2: %d\n", Part2(data))
}

const (
	red      = "red"
	green    = "green"
	blue     = "blue"
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func Part1(data []byte) int {
	buf := bytes.Buffer{}
	_, _ = buf.Write(data)
	scanner := bufio.NewScanner(&buf)
	scanner.Split(bufio.ScanLines)

	var possibleSum int
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ":")
		// Get game ID
		gameFields := strings.Fields(fields[0])
		gameID, err := strconv.Atoi(gameFields[1])
		if err != nil {
			log.Fatal(err)
		}
		// Get moves
		validGame := true
		moves := strings.Split(fields[1], ";")
		for _, move := range moves {
			blocks := strings.Split(move, ",")
			for _, block := range blocks {
				blockFields := strings.Fields(block)
				color := blockFields[1]
				count, err := strconv.Atoi(blockFields[0])
				if err != nil {
					log.Fatal(err)
				}
				var threshold int
				switch color {
				case red:
					threshold = maxRed
				case green:
					threshold = maxGreen
				case blue:
					threshold = maxBlue
				default:
					log.Fatalf("invalid color %s", color)
				}
				if count > threshold {
					validGame = false
					goto Done
				}
			}
		}

	Done:
		if validGame {
			possibleSum += gameID
		}
	}
	return possibleSum
}

func Part2(data []byte) int {
	buf := bytes.Buffer{}
	_, _ = buf.Write(data)
	scanner := bufio.NewScanner(&buf)
	scanner.Split(bufio.ScanLines)

	var possibleSum int
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ":")
		// Get moves
		moves := strings.Split(fields[1], ";")
		// Keep counts
		var blueCount int
		var redCount int
		var greenCount int
		for _, move := range moves {
			blocks := strings.Split(move, ",")
			for _, block := range blocks {
				blockFields := strings.Fields(block)
				color := blockFields[1]
				count, err := strconv.Atoi(blockFields[0])
				if err != nil {
					log.Fatal(err)
				}

				switch color {
				case red:
					if count > redCount {
						redCount = count
					}
				case green:
					if count > greenCount {
						greenCount = count
					}
				case blue:
					if count > blueCount {
						blueCount = count
					}
				default:
					log.Fatalf("invalid color %s", color)
				}
			}
		}
		possibleSum += redCount * blueCount * greenCount
	}
	return possibleSum
}
