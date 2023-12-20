package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	got := Solve(testData, false)
	want1 := 62
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, false))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, true)
	want1 := 952408144115
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve(data, true))
}

type Point [2]int

type Node struct {
	Pos Point
}

func Solve(data []byte, doHex bool) int {
	var nodes []Node
	scanner := bufio.NewScanner(bytes.NewReader(data))
	boundaryCount := 0
	for scanner.Scan() {
		dir, steps, hexString := parseRow(scanner.Text())
		var previous Node
		if len(nodes) > 0 {
			previous = nodes[len(nodes)-1]
		}
		if doHex {
			dir, steps = parseHex(hexString)
		}
		newNode := Node{
			Pos: Point{
				previous.Pos[0] + steps*dir[0],
				previous.Pos[1] + steps*dir[1],
			},
		}
		nodes = append(nodes, newNode)
		boundaryCount += steps
	}

	surface := shoelace(nodes)
	interior := surface - boundaryCount/2 + 1 // Pick's theorem
	return interior + boundaryCount
}

func parseRow(text string) ([2]int, int, string) {
	fields := strings.Fields(text)
	steps, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	dir := getDirFromChar(fields[0])
	return dir, steps, strings.Trim(fields[2], "()")
}

func getDirFromChar(char string) [2]int {
	switch char {
	case "R":
		return right
	case "L":
		return left
	case "U":
		return up
	case "D":
		return down
	default:
		panic("not implemented")
	}
}

func getDirFromNum(num uint64) [2]int {
	switch num {
	case 0:
		return right
	case 1:
		return down
	case 2:
		return left
	case 3:
		return up
	default:
		panic("not implemented")
	}
}

func parseHex(data string) ([2]int, int) {
	numberStr := strings.Replace(data, "#", "", -1)
	n, err := strconv.ParseUint(numberStr[0:5], 16, 64)
	if err != nil {
		panic(err)
	}
	d, err := strconv.ParseUint(string(numberStr[len(numberStr)-1]), 16, 64)
	if err != nil {
		panic(err)
	}

	return getDirFromNum(d), int(n)
}

var (
	right = Point{0, 1}
	down  = Point{1, 0}
	left  = Point{0, -1}
	up    = Point{-1, 0}
)

func shoelace(nodes []Node) int {
	n := len(nodes)
	if n < 3 {
		// At least 3 vertices are required to form a polygon
		return 0.0
	}

	// Initialize the sum variables
	sumX, sumY := 0, 0

	// Compute the sum of products
	for i := 0; i < n-1; i++ {
		sumX += nodes[i].Pos[0] * nodes[i+1].Pos[1]
		sumY += nodes[i].Pos[1] * nodes[i+1].Pos[0]
	}

	// Add the products of the last and first vertices
	sumX += nodes[n-1].Pos[0] * nodes[0].Pos[1]
	sumY += nodes[n-1].Pos[1] * nodes[0].Pos[0]

	// Compute the area using the Shoelace formula
	area := 0.5 * float64(sumX-sumY)
	if area < 0 {
		// Take the absolute value since area should be non-negative
		area = -area
	}

	return int(area)
}
