package main

import (
	"testing"
)

func Test_stripUrl(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripUrl(tt.args.link); got != tt.want {
				t.Errorf("stripUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePass(t *testing.T) {
	type args struct {
		key   string
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePass(tt.args.key, tt.args.input...); got != tt.want {
				t.Errorf("generatePass() = %v, want %v", got, tt.want)
			}
		})
	}
}
