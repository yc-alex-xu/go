package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "tc1",
			args: args{num: 3},
			want: "III",
		},
		{name: "tc2",
			args: args{num: 4},
			want: "IV",
		},
		{name: "tc3",
			args: args{num: 9},
			want: "IX",
		},
		{name: "tc4",
			args: args{num: 58},
			want: "LVIII",
		},
		{name: "tc5",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
