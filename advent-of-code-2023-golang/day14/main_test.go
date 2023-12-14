package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

var data []byte
var testData []byte

func TestMain(m *testing.M) {
	var err error
	data, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func Test_Part1(t *testing.T) {
	got := Solve(testData, 1)
	want1 := 136
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, 1))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, 1e3)
	want1 := 64
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve(data, 1e3))
}

func readPlatform(data []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)

	var platform []string
	for scanner.Scan() {
		platform = append(platform, scanner.Text())
	}
	return platform
}

func Solve(data []byte, cycles int) int {
	platform := readPlatform(data)
	modes := 0
	if cycles != 1 {
		modes = 3
	}
	for c := 0; c < cycles; c++ {
		for m := 0; m <= modes; m++ {
			var transform, invert bool
			switch m % 4 {
			case 0: // North
				transform = true
			case 1: // West
				// do nothing
			case 2: // South
				transform = true
				invert = true
			case 3: // East
				invert = true
			}

			var platformT []string
			if transform {
				platformT = transpose(platform)
			} else {
				platformT = platform
			}

			for i, row := range platformT {
				platformT[i] = moveUpDown(row, invert)
			}

			if transform {
				platform = transpose(platformT)
			} else {
				platform = platformT
			}
		}
	}

	// Calculate weight
	platformT := transpose(platform)
	weight := 0
	for _, row := range platformT {
		weight += calculateWeight(row)
	}

	return weight
}

func Test_MoveUpDown(t *testing.T) {
	type test struct {
		row    string
		want   string
		invert bool
	}
	tests := []test{
		{".O.O..#.O.##", "OO....#O..##", false},
		{".O.O..#.O.##", "....OO#..O##", true},
	}

	for _, tt := range tests {
		t.Run(tt.row, func(t *testing.T) {
			got := moveUpDown(tt.row, tt.invert)
			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func calculateWeight(line string) int {
	weight := 0
	for i, char := range line {
		if char == 'O' {
			weight += len(line) - i
		}
	}
	return weight
}

func transpose(input []string) []string {
	output := make([]string, len(input[0]))

	for _, row := range input {
		for c, char := range row {
			output[c] += string(char)
		}
	}
	return output
}

func moveUpDown(line string, inverse bool) string {
	if inverse {
		line = reverseString(line)
	}
	parts := strings.Split(line, "#")

	for i := 0; i < len(parts); i++ {
		movingRocks := strings.Count(parts[i], "O")
		parts[i] = strings.Replace(parts[i], "O", ".", -1)
		parts[i] = strings.Replace(parts[i], ".", "O", movingRocks)
	}
	output := strings.Join(parts, "#")
	if len(output) != len(line) {
		panic("invalid length")
	}
	if inverse {
		output = reverseString(output)
	}
	return output
}

func reverseString(str string) string {
	var out string
	for i := len(str) - 1; i >= 0; i-- {
		out += string(str[i])
	}
	return out
}
