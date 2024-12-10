package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"maps"
)

//go:embed input.txt
var input []byte

func main() {
	data := parseInput(bytes.NewReader(input))

	got1 := part1(data)
	fmt.Println("Part1: ", got1)
	got2 := part2(data)
	fmt.Println("Part2: ", got2)

	want1 := 510
	if got1 != want1 {
		log.Fatalf("Part1: want %d, got %d\n", want1, got1)
	}
	want2 := 1058
	if got2 != want2 {
		log.Fatalf("Part2: want %d, got %d\n", want2, got2)
	}
}

func parseInput(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var rows [][]int
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]int, 0, len(text))
		for _, char := range text {
			row = append(row, runeToInt(char))
		}
		rows = append(rows, row)
	}
	return rows
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func part1(m [][]int) int {
	peaks, _ := solve(m)

	return peaks
}

type Puzzle struct {
	m     [][]int
	heads map[Point]struct{}
}

func (p *Puzzle) reset() {
	p.heads = make(map[Point]struct{})
}

func solve(m [][]int) (int, int) {
	puzzle := Puzzle{
		m: m,
	}

	var (
		trails int
		peaks  int
	)

	valleys := puzzle.findAll(0)
	for _, head := range valleys {
		puzzle.reset()
		trailsI := puzzle.DFS(head, map[Point]struct{}{})
		peaks += len(puzzle.heads)
		trails += trailsI
	}

	return peaks, trails
}

func part2(m [][]int) int {
	_, trails := solve(m)
	return trails
}

func (p *Puzzle) Get(loc Point) int {
	return p.m[loc.y][loc.x]
}

func (p *Puzzle) GetNextUnvisited(visited map[Point]struct{}, node Point) []Point {
	nextPositions := make([]Point, 0, 4)
	for _, next := range p.Next(node) {
		_, ok := visited[next]
		if !ok {
			nextPositions = append(nextPositions, next)
		}
	}
	return nextPositions
}

func (p *Puzzle) DFS(node Point, visited map[Point]struct{}) int {
	var trails int
	for {
		if p.Get(node) == 9 {
			if _, ok := p.heads[node]; !ok {
				p.heads[node] = struct{}{}
			}
		}

		visited[node] = struct{}{}
		nextPositions := p.GetNextUnvisited(visited, node)

		if len(nextPositions) == 1 {
			node = nextPositions[0]
			continue
		} else if len(nextPositions) > 1 {
			for _, next := range nextPositions {
				visitedCp := make(map[Point]struct{})
				maps.Copy(visitedCp, visited)
				trailsI := p.DFS(next, visitedCp)
				trails += trailsI
			}
			return trails
		} else {
			if p.Get(node) == 9 {
				trails++
			}
			return trails
		}
	}
}

func (p *Puzzle) Next(last Point) []Point {
	nexts := make([]Point, 0, 4)
	if next, ok := p.MoveOK(last, right); ok {
		nexts = append(nexts, next)
	}
	if next, ok := p.MoveOK(last, left); ok {
		nexts = append(nexts, next)
	}
	if next, ok := p.MoveOK(last, up); ok {
		nexts = append(nexts, next)
	}
	if next, ok := p.MoveOK(last, down); ok {
		nexts = append(nexts, next)
	}

	return nexts
}

func (p *Puzzle) MoveOK(last Point, dir Point) (Point, bool) {
	next := last.Move(dir)
	if next.x < 0 || next.x >= len(p.m[0]) || next.y < 0 || next.y >= len(p.m) {
		return next, false
	}
	delta := p.Get(next) - p.Get(last)
	if delta > 1 || delta <= 0 {
		return next, false
	}
	return next, true
}

var (
	right = Point{1, 0}
	down  = Point{0, 1}
	left  = Point{-1, 0}
	up    = Point{0, -1}
)

type Point struct {
	x, y int
}

func (p *Point) Move(dir Point) Point {
	return Point{
		x: p.x + dir.x,
		y: p.y + dir.y,
	}
}

func (p *Puzzle) findAll(target int) []Point {
	var heads []Point
	for y, row := range p.m {
		for x, cell := range row {
			if cell == target {
				heads = append(heads, Point{x, y})
			}
		}
	}
	if len(heads) == 0 {
		panic("not found")
	}
	return heads
}
