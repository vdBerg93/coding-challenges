package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	readData()

	// Test data
	got := Part1(testData)
	want1 := 13
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Solution part 1: %d\n", Part1(data))

	got = Part2(testData)
	want2 := 30
	if got != want2 {
		log.Fatalf("expected %d, got %d", want2, got)
	}
	got = Part2(data)
	if got != 5923918 {
		log.Fatalf("expected 5923918, got %d", got)
	}
	fmt.Printf("Solution part 2: %d\n", got)
}

var testData []byte

var data []byte

var lines [][]byte

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
	//printCounts()

	lines = bytes.Split(data, []byte("\n"))

	var points int
	for _, line := range lines {
		got, win := parseLine(line)
		matchCount := getMatchingCount(got, win)
		points += getScore(matchCount)
	}

	return points
}

func getMatchingCount(got, win []string) int {
	count := 0
	for _, g := range got {
		for _, w := range win {
			if g == w {
				count++
				break
			}
		}
	}
	return count
}

func getScore(matchCount int) int {
	if matchCount == 0 {
		return 0
	}
	score := 1
	for i := 1; i < matchCount; i++ {
		score *= 2
	}
	return score
}

func parseLine(data []byte) (got, winning []string) {
	parts := strings.Split(string(data), ":")

	scores := strings.Split(parts[1], "|")

	got = strings.Fields(scores[1])
	winning = strings.Fields(scores[0])

	return
}

func Part2(data []byte) int {

	lines = bytes.Split(data, []byte("\n"))

	cc := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cc[i] = 1
	}

	for cardId := range lines {
		got, win := parseLine(lines[cardId])
		matchCount := getMatchingCount(got, win)

		// for the number of wins, update any cards with extras.
		for i := 0; i < matchCount; i++ {
			cc[cardId+1+i] += cc[cardId]
		}
		fmt.Println(cc)

	}

	os.Exit(1)

	sum := 0
	for _, n := range cc {
		sum += n
	}
	return sum
}

var wg sync.WaitGroup
var cardCount atomic.Int64

func process(lineIdx int) {
	defer wg.Done()

	got, win := parseLine(lines[lineIdx])
	matchCount := getMatchingCount(got, win)
	if matchCount == 0 {
		return
	}

	for i := lineIdx + 1; i <= lineIdx+matchCount; i++ {
		cardCount.Add(1)
		wg.Add(1)
		process(i)
	}
}

func Part2_old(data []byte) int {

	lines = bytes.Split(data, []byte("\n"))
	cardCount.Add(int64(len(lines)))

	for i := range lines {
		wg.Add(1)
		go process(i)
	}
	wg.Wait()

	return int(cardCount.Load())
}
