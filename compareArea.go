//1) Write a program that supports following features :
//a) Compare area of two rectangles and returns an appropriate value (i.e comparison result).
//b) Compare area of two circles and returns an appropriate value (i.e comparison result).
//c) Compare area of a rectangle and a circle and returns an appropriate value (i.e comparison result).

package main

import (
	"fmt"
	"math"
	"reflect"
)

// Shape is an interface for calculating the area of different shapes
type Shape interface {
	Area() float64
}

// Rectangle holds W and H
type Rectangle struct {
	Width, Height float64
}

// NewRectangle initializes the Rectangle with given W and H
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}

// Area here calculates the area of rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle holds the radius of the circle
type Circle struct {
	Radius float64
}

// NewCircle initializes the circle with given radius r
func NewCircle(r float64) *Circle {
	return &Circle{Radius: r}
}

// Area here calculates the area of the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// getShapeName returns the name of the shape
func getShapeName(s Shape) string {
	return reflect.TypeOf(s).Elem().Name()
}

func CompareArea(a, b Shape) {
	aName := getShapeName(a)
	bName := getShapeName(b)
	aArea := a.Area()
	bArea := b.Area()

	if aArea > bArea {
		fmt.Printf("The area of the %s (%.2f) is greater than the area of the %s (%.2f).\n", aName, aArea, bName, bArea)
	} else if aArea < bArea {
		fmt.Printf("The area of the %s (%.2f) is less than the area of the %s (%.2f).\n", aName, aArea, bName, bArea)
	} else {
		fmt.Printf("The area of the %s (%.2f) is equal to the area of the %s (%.2f).\n", aName, aArea, bName, bArea)
	}
}

func main() {
	rec1 := NewRectangle(25.0, 50.0)
	rec2 := NewRectangle(32.0, 25.0)
	CompareArea(rec1, rec2)

	cir1 := NewCircle(25.0)
	cir2 := NewCircle(32.0)
	CompareArea(cir1, cir2)

	rec3 := NewRectangle(23.0, 36.0)
	cir3 := NewCircle(23.0)
	CompareArea(rec3, cir3)
}
