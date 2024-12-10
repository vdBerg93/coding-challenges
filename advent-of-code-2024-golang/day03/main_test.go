package main

import (
	"reflect"
	"testing"
)

func Test_getStatements(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example",
			args: args{
				s: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatements(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStatements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.data); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterConditions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				s: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!^do()?mul(8,5))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterConditions(tt.args.s); got != tt.want {
				t.Errorf("filterConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}
