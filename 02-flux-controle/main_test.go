package controlflow

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
	}
	for _, tt := range tests {
		if got := IsEven(tt.n); got != tt.want {
			t.Errorf("IsEven(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}

func TestGetGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{82, "B"},
		{70, "C"},
		{65, "D"},
		{50, "F"},
	}
	for _, tt := range tests {
		if got := GetGrade(tt.score); got != tt.want {
			t.Errorf("GetGrade(%d) = %q; want %q", tt.score, got, tt.want)
		}
	}
}

func TestSumUpTo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 6},   // 1+2+3
		{5, 15},  // 1+2+3+4+5
		{0, 0},
		{1, 1},
	}
	for _, tt := range tests {
		if got := SumUpTo(tt.n); got != tt.want {
			t.Errorf("SumUpTo(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestGetDayName(t *testing.T) {
	tests := []struct {
		day  int
		want string
	}{
		{1, "Lundi"},
		{7, "Dimanche"},
		{4, "Jeudi"},
		{9, "Inconnu"},
	}
	for _, tt := range tests {
		if got := GetDayName(tt.day); got != tt.want {
			t.Errorf("GetDayName(%d) = %q; want %q", tt.day, got, tt.want)
		}
	}
}
