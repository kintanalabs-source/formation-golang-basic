package collections

import (
	"reflect"
	"testing"
)

func TestGetFruits(t *testing.T) {
	got := GetFruits()
	want := []string{"pomme", "banane", "orange"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetFruits() = %v; want %v", got, want)
	}
}

func TestAddFruit(t *testing.T) {
	initial := []string{"pomme"}
	got := AddFruit(initial, "banane")
	want := []string{"pomme", "banane"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddFruit() = %v; want %v", got, want)
	}
}

func TestGetScores(t *testing.T) {
	got := GetScores()
	want := map[string]int{"Alice": 10, "Bob": 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetScores() = %v; want %v", got, want)
	}
}

func TestCheckScore(t *testing.T) {
	scores := map[string]int{"Alice": 10}
	
	t.Run("Existing key", func(t *testing.T) {
		val, ok := CheckScore(scores, "Alice")
		if !ok || val != 10 {
			t.Errorf("CheckScore(Alice) = %d, %v; want 10, true", val, ok)
		}
	})

	t.Run("Missing key", func(t *testing.T) {
		val, ok := CheckScore(scores, "Charlie")
		if ok || val != 0 {
			t.Errorf("CheckScore(Charlie) = %d, %v; want 0, false", val, ok)
		}
	})
}
