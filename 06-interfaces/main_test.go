package interfaces

import (
	"math"
	"testing"
)

func TestInterfaces(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 10}

	t.Run("Rectangle Area", func(t *testing.T) {
		got := rect.Area()
		want := 50.0
		if got != want {
			t.Errorf("Rectangle.Area() = %v; want %v", got, want)
		}
	})

	t.Run("Circle Area", func(t *testing.T) {
		got := circle.Area()
		want := math.Pi * 100
		if got != want {
			t.Errorf("Circle.Area() = %v; want %v", got, want)
		}
	})

	t.Run("Total Area", func(t *testing.T) {
		shapes := []Shape{rect, circle}
		got := GetTotalArea(shapes)
		want := 50.0 + (math.Pi * 100)
		if got != want {
			t.Errorf("GetTotalArea() = %v; want %v", got, want)
		}
	})
}
