package structsmethodsinterfaces

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle: r.Height * r.Width
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Area returns the area of a circle: Pi * c.Radius**2
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perimeter calculates the perimeter of a rectangle: 2 * (width + height)
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle, given by height * width
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
