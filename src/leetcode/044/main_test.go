package main

import (
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{s: "aa", p: "a"},
			want: false,
		},
		{
			name: "tc2",
			args: args{s: "aa", p: "*"},
			want: true,
		},
		{
			name: "tc3",
			args: args{s: "cb", p: "?a"},
			want: false,
		},
		{
			name: "tc4",
			args: args{s: "adceb", p: "*a*b"},
			want: true,
		},
		{
			name: "tc5",
			args: args{s: "acdcb", p: "a*c?b"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
