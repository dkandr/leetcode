package main

import "testing"

func Test_numSpecial(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				mat: [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				mat: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				mat: [][]int{{0, 0, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 1, 0, 0}},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				mat: [][]int{{0, 0, 0, 1}, {1, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecial(tt.args.mat); got != tt.want {
				t.Errorf("numSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
