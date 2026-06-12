package algo

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("FizzBuzz up to 5", func(t *testing.T) {
		got := FizzBuzz(5)
		want := []string{"1", "2", "Fizz", "4", "Buzz"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FizzBuzz(5) = %v; want %v", got, want)
		}
	})

	t.Run("FizzBuzz up to 15", func(t *testing.T) {
		got := FizzBuzz(15)
		if got[14] != "FizzBuzz" {
			t.Errorf("FizzBuzz(15)[14] = %q; want \"FizzBuzz\"", got[14])
		}
		if got[2] != "Fizz" {
			t.Errorf("FizzBuzz(15)[2] = %q; want \"Fizz\"", got[2])
		}
		if got[4] != "Buzz" {
			t.Errorf("FizzBuzz(15)[4] = %q; want \"Buzz\"", got[4])
		}
	})
}
