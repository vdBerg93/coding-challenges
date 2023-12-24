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

var dataInput []byte
var dataSample1 []byte

func TestMain(m *testing.M) {
	var err error
	dataInput, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	dataSample1, err = os.ReadFile("sample1")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

const debug = false

func Test_Part1(t *testing.T) {
	got := Solve(dataSample1, 7, 27, false)
	want1 := 2
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}

	fmt.Printf("Part 1 test succeeded.\n")
	got = Solve(dataInput, 2e14, 4e14, false)
	fmt.Printf("Solution part 1: %d\n", got)
}

type Hailstone struct {
	x, y, z    float64
	vx, vy, vz float64
}

func Solve(data []byte, minArea, maxArea int, addZ bool) int {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var stones []Hailstone
	for scanner.Scan() {
		stones = append(stones, NewStone(scanner.Text()))
	}
	return TraceTrajectories(stones, minArea, maxArea)
}

func NewStone(row string) Hailstone {
	row = strings.Replace(row, " ", "", -1)
	parts := strings.Split(row, "@")
	positions := strings.Split(parts[0], ",")
	velocities := strings.Split(parts[1], ",")

	return Hailstone{
		x:  getFloat(positions[0]),
		y:  getFloat(positions[1]),
		z:  getFloat(positions[2]),
		vx: getFloat(velocities[0]),
		vy: getFloat(velocities[1]),
		vz: getFloat(velocities[2]),
	}
}

func getFloat(str string) float64 {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return float64(val)
}

func TraceTrajectories(stones []Hailstone, xmin, xmax int) int {
	intersectCount := 0
	seen := make(map[[2]int]struct{})
	for i1, stone1 := range stones {
		for i2, stone2 := range stones {
			if i1 == i2 {
				continue
			}
			key := [2]int{min(i1, i2), max(i1, i2)}
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			if Intersect(stone1, stone2, float64(xmin), float64(xmax)) {
				intersectCount++
			}
		}
	}
	return intersectCount
}

func Intersect(s1, s2 Hailstone, xmin, xmax float64) bool {
	a1, b1 := getLineParam(s1)
	a2, b2 := getLineParam(s2)
	xIts := (b2 - b1) / (a1 - a2)
	yIts := a1*xIts + b1
	intersect := xIts >= xmin && xIts <= xmax && yIts >= xmin && yIts <= xmax
	t1x := (xIts - s1.x) / s1.vx
	t2x := (xIts - s2.x) / s2.vx
	if t1x < 0 || t2x < 0 {
		return false
	}
	if intersect {
		fmt.Print()
	}
	return intersect
}

func getLineParam(s Hailstone) (float64, float64) {
	a := s.vy / s.vx
	b := s.y - s.x*(s.vy/s.vx)
	return a, b
}
