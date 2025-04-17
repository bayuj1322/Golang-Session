package session4

import (
	"fmt"
	"math"
)

// Interface shape
type shape interface {
	area() float64
	perimeter() float64
}

// Struct circle
type circle struct {
	radius float64
}

// Struct rectangle
type rectangle struct {
	width  float64
	height float64
}

// Implementasi method area untuk circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implementasi method area untuk rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}

// Implementasi method perimeter untuk circle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Implementasi method perimeter untuk rectangle
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Fungsi untuk mencetak informasi shape
func printShape(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())
}

func Sesi4() {
	c := circle{radius: 5}
	r := rectangle{width: 4, height: 6}

	fmt.Println("Circle:")
	printShape(c)

	fmt.Println("\nRectangle:")
	printShape(r)
}
