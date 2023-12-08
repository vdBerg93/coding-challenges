package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func run() {
	// Test data
	got := Part1(testData)
	want1 := 6440
	if got != want1 {
		log.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Solution part 1: %d\n", Part1(data))

	//got = Part2(testData)
	//want2 := 71503
	//if got != want2 {
	//	log.Fatalf("expected %d, got %d", want2, got)
	//}
	//got = Part2(data)
	//fmt.Printf("Solution part 2: %d\n", got)
}

var testData []byte

var data []byte

func Part1(data []byte) int {
	rows := strings.Split(string(data), "\r\n")

	sort.Slice(rows, func(i, j int) bool {
		return isLess(rows[i][0:6], rows[j][0:6])
	})

	return calculateWinnings(rows)
}

func isLess(hand1, hand2 string) bool {
	r1 := getTypeRank(hand1)
	r2 := getTypeRank(hand2)

	if r1 < r2 {
		return true
	} else if r1 > r2 {
		return false
	} else {
		return hand1LosesTie(hand1, hand2)
	}
}

func getTypeRank(hand string) int {
	cards := countCards(hand)
	if countKind(5, cards) {
		return 7
	}
	if countKind(4, cards) {
		return 6
	}
	if fullhouse(cards) {
		return 5
	}
	if countKind(3, cards) {
		return 4
	}
	pairs := countPairs(cards)
	if pairs == 2 {
		return 3
	}
	if pairs == 1 {
		return 2
	}
	return 1
}

func countCards(hand string) map[rune]int {
	cnt := make(map[rune]int)
	for _, card := range hand {
		cnt[card]++
	}
	return cnt
}

func countPairs(cards map[rune]int) int {
	pairs := 0
	for _, count := range cards {
		if count == 2 {
			pairs++
		}
	}
	return pairs
}

func fullhouse(cards map[rune]int) bool {
	var gotTwos, gotThrees bool
	for _, cardCount := range cards {
		switch cardCount {
		case 2:
			gotTwos = true
		case 3:
			gotThrees = true
		}
	}
	return gotTwos && gotThrees
}

func countKind(count int, cards map[rune]int) bool {
	for _, cardCount := range cards {
		if cardCount == count {
			return true
		}
	}
	return false
}

func hand1LosesTie(hand1, hand2 string) bool {
	hand1r := []rune(hand1)
	hand2r := []rune(hand2)

	for i, c1 := range hand1r {
		c2 := hand2r[i]
		if getCardRank(c1) < getCardRank(c2) {
			return true
		}
		if getCardRank(c1) > getCardRank(c2) {
			return false
		}
	}
	panic("shouldn't be possible")
	return true
}

func getCardRank(card rune) int {
	switch card {
	case 'A':
		return 13
	case 'K':
		return 12
	case 'Q':
		return 11
	case 'J':
		return 10
	case 'T':
		return 9
	case '9':
		return 8
	case '8':
		return 7
	case '7':
		return 6
	case '6':
		return 5
	case '5':
		return 4
	case '4':
		return 3
	case '3':
		return 2
	case '2':
		return 1
	}
	return 0
}

func calculateWinnings(rows []string) int {
	score := 0
	for rank, row := range rows {
		bid, err := strconv.Atoi(strings.Fields(row)[1])
		if err != nil {
			panic(err)
		}
		score += bid * (rank + 1)
	}
	return score
}
