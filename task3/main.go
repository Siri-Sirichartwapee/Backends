package main

import (
	"fmt"
	"math"
)

// Shape interface defines the Area method.
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// PrintArea prints the area of any Shape.
func PrintArea(s Shape) {
	fmt.Printf("Shape area: %.2f\n", s.Area())
}

func main() {
	fmt.Println("--- Task 3: Interfaces ---")
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Print("Rectangle: ")
	PrintArea(rect)
	fmt.Print("Circle: ")
	PrintArea(circ)
}
