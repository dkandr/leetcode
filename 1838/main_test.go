package main

import "testing"

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				nums: []int{1, 2, 2, 5, 5, 10},
				k:    1,
			},
			want: 3,
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 4},
				k:    5,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    5,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 9, 6},
				k:    2,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				nums: []int{9940, 9995, 9944, 9937, 9941, 9952, 9907, 9952, 9987, 9964, 9940, 9914, 9941, 9933, 9912, 9934, 9980, 9907, 9980, 9944, 9910, 9997},
				k:    7925,
			},
			want: 22,
		},
		{
			name: "5",
			args: args{
				nums: []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966},
				k:    3056,
			},
			want: 73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
