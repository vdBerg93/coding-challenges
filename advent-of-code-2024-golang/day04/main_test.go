package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var testData []byte

func Test_part1(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: testData,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readInput(tt.args.data)
			if got := part1(data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: testData,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readInput(tt.args.data)
			if got := part2(data); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
