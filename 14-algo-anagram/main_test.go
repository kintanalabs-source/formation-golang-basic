package algo

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"gare", "rage", true},
		{"Gare", "Rage", true},
		{"hello", "world", false},
		{"listen", "silent", true},
		{"rat", "car", false},
		{"", "", true},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.s1, tt.s2); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}
