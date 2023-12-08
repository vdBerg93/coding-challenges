package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	readData()

	// Test data
	got := Part1(testData)
	want1 := 35
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Solution part 1: %d\n", Part1(data))

	//got = Part2(testData)
	//want2 := 30
	//if got != want2 {
	//	log.Fatalf("expected %d, got %d", want2, got)
	//}
	//got = Part2(data)
	//if got != 5923918 {
	//	log.Fatalf("expected 5923918, got %d", got)
	//}
	//fmt.Printf("Solution part 2: %d\n", got)
}

var testData []byte

var data []byte

func readData() {
	var err error
	data, err = os.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("day5/sample.txt")
	if err != nil {
		panic(err)
	}
}

func Part1(data []byte) int {
	text := string(data)
	blocks := strings.Split(text, "\r\n\r\n")

	var sources []int

	for iBlock, block := range blocks {
		if iBlock == 0 {
			sources = getSeeds(block)
			continue
		}
		lines := strings.Split(block, "\r\n")

		ranges := getRanges(lines[1:])

		for i, src := range sources {
			for _, rng := range ranges {
				if rng.InRange(src) {
					sources[i] = rng.GetMapping(src)
				}
			}
		}
	}

	minLocation := math.MaxInt
	for _, result := range sources {
		if result < minLocation {
			minLocation = result
		}
	}
	return minLocation
}

func getRanges(lines []string) []Mapping {
	var ranges []Mapping
	for _, line := range lines {
		fields := strings.Fields(line)
		r := Mapping{}
		for i, val := range fields {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			switch i {
			case 0:
				r.Dst = num
			case 1:
				r.Src = num
			case 2:
				r.Len = num
			default:
				panic("impossible")
			}
		}
		ranges = append(ranges, r)
	}
	return ranges
}

type Mapping struct {
	Src int
	Dst int
	Len int
}

func (m *Mapping) InRange(src int) bool {
	return src >= m.Src && src < m.Src+m.Len
}

func (m *Mapping) GetMapping(src int) (dest int) {
	return m.Dst - m.Src + src
}

func getSeeds(block string) []int {
	var seeds []int
	s := strings.Split(block, ":")
	fields := strings.Fields(s[1])
	for _, f := range fields {
		seed, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}
