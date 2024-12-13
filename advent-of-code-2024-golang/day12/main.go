package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"maps"
)

//go:embed input.txt
var input []byte

func main() {

	got1 := part1(input)

	fmt.Println("Part1: ", got1)
	want1 := 1546338
	if want1 != got1 {
		log.Fatalf("want %d, got %d", want1, got1)
	}
	got2 := part2(input)
	fmt.Println("Part2: ", got2)
	want2 := 978590
	if got2 != want2 {
		log.Fatalf("want %d, got %d", want2, got2)
	}

}

func parseInput(input []byte) [][]rune {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)
	var M [][]rune
	for scanner.Scan() {
		text := scanner.Text()
		M = append(M, []rune(text))
	}
	return M
}

func get(m [][]rune, p [2]int) rune {
	return m[p[0]][p[1]]
}

func part1(input []byte) int {
	m := parseInput(input)

	sortedAreas := getAreas(m)
	var total int
	for plant, areas := range sortedAreas {
		plant = plant
		for _, area := range areas {
			pi := perimeter(area)
			total += len(area) * pi
		}
	}
	return total
}

func getAreas(m [][]rune) map[rune]map[int]map[[2]int]struct{} {
	var uid int

	sortedAreas := make(map[rune]map[int]map[[2]int]struct{})

	for i, _ := range m {
		for j, _ := range m[i] {
			point := [2]int{i, j}
			plant := get(m, point)
			if plant == '.' {
				continue
			}
			areas, _ := sortedAreas[plant]
			var keys []int
			for _, n := range getNeighbours(m, [2]int{i, j}) {
				for key, val := range areas {
					if pointInArea(val, n) {
						keys = append(keys, key)
					}
				}
			}
			keys = unique(keys)
			if areas == nil {
				areas = make(map[int]map[[2]int]struct{})
			}
			if len(keys) == 0 {
				areas[uid] = map[[2]int]struct{}{point: {}}
				uid++
				sortedAreas[plant] = areas
				continue
			} else if len(keys) == 1 {
				areas[keys[0]][point] = struct{}{}
				sortedAreas[plant] = areas
			} else {
				for idx := len(keys) - 1; idx > 0; idx-- {
					maps.Copy(areas[keys[0]], areas[keys[idx]])
					delete(areas, keys[idx])
				}
				areas[keys[0]][point] = struct{}{}
				sortedAreas[plant] = areas
			}
		}
	}
	return sortedAreas
}

func perimeter(area map[[2]int]struct{}) int {
	var p int
	for p1 := range area {
		var adj int
		for p2 := range area {
			if isAdjacent(p1, p2) {
				adj++
			}
		}
		p += 4 - adj
	}
	return p
}

// isAdjacent checks if two points are adjacent on left, right, top, or bottom.
func isAdjacent(p1, p2 [2]int) bool {
	// Calculate the differences in x and y coordinates
	dx := abs(p1[0] - p2[0])
	dy := abs(p1[1] - p2[1])

	// Check adjacency condition: either directly horizontal or vertical
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

// abs returns the absolute value of an integer.
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func unique(s []int) []int {
	vals := make(map[int]struct{})
	for _, val := range s {
		vals[val] = struct{}{}
	}
	out := make([]int, 0, len(vals))
	for val := range vals {
		out = append(out, val)
	}
	return out
}

func pointInArea(perimeter map[[2]int]struct{}, p [2]int) bool {
	for pi := range perimeter {
		if pi == p {
			return true
		}
	}
	return false
}

func getNeighbours(m [][]rune, p [2]int) [][2]int {
	nexts := make([][2]int, 0, 4)
	if next, ok := isNeighbor(m, p, [2]int{p[0], p[1] + 1}); ok {
		nexts = append(nexts, next)
	}
	if next, ok := isNeighbor(m, p, [2]int{p[0], p[1] - 1}); ok {
		nexts = append(nexts, next)
	}
	if next, ok := isNeighbor(m, p, [2]int{p[0] + 1, p[1]}); ok {
		nexts = append(nexts, next)
	}
	if next, ok := isNeighbor(m, p, [2]int{p[0] - 1, p[1]}); ok {
		nexts = append(nexts, next)
	}

	return nexts
}

func isNeighbor(m [][]rune, last, next [2]int) ([2]int, bool) {
	if next[1] < 0 || next[1] >= len(m[0]) || next[0] < 0 || next[0] >= len(m) {
		return next, false
	}
	return next, get(m, last) == get(m, next)
}

func part2(input []byte) int {
	m := parseInput(input)
	sortedAreas := getAreas(m)
	var total int
	for plant, areas := range sortedAreas {
		for _, area := range areas {
			corners := 0
			for p := range area {
				i := getCorners(m, p)
				corners += i
			}
			fmt.Printf("Plant: %v found %d fence corners\n", string(plant), corners)
			total += len(area) * corners
		}
	}

	return total
}

func topLeft(p [2]int) int {
	if p[0] == 0 && p[1] == 0 {
		return 1
	}
	return 0
}

func bottomLeft(m [][]rune, x, y int) int {
	if x == 0 && y == len(m)-1 {
		return 1
	}
	return 0
}

func bottomRight(m [][]rune, x, y int) int {
	if x == len(m[0])-1 && y == len(m)-1 {
		return 1
	}
	return 0
}

func topRight(m [][]rune, x, y int) int {
	if x == len(m[0])-1 && y == 0 {
		return 1
	}
	return 0
}

func topLeftOutside(m [][]rune, x, y int, T rune) int {
	if (x > 0 && y > 0 && m[y][x-1] != T && m[y-1][x] != T) ||
		(x > 0 && y == 0 && m[y][x-1] != T) ||
		(x == 0 && y > 0 && m[y-1][x] != T) {
		return 1
	}
	return 0
}

func topLeftInside(m [][]rune, x, y int, T rune) int {
	if x < len(m[0])-1 && y < len(m)-1 && m[y][x+1] == T && m[y+1][x] == T && m[y+1][x+1] != T {
		return 1
	}
	return 0
}

func topRightOutsideCorner(m [][]rune, x, y int, T rune) int {
	if (x < len(m[0])-1 && y > 0 && m[y][x+1] != T && m[y-1][x] != T) ||
		(x < len(m[0])-1 && y == 0 && m[y][x+1] != T) ||
		(x == len(m[0])-1 && y > 0 && m[y-1][x] != T) {
		return 1
	}
	return 0
}

func topRightInsideCorner(m [][]rune, x, y int, T rune) int {
	if x > 0 && y < len(m)-1 && m[y][x-1] == T && m[y+1][x] == T && m[y+1][x-1] != T {
		return 1
	}
	return 0
}

func bottomLeftOutsideCorner(m [][]rune, x, y int, T rune) int {
	if (x > 0 && y < len(m)-1 && m[y][x-1] != T && m[y+1][x] != T) ||
		(x > 0 && y == len(m)-1 && m[y][x-1] != T) ||
		(x == 0 && y < len(m)-1 && m[y+1][x] != T) {
		return 1
	}
	return 0
}

func bottomLeftInsideCorner(m [][]rune, x, y int, T rune) int {
	if x < len(m[0])-1 && y > 0 && m[y][x+1] == T && m[y-1][x] == T && m[y-1][x+1] != T {
		return 1
	}
	return 0
}

func bottomRightOutsideCorner(m [][]rune, x, y int, T rune) int {
	if (x < len(m[0])-1 && y < len(m)-1 && m[y][x+1] != T && m[y+1][x] != T) ||
		(x < len(m[0])-1 && y == len(m)-1 && m[y][x+1] != T) ||
		(x == len(m[0])-1 && y < len(m)-1 && m[y+1][x] != T) {
		return 1
	}
	return 0
}

func bottomRightInsideCorner(m [][]rune, x, y int, T rune) int {
	if x > 0 && y > 0 && m[y][x-1] == T && m[y-1][x] == T && m[y-1][x-1] != T {
		return 1
	}
	return 0
}

func getCorners(m [][]rune, current [2]int) int {
	count := 0
	x, y := current[1], current[0]
	T := m[y][x]

	count += topLeft(current)
	count += topRight(m, x, y)
	count += bottomLeft(m, x, y)
	count += bottomRight(m, x, y)
	count += topLeftOutside(m, x, y, T)
	count += topLeftInside(m, x, y, T)
	count += topRightOutsideCorner(m, x, y, T)
	count += topRightInsideCorner(m, x, y, T)
	count += bottomLeftOutsideCorner(m, x, y, T)
	count += bottomLeftInsideCorner(m, x, y, T)
	count += bottomRightOutsideCorner(m, x, y, T)
	count += bottomRightInsideCorner(m, x, y, T)

	return count
}
