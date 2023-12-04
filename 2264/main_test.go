package main

import "testing"

func Test_largestGoodInteger(t *testing.T) {
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
				num: "6777133339",
			},
			want: "777",
		},
		{
			name: "2",
			args: args{
				num: "2300019",
			},
			want: "000",
		},
		{
			name: "3",
			args: args{
				num: "23019",
			},
			want: "",
		},
		{
			name: "4",
			args: args{
				num: "7678222622241118390785777474281834906756431393782326744172075725179542796491876218340",
			},
			want: "777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
