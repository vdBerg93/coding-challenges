package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
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
	got := Solve(testData, false)
	want1 := 114
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, false))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, true)
	want1 := 2
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	got2 := Solve(data, true)
	if got2 != 1031 {
		log.Fatalf("expected 1031, got %d", got2)
	}
	fmt.Printf("Solution part 2: %d\n", got2)
}

func Solve(text []byte, reverse bool) int {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	scanner.Split(bufio.ScanLines)

	var data [][]int
	for scanner.Scan() {
		data = append(data, parseInts(scanner.Text()))
	}

	result := 0
	for _, row := range data {
		if reverse {
			sort.Sort(sort.Reverse(sort.IntSlice(row)))
		}
		result += rowPredictionSum(row, reverse)
	}

	return result
}

func rowPredictionSum(data []int, reverse bool) int {

	stack := [][]int{}
	stack = append(stack, data)
	for {
		diff := getDifference(stack[len(stack)-1])
		stack = append(stack, diff)
		if allZeros(diff) {
			break
		}
	}
	//printStack(stack)
	// Backtracking
	sum := 0
	for i := len(stack) - 1; i >= 0; i-- {
		r := stack[i]
		if len(r) == 0 {
			continue
		}
		sum += r[len(r)-1]
	}
	return sum
}

func printStack(stack [][]int) {
	fmt.Println("----------")
	for _, row := range stack {
		fmt.Println(row)
	}

}

func getDifference(data []int) []int {
	diff := make([]int, len(data)-1)
	for i := 0; i < len(data)-1; i++ {
		diff[i] = data[i+1] - data[i]
	}
	return diff
}

func allZeros(data []int) bool {
	for _, val := range data {
		if val != 0 {
			return false
		}
	}
	return true
}

func parseInts(row string) (values []int) {
	for _, num := range strings.Fields(row) {
		val, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	return values
}
