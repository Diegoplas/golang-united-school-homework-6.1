package main

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (t Triangle) CalcPerimeter() float64 {
	trianglePerim := 3 * t.Side
	return trianglePerim
}

// CalcArea returns calculation result of area
func (t Triangle) CalcArea() float64 {
	triangleArea := (math.Sqrt(3.0) / 4) * (t.Side * t.Side)
	return triangleArea
}
