package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle: 2 * (width + height)
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle, given by height * width
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
