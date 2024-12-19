package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("Solution:", part1(input))
}

var designs []string
var patterns []string
var memo = map[string]int{}

func part1(input []byte) string {
	patterns, designs = parseInput(input)
	possible, cnt := countFeasible(designs)
	return fmt.Sprintf("%d,%d", possible, cnt)
}

const lineSep = "\r\n"

func parseInput(data []byte) ([]string, []string) {
	b := strings.Split(string(data), lineSep+lineSep)

	var pat []string
	for _, row := range strings.Split(b[0], lineSep) {
		pat = append(pat, strings.Split(row, ", ")...)
	}

	des := strings.Split(b[1], lineSep)

	return pat, des
}

func countFeasible(designs []string) (int, int) {
	var (
		boolCount int
		intCount  int
	)

	for _, d := range designs {
		result := count(d)
		if result > 0 {
			boolCount++
		}
		intCount += result
	}
	return boolCount, intCount
}

func count(d string) int {
	if d == "" {
		return 1
	}
	if val, ok := memo[d]; ok {
		return val
	}

	var sum int
	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			sum += count(strings.TrimPrefix(d, p))
		}
	}

	memo[d] = sum

	return sum
}
