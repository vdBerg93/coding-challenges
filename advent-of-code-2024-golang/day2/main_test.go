package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed test.txt
var testData []byte

func Test_main(t *testing.T) {
	reports := parseReports(testData)

	want1 := 2
	if got := part1(reports); got != want1 {
		t.Fatalf("part 1: got %v, want %v", got, want1)
	}
	want2 := 4
	if got := part2(reports); got != want2 {
		t.Fatalf("part 2: got %v, want %v", got, want2)
	}
}

func Test_dropIndex(t *testing.T) {
	type args struct {
		report []int
		idx    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok",
			args: args{
				report: []int{1, 2, 3, 4, 5, 6},
				idx:    1,
			},
			want: []int{1, 3, 4, 5, 6},
		},
		{
			name: "last",
			args: args{
				report: []int{1, 2, 3, 4, 5, 6},
				idx:    5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dropIndex(tt.args.report, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dropIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
