package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global variables
var patterns2 []string
var memo2 = map[string]int{}

// count2 recursively checks how many ways a string can be formed using prefixes from patterns
func count2(d string) int {
	// Base case: if the string is empty, there's one way (it's already complete)
	if d == "" {
		return 1
	}

	// Check memoization cache
	if val, exists := memo2[d]; exists {
		return val
	}

	// Calculate the number of ways recursively
	total := 0
	for _, p := range patterns2 {
		if strings.HasPrefix(d, p) {
			total += count2(strings.TrimPrefix(d, p))
		}
	}

	// Cache the result
	memo2[d] = total
	return total
}

func main() {
	// Open the input file
	file, err := os.Open("day19/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file lines
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the patterns and data
	patterns2 = strings.Split(lines[0], ", ")
	data := lines[2:] // Skip the second empty line and process the rest

	// Compute boolean and integer results
	booleanSum := 0
	integerSum := 0
	for _, d := range data {
		result := count2(d)
		if result > 0 {
			booleanSum++
		}
		integerSum += result
	}

	// Print the results
	fmt.Println(booleanSum)
	fmt.Println(integerSum)
}
