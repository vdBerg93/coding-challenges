package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	input    = "475449 2599064 213 0 2 65 5755 51149"
	example1 = "0 1 10 99 999"
	example2 = "125 17"
)

func main() {
	fmt.Println("Example1: ", part1(example1, 1))
	fmt.Println("Example2: ", part1(example2, 5))
	fmt.Println("Part1: ", part1(input, 25))
	fmt.Println("Part2: ", part1(input, 75))
}

func part1(input string, N int) int {
	stones := strings.Fields(input)

	// Use a map to track stones and their counts, reducing unnecessary duplication
	stoneCounts := make(map[string]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 1; i <= N; i++ {
		newStoneCounts := make(map[string]int, len(stoneCounts)*2)
		for stone, count := range stoneCounts {
			processed := do(stone)
			for _, newStone := range processed {
				newStoneCounts[newStone] += count
			}
		}
		stoneCounts = newStoneCounts
	}

	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}

	return totalStones
}

func do(stone string) []string {
	l := len(stone)
	if stone == "0" {
		num, _ := strconv.Atoi(stone)
		return []string{strconv.Itoa(num + 1)}
	} else if l%2 == 0 {
		return []string{
			rmTrailZero(stone[0 : l/2]),
			rmTrailZero(stone[l/2:]),
		}
	} else {
		num, _ := strconv.Atoi(stone)
		return []string{strconv.Itoa(num * 2024)}
	}
}

func rmTrailZero(s string) string {
	s2 := strings.TrimLeft(s, "0")
	if len(s2) == 0 {
		return "0"
	}
	return s2
}
