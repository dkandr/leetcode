package main

import "testing"

func Test_minPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 5, 2, 3},
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 5, 4, 2, 4, 6},
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				nums: []int{4, 1, 5, 1, 2, 5, 1, 5, 5, 4},
			},
			want: 8,
		},
		{
			name: "4",
			args: args{
				nums: []int{5, 3, 5, 2, 1, 5, 5, 2, 3, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPairSum(tt.args.nums); got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
