package golang_united_school_homework

import (
	"errors"
	"fmt"
)

type ValidIndexError struct{}

func (m *ValidIndexError) Error() string {
	return "please insert a valid index"
}

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	lenShapes := len(b.shapes)
	if lenShapes < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return errors.New("exceeded shapes capacity")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	lastIndex := len(b.shapes) - 1
	if i <= lastIndex {
		fmt.Println("Got Shape: ", b.shapes[i])
		return b.shapes[i], nil
	} else {
		return nil, &ValidIndexError{}
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	lastIndex := len(b.shapes) - 1
	if i <= lastIndex {
		removedShape := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return removedShape, nil
	} else {
		return nil, &ValidIndexError{}
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	lastIndex := len(b.shapes) - 1
	if i <= lastIndex {
		removedShape := b.shapes[i]
		b.shapes[i] = shape
		return removedShape, nil
	} else {
		return nil, &ValidIndexError{}
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	perimeterSum := 0.0
	for _, shape := range b.shapes {
		perimeterSum += shape.CalcPerimeter()
	}
	return perimeterSum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	areaSum := 0.0
	for _, singleShape := range b.shapes {
		areaSum += singleShape.CalcArea()
	}
	return areaSum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circles := 0
	noCircleShapes := []Shape{}
	for _, shape := range b.shapes {
		switch shape.(type) {
		case Circle:
			fmt.Println("Corcñe")
			circles += 1
		case Triangle:
			fmt.Println("Triangle")
			noCircleShapes = append(noCircleShapes, shape)
		case Rectangle:
			fmt.Println("Rectangle")
			noCircleShapes = append(noCircleShapes, shape)
		}
	}
	if circles == 0 {
		return errors.New("no circles to remove")
	}
	b.shapes = noCircleShapes
	return nil
}
