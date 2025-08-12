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
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("got %0.1f want %0.1f", got, want)
	}
}

func ExamplePerimeter() {
	r := Rectangle{10.0, 10.0}
	fmt.Printf("%.2f", Perimeter(r))
	// Output: 40.00
}
