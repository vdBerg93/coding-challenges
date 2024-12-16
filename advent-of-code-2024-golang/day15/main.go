package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//go:embed input.txt
var input []byte

func main() {
	//fmt.Println("Part1:", part1(input)) // 1318523
	fmt.Println("Part2:", part2(input)) // 1337648
}

func part1(input []byte) int {
	m, moves := parseInput(input)
	m.findStart()
	for _, move := range moves {
		m.move2(move)
	}

	return m.score()
}

type Point struct {
	x, y int
}

type Map struct {
	m   [][]rune
	pos Point
}

func (m *Map) findStart() {
	for y, line := range m.m {
		for x, c := range line {
			if c == '@' {
				m.pos = Point{x, y}
			}
		}
	}
}

func (m *Map) get(p Point) rune {
	return m.m[p.y][p.x]
}

func (m *Map) set(p Point, r rune) {
	m.m[p.y][p.x] = r
}

func (m *Map) score() int {
	var s int
	for y, row := range m.m {
		for x, c := range row {
			if c == 'O' || c == '[' {
				si := 100*y + x
				s += si
			}
		}
	}
	return s
}

func (m *Map) occupied(p Point) (bool, bool) {
	occupied, fixed := false, false
	v := m.get(p)
	switch v {
	case 'O':
		occupied = true
	case '#':
		occupied = true
		fixed = true
	case '[', ']':
		occupied = true
	case '.':
		return false, false
	default:
		log.Panicf("invalid char %v", v)
	}
	return occupied, fixed
}

func step(pos, dir Point) Point {
	return Point{pos.x + dir.x, pos.y + dir.y}
}

func (m *Map) String() string {
	b := strings.Builder{}
	for _, line := range m.m {
		b.WriteString(string(line) + "\n")
	}
	return b.String()
}

func (m *Map) move(pos, dir Point) bool {
	previous := m.get(pos)
	next := step(m.pos, dir)
	occupied, fixed := m.occupied(next)
	if fixed {
		return false
	}

	if occupied && !m.moveBox(next, dir) {
		return false
	}

	m.set(next, previous)
	m.set(m.pos, '.')
	m.pos = next
	return true
}

func (m *Map) upgradeToPart2() {
	type filter struct {
		old, new string
	}
	filters := []filter{
		{"O", "[]"},
		{"#", "##"},
		{".", ".."},
		{"@", "@."},
	}

	for i, row := range m.m {
		r := string(row)
		for _, f := range filters {
			r = strings.ReplaceAll(r, f.old, f.new)
		}
		m.m[i] = []rune(r)
	}
}

const linesep = "\r\n"

func parseInput(input []byte) (Map, []rune) {
	f := strings.Split(string(input), linesep+linesep)
	if len(f) < 2 {
		log.Printf("invalid length %v", len(f))
	}
	return parseMap(f[0]), parseMoves(f[1])
}

func parseMap(data string) Map {
	rows := strings.Split(data, linesep)
	m := make([][]rune, 0, len(rows))
	for _, r := range rows {
		m = append(m, []rune(r))
	}
	return Map{
		m: m,
	}
}

func parseMoves(data string) []rune {
	var moves []rune
	for _, r := range strings.Split(data, linesep) {
		moves = append(moves, []rune(r)...)
	}
	return moves
}

var (
	right = Point{1, 0}
	left  = Point{-1, 0}
	up    = Point{0, -1}
	down  = Point{0, 1}
)

func getMove(c rune) Point {
	switch c {
	case '>':
		return right
	case '<':
		return left
	case '^':
		return up
	case 'v':
		return down
	default:
		log.Printf("invalid char %v", string(c))
		return Point{}
	}
}

func part2(input []byte) int {
	m, moves := parseInput(input)
	m.upgradeToPart2()
	m.findStart()

	for _, move := range moves {
		m.move2(move)
		os.WriteFile("output.txt", []byte(m.String()), 744)
		time.Sleep(time.Second / 10)
	}

	return m.score()
}

func (m *Map) move2(move rune) {
	dir := getMove(move)
	switch dir {
	case right, left:
		m.move(m.pos, dir)
	case up, down:
		if m.moveOK(m.pos, dir) {
			m.moveExec(m.pos, dir)
			m.pos = step(m.pos, dir)
		}
	}
}

func (m *Map) isDoubleBox(p Point) bool {
	v := m.get(p)
	return v == '[' || v == ']'
}

// moveBox - returns moveable
func (m *Map) moveBox(a, dir Point) bool {
	next := step(a, dir)
	occupied, fixed := m.occupied(next)
	if fixed {
		return false
	}

	if occupied && !m.moveBox(next, dir) {
		return false
	}

	m.set(next, m.get(a))
	m.set(a, '.')

	return true
}

func (m *Map) getOtherHalf(b Point) Point {
	switch m.get(b) {
	case '[':
		return Point{b.x + 1, b.y}
	case ']':
		return Point{b.x - 1, b.y}
	case '.':
		log.Panicf("impossible")
		return Point{}
	default:
		log.Panicf("invalid character %v for point %+v", m.get(b), b)
		return Point{}
	}
}

func (m *Map) moveExec(pos, dir Point) {
	next := step(pos, dir)
	if m.isDoubleBox(next) {
		m.moveExec(m.getOtherHalf(next), dir)
		m.moveExec(next, dir)
	}
	if m.get(next) == 'O' {
		m.moveExec(next, dir)
	}
	m.set(next, m.get(pos))
	m.set(pos, '.')
}

func (m *Map) moveOK(pos, dir Point) bool {
	next := step(pos, dir)
	occupied, fixed := m.occupied(next)
	if fixed {
		return false
	}
	if !occupied {
		return true
	}

	if m.isDoubleBox(next) {
		other := m.getOtherHalf(next)
		ok := m.moveOK(other, dir)
		if !ok {
			return false
		}
	}

	return m.moveOK(next, dir)
}
