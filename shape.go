package main

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter
	CalcPerimeter() float64
	// CalcArea returns calculation result of area
	CalcArea() float64
}
