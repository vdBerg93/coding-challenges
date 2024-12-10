package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"reflect"
)

//go:embed input.txt
var input []byte

//go:embed sample3.txt
var example []byte

//go:embed sample.txt
var sample []byte

//go:embed sample2.txt
var sample2 []byte

func main() {
	want1 := 36
	data := parseInput(bytes.NewReader(example))
	if got := part1(data); got != want1 {
		log.Fatalf("part 1 example: want %v, got %v", want1, got)
	}

	data = parseInput(bytes.NewReader(input))
	fmt.Println("Part1: ", part1(data))
}

func parseInput(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var rows [][]int
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]int, 0, len(text))
		for _, char := range text {
			if char == '.' {
				row = append(row, 99)
			} else {
				row = append(row, runeToInt(char))
			}
		}
		rows = append(rows, row)
	}
	return rows
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func intToRune(i int) rune {
	return rune(i + '0')
}

func part1(m [][]int) int {
	puzzle := Puzzle{
		m: m,
	}
	puzzle.Save(map[Point]struct{}{})

	var peakCount []int

	valleys := puzzle.findAll(0)
	for i, head := range valleys {
		puzzle.Save(map[Point]struct{}{})
		visited, peaks := puzzle.DFS(head, map[Point]struct{}{})
		puzzle.Save(visited)
		fmt.Printf("Valley %d (x:%d,y:%d) found %d peaks\n", i, head.x, head.y, len(peaks))
		peakCount = append(peakCount, len(peaks))
	}

	var score int
	for _, sol := range peakCount {
		score += sol
	}
	return score
}

type Puzzle struct {
	m [][]int
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

func (p *Puzzle) DFS(node Point, visited map[Point]struct{}) (map[Point]struct{}, map[Point]struct{}) {
	//fmt.Println("start recursion")

	heads := make(map[Point]struct{})
	for {
		if p.Get(node) == 9 {
			if _, ok := heads[node]; !ok {
				heads[node] = struct{}{}
				fmt.Printf("found peak at (x:%d, y:%d)\n", node.x, node.y)
				print()
			}
		}

		visited[node] = struct{}{}
		nextPositions := p.GetNextUnvisited(visited, node)
		p.Save(visited)
		if len(nextPositions) == 1 {
			node = nextPositions[0]
			continue
		} else if len(nextPositions) > 1 {
			for _, next := range nextPositions {
				visitI, headsI := p.DFS(next, visited)
				maps.Copy(heads, headsI)
				maps.Copy(visitI, visited)
				p.Save(visitI)
				//fmt.Printf("Found %d heads\n", len(headsI))
			}
			return visited, heads
		} else {
			p.Save(visited)
			return visited, heads
		}
	}
}

func (p *Puzzle) Save(visited map[Point]struct{}) {
	cp := make([][]rune, 0, len(p.m))
	for _, row := range p.m {
		rowCp := make([]rune, 0, len(row))
		for _, val := range row {
			if val == 99 {
				rowCp = append(rowCp, '.')
			} else {
				rowCp = append(rowCp, intToRune(val))
			}
		}
		cp = append(cp, rowCp)
	}
	for point := range visited {
		cp[point.y][point.x] = 'o'
	}

	file, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, row := range cp {
		_, _ = file.WriteString(string(row) + "\n")
	}

}

func (p *Puzzle) Next(last Point) []Point {
	if reflect.DeepEqual(last, Point{5, 13}) {
		print()
	}
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

	for _, pt := range nexts {
		if reflect.DeepEqual(pt, Point{4, 6}) {
			print()
		}
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
