package main

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		arg  Point
		want Point
	}{
		{arg: Point{x: 1, y: 0}, want: Point{x: 0, y: 1}},
		{arg: Point{x: 0, y: -1}, want: Point{x: 1, y: 0}},
		{arg: Point{x: -1, y: 0}, want: Point{x: 0, y: -1}},
		{arg: Point{x: 0, y: 1}, want: Point{x: -1, y: 0}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := Pose{Heading: tt.arg}
			want := Pose{Heading: tt.want}
			if got := rotate(a); !reflect.DeepEqual(got, want) {
				t.Errorf("rotate() = %v, want %v", got, want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	data, err := os.OpenFile("example.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	want := 41
	if got := part1(parseInput(data)); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}

}

func Test_part2(t *testing.T) {
	data, err := os.OpenFile("example.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	if got := part2(parseInput(data)); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}

}
