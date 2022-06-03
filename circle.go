package golang_united_school_homework

const pi = 3.14159265

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

// CalcPerimeter returns calculation result of perimeter
func (c Circle) CalcPerimeter() float64 {
	circlePerim := 2 * pi * c.Radius
	return circlePerim
}

// CalcArea returns calculation result of area
func (c Circle) CalcArea() float64 {
	circleArea := pi * c.Radius * c.Radius
	return circleArea
}
