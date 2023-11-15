package main

import "testing"

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{100, 1, 1000},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				arr: []int{2, 2, 1, 2, 1},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "4",
			args: args{
				arr: []int{63, 94, 4, 61, 6, 40},
			},
			want: 6,
		},
		{
			name: "5",
			args: args{
				arr: []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
			},
			want: 8,
		},
		{
			name: "6",
			args: args{
				arr: []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 8, 9},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.args.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
