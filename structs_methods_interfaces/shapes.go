package structs_methods_interfaces

import "math"

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the circle triangle
func (triangle Triangle) Area() float64 {
	return (triangle.Base * triangle.Height) * 0.5
}
