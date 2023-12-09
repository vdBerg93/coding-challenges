package main

import "testing"

func Test_replaceFirst(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "1",
			in:   "two1nine",
			want: "21nine",
		},
		{
			"2",
			"abcone2threexyz",
			"abc12threexyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceFirst(tt.in); got != tt.want {
				t.Errorf("replaceFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceLast(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "1",
			in:   "two1nine",
			want: "219",
		},
		{
			"2",
			"abc12threexyz",
			"abc123xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceLast(tt.in); got != tt.want {
				t.Errorf("replaceFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay1_Ex2(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     int64
	}{
		{
			"test input",
			"day1-input-sample.txt",
			281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day1_Ex2(tt.filePath); got != tt.want {
				t.Errorf("Day1_Ex2() = %v, want %v", got, tt.want)
			}
		})
	}
}
