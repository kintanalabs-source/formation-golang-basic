package algo

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
	}
	for _, tt := range tests {
		// On travaille sur une copie pour ne pas corrompre les tests suivants si besoin
		input := make([]int, len(tt.nums))
		copy(input, tt.nums)
		got := BubbleSort(input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSort(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
