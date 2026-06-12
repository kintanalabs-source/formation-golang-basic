package algo

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	
	tests := []struct {
		target int
		want   int
	}{
		{7, 3},
		{1, 0},
		{15, 7},
		{8, -1},
		{20, -1},
		{-5, -1},
	}
	for _, tt := range tests {
		if got := BinarySearch(nums, tt.target); got != tt.want {
			t.Errorf("BinarySearch(nums, %d) = %d; want %d", tt.target, got, tt.want)
		}
	}
}
