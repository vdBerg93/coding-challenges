package main

import (
	"reflect"
	"testing"
)

//
//func Test_day1(t *testing.T) {
//	tests := []struct {
//		name string
//		arg  []byte
//		want int
//	}{
//		{
//			name: "example",
//			arg:  example,
//			want: 0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := part1(tt.arg); got != tt.want {
//				t.Errorf("part1() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_do(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{
			name: "zr",
			arg:  "0",
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := do(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("do() = %v, want %v", got, tt.want)
			}
		})
	}
}
