package functions

import (
	"testing"
)

func TestSwap(t *testing.T) {
	s1, s2 := Swap("hello", "world")
	if s1 != "world" || s2 != "hello" {
		t.Errorf("Swap(\"hello\", \"world\") = %q, %q; want \"world\", \"hello\"", s1, s2)
	}
}

func TestDivide(t *testing.T) {
	t.Run("Valid division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) returned unexpected error: %v", err)
		}
		if got != 5.0 {
			t.Errorf("Divide(10, 2) = %v; want 5.0", got)
		}
	})

	t.Run("Division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) expected error, got nil")
		}
	})
}

func TestGetRectProps(t *testing.T) {
	area, perimeter := GetRectProps(10, 5)
	if area != 50.0 {
		t.Errorf("GetRectProps(10, 5) area = %v; want 50.0", area)
	}
	if perimeter != 30.0 {
		t.Errorf("GetRectProps(10, 5) perimeter = %v; want 30.0", perimeter)
	}
}
