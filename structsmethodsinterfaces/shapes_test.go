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

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v: got %g want %g", tt, got, tt.want)
			}
		})
	}
}

func ExamplePerimeter() {
	r := Rectangle{10.0, 10.0}
	fmt.Printf("%.2f", Perimeter(r))
	// Output: 40.00
}
