package main

import (
	_ "embed"
	"testing"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

//go:embed example3.txt
var example3 []byte

//go:embed example4.txt
var example4 []byte

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want int
	}{
		{
			name: "one",
			arg:  example1,
			want: 80,
		},
		{
			name: "two",
			arg:  example2,
			want: 436,
		},
		{
			name: "three",
			arg:  example3,
			want: 236,
		},
		{
			name: "four",
			arg:  example4,
			want: 368,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.arg); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
