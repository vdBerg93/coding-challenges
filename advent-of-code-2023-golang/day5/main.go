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

	got = Part2(testData)
	want2 := 46
	if got != want2 {
		log.Fatalf("expected %d, got %d", want2, got)
	}
	got = Part2(data)
	fmt.Printf("Solution part 2: %d\n", got)
}

var testData []byte

var data []byte

func readData() {
	var err error
	data, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample")
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
				if rng.SeedInRange(src) {
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

func (m *Mapping) SeedInRange(src int) bool {
	return src >= m.Src && src < m.Src+m.Len
}

func (m *Mapping) RangeInRange(start, rng int) bool {
	return m.Src <= start && start < m.Src+m.Len || // Start is in range
		m.Src <= start+rng && start+rng < m.Src+m.Len || // End is in range
		start <= m.Src && m.Src <= start+rng ||
		start < m.Src+m.Len && m.Src+m.Len <= start+rng
}

func (m *Mapping) GetMapping(src int) (dest int) {
	return m.Dst - m.Src + src
}

func getSeeds(block string) []int {
	s := strings.Fields(strings.Split(block, ":")[1])
	seeds := make([]int, len(s))

	for i, f := range s {
		seed, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		seeds[i] = seed
	}
	return seeds
}

func getSeedsFromRanges(seedRanges []int) []int {
	totalSeeds := 0

	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		stop := start + seedRanges[i+1]

		totalSeeds += stop - start + 1
	}

	seeds := make([]int, 0, totalSeeds)
	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		stop := start + seedRanges[i+1]

		for seed := start; seed <= stop; seed++ {
			seeds = append(seeds, seed)
		}
	}

	return seeds
}

func Part2(data []byte) int {
	text := string(data)
	blocks := strings.Split(text, "\r\n\r\n")

	seedRanges := getSeeds(blocks[0])
	overallMin := math.MaxInt
	for i := 0; i < len(seedRanges); i += 2 {
		seeds := getSeedsFromRanges(seedRanges[i : i+2])

		for _, block := range blocks[1:] {
			lines := strings.Split(block, "\r\n")

			ranges := getRanges(lines[1:])

			for i, src := range seeds {
				for _, rng := range ranges {
					if rng.SeedInRange(src) {
						seeds[i] = rng.GetMapping(src)
					}
				}
			}
		}
		for _, loc := range seeds {
			overallMin = min(loc, overallMin)
		}
	}

	return overallMin
}
