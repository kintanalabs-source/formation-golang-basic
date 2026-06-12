package algo

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"radar", true},
		{"Radar", true},
		{"hello", false},
		{"kayak", true},
		{"Go", false},
		{"", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.s); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tt.s, got, tt.want)
		}
	}
}
