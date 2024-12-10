package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
)

//go:embed input.txt
var input []byte

func main() {
	m := parseInput(bytes.NewReader(input))
	fmt.Println("Part1: ", solve(m, getAntiZones))
	//2500 is too high
	m = parseInput(bytes.NewReader(input))
	fmt.Println("Part2: ", solve(m, getAntiZones2))
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

func solve(m [][]rune, zoneLocator func([][]rune, Point, Point) []Point) int {
	sortedAntennas := getAntennas(m)
	allAz := make(map[Point]struct{})
	for _, locations := range sortedAntennas {
		for _, a := range locations {
			for _, b := range locations {
				if a.Equal(b) {
					continue
				}
				azs := zoneLocator(m, a, b)
				for _, az := range azs {
					if isOnMap(m, az) {
						allAz[az] = struct{}{}
						m[az.y][az.x] = '#'
					}
				}
			}
		}
	}
	return len(allAz)
}

type Point struct {
	x, y int
}

func (p *Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func getAntennas(m [][]rune) map[rune][]Point {
	antennas := make(map[rune][]Point)
	for y, row := range m {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			antennas[cell] = append(antennas[cell], Point{x, y})
		}
	}
	return antennas
}

func getAntiZones(_ [][]rune, a, b Point) []Point {
	dx := b.x - a.x
	dy := b.y - a.y

	az := make([]Point, 0, 2)
	for _, p := range []Point{a, b} {
		loc1 := Point{x: p.x + dx, y: p.y + dy}
		if notAtAntenna(a, b, loc1) {
			az = append(az, loc1)
		}
		loc2 := Point{x: p.x - dx, y: p.y - dy}
		if notAtAntenna(a, b, loc2) {
			az = append(az, loc2)
		}
	}

	return az
}

func notAtAntenna(a, b, az Point) bool {
	return !az.Equal(a) && !az.Equal(b)
}

func isOnMap(m [][]rune, loc Point) bool {
	return loc.x >= 0 && loc.x < len(m[0]) && loc.y >= 0 && loc.y < len(m)
}

func getAntiZones2(m [][]rune, a, b Point) []Point {
	dx := b.x - a.x
	dy := b.y - a.y

	az := make([]Point, 0, 2)
	for _, p := range []Point{a, b} {
		az = append(az, searchDirection(m, p, dx, dy, 1)...)
		az = append(az, searchDirection(m, p, dx, dy, -1)...)
	}

	return az
}

func searchDirection(m [][]rune, p Point, dx, dy int, dir int) []Point {
	var az []Point
	for i := dir; ; i += dir {
		xi := p.x + i*dx
		yi := p.y + i*dy
		if !isOnMap(m, Point{xi, yi}) {
			break
		}
		az = append(az, Point{xi, yi})
	}
	return az
}
