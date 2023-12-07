package main

import "testing"

func Test_getGearRatio(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"r3-c37",
			args{
				row: 3,
				col: 37,
			},
			203 * 917,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGearRatio(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("getGearRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
