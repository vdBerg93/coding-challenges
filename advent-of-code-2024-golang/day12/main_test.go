package main

import "testing"

func Test_countParts(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "a",
			arg:  []int{0, 1, 3, 4},
			want: 2,
		},
		{
			name: "b",
			arg:  []int{0, 1, 1, 2},
			want: 2,
		},
		{
			name: "c",
			arg:  []int{-1, 1, 1, 1, 2, 2, 3, 3, 4, 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parts := split(tt.arg)
			got := len(parts)
			if got != tt.want {
				t.Errorf("countParts() = %v, want %v", got, tt.want)
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
