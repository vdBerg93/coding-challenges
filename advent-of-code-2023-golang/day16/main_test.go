package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"unicode"
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
	got := Solve(testData, false)
	want1 := 46
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, false))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, true)
	want1 := 51
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve(data, true))
}

func Solve(data []byte, tune bool) int {
	grid := NewGrid(data)

	var startLasers []Laser
	if tune {
		startLasers = grid.StartLocations()
	} else {
		startLasers = append(startLasers, Laser{
			Pos:     Point{-1, 0},
			Heading: Point{1, 0},
		})
	}

	mostIlluminated := 0
	// Brute force all start locations
	for _, startLaser := range startLasers {
		lasers := []Laser{startLaser}
		for len(lasers) > 0 {
			var newLasers []Laser
			for _, l := range lasers {
				newLasers = append(newLasers, grid.Move(l)...)
			}
			lasers = newLasers
		}
		illuminated := grid.GetIlluminatedCount()
		if illuminated > mostIlluminated {
			mostIlluminated = illuminated
		}
		grid.Reset()
	}

	return mostIlluminated
}

type Grid struct {
	mirrors [][]Tile
	starts  []Laser
}

type Tile struct {
	Char                     rune
	Energy                   int
	North, East, South, West bool
}

type Laser struct {
	Pos, Heading Point
}

type Point struct {
	x, y int
}

// Move returns the lasers after the next step
func (g *Grid) Move(laser Laser) []Laser {
	newPos := Point{
		x: laser.Pos.x + laser.Heading.x,
		y: laser.Pos.y + laser.Heading.y,
	}

	if newPos.x < 0 || newPos.x >= len(g.mirrors[0]) || newPos.y < 0 || newPos.y >= len(g.mirrors) {
		return nil // Laser went off grid
	}

	ok := g.UpdateTile(newPos, laser.Heading)
	if !ok {
		return nil
	}

	char := g.mirrors[newPos.y][newPos.x].Char
	if char == '-' && laser.Heading.y != 0 {
		return []Laser{
			{Pos: newPos, Heading: Point{-1, 0}},
			{Pos: newPos, Heading: Point{1, 0}},
		}
	}

	if char == '|' && laser.Heading.x != 0 {
		return []Laser{
			{Pos: newPos, Heading: Point{0, -1}},
			{Pos: newPos, Heading: Point{0, 1}},
		}
	}

	if char == '\\' {
		if laser.Heading.x == 1 {
			return []Laser{{newPos, Point{0, 1}}}
		} else if laser.Heading.x == -1 {
			return []Laser{{newPos, Point{0, -1}}}
		} else if laser.Heading.y == -1 {
			return []Laser{{newPos, Point{-1, 0}}}
		} else if laser.Heading.y == 1 {
			return []Laser{{newPos, Point{1, 0}}}
		} else {
			panic("impossible")
		}
	}

	if char == '/' {
		if laser.Heading.x == 1 {
			return []Laser{{newPos, Point{0, -1}}}
		} else if laser.Heading.x == -1 {
			return []Laser{{newPos, Point{0, 1}}}
		} else if laser.Heading.y == -1 {
			return []Laser{{newPos, Point{1, 0}}}
		} else if laser.Heading.y == 1 {
			return []Laser{{newPos, Point{-1, 0}}}
		} else {
			panic("impossible")
		}
	}

	return []Laser{{newPos, laser.Heading}}

}

// UpdateTile - returns false if already passed from this heading
func (g *Grid) UpdateTile(p Point, h Point) bool {
	t := &g.mirrors[p.y][p.x]
	var visited bool
	switch h {
	case Point{1, 0}:
		visited = t.East
		t.East = true
	case Point{-1, 0}:
		visited = t.West
		t.West = true
	case Point{0, 1}:
		visited = t.South
		t.South = true
	case Point{0, -1}:
		visited = t.North
		t.North = true
	}
	if visited {
		return false
	}
	t.Energy++

	if t.Char == '.' {
		g.mirrors[p.y][p.x].Char = '1'
		return true
	}
	if unicode.IsNumber(t.Char) {
		g.mirrors[p.y][p.x].Char++
		return true
	}
	return true
}

func NewGrid(data []byte) Grid {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var grid Grid
	for scanner.Scan() {
		row := scanner.Text()
		var tiles []Tile
		for _, char := range row {
			tiles = append(tiles, Tile{char, 0, false, false, false, false})
		}
		grid.mirrors = append(grid.mirrors, tiles)
	}

	return grid
}

func (g *Grid) StartLocations() []Laser {
	var starts []Laser
	for x := 0; x < len(g.mirrors[0]); x++ {
		starts = append(starts, Laser{Point{x, -1}, Point{0, 1}})
		starts = append(starts, Laser{Point{x, len(g.mirrors)}, Point{0, -1}})
	}
	for y := 0; y < len(g.mirrors[0]); y++ {
		starts = append(starts, Laser{Point{-1, y}, Point{1, 0}})
		starts = append(starts, Laser{Point{len(g.mirrors[0]), y}, Point{-1, 0}})
	}
	return starts
}

func (g *Grid) GetIlluminatedCount() int {
	energy := 0
	for _, row := range g.mirrors {
		for _, tile := range row {
			if tile.Energy > 0 {
				energy++
			}
		}
	}
	return energy
}

func (g *Grid) Reset() {
	for i := 0; i < len(g.mirrors); i++ {
		for j := 0; j < len(g.mirrors[0]); j++ {
			g.mirrors[i][j].Energy = 0
			g.mirrors[i][j].East = false
			g.mirrors[i][j].West = false
			g.mirrors[i][j].North = false
			g.mirrors[i][j].South = false
		}
	}
}
