package main

import (
	"fmt"
	"log"
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
	got := Solve(testData)
	want1 := 1320
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data))
}

func Test_Part2(t *testing.T) {
	got := SolvePart2(testData)
	want1 := 145
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", SolvePart2(data))
}

func Solve(data []byte) int {
	parts := strings.Split(string(data), ",")
	solution := 0
	for _, word := range parts {
		solution += int(hash(word))
	}
	return solution
}

func hash(word string) uint8 {
	h := 0
	for _, r := range word {
		h += int(r) // ascii
		h *= 17
		h = h % 256
	}
	return uint8(h)
}

type Lens struct {
	label string
	fl    int
}

type BoxSeries struct {
	boxes [256][]Lens
}

func (B *BoxSeries) Remove(box int, label string) {
	for i, lens := range B.boxes[box] {
		if lens.label == label {
			B.boxes[box] = append(B.boxes[box][:i], B.boxes[box][i+1:]...)
			return
		}
	}
}

func (B *BoxSeries) Add(box int, label string, mag int) {
	if label == "ot" {
		fmt.Print()
	}
	lens := Lens{label: label, fl: mag}
	for i, li := range B.boxes[box] {
		if li.label == label {
			B.boxes[box][i] = lens
			return
		}
	}
	B.boxes[box] = append(B.boxes[box], lens)
}

func SolvePart2(data []byte) int {
	parts := strings.Split(string(data), ",")

	boxes := BoxSeries{}
	for _, word := range parts {

		box, label, op, mag := parseStep(word)
		switch op {
		case '=':
			boxes.Add(box, label, mag)
		case '-':
			boxes.Remove(box, label)
		default:
			log.Panicf("unkown operation %v", op)
		}
		fmt.Print()
	}
	return calculateFocusPower(boxes)
}

func parseStep(word string) (int, string, rune, int) {
	var operation rune
	if strings.ContainsRune(word, '=') {
		operation = '='
	} else {
		operation = '-'
	}

	parts := strings.Split(word, string(operation))
	label := parts[0]
	box := int(hash(parts[0]))

	if operation == '-' {
		return box, label, operation, 0
	} else {
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		return box, label, operation, val
	}
}

func calculateFocusPower(boxes BoxSeries) int {
	power := 0
	for i, box := range boxes.boxes {
		for j, lens := range box {
			p := (i + 1) * (j + 1) * lens.fl
			power += p
		}
	}
	return power
}
