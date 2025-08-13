package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.0, 5.0}
	got := Perimeter(rectangle)
	want := 16.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("area of a triangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %0.1f want %0.1f", got, want)
		}
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func ExamplePerimeter() {
	r := Rectangle{10.0, 10.0}
	fmt.Printf("%.2f", Perimeter(r))
	// Output: 40.00
}
