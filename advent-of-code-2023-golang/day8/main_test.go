package main

import (
	"bufio"
	"bytes"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err := os.ReadFile("sample")
	if err != nil {
		panic(err)
	}
	//Test data
	got := Part1(testData, "AAA", "ZZZ")
	want1 := int64(2)
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Part1(data, "AAA", "ZZZ"))

	testData2, err := os.ReadFile("sample2")
	if err != nil {
		panic(err)
	}

	got2 := Part1(testData2, "A", "Z")
	want2 := int64(6)
	if got2 != want2 {
		log.Fatalf("expected %d, got %d", want2, got2)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	got2 = Part1(data, "A", "Z")
	fmt.Printf("Solution part 2: %d\n", got2)

}

func Part1(data []byte, startMode, finishMode string) int64 {
	m := newMap(data)
	m.setStartFinish(startMode, finishMode)
	ret := 1
	for _, start := range m.starts {
		steps := m.Solve(start, finishMode)
		ret = lcm(ret, int(steps))
	}
	return int64(ret)
}

// Function to calculate the greatest common divisor (GCD) using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM) using the GCD
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

type Map struct {
	starts   []string
	finishes map[string]bool
	m        map[string][2]string
	lr       *ring.Ring
}

func newMap(data []byte) *Map {
	m := Map{
		starts:   make([]string, 0),
		finishes: make(map[string]bool),
		m:        make(map[string][2]string),
	}
	s := bufio.NewScanner(bytes.NewReader(data))
	s.Split(bufio.ScanLines)

	s.Scan()
	m.initializeRing(s.Text())

	var loc, left, right string
	for s.Scan() {
		text := s.Text()
		if text == "" {
			continue
		}

		loc = text[0:3]
		left = text[7:10]
		right = text[12:15]
		m.m[loc] = [2]string{left, right}
	}

	return &m
}

// initializeRing for left-right
func (m *Map) initializeRing(text string) {
	lr := []rune(text)
	m.lr = ring.New(len(lr))
	for _, r := range lr {
		m.lr.Value = r
		m.lr = m.lr.Next()
	}
}

func (m *Map) setStartFinish(startMode, finishMode string) {
	for loc := range m.m {
		if strings.HasSuffix(loc, startMode) {
			m.starts = append(m.starts, loc)
		}
		if strings.HasSuffix(loc, finishMode) {
			m.finishes[loc] = true
		}
	}
	fmt.Printf("Starts: %v,\t Finishes: %v\n", m.starts, m.finishes)
}

func (m *Map) NextLeftRight() rune {
	lr := m.lr.Value
	m.lr = m.lr.Next()
	return lr.(rune)
}

func (m *Map) NextLocation(location string, lr rune) string {
	step := m.m[location]
	if lr == 'L' {
		return step[0]
	} else {
		return step[1]
	}
}

func (m *Map) Solve(pos, finishMode string) int64 {
	steps := int64(0)
	for {
		steps++
		lr := m.NextLeftRight()
		pos = m.NextLocation(pos, lr)
		if strings.HasSuffix(pos, finishMode) {
			return steps
		}
	}
}
