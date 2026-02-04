package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6, SideA: 8, SideB: 18}, hasPerimeter: 38.0},
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasPerimeter: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()

		if got != tt.hasPerimeter {
			t.Errorf("got %g hasPerimeter %g", got, tt.hasPerimeter)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 20.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g hasArea %g", got, tt.hasArea)
			}
		})
	}
}
