package main

import "testing"

func Test_countHomogenous(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "abbcccaa",
			},
			want: 13,
		},
		{
			name: "2",
			args: args{
				s: "xy",
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				s: "zzzzz",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHomogenous(tt.args.s); got != tt.want {
				t.Errorf("countHomogenous() = %v, want %v", got, tt.want)
			}
		})
	}
}
