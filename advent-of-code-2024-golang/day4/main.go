package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	data := readInput(input)
	fmt.Println("part1 solution: ", part1(data))
	fmt.Println("part2 solution: ", part2(data))
}

func readInput(in []byte) [][]rune {
	rows := strings.Split(string(in), "\r\n")
	data := make([][]rune, len(rows))
	for i, row := range rows {
		data[i] = []rune(row)
	}
	return data
}

type point struct {
	x, y int
}

type Puzzle struct {
	Evaluate      func(x, y int)
	Grid          [][]rune
	Directions    []point
	CenterOffsets []int
	Target        string
	CenterIdx     int
	total         int
}

func (P *Puzzle) get(x, y int) rune {
	return P.Grid[y][x]
}

func part1(data [][]rune) int {
	puzzle := Puzzle{
		Grid: data,
		Directions: []point{
			{1, 0},   // right
			{-1, 0},  // left
			{0, -1},  // up
			{0, 1},   // down
			{1, -1},  // up and right
			{-1, -1}, // up and left
			{-1, 1},  // down and right
			{1, 1},   // down and left
		},
		Target:        "XMAS",
		CenterIdx:     0,
		CenterOffsets: []int{0, 1, 2, 3},
	}
	puzzle.Evaluate = puzzle.evaluate
	solution := puzzle.solve()
	return solution
}

func part2(data [][]rune) int {
	puzzle := Puzzle{
		Grid: data,
		Directions: []point{
			{1, -1},  // up and right
			{-1, -1}, // up and left
			{-1, 1},  // down and right
			{1, 1},   // down and left
		},
		Target:        "MAS",
		CenterIdx:     1,
		CenterOffsets: []int{-1, 0, 1},
	}
	puzzle.Evaluate = puzzle.evaluatePart2
	solution := puzzle.solve()
	return solution
}

func (P *Puzzle) solve() int {
	size_y := len(P.Grid)
	size_x := len(P.Grid[0])
	for y := 0; y < size_y; y++ {
		for x := 0; x < size_x; x++ {
			P.Evaluate(x, y)
		}
	}
	return P.total
}

func (P *Puzzle) evaluate(x, y int) {
	for _, d := range P.Directions {
		P.total += P.checkDirection(x, y, d)
	}
}

func (P *Puzzle) evaluatePart2(x, y int) {
	var count int
	for _, d := range P.Directions {
		count += P.checkDirection(x, y, d)
	}
	if count >= 2 {
		P.total++
	}
}

func (P *Puzzle) checkDirection(x, y int, dir point) int {
	var ok int
	for _, offset := range P.CenterOffsets {
		xi, yi := x+dir.x*offset, y+dir.y*offset
		if !P.checkBounds(xi, yi) {
			return 0
		}
		got := P.get(xi, yi)
		want := rune(P.Target[P.CenterIdx+offset])
		if got != want {
			return 0
		}
		ok++
	}
	if ok < len(P.Target) {
		return 0
	}
	return 1
}

func (P *Puzzle) checkBounds(xi, yi int) bool {
	if xi < 0 || xi >= len(P.Grid[0]) {
		return false
	}
	if yi < 0 || yi >= len(P.Grid) {
		return false
	}
	return true
}
