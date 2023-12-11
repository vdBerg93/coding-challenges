package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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
	got := Solve(testData, 2)
	want1 := 374
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, 2))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, 100)
	want1 := 8410
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve(data, 1e6))
}

type Space struct {
	content        []string
	columnDistance []int
	rowDistance    []int
	galaxies       []Galaxy
}

type Galaxy struct {
	row, col int
}

func Solve(data []byte, expansionRate int) int {
	space := NewSpace(data, expansionRate)
	return space.GetDistanceSum()
}

func NewSpace(data []byte, expansionRate int) *Space {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)

	var space []string
	for scanner.Scan() {
		row := scanner.Text()
		space = append(space, row)
	}

	rowDistances := getExpansionRates(space, expansionRate)

	spaceTransformed := transpose(space)
	columnDistances := getExpansionRates(spaceTransformed, expansionRate)

	return &Space{
		content:        space,
		columnDistance: columnDistances,
		rowDistance:    rowDistances,
		galaxies:       getGalaxies(space),
	}
}

func getExpansionRates(space []string, expansionRate int) []int {
	var distances []int
	for _, row := range space {
		if emptyRow(row) {
			distances = append(distances, expansionRate)
		} else {
			distances = append(distances, 1)
		}
	}
	return distances
}

func getGalaxies(space []string) []Galaxy {
	var planets []Galaxy
	for r, row := range space {
		for c, char := range row {
			if char == '#' {
				planets = append(planets, Galaxy{r, c})
			}
		}
	}
	return planets
}

func (s *Space) GetDistanceSum() int {
	total := 0
	for i1 := range s.galaxies {
		for i2 := range s.galaxies {
			total += s.getShortestPath(i1, i2)
		}
	}
	return total / 2 // Calculated every distance twice
}

func (s *Space) getShortestPath(i, j int) int {
	steps := 0
	c1 := s.galaxies[i].col
	c2 := s.galaxies[j].col
	r1 := s.galaxies[i].row
	r2 := s.galaxies[j].row
	// Horizontal steps
	for i := min(c1, c2); i < max(c1, c2); i++ {
		steps += s.columnDistance[i]
	}
	// Vertical steps
	for j := min(r1, r2); j < max(r1, r2); j++ {
		steps += s.rowDistance[j]

	}
	return steps
}

func emptyRow(row string) bool {
	for _, char := range row {
		if char != '.' {
			return false
		}
	}
	return true
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
