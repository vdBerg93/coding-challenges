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
	fmt.Println(strings.Join(stones, " "))

	for i := 1; i <= N; i++ {
		newStones := make([]string, 0, len(stones)*2)
		for _, stone := range stones {
			newStones = append(newStones, do(stone)...)
		}
		stones = newStones
		fmt.Printf("i:%d, %s\n", i, strings.Join(stones, " "))
	}

	return len(stones)
}

func do(stone string) []string {
	l := len(stone)
	if stone == "0" {
		num, _ := strconv.Atoi(stone)
		return []string{strconv.Itoa(num + 1)}
	} else if l%2 == 0 {
		return []string{
			rmTrailZero(stone[0 : l/2]),
			rmTrailZero(stone[l/2 : l]),
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
