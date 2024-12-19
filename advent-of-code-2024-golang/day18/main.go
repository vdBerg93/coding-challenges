package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"log"
	"maps"
	"math"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

func main() {
	fmt.Println("Part1 example:", part1(example))
	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2 example:", part2(example))
	fmt.Println("Part2:", part2(input))
}

const visualize = false

func part1(input []byte) int {
	m := parseMap(input)
	start := m.Find('S')
	end := m.Find('E')
	cost, _ := Dijkstra(m, start, end)
	return cost
}

func part2(input []byte) int {
	m := parseMap(input)
	start := m.Find('S')
	end := m.Find('E')
	_, path := Dijkstra(m, start, end)
	return len(path)
}

const linesep = "\r\n"

type Map struct {
	m [][]rune
}

func (m *Map) Find(r rune) Point {
	for y, row := range m.m {
		for x, cell := range row {
			if cell == r {
				return Point{x, y}
			}
		}
	}
	log.Panicf("solution not found for %v", string(r))
	return Point{}
}

func (m *Map) Get(p Point) rune {
	return m.m[p.y][p.x]
}

func (m *Map) occupied(p Point) bool {
	if p.x < 0 || p.x >= len(m.m[0]) || p.y < 0 || p.y >= len(m.m) {
		return true
	}
	if m.Get(p) == '#' {
		return true
	}
	return false
}

func parseMap(input []byte) Map {
	rows := strings.Split(string(input), linesep)
	m := make([][]rune, 0, len(rows))
	for _, r := range rows {
		m = append(m, []rune(r))
	}
	return Map{
		m: m,
	}
}

type VisitedKey struct {
	p Point
	d Point
}

// Dijkstra finds the cheapest route from top left to bottom right
func Dijkstra(track Map, start, end Point) (int, map[Point]struct{}) {

	visited := make(map[VisitedKey][]VisitedKey)
	visitedCost := make(map[VisitedKey]int)
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)

	startState := State{&State{}, start, right, 0}
	heap.Push(&priorityQueue, startState)

	best := math.MaxInt
	var solutions []VisitedKey

	for priorityQueue.Len() > 0 {
		current := heap.Pop(&priorityQueue).(State)
		if current.cost > best {
			paths := backTrack(visited, startState.Key(), solutions, 0)
			return best, paths
		}

		if current.pos == end && current.cost <= best {
			best = min(best, current.cost)
			solutions = append(solutions, current.Key())
		}

		cost, costOk := visitedCost[current.Key()]
		if costOk && current.cost > cost {
			continue
		}

		parents, ok := visited[current.Key()]
		if current != startState {
			parents = append(parents, current.parent.Key())
		}
		visited[current.Key()] = parents
		visitedCost[current.Key()] = current.cost

		if ok {
			continue
		}
		if visualize {
			fmt.Printf("Exploring %s: %+v\n", getHeading(current.dir), current)
		}

		expandFrontier(track, &priorityQueue, current, best)
	}

	return -1, nil // No valid path found
}

func backTrack(visited map[VisitedKey][]VisitedKey, start VisitedKey, solution []VisitedKey, i int) map[Point]struct{} {
	path := make(map[Point]struct{})

	for _, s := range solution {
		path[s.p] = struct{}{}
		if s.p == start.p {
			break
		}
		locs := backTrack(visited, start, visited[s], i+1)
		maps.Copy(path, locs)
	}

	return path
}

func expandFrontier(track Map, queue *PriorityQueue, current State, best int) {
	for _, i := range []int{-1, 0, 1} {
		next := doStep(current, i)
		if next.cost > best {
			continue
		}
		if visualize {
			fmt.Printf("Expanding with %+v\n", next)
		}
		if track.occupied(next.pos) {
			if visualize {
				fmt.Println("occupied")
			}
			continue
		}
		heap.Push(queue, next)
	}
}

const (
	costStraight = 1
	costRotation = 1000
)

func (s *State) forward() State {
	return State{
		pos: Point{
			x: s.pos.x + s.dir.x,
			y: s.pos.y + s.dir.y,
		},
		dir:    s.dir,
		cost:   s.cost + costStraight,
		parent: s,
	}
}

func (s *State) rotate(dir int) State {
	return State{
		pos: s.pos,
		dir: Point{
			x: -s.dir.y * dir,
			y: s.dir.x * dir,
		},
		cost:   s.cost + costRotation,
		parent: s,
	}
}

// doStep: dir=-1 -> CCW, dir=1 -> CW, 0=straight
func doStep(pose State, dir int) State {
	if dir == 0 {
		return pose.forward()
	}
	return pose.rotate(dir)
}

type State struct {
	parent *State
	pos    Point
	dir    Point
	cost   int
}

func (s *State) Key() VisitedKey {
	return VisitedKey{
		p: s.pos,
		d: s.dir,
	}
}

var (
	right = Point{1, 0}
	down  = Point{0, 1}
	left  = Point{-1, 0}
	up    = Point{0, -1}
)

func getHeading(dir Point) string {
	switch dir {
	case right:
		return "right"
	case left:
		return "left"
	case up:
		return "up"
	case down:
		return "down"
	default:
		return "unknown"
	}
}

type Point struct {
	x, y int
}

// PriorityQueue is a min-heap implementation for A* algorithm
type PriorityQueue []State

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].cost < (*pq)[j].cost
}
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
