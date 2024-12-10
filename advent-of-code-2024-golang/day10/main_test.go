package main

import (
	"bytes"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want int
	}{
		{
			name: "one",
			arg:  sample,
			want: 3,
		},
		{
			name: "two",
			arg:  sample2,
			want: 4,
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
