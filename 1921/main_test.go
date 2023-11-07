package main

import "testing"

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				dist:  []int{1, 3, 4},
				speed: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				dist:  []int{1, 1, 2, 3},
				speed: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				dist:  []int{3, 2, 4},
				speed: []int{5, 3, 2},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				dist:  []int{3},
				speed: []int{5},
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				dist:  []int{4, 2, 3},
				speed: []int{2, 1, 1},
			},
			want: 3,
		},
		{
			name: "6",
			args: args{
				dist:  []int{4, 2, 8},
				speed: []int{2, 1, 4},
			},
			want: 2,
		},
		{
			name: "7",
			args: args{
				dist:  []int{3, 5, 7, 4, 5},
				speed: []int{2, 3, 6, 3, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
