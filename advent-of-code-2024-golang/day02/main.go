package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	reports := parseReports(input)
	fmt.Println("Part 1 solution:", part1(reports))
	fmt.Println("Part 2 solution:", part2(reports))
}

func part1(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if ok := evaluateReport(report, false); ok {
			safe++
		}
	}
	return safe
}

func part2(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if evaluateReport(report, true) {
			safe++
		}
	}
	return safe
}

func evaluateReport(report []int, withDamping bool) bool {
	ok, _ := isSafe(report)
	if ok {
		return true
	}
	if !withDamping {
		return false
	}
	for idx := range report {
		subReport := dropIndex(report, idx)
		ok, _ = isSafe(subReport)
		if ok {
			return true
		}
	}
	return false
}

func dropIndex(report []int, idx int) []int {
	out := make([]int, 0, len(report)-1)
	for i, v := range report {
		if i == idx {
			continue
		}
		out = append(out, v)
	}
	return out
}

func isSafe(report []int) (bool, int) {
	if len(report) == 0 {
		panic("empty")
	}
	if len(report) == 1 {
		return true, 0
	}
	var sign int
	for i := 1; i < len(report); i++ {
		delta := report[i] - report[i-1]
		if delta == 0 || math.Abs(float64(delta)) > 3 {
			return false, i
		}
		if i == 1 {
			sign = getSign(delta)
			continue
		}
		if getSign(delta) != sign {
			return false, i
		}
	}

	return true, 0
}

func getSign(v int) int {
	if v > 0 {
		return 1
	} else if v < 0 {
		return -1
	} else {
		return 0
	}
}

func parseReports(data []byte) [][]int {
	rows := strings.Split(string(data), "\n")
	reports := make([][]int, 0, len(rows))
	for _, row := range rows {
		reports = append(reports, parseReport(row))
	}
	return reports
}

func parseReport(row string) []int {
	f := strings.Fields(row)
	report := make([]int, 0, len(f))
	for _, fi := range f {
		val, err := strconv.Atoi(fi)
		if err != nil {
			panic(err)
		}
		report = append(report, val)
	}
	return report
}
