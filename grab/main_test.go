package main

import "testing"

func Test_genIpaddr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genIpaddr()
			if got != tt.want {
				t.Errorf("genIpaddr() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
