package advent_of_code_2024_golang

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

const token = "53616c7465645f5ff7e3d1be244120f58ab701782eaf29624c6dadfeb9d769e85f4b8747dfd9df1ce7f16cb1abe1a8dcd4022f659c2412e9f29f5e0d12532be8"

func getInput(token string) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/day/11/input", bytes.NewReader([]byte{}))
	if err != nil {
		log.Fatalf("NewRequest: %v", err)
	}
	client := http.Client{Timeout: 10 * time.Second}
	req.Header.Set("session", token)
	//req.AddCookie(&http.Cookie{Name: "session", Value: token})
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Do: %v", err)
	}

	fmt.Printf("%+v", resp)
}
