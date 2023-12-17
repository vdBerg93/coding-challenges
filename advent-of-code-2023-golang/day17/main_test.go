package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"testing"
)

var data []byte
var testData []byte

func TestMain(m *testing.M) {
	var err error
	data, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func Test_Part1(t *testing.T) {
	got := Solve(testData, 1, 3)
	want1 := 102
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, 1, 3))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, 4, 10)
	want1 := 71
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, 4, 10))
}

func readHeatMap(data []byte) [][]int {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var heatmap [][]int
	for scanner.Scan() {
		var row []int
		for _, rune := range scanner.Text() {
			heat, err := strconv.Atoi(string(rune))
			if err != nil {
				panic(err)
			}
			row = append(row, heat)
		}
		heatmap = append(heatmap, row)
	}
	return heatmap
}

func Solve(data []byte, minConSec, maxConSex int) int {

	costMap := readHeatMap(data)

	result := Dijkstra(costMap, minConSec, maxConSex)
	fmt.Printf("Cheapest route cost: %d\n", result)
	return result
}

var (
	right = Point{0, 1}
	down  = Point{1, 0}
	left  = Point{0, -1}
	up    = Point{-1, 0}
)

type VisitedKey struct {
	p Point
	d Point
	c int
}

// dijkstra finds the cheapest route from top left to bottom right
func Dijkstra(costMap [][]int, minConSec, maxConSec int) int {
	rows, cols := len(costMap), len(costMap[0])
	start := Point{0, 0}
	end := Point{rows - 1, cols - 1}

	visited := make(map[VisitedKey]bool)
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)

	heap.Push(&priorityQueue, State{start, 0, right, 1})
	heap.Push(&priorityQueue, State{start, 0, down, 1})

	directions := []Point{right, down, left, up}

	for priorityQueue.Len() > 0 {
		current := heap.Pop(&priorityQueue).(State)

		key := VisitedKey{current.pos, current.dir, current.dirCount}
		if visited[key] {
			continue
		}
		visited[key] = true

		nextPos := Point{current.pos[0] + current.dir[0], current.pos[1] + current.dir[1]}

		if !isValid(nextPos, rows, cols) {
			continue
		}

		nextCost := current.cost + costMap[nextPos[0]][nextPos[1]]

		if minConSec <= current.dirCount && current.dirCount <= maxConSec {
			if nextPos == end {
				return nextCost
			}
		}

		for _, dir := range directions {
			if current.dir[0]+dir[0] == 0 && current.dir[1]+dir[1] == 0 {
				continue // No reverse
			}

			var nextDirCount int
			if dir == current.dir {
				nextDirCount = current.dirCount + 1
			} else {
				nextDirCount = 1
			}

			if (dir != current.dir && current.dirCount < minConSec) || nextDirCount > maxConSec {
				continue
			}

			next := State{
				pos:      nextPos,
				cost:     nextCost,
				dir:      dir,
				dirCount: nextDirCount,
			}

			heap.Push(&priorityQueue, next)

		}
	}

	return -1 // No valid path found
}

type State struct {
	pos      Point
	cost     int
	dir      Point
	dirCount int
}
type Point [2]int

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

// isValid checks if the cell is within the bounds of the map
func isValid(cell Point, rows, cols int) bool {
	return cell[0] >= 0 && cell[0] < rows && cell[1] >= 0 && cell[1] < cols
}
