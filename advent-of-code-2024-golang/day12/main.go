package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

//go:embed example3.txt
var example3 []byte

//go:embed example4.txt
var example4 []byte

//go:embed input.txt
var input []byte

func main() {
	//
	//got1 := part1(input)
	//
	//fmt.Println("Part1: ", got1)
	//want1 := 1546338
	//if want1 != got1 {
	//	log.Fatalf("want %d, got %d", want1, got1)
	//}
	fmt.Println("Part2: ", part2(example1))

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
			//fmt.Printf("Area %s: perimeter: %d, points: %d | %v\n", string(plant), pi, len(area), area)
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
				//fmt.Printf("Initialized area uid %d for plant %v point %v\n", uid, string(plant), point)
				continue
			} else if len(keys) == 1 {
				//fmt.Printf("Appending area for plant %v point %v\n", string(plant), point)
				areas[keys[0]][point] = struct{}{}
				sortedAreas[plant] = areas
			} else {
				//fmt.Printf("Merging %d area for plant %v point %v\n", len(keys), string(plant), point)
				for idx := len(keys) - 1; idx > 0; idx-- {
					maps.Copy(areas[keys[0]], areas[keys[idx]])
					delete(areas, keys[idx])
					//fmt.Printf("merge and delete key: %d\n", keys[idx])
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
	if len(out) != len(s) {
		//fmt.Printf("Dropped %d keys", len(s)-len(out))
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

func getAdjecent(p [2]int) [][2]int {
	return [][2]int{
		{p[0], p[1] + 1},
		{p[0], p[1] - 1},
		{p[0] + 1, p[1]},
		{p[0] - 1, p[1]},
	}
}

func isNeighbor(m [][]rune, last, next [2]int) ([2]int, bool) {
	if next[1] < 0 || next[1] >= len(m[0]) || next[0] < 0 || next[0] >= len(m) {
		return next, false
	}
	return next, get(m, last) == get(m, next)
}

type Point struct {
	x, y    int
	visited bool
}

func getFencePoints(m [][]rune, area map[[2]int]struct{}) [][2]int {
	var fencePoints [][2]int
	for p1 := range area {
		N := getAdjecent(p1)
		for _, n := range N {
			_, ok := area[n]
			if ok {
				continue
			}
			fencePoints = append(fencePoints, n)
		}
	}

	return fencePoints
}

func getFencePoints2(area map[[2]int]struct{}) [][2]int {
	var fencePoints [][2]int
	for p1 := range area {
		N := getAdjecent(p1)
		for _, n := range N {
			_, ok := area[n]
			if ok {
				continue
			}
			fencePoints = append(fencePoints, p1)
		}
	}

	return fencePoints
}

func countMergedFences(fences [][2]int) int {
	return 0
}

func printMap(m [][]rune) {
	builder := strings.Builder{}
	for _, row := range m {
		builder.WriteString(string(row) + "\n")
	}
	fmt.Println(builder.String())
}

func getEmptyMap(rows, cols int) [][]rune {
	m := make([][]rune, 0, rows)
	for i := 0; i < rows; i++ {
		row := make([]rune, 0, cols)
		for j := 0; j < cols; j++ {
			row = append(row, '.')
		}
		m = append(m, row)
	}
	return m
}

func part2(input []byte) int {
	m := parseInput(input)
	sortedAreas := getAreas(m)
	var total int
	for plant, areas := range sortedAreas {
		for _, area := range areas {
			fmt.Println("Plant: ", string(plant))
			fensePoints := getFencePoints(m, area)
			if plant == 'C' {
				//printPoints(fensePoints)
				print()
			}
			hor, ver, poi := groupPoints2(fensePoints)
			sum := (hor + ver + poi)
			fmt.Printf("Plant: %v found %d fence parts. Hor: %d, Vert: %d, Points: %d\n", string(plant), sum, hor, ver, poi)
			total += len(area) * sum
		}
	}

	return total
}

//func printPoints(p [][2]int) {
//	slices.SortFunc(p, func(a, b [2]int) int {
//		return a[0]-b[0]
//			return -1
//		}else{
//			return 1
//		}
//	})
//	for _, pi := range p {
//		fmt.Printf("(%d,%d)\n", pi[0], pi[1])
//	}
//}

//func part2(input []byte) int {
//	m := parseInput(input)
//	sortedAreas := getAreas(m)
//	var total int
//	for plant, areas := range sortedAreas {
//		for _, area := range areas {
//			fensePoints := getFencePoints(m, area)
//			fenses := calculateFences(fensePoints)
//			fmt.Printf("Plant: %v found %d fence groups\n", string(plant), fenses)
//			total += len(area) * fenses
//		}
//	}
//
//	return total
//}
/*
func calculateFences(fencePoints []*Point) int {
	var parts int
	// Sort vertical
	slices.SortFunc(fencePoints, func(a, b *Point) int {
		return a.y - b.y
	})
	// Sort horizontal
	slices.SortFunc(fencePoints, func(a, b *Point) int {
		return a.x - b.x
	})

	var part []int
	for _, p := range fencePoints {

	}

	return parts
}


*/
// FROM CHATGPT

// GroupPoints groups points into sets that are directly adjacent and form a straight line.
func GroupPoints(points [][2]int) [][][2]int {
	visited := make(map[[2]int]bool)
	var result [][][2]int

	for _, p := range points {
		if !visited[p] {
			// Start a new group from this point
			group := [][2]int{p}
			visited[p] = true
			expandGroup(points, &group, visited)
			if isLinear(group) {
				result = append(result, group)
			}
		}
	}

	return result
}

// expandGroup recursively adds adjacent points to the current group
func expandGroup(points [][2]int, group *[][2]int, visited map[[2]int]bool) {
	current := (*group)[len(*group)-1]
	adjacentPoints := findAdjacent(current, points)

	for _, adj := range adjacentPoints {
		if !visited[adj] {
			*group = append(*group, adj)
			visited[adj] = true
			expandGroup(points, group, visited)
		}
	}
}

// findAdjacent finds all points directly adjacent to the given point
func findAdjacent(p [2]int, points [][2]int) [][2]int {
	var adjacent [][2]int
	for _, q := range points {
		if (p[0] == q[0] && abs(p[1]-q[1]) == 1) || (p[1] == q[1] && abs(p[0]-q[0]) == 1) {
			adjacent = append(adjacent, q)
		}
	}
	return adjacent
}

// isLinear checks if all points in the group form a straight line
func isLinear(group [][2]int) bool {
	if len(group) <= 2 {
		return true
	}

	dx := group[1][0] - group[0][0]
	dy := group[1][1] - group[0][1]

	for i := 2; i < len(group); i++ {
		ndx := group[i][0] - group[i-1][0]
		ndy := group[i][1] - group[i-1][1]
		if ndx*dy != ndy*dx { // Cross product is zero for collinear points
			return false
		}
	}
	return true
}

func groupPoints(points [][2]int) [][][2]int {
	// Maps to store points by x and y values
	verticalMap := make(map[int][][2]int)
	horizontalMap := make(map[int][][2]int)

	// Group points by x for vertical lines and by y for horizontal lines
	for _, p := range points {
		x, y := p[0], p[1]
		verticalMap[x] = append(verticalMap[x], p)
		horizontalMap[y] = append(horizontalMap[y], p)
	}

	// Function to check if the points in a group form an uninterrupted line
	isUninterrupted := func(pts [][2]int, index int) bool {
		// Sort the points based on their coordinate
		// This will help us check if the line is uninterrupted
		for i := 0; i < len(pts)-1; i++ {
			if pts[i][index]+1 != pts[i+1][index] {
				return false
			}
		}
		return true
	}

	var result [][][2]int

	// Check vertical lines
	for _, pts := range verticalMap {
		// Sort points by y (for vertical lines, we check y values)
		// You can use your custom sorting method here, assuming points are unsorted
		sort.Slice(pts, func(i, j int) bool {
			return pts[i][1] < pts[j][1] // sort by y-value for vertical lines
		})
		if isUninterrupted(pts, 1) {
			result = append(result, pts)
		}
	}

	// Check horizontal lines
	for _, pts := range horizontalMap {
		// Sort points by x (for horizontal lines, we check x values)
		// You can use your custom sorting method here, assuming points are unsorted
		sort.Slice(pts, func(i, j int) bool {
			return pts[i][0] < pts[j][0] // sort by x-value for horizontal lines
		})
		if isUninterrupted(pts, 0) {
			result = append(result, pts)
		}
	}

	return result
}

var debug bool

func groupPoints2(points [][2]int) (int, int, int) {
	// Maps to store points by x and y values
	verticalMap := make(map[int][]int)
	horizontalMap := make(map[int][]int)

	// Group points by x for vertical lines and by y for horizontal lines
	for _, p := range points {
		y, x := p[0], p[1]
		verticalMap[x] = append(verticalMap[x], y)
		horizontalMap[y] = append(horizontalMap[y], x)
	}

	togo := make(map[[2]int]int)
	for _, p := range points {
		togo[p]++
	}

	var verticalLines int
	for x, line := range verticalMap {
		fmt.Println("detecting vertical ", x)
		slices.Sort(line)
		parts := split(line)
		for _, part := range parts {
			verticalLines++
			for _, y := range part {
				point := [2]int{y, x}
				if cnt, ok := togo[point]; ok && cnt >= 1 {
					togo[point]--
				}
			}
		}
	}
	debug = true

	var horizontalLines int
	for y, line := range horizontalMap {
		fmt.Println("detecting horizontal ", y)
		slices.Sort(line)
		parts := split(line)
		for _, part := range parts {
			horizontalLines++
			for _, x := range part {
				point := [2]int{y, x}
				if cnt, ok := togo[point]; ok && cnt >= 1 {
					togo[point]--
				}
			}
		}
	}

	var fensePoints int
	for _, cnt := range togo {
		fensePoints += cnt
	}

	//fmt.Printf("hor: %d, vert: %d, points: %d\n", horizontalLines, verticalLines, fensePoints)

	return horizontalLines, verticalLines, fensePoints
}

func split(s []int) [][]int {
	// Duplicate separation
	var other []int
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			other = append(other, s[i])
		}
	}
	sub := unique(s)
	slices.Sort(sub)

	var out [][]int
	out = append(out, getParts(sub)...)
	out = append(out, getParts(other)...)

	return out
}

func getParts(s []int) [][]int {
	if len(s) <= 1 {
		return nil
	}
	// Count parts
	var out [][]int
	part := []int{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] > 1 || s[i]-s[i-1] == 0 {
			if len(part) > 1 {
				out = append(out, part)
			}
			part = []int{s[i]}
		} else {
			part = append(part, s[i])
		}
	}
	if len(part) > 1 {
		out = append(out, part)
	}
	return out
}
