package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1. Two Sum", args: args{origin: "1. Two Sum"}, want: "01-TwoSum"},
		{name: "49. Group Anagrams", args: args{origin: "49. Group Anagrams"}, want: "49-GroupAnagrams"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.origin); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
