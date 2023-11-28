package main

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		corridor string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				corridor: "SSPPSPS",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				corridor: "PPSPSP",
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				corridor: "S",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				corridor: "SPPSSSSPPS",
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				corridor: "SPSPPSSPSSSS",
			},
			want: 6,
		},
		{
			name: "6",
			args: args{
				corridor: "SPPPPPPPSPPPSPSSSPPPPPPPPPPPPPPPPPSPPPPPPPPPPPPPPPPSPPPPPSPSPPPPPPSPSPPSPSPPPSPSPPSSPPPPPSPPSSPP",
			},
			want: 0,
		},
		{
			name: "7",
			args: args{
				corridor: "SSSSPPPPPSSPSPSSPPPPPPSPSPSSPPPSPSPSPSSPPPSSSSPSSPSPSPSSSSPPSSPSSSSSSPPPSPSSPPSSPSPSPSSPSSSSSSPSSSPSSSPSSPSSPPPSSPSSPPSSSPSSPSPSSSSPSPPSSPSSPPSSPPSPSPPPSPSPPSPPSPSPPPPPPPPSPSSSPPSPSSPSSSSPPSPSPPPPSPPSSPSPSSPSSPSPPPPPPSSPSPPPPSSPSPPPPPSSSSPSSSSSPPSPPPSPSSPPSPSPSSSSPPPSPPSSSPSPPSSSSSSSSPSSSSPSPSSSPSPPPPPSPSPSPPPSSSPSSPSPSSSPPPPSSPSPSSSSPSSSSSPPSPSSSSPSSSSSSPPPSSPPPPPPSPSSSPPSPPSSSPPPSSPPPSPPSSPPSPSSSSSSPPSSSPPSPSPSPPSPSSPPSPSSPSPPSPSPPSSPSSPPSPPSPSPPSPPSSSSPSPPPSSPPSSPPPSPPSSPSPPSPSPSSPSPPSSSPPSSSPSSPSPPSPSPSPSPSPSSPPPSPSSSSSPPSPSSSPSSSPSSSSSSPSPPPPPPPSSPPSPPSPPSPSSSSPSSSPPPPSPSSPSSPSSPPSSSSPPSPSPSPPPSSSPSPSSPPSSSPPSSPSPPPPPSSSPPSSPSPPPSSSPPPPPSPPPPPSPSPSSSPPSPSSPPSPPSSSPSSSPPSSPPSPSSSPSPPSSSSPSPPSSSPSSPSPPPSPPPSPSSPSPPSSSPPSPSSPSPPPPSSSPSSPPSPSPSPSPSSSPSPSPPPPSSSSSSPPPPPPSPPPSPSPPSPSPPSPPPPPSSSPSPPPSPPSPSPPSSSPPSPSSPPPPPSSPPPPPSSPSSPPSPSPSSSPPPSSSSPPPSPPPPSSSPSPPPSPPPSPPSPSSSPPPPPSSSSPPPPSSSSPSPSPSPPSPSPPSPSPSPSSPPPSSPPPPPSSPSPPPPPPSSPSSPSPSSPPSPPSSSPPSPSSSSSSSSSSSPSSPSPSPPPSPSPPPPSPSSSPPPSSPSPSSSSPSPPPSPSPPSPSPPPSPSPPSSSPSPPSSSPPPPPPSPPSSPPPPSPPPPPSPPPPPSPSPSSPPPSPSSSSSPSSSSPPPSSSPPSPPPPPSSPSSSPSPSPSSSPPSSPSPSPPSSPSPSPSPSPPSSSSPPPSSPPPSSSPPSPSPPPPSPSPSPPPSSPSSPSSPPPPSSSPPPPSSSPPPPPPSPSSSPSSPSPSSSSPSPSPSPPSPSSSPPSSSSSPSSPPPPSSSSPPPSPSSPPSPSSPPPPPPPSSPSPPPSSPPPPSSSSSSPSPPPSPPSSSPSPSSSPSSPSSSSSPSSPSPSSPSSPPSSSSPPSPSSPPSSSPSPPSSSPSSPSSSPSPSPPSSPPPPSSPSPSPPSPPPPPSPSSSPPSPPPPSSPPSPSPSSPSPSPSSPPPPPPSSSPPPPSPSPPSSPSPSSSPSSPPPPSPSSSSPPSPPSSSPPSPPPSSPSPSPPPSPSSSSSPPSPSPPPSSPPPSSSPPSSSPPSPSSPPSSSPPSPPPPPPSSSPSSPSSPSPPPSPSPPPSPSSPSSSSSSSSPPPSSSSPPSSSPSSSPSSPPPSSPSPPPPPPPPSSPPSPPPSPPPSSPSSPPPPPSSSSPPSSPPPPSS",
			},
			want: 596333475,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.corridor); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
