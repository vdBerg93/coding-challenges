package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"reflect"
)

//go:embed input.txt
var input []byte

func main() {
	data := parseInput(bytes.NewReader(input))
	fmt.Println("part1: ", part1(data))
	data = parseInput(bytes.NewReader(input))
	fmt.Println("part2: ", part2(data))
}

func parseInput(r io.Reader) [][]rune {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var rows [][]rune
	for scanner.Scan() {
		rows = append(rows, []rune(scanner.Text()))
	}
	return rows
}

type Pose struct {
	Position Point
	Heading  Point
}

type Point struct {
	x, y int
}

type Puzzle struct {
	Map       [][]rune
	StartPose Pose
}

func (p *Puzzle) Get(pos Point) rune {
	return p.Map[pos.y][pos.x]
}

func (p *Puzzle) Set(pos Point, v rune) {
	p.Map[pos.y][pos.x] = v
}

func move(pose Pose) Pose {
	return Pose{
		Position: Point{
			x: pose.Position.x + pose.Heading.x,
			y: pose.Position.y + pose.Heading.y,
		},
		Heading: pose.Heading,
	}
}

func rotate(pose Pose) Pose {
	return Pose{
		Position: pose.Position,
		Heading: Point{
			x: -pose.Heading.y,
			y: pose.Heading.x,
		},
	}
}

func (p *Puzzle) LeftMap(point Point) bool {
	if point.x < 0 || point.x >= len(p.Map[0]) || point.y < 0 || point.y >= len(p.Map) {
		return true
	}
	return false
}

func (p *Puzzle) Blocked(point Point) bool {
	return p.Get(point) == '#'
}

func (p *Puzzle) NextPose(pose Pose) (Pose, bool) {
	next := move(pose)
	if p.LeftMap(next.Position) {
		return next, false
	}
	if p.Blocked(next.Position) {
		next = rotate(pose)
	}

	return next, true
}

func part1(data [][]rune) int {
	puzzle := &Puzzle{
		Map: data,
	}

	positions := puzzle.Solve1()

	return len(positions)
}

func (p *Puzzle) FindStart() {
	for y, row := range p.Map {
		for x, cell := range row {
			if cell != '^' {
				continue
			}
			p.StartPose = Pose{
				Position: Point{x: x, y: y},
				Heading:  Point{x: 0, y: -1},
			}
			return
		}
	}
	panic("start not found")
}

func (p *Puzzle) Solve1() map[Point]struct{} {
	p.FindStart()
	path := make(map[Point]struct{})
	pose := p.StartPose
	for {
		next, ok := p.NextPose(pose)
		if !ok {
			break
		}
		pose = next
		path[pose.Position] = struct{}{}
	}
	return path
}

func part2(data [][]rune) int {
	puzzle := &Puzzle{
		Map: data,
	}
	path := puzzle.Solve1()
	return puzzle.Solve2(path)
}

func (p *Puzzle) Solve2(path map[Point]struct{}) int {
	delete(path, p.StartPose.Position)

	var count int

	for point := range path {
		p.Set(point, '#')
		if p.DetectLoop(len(path)) {
			count++
		}
		p.Set(point, '.')
	}

	return count
}

func (p *Puzzle) DetectLoop(max int) bool {
	pose := p.StartPose

	for i := 0; i < 2*max; i++ {
		next, leftMap := p.NextPose(pose)
		if !leftMap {
			return false
		}

		if reflect.DeepEqual(next, p.StartPose) {
			return true
		}
		pose = next
	}
	return true
}
