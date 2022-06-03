package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

// CalcPerimeter returns calculation result of perimeter
func (r Rectangle) CalcPerimeter() float64 {
	rectanglePerim := 2*r.Height + 2*r.Weight
	return rectanglePerim
}

// CalcArea returns calculation result of area
func (r Rectangle) CalcArea() float64 {
	rectangleArea := r.Height * r.Weight
	return rectangleArea
}
