package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	data := getInput(input)
	fmt.Println("Part1: ", part1(data))
	fmt.Println("Part2: ", part2(data))
}

func part1(data string) int {
	statements := getStatements(data)
	var sum int
	for _, statement := range statements {
		sum += executeStatement(statement)
	}
	return sum
}

func part2(data string) int {
	data = strings.Replace(data, "\n", "", -1)
	filtered := filterConditions(data)
	statements := getStatements(filtered)

	var sum int
	for _, statement := range statements {
		sum += executeStatement(statement)
	}
	return sum
}

func getInput(data []byte) string {
	return string(data)
}

func getStatements(s string) []string {
	exp, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}
	return exp.FindAllString(s, -1)
}

func executeStatement(s string) int {
	s = strings.TrimPrefix(s, "mul(")
	s = strings.TrimSuffix(s, ")")
	f := strings.Split(s, ",")
	a, _ := strconv.Atoi(f[0])
	b, _ := strconv.Atoi(f[1])

	return a * b
}

func filterConditions(s string) string {
	s = strings.ReplaceAll(s, "do()", "\ndo()")
	s = strings.ReplaceAll(s, "don't()", "\ndon't()")
	exp := regexp.MustCompile(`(?m)^don't\(\).*$`)
	s = exp.ReplaceAllString(s, "")
	return strings.ReplaceAll(s, "\n", "")
}
