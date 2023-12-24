package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

var dataInput []byte
var dataSample1 []byte

func TestMain(m *testing.M) {
	var err error
	dataInput, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	dataSample1, err = os.ReadFile("sample1")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

const debug = false

func Test_Part1(t *testing.T) {
	got := Solve(dataSample1, 6, false)
	want1 := 16
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}

	fmt.Printf("Part 1 test succeeded.\n")
	got = Solve(dataInput, 64, false)
	fmt.Printf("Solution part 1: %d\n", got)
}

func Test_Part2(t *testing.T) {
	got := Solve(dataSample1, 5000, true)
	want1 := 16733044
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}

	fmt.Printf("Part 2 test succeeded.\n")
	got = Solve(dataInput, 26501365, true)
	fmt.Printf("Solution part 1: %d\n", got)
}

type Location [2]int

func Solve(data []byte, steps int, infinite bool) int {
	m := NewGardenMap(data, infinite)
	return m.Solve(steps)
}

var (
	up         = Location{1, 0}
	down       = Location{-1, 0}
	left       = Location{0, -1}
	right      = Location{0, 1}
	directions = []Location{up, down, left, right}
)

type GardenMap struct {
	Map      [][]rune
	start    Location
	infinite bool
}

func NewGardenMap(data []byte, infinite bool) *GardenMap {
	m := &GardenMap{
		Map:      nil,
		infinite: infinite,
	}
	var start Location
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		m.Map = append(m.Map, []rune(scanner.Text()))
		s := strings.IndexRune(scanner.Text(), 'S')
		if s != -1 {
			start = Location{len(m.Map) - 1, s}
		}
	}
	m.start = start
	return m
}

func (m *GardenMap) Solve(steps int) int {
	toExplore := map[Location]struct{}{m.start: {}}
	for s := 1; s <= steps; s++ {
		newExplore := make(map[Location]struct{})
		for loc := range toExplore {
			next := m.explore(loc)
			for _, l := range next {
				newExplore[l] = struct{}{}
			}
		}
		toExplore = newExplore
	}
	return len(toExplore)
}

func (m *GardenMap) explore(loc Location) []Location {
	var next []Location
	for _, dir := range directions {
		newLoc := Location{loc[0] + dir[0], loc[1] + dir[1]}
		newLoc, ok := m.inBounds(newLoc)
		if ok == false {
			continue
		}
		if m.isRock(newLoc) {
			continue
		}
		next = append(next, newLoc)
	}
	return next
}

func (m *GardenMap) inBounds(loc Location) (Location, bool) {
	if m.infinite == false {
		return loc, loc[0] >= 0 && loc[1] >= 0 && loc[0] <= len(m.Map)-1 && loc[1] <= len(m.Map[0])-1
	}
	height := len(m.Map)
	width := len(m.Map[0])

	// Out of map y
	var addRow int
	if loc[0] < 0 {
		addRow = height
	}
	if loc[0] >= height {
		addRow = -height
	}

	// Out of map x
	var addCol int
	if loc[1] < 0 {
		addCol = width
	}
	if loc[1] >= width {
		addCol = -width
	}

	loc[0] += addRow
	loc[1] += addCol

	if loc[1] >= len(m.Map[0]) || loc[0] >= len(m.Map) || loc[1] < 0 || loc[0] < 0 {
		fmt.Print()
	}

	return loc, true
}

func (m *GardenMap) isRock(loc Location) bool {
	return m.Map[loc[0]][loc[1]] == '#'
}
