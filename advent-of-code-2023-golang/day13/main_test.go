package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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
	want1 := 405
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, false))
}

func Solve(data []byte, expand bool) int {
	blocks := bytes.Split(data, []byte("\r\n\r\n"))

	solution := 0
	for _, dataBlock := range blocks {
		block := readBlock(dataBlock)
		hor := findMirror(block)
		if hor != nil {
			solution += *hor
			continue
		}

		trBlock := transpose(block)
		ver := findMirror(trBlock)
		if ver == nil {
			panic("no solution")
		}
		solution += 100 * (*ver)
	}

	return solution
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

func findMirror(block []string) *int {

	solutions := map[int]int{}
	for _, row := range block {
		middles := getMiddles(row)
		for _, middle := range middles {
			solutions[middle]++
		}
	}

	for col, cnt := range solutions {
		if cnt == len(block) {
			return &col
		}
	}

	return nil
}

func getMiddles(row string) []int {
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
			got := getMiddles(tt.row)

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
