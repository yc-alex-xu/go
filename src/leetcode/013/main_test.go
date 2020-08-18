package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "tc1",
			args: args{s: "III"},
			want: 3,
		},
		{name: "tc2",
			args: args{s: "IV"},
			want: 4,
		},
		{name: "tc3",
			args: args{s: "IX"},
			want: 9,
		},
		{name: "tc4",
			args: args{s: "LVIII"},
			want: 58,
		},
		{name: "tc5",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
