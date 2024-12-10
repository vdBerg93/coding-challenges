package main

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed sample1_1.txt
var sample1_1 []byte

//go:embed sample1_2.txt
var sample1_2 []byte

//go:embed sample1_3.txt
var sample1_3 []byte

//go:embed sample2_1.txt
var sample2_1 []byte

//go:embed sample2_2.txt
var sample2_2 []byte

//go:embed sample2_3.txt
var sample2_3 []byte

//go:embed sample2_4.txt
var sample2_4 []byte

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want int
	}{
		{
			name: "one",
			arg:  sample1_1,
			want: 3,
		},
		{
			name: "two",
			arg:  sample1_2,
			want: 4,
		},
		{
			name: "example",
			arg:  sample1_3,
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := parseInput(bytes.NewReader(tt.arg))
			if got := part1(data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want int
	}{
		{
			name: "one",
			arg:  sample2_1,
			want: 3,
		},
		{
			name: "two",
			arg:  sample2_2,
			want: 13,
		},
		{
			name: "three",
			arg:  sample2_3,
			want: 227,
		},
		{
			name: "four",
			arg:  sample2_4,
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := parseInput(bytes.NewReader(tt.arg))
			if got := part2(data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
