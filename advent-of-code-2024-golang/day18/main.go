package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"maps"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

func main() {
	fmt.Println("Part1 example:", part1(example, 7, 7, 12))
	fmt.Println("Part1:", part1(input, 71, 71, 1024))
	fmt.Println("Part2 example:", part2(example, 7, 7))
	fmt.Println("Part1:", part2(input, 71, 71))
}

const visualize = false

func part1(input []byte, szX, szY, bytes int) int {
	obstacles := parseInput(input)
	D := newDijkstra(szX, szY, obstacles[0:bytes])
	start := Point{0, 0}
	end := Point{szX - 1, szY - 1}
	cost, _ := D.Run(start, end)
	return cost
}

func part2(input []byte, szX, szY int) Point {
	obstacles := parseInput(input)
	start := Point{0, 0}
	end := Point{szX - 1, szY - 1}
	for bytes := len(obstacles); bytes >= 0; bytes-- {
		D := newDijkstra(szX, szY, obstacles[0:bytes])
		cost, _ := D.Run(start, end)
		if cost != -1 {
			return obstacles[bytes]
		}
	}
	return Point{}
}

const linesep = "\r\n"

type Dijkstra struct {
	visited     map[VisitedKey][]VisitedKey
	visitedCost map[VisitedKey]int
	queue       PriorityQueue
	obs         []Point
	sizeX       int
	sizeY       int
}

func newDijkstra(szX, szY int, obstacles []Point) Dijkstra {
	m := Dijkstra{
		visited:     make(map[VisitedKey][]VisitedKey),
		visitedCost: make(map[VisitedKey]int),
		queue:       make(PriorityQueue, 0),
		obs:         obstacles,
		sizeX:       szX,
		sizeY:       szY,
	}
	heap.Init(&m.queue)
	return m
}

func parseInput(input []byte) []Point {
	rows := strings.Split(string(input), linesep)
	obs := make([]Point, 0, len(rows))
	for _, r := range rows {
		f := strings.Split(r, ",")
		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		obs = append(obs, Point{x: x, y: y})
	}
	return obs
}

type VisitedKey struct {
	pos Point
}

// Run finds the cheapest route from top left to bottom right
func (d *Dijkstra) Run(start, end Point) (int, map[Point]struct{}) {
	startState := State{&State{}, start, 0}
	heap.Push(&d.queue, startState)

	best := math.MaxInt
	var solutions []VisitedKey

	for d.queue.Len() > 0 {
		current := heap.Pop(&d.queue).(State)
		if current.cost > best {
			break
		}
		d.Print(current)

		if current.pos == end && current.cost <= best {
			best = min(best, current.cost)
			solutions = append(solutions, current.Key())
		}

		costPrevious, costOk := d.visitedCost[current.Key()]
		if costOk && current.cost > costPrevious {
			if visualize {
				fmt.Println("same pos higher cost")
			}
			continue
		}

		parents, ok := d.visited[current.Key()]
		if current != startState {
			parents = append(parents, current.parent.Key())
		}
		d.visited[current.Key()] = parents
		d.visitedCost[current.Key()] = current.cost

		if ok {
			continue
		}
		if visualize {
			fmt.Printf("Exploring %+v\n", current)
		}

		d.expandFrontier(current, best)
	}

	if best != math.MaxInt {
		return best, d.backTrack(startState.Key(), solutions, 0)
	}

	return -1, nil // No valid path found
}

func (d *Dijkstra) expandFrontier(current State, best int) {
	for _, dir := range []Point{left, right, up, down} {
		next := current.move(dir)
		if next.cost > best {
			continue
		}
		if visualize {
			fmt.Printf("Expanding with %+v\n", next)
		}
		if d.occupied(next) {
			if visualize {
				fmt.Println("occupied")
			}
			continue
		}
		heap.Push(&d.queue, next)
	}
}

func (s *State) move(dir Point) State {
	return State{
		pos: Point{
			x: s.pos.x + dir.x,
			y: s.pos.y + dir.y,
		},
		cost:   s.cost + 1,
		parent: s,
	}
}

func (d *Dijkstra) getObstacles(t int) []Point {
	//last := min(len(d.obs), t)
	//return d.obs[:last]
	return d.obs
}

func (d *Dijkstra) occupied(p State) bool {
	if d.outsideGrid(p.pos) {
		return true
	}
	for _, o := range d.getObstacles(p.cost) {
		if p.pos.x == o.x && p.pos.y == o.y {
			return true
		}
	}
	return false
}

func (d *Dijkstra) outsideGrid(p Point) bool {
	return p.x < 0 || p.x >= d.sizeX || p.y < 0 || p.y >= d.sizeY
}

func (d *Dijkstra) backTrack(start VisitedKey, solution []VisitedKey, i int) map[Point]struct{} {
	path := make(map[Point]struct{})

	for _, s := range solution {
		path[s.pos] = struct{}{}
		if s.pos == start.pos {
			break
		}
		locs := d.backTrack(start, d.visited[s], i+1)
		maps.Copy(path, locs)
	}

	return path
}

func (d *Dijkstra) Print(state State) {
	if !visualize {
		return
	}
	fmt.Printf("Printing %+v\n", state)
	m := make([][]rune, 0, d.sizeY)
	for y := 0; y < d.sizeY; y++ {
		row := make([]rune, 0, d.sizeX)
		for x := 0; x < d.sizeX; x++ {
			row = append(row, '.')
		}
		m = append(m, row)
	}

	m[state.pos.y][state.pos.x] = '@'

	obs := d.getObstacles(state.cost)
	for _, o := range obs {
		m[o.y][o.x] = '#'
	}

	b := strings.Builder{}
	for _, row := range m {
		b.WriteString(string(row) + "\n")
	}
	fmt.Println(b.String())
	time.Sleep(100 * time.Millisecond)
}

type State struct {
	parent *State
	pos    Point
	cost   int
}

func (s *State) Key() VisitedKey {
	return VisitedKey{
		pos: s.pos,
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
