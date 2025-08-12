package structsmethodsinterfaces

// Perimeter calculates the perimeter of a rectangle: 2 * (width + height)
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area calculates the area of a rectangle, given by height * width
func Area(width, height float64) float64 {
	return width * height
}
