package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Inspired by:
https://github.com/ayoubzulfiqar/advent-of-code/blob/main/Go/Day12/
https://gist.github.com/Nathan-Fenner/781285b77244f06cf3248a04869e7161
*/

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
	want1 := 21
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data, false))
}

func Test_Part2(t *testing.T) {
	got := Solve(testData, true)
	want1 := 525152
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 2 test succeeded.\n")
	fmt.Printf("Solution part 2: %d\n", Solve(data, true))
}

func Solve(data []byte, expand bool) int {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	sum := atomic.Int64{}
	wg := sync.WaitGroup{}
	finished := atomic.Int64{}
	for _, row := range rows {
		counter := NewMemoizeCounter()
		gears, runs := parseRow(row)
		if expand {
			gears = unfoldGears(gears, 5)
			runs = unfoldGroups(runs, 5)
		}
		wg.Add(1)
		go func() {
			cnt := counter.CountWays(gears, runs)
			sum.Add(int64(cnt))
			finished.Add(1)
			wg.Done()
			fmt.Printf("Finished %d %%\n", 100*int(finished.Load())/len(rows))
		}()
	}
	wg.Wait()

	return int(sum.Load())
}

func unfoldGroups(groups []int, count int) []int {
	var out []int
	for i := 1; i <= count; i++ {
		out = append(out, groups...)
	}
	return out
}

func unfoldGears(gears string, count int) string {
	var out string
	for i := 1; i <= count; i++ {
		out += gears
		if i != count {
			out += "?"
		}
	}
	return out
}

func parseRow(row string) (string, []int) {
	fields := strings.Fields(row)
	gears := fields[0]
	groupSizes := parseInts(strings.Split(fields[1], ","))

	return gears, groupSizes
}

func (m *Memoize) countWaysNew(gears string, groupSize []int) int {
	if len(gears) == 0 {
		if len(groupSize) == 0 {
			return 1
		}
		return 0
	}
	if len(groupSize) == 0 {
		if strings.Contains(gears, "#") {
			return 0
		}
		return 1
	}

	result := 0

	if gears[0] == '.' || gears[0] == '?' {
		result += m.CountWays(gears[1:], groupSize)
	}

	if gears[0] == '#' || gears[0] == '?' {
		if groupSize[0] <= len(gears) && !strings.Contains(gears[:groupSize[0]], ".") &&
			(groupSize[0] == len(gears) || gears[groupSize[0]] != '#') {
			if groupSize[0] == len(gears) {
				result += m.CountWays("", groupSize[1:])
			} else {
				result += m.CountWays(gears[groupSize[0]+1:], groupSize[1:])
			}
		}
	}
	return result
}

func parseInts(text []string) []int {
	var ints []int
	for _, char := range text {
		val, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		ints = append(ints, val)
	}
	return ints
}

type Memoize struct {
	cache map[string]int
}

type cacheKey struct {
	Row  string
	Runs []int
}

func NewMemoizeCounter() *Memoize {
	return &Memoize{
		cache: make(map[string]int),
	}
}

func (m *Memoize) CountWays(row string, runs []int) int {
	k, err := json.Marshal(cacheKey{row, runs})
	if err != nil {
		panic(err)
	}

	key := string(k)
	if result, ok := m.cache[key]; ok {
		return result
	}

	result := m.countWaysNew(row, runs)
	m.cache[key] = result
	return result
}
