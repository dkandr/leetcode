package main

import "testing"

func Test_sortVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				s: "lYmpH",
			},
			want: "lYmpH",
		},
		{
			name: "1",
			args: args{
				s: "lEetcOde",
			},
			want: "lEOtcede",
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				s: "uoieaUOIEA",
			},
			want: "AEIOUaeiou",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
