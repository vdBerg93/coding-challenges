package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var data []byte

func main() {
	set1, set2 := parseInput()
	part1(set1, set2)
	part2(set1, set2)
}

func part1(set1, set2 []int) {
	slices.Sort(set1)
	slices.Sort(set2)
	fmt.Println("Part 1 solution: ", calculatePairwiseDistanceSum(set1, set2))
}
func part2(set1, set2 []int) {
	fmt.Println("Part 2 solution: ", calculateSimilarity(set1, set2))
}

func parseInput() ([]int, []int) {
	lines := strings.Split(string(data), "\n")
	one := make([]int, 0, len(lines))
	two := make([]int, 0, len(lines))
	for _, line := range lines {
		v1, v2 := parseLine(line)
		one = append(one, v1)
		two = append(two, v2)
	}
	return one, two
}

func parseLine(line string) (int, int) {
	f := strings.Fields(line)
	v1, _ := strconv.Atoi(f[0])
	v2, _ := strconv.Atoi(f[1])
	return v1, v2
}

func calculatePairwiseDistanceSum(v1, v2 []int) int {
	var d int
	for i := 0; i < len(v1); i++ {
		d += distance(v1[i], v2[i])
	}
	return d
}

func distance(one, two int) int {
	di := one - two
	if di < 0 {
		return -di
	}
	return di
}

func calculateSimilarity(set1, set2 []int) int {
	count := countNumbers(set2)

	var similarity int
	for _, val := range set1 {
		if c, ok := count[val]; ok {
			similarity += c * val
		}
	}
	return similarity
}

func countNumbers(v []int) map[int]int {
	f := make(map[int]int)
	for _, val := range v {
		f[val]++
	}
	return f
}
