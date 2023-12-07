package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed sample.txt
var testData []byte

//go:embed input.txt
var data []byte

var lines [][]byte

func printCounts() {
	chars := make(map[rune]int)
	for _, r := range string(data) {
		chars[r]++
	}
	for r, cnt := range chars {
		fmt.Printf("%s: %d\n", string(r), cnt)
	}

}

func Part1(data []byte) int {
	//printCounts()

	lines = bytes.Split(data, []byte("\n"))

	var sum int
	var word string
	var adjacent bool

	reset := func() {
		adjacent = false
		word = ""
	}

	store := func() {
		if adjacent == false {
			log.Fatal("shouldn't be possible")
		}
		val, err := strconv.Atoi(word)
		if err != nil {
			log.Fatalf("parsing int; %v", err)
		}
		sum += val
	}

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			char := rune(lines[r][c])
			if unicode.IsNumber(char) {
				word += string(char)
				if !adjacent && isAdjacent(r, c) {
					adjacent = true
				}
			} else {
				if adjacent {
					store()
				}
				reset()
			}
		}
	}

	return sum
}

func isAdjacent(row, col int) (isAdjacent bool) {
	rPrev, rNext := getOffset(row, len(lines))
	cPrev, cNext := getOffset(col, len(lines[0]))
	for r := row + rPrev; r <= row+rNext; r++ {
		for c := col + cPrev; c <= col+cNext; c++ {
			if r == row && c == col {
				continue
			}
			char := rune(lines[r][c])
			if strings.ContainsAny(string(lines[r][c]), "/=!@#$%^&*():;-+") || unicode.IsSymbol(char) {

				return true
			}
		}
	}
	return false
}

func getOffset(i, length int) (previous, next int) {
	if i > 0 {
		previous = -1
	}
	if i < length-1 {
		next = 1
	}
	return previous, next
}

func Part2(data []byte) int {
	//printCounts()

	lines = bytes.Split(data, []byte("\n"))

	var sum int

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			char := string(lines[r][c])
			if char == "*" {
				ratio := getGearRatio(r, c)
				sum += ratio
			}
		}
	}

	return sum
}

// getGearRatio - Loop over a box of max 3x3 characters and check for parts to calculate the gear ratio
func getGearRatio(row, col int) int {
	parts := make(map[int]bool) // Part numbers are unique so store once

	// Get box range, corrected for min/max column idx
	rPrev, rNext := getOffset(row, len(lines))
	cPrev, cNext := getOffset(col, len(lines[0]))

	for r := row + rPrev; r <= row+rNext; r++ {
		for c := col + cPrev; c <= col+cNext; c++ {
			if r == row && c == col {
				// Gear itself
				continue
			}

			char := rune(lines[r][c])
			if unicode.IsNumber(char) {
				part := getPart(r, c)
				parts[part] = true
			}
		}
	}
	if len(parts) < 2 {
		// Gear is not connecting at least two parts
		return 0
	}

	// Calculate gear ratio
	gearRatio := 1
	for part, _ := range parts {
		gearRatio *= part
	}

	return gearRatio
}

func getPart(row, col int) int {
	// Get part start
	start := col
	for i := start; i >= 0; i-- {
		if unicode.IsNumber(rune(lines[row][i])) {
			start = i
		} else {
			break
		}
	}
	// Get part end
	end := col
	if end < len(lines[row]) {
		for i := end; i < len(lines[row]); i++ {
			if unicode.IsNumber(rune(lines[row][i])) {
				end++
			} else {
				break
			}
		}
	}
	// Get part number
	part, err := strconv.Atoi(string(lines[row][start:end]))
	if err != nil {
		log.Fatalf("parsing part; %v", err)
	}
	return part
}
