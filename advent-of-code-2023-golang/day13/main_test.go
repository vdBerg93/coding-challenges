package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
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
	got := SolvePart1(testData)
	want1 := 405
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", SolvePart1(data))
}

func Test_Part2(t *testing.T) {
	got := Solve2(testData)
	want1 := 400
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve2(data))

}

func SolvePart1(data []byte) int {
	blocks := bytes.Split(data, []byte("\r\n\r\n"))

	solution := 0
	for _, dataBlock := range blocks {
		block := readBlock(dataBlock)
		solution += getScore(findReflections2D(block))
	}

	return solution
}

func findReflections2D(block []string) ([]int, []int) {
	return findReflections(block), findReflections(transpose(block))
}

func findReflections(block []string) []int {
	// Try to find a reflection for each row
	solutions := make(map[int]int)
	for _, row := range block {
		middles := getReflectionLine(row)
		for _, middle := range middles {
			solutions[middle]++
		}
	}

	// Reflection is valid if a reflection index exists for all rows
	var reflections []int
	for col, cnt := range solutions {
		if cnt == len(block) {
			reflections = append(reflections, col)
		}
	}

	return reflections
}

func getScore(hor, vert []int) int {
	score := 0
	for _, h := range hor {
		score += h
	}
	for _, v := range vert {
		score += 100 * v
	}
	return score
}

func Solve2(data []byte) int {
	blocks := bytes.Split(data, []byte("\r\n\r\n"))

	solution := 0
	for _, dataBlock := range blocks {
		block := readBlock(dataBlock)

		// Find initial reflections
		horizontal, vertical := findReflections2D(block)
		// Use them to determine which new reflection mode we can create
		horS, vertS := findReflection2DAlternatives(block, horizontal, vertical)
		score := getScore(horS, vertS)
		solution += score
	}

	return solution
}

func newReflection(old []int, new int) bool {
	for _, h := range old {
		if new == h {
			return false
		}
	}
	return true
}

func findReflection2DAlternatives(block []string, oldHor, oldVert []int) ([]int, []int) {
	smudges := len(block) * len(block[0])

	// Brute force all smudge combinations
	for i := 0; i < smudges; i++ {
		var newHor, newVert []int

		smudged := nextSmudge(block, i)

		hor := findReflections(smudged)
		for _, h := range hor {
			if newReflection(oldHor, h) {
				newHor = append(newHor, h)
			}
		}

		trBlock := transpose(smudged)
		ver := findReflections(trBlock)

		for _, v := range ver {
			if newReflection(oldVert, v) {
				newVert = append(newVert, v)
			}
		}

		if len(newHor)+len(newVert) > 0 {
			return newHor, newVert
		}
	}
	return nil, nil
}

func nextSmudge(data []string, last int) []string {
	width := len(data[0])
	idxRow := last / width
	idxCol := last - idxRow*width
	if idxCol < 0 {
		fmt.Print()
	}
	row := []rune(data[idxRow])
	row[idxCol] = smudge(row[idxCol])
	out := slices.Clone(data)
	out[idxRow] = string(row)
	return out
}

func readBlock(data []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)

	var block []string
	for scanner.Scan() {
		block = append(block, scanner.Text())
	}
	return block
}

func transpose(input []string) []string {
	output := make([]string, len(input[0]))

	for _, row := range input {
		for c, char := range row {
			output[c] += string(char)
		}
	}
	return output
}

func getReflectionLine(row string) []int {
	var middles []int

	for i := 1; i < len(row); i++ {
		var left, right string
		if 2*i <= len(row) {
			left = reverseString(row[0:i])
			right = row[i : i+len(left)]
		} else {
			right = row[i:]
			left = reverseString(row[i-len(right) : i])
		}

		if left == right {
			middles = append(middles, i)
		}
	}

	return middles
}

func Test_GetMiddles(t *testing.T) {
	type test struct {
		row  string
		want int
	}
	tests := []test{
		{"..##...", 5},
		{"#....#.", 3},
	}

	for _, tt := range tests {
		t.Run(tt.row, func(t *testing.T) {
			got := getReflectionLine(tt.row)

			for _, g := range got {
				if g == tt.want {
					return
				}
			}
			t.Errorf("want %v, got %v", tt.want, got)
		})
	}
}

func reverseString(in string) string {
	var out string
	for i := len(in) - 1; i >= 0; i-- {
		out += string(in[i])
	}
	return out
}

func smudge(char rune) rune {
	switch char {
	case '#':
		return '.'
	case '.':
		return '#'
	default:
		panic("not implemented rune")
	}
}
