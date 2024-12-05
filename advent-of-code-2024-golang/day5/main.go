package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readInput(input)
	part1, part2 := solve(rules, updates)
	fmt.Println("Answer part1:", part1)
	fmt.Println("Answer part2:", part2)
}

//go:embed input.txt
var input []byte

func readInput(data []byte) ([][2]int, [][]int) {
	blocks := strings.Split(string(data), "\r\n\r\n")
	if len(blocks) != 2 {
		panic("invalid blocks")
	}

	return parseRules(blocks[0]), parseUpdates(blocks[1])
}

func parseRules(s string) [][2]int {
	lines := splitLines(s)
	rules := make([][2]int, 0, len(lines))
	for _, row := range lines {
		fields := strings.Split(row, "|")
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])
		rules = append(rules, [2]int{a, b})
	}
	return rules
}

func parseUpdates(s string) [][]int {
	lines := splitLines(s)
	pages := make([][]int, 0, len(lines))
	for _, row := range lines {
		var update []int
		fields := strings.Split(row, ",")
		for _, field := range fields {
			value, _ := strconv.Atoi(field)
			update = append(update, value)
		}
		pages = append(pages, update)
	}
	return pages
}

func splitLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func solve(rules [][2]int, updates [][]int) (int, int) {
	var part1, part2 int

	for _, update := range updates {
		updateOK := true
		for _, rule := range rules {
			if !orderOK(rule, update) {
				update = orderUpdate(rules, update)
				updateOK = false
				break
			}
		}
		middle := update[len(update)/2]
		if updateOK {
			part1 += middle
		} else {
			part2 += middle
		}
	}

	return part1, part2
}

func orderOK(rule [2]int, update []int) bool {
	before := lookupIndex(rule[0], update)
	after := lookupIndex(rule[1], update)
	if before == -1 || after == -1 {
		return true
	}
	return before < after
}

func lookupIndex(p int, update []int) int {
	for idx, val := range update {
		if val == p {
			return idx
		}
	}
	return -1
}

func orderUpdate(rules [][2]int, update []int) []int {
	slices.SortFunc(update, func(a, b int) int {
		for _, rule := range rules {
			if rule[0] == a && rule[1] == b {
				return -1
			}
			if rule[0] == b && rule[1] == a {
				return 1
			}
		}
		return 1
	})
	return update
}
