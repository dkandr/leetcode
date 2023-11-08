package main

import "testing"

func Test_isReachableAtTime(t *testing.T) {
	type args struct {
		sx int
		sy int
		fx int
		fy int
		t  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				sx: 2,
				sy: 4,
				fx: 7,
				fy: 7,
				t:  6,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				sx: 3,
				sy: 1,
				fx: 7,
				fy: 3,
				t:  3,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				sx: 1,
				sy: 1,
				fx: 1,
				fy: 2,
				t:  0,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				sx: 1,
				sy: 1,
				fx: 2,
				fy: 2,
				t:  1,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				sx: 1,
				sy: 2,
				fx: 1,
				fy: 2,
				t:  1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReachableAtTime(tt.args.sx, tt.args.sy, tt.args.fx, tt.args.fy, tt.args.t); got != tt.want {
				t.Errorf("isReachableAtTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
