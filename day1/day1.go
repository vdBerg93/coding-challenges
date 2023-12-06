package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Day1_Ex1(filePath string) int64 {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	_, _ = buf.Write(data)
	scanner := bufio.NewScanner(&buf)
	scanner.Split(bufio.ScanLines)

	var sum int64
	for scanner.Scan() {
		row := scanner.Text()

		r := regexp.MustCompile("[^0-9]")
		numbers := r.ReplaceAllString(row, "")

		if len(numbers) == 0 {
			continue
		}

		str := string(numbers[0]) + string(numbers[len(numbers)-1])

		valInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if valInt > 99 {
			log.Fatal(">99")
		}
		sum += valInt
	}

	return sum
}

func replaceFirst(s string) string {
	if unicode.IsDigit(rune(s[0])) {
		return s
	}

	for i := 1; i < len(s); i++ {
		for _, pair := range replacementPairs {
			if strings.Contains(s[0:i], pair.from) {
				r := strings.Replace(s[0:i], pair.from, pair.to, -1)
				return r + s[i:]
			}
		}
		if unicode.IsDigit(rune(s[i])) {
			return s
		}
	}

	return s
}

func replaceLast(s string) string {
	if unicode.IsDigit(rune(s[len(s)-1])) {
		return s
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return s
		}
		for _, pair := range replacementPairs {
			if strings.Contains(s[i:], pair.from) {
				r := strings.Replace(s[i:], pair.from, pair.to, -1)
				return s[0:i] + r
			}
		}
	}

	return s
}

var replacementPairs = []struct{ from, to string }{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func Day1_Ex2(filePath string) int64 {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	_, _ = buf.Write(data)
	scanner := bufio.NewScanner(&buf)
	scanner.Split(bufio.ScanLines)

	var sum int64
	for scanner.Scan() {
		row := scanner.Text()

		rowNumbers := replaceFirst(row)
		rowNumbers = replaceLast(rowNumbers)

		r := regexp.MustCompile("[^0-9]")
		numbers := r.ReplaceAllString(rowNumbers, "")

		if len(numbers) == 0 {
			continue
		}

		str := string(numbers[0]) + string(numbers[len(numbers)-1])

		valInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if valInt > 99 {
			log.Fatal(">99")
		}
		sum += valInt
	}

	return sum
}

// too high: 56345
// too low: 55100
