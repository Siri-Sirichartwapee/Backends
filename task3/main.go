package main

import (
	"fmt"
)

// 1) สร้าง interface
type Shape interface {
	Area() float64
}

// 2) สร้าง Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// implement Area ของ Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3) สร้าง Circle
type Circle struct {
	Radius float64
}

// implement Area ของ Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 4) ฟังก์ชันรับ Shape อะไรก็ได้
func PrintArea(s Shape) {
	fmt.Println("Area =", s.Area())
}

func main() {
	// สร้าง object
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	// เรียกใช้
	PrintArea(rect)
	PrintArea(circle)
}