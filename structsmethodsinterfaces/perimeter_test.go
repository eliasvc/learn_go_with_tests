package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(3.0, 5.0)
	want := 16.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	width := 10.0
	height := 10.0

	fmt.Printf("Perimeter: %.1f", Perimeter(width, height))
	// Output: Perimeter: 40.0
}
