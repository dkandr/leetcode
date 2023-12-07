package main

import "testing"

func Test_largestOddNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: "52",
			},
			want: "5",
		},
		{
			name: "2",
			args: args{
				num: "4206",
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				num: "35427",
			},
			want: "35427",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOddNumber(tt.args.num); got != tt.want {
				t.Errorf("largestOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
