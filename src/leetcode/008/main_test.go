package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{str: "42"},
			want: 42,
		},
		{
			name: "tc2",
			args: args{str: "-42"},
			want: -42,
		},
		{
			name: "tc3",
			args: args{str: "4193 with words"},
			want: 4193,
		},
		{
			name: "tc4",
			args: args{str: "words and 987"},
			want: 0,
		},
		{
			name: "tc5",
			args: args{str: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "tc6",
			args: args{str: "9223372036854775808"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
